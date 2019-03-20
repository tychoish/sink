package certdepot

import (
	"context"
	"time"

	"github.com/mongodb/grip"
	"github.com/mongodb/grip/message"
	"github.com/pkg/errors"
	"github.com/square/certstrap/depot"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

type mongoDepot struct {
	ctx            context.Context
	client         *mongo.Client
	databaseName   string
	collectionName string
	expireAfter    time.Duration
}

func NewMongoDBCertDepot(ctx context.Context, opts MongoDBOptions) (depot.Depot, error) {
	if err := opts.validate(); err != nil {
		return nil, errors.Wrap(err, "invalid options")
	}

	connctx, cancel := context.WithTimeout(ctx, opts.MongoDBDialTimeout)
	defer cancel()
	client, err := mongo.Connect(connctx, options.Client().ApplyURI(opts.MongoDBURI))
	if err != nil {
		return nil, errors.Wrap(err, "problem connecting to database")
	}

	return &mongoDepot{
		ctx:            ctx,
		client:         client,
		databaseName:   opts.DatabaseName,
		collectionName: opts.CollectionName,
		expireAfter:    opts.ExpireAfter,
	}, nil
}

func NewMongoDBCertDepotWithClient(ctx context.Context, client *mongo.Client, opts MongoDBOptions) (depot.Depot, error) {
	if client == nil {
		return nil, errors.New("must specify a non-nil client")
	}

	if err := opts.validate(); err != nil {
		return nil, errors.Wrap(err, "invalid options")
	}

	return &mongoDepot{
		ctx:            ctx,
		client:         client,
		databaseName:   opts.DatabaseName,
		collectionName: opts.CollectionName,
		expireAfter:    opts.ExpireAfter,
	}, nil
}

// Put inserts the data into the document specified by the tag.
func (m *mongoDepot) Put(tag *depot.Tag, data []byte) error {
	if data == nil {
		return errors.New("data is nil")
	}

	name, key := getNameAndKey(tag)

	update := bson.M{"$set": bson.M{key: string(data)}}
	if key == userCertKey {
		update["$set"].(bson.M)[userTTLKey] = time.Now()
	}

	res, err := m.client.Database(m.databaseName).Collection(m.collectionName).UpdateOne(m.ctx,
		bson.D{{"_id", name}},
		update,
		options.Update().SetUpsert(true))
	if err != nil {
		return errors.Wrap(err, "problem adding data to the database")
	}
	grip.Debug(message.Fields{
		"db":     m.databaseName,
		"coll":   m.collectionName,
		"id":     name,
		"change": res,
		"op":     "put",
	})

	return nil
}

// Check returns whether the user and data specified by the tag exists.
func (m *mongoDepot) Check(tag *depot.Tag) bool {
	name, key := getNameAndKey(tag)

	u := &User{}

	err := m.client.Database(m.databaseName).Collection(m.collectionName).FindOne(m.ctx, bson.D{{"_id", name}}).Decode(u)
	grip.WarningWhen(errNotNotFound(err), message.Fields{
		"db":   m.databaseName,
		"coll": m.collectionName,
		"id":   name,
		"err":  err,
		"op":   "check",
	})

	switch key {
	case userCertKey:
		return u.Cert != ""
	case userPrivateKeyKey:
		return u.PrivateKey != ""
	case userCertReqKey:
		return u.CertReq != ""
	case userCertRevocListKey:
		return u.CertRevocList != ""
	default:
		return false
	}
}

// Get reads the data for the user specified by tag. Returns an error if the
// user does not exist, if the TTL has expired (for certs), or if the data is
// empty.
func (m *mongoDepot) Get(tag *depot.Tag) ([]byte, error) {
	name, key := getNameAndKey(tag)

	u := &User{}
	err := m.client.Database(m.databaseName).Collection(m.collectionName).FindOne(m.ctx, bson.D{{"_id", name}}).Decode(u)
	if err == mongo.ErrNoDocuments {
		return nil, errors.Errorf("could not find %s in the database", name)
	}
	if err != nil {
		return nil, errors.Wrapf(err, "problem looking up %s in the database", name)
	}

	var data []byte
	switch key {
	case userCertKey:
		data = []byte(u.Cert)
		if len(data) > 0 && time.Since(u.TTL) > m.expireAfter {
			return nil, errors.Errorf("certificate for %s has expired", name)
		}
	case userPrivateKeyKey:
		data = []byte(u.PrivateKey)
	case userCertReqKey:
		data = []byte(u.CertReq)
	case userCertRevocListKey:
		data = []byte(u.CertRevocList)
		if len(data) > 0 && time.Since(u.TTL) > m.expireAfter {
			return nil, errors.Errorf("certificate revocation list for %s has expired", name)
		}
	}

	if len(data) == 0 {
		return nil, errors.New("no data available")
	}
	return data, nil
}

// Delete removes the data from a user specified by the tag.
func (m *mongoDepot) Delete(tag *depot.Tag) error {
	name, key := getNameAndKey(tag)

	_, err := m.client.Database(m.databaseName).Collection(m.collectionName).UpdateOne(m.ctx,
		bson.D{{"_id", name}},
		bson.M{"$unset": bson.M{key: ""}})

	if errNotNotFound(err) {
		return errors.Wrapf(err, "problem deleting %s.%s from the database", name, key)
	}

	return nil
}
