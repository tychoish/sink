package cedar

import (
	"context"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testDatabaseName = "cedar_test"

func TestGlobalEnvironment(t *testing.T) {
	assert.Exactly(t, globalEnv, GetEnvironment())

	first := GetEnvironment()
	first.(*envState).name = "cedar-init"
	assert.Exactly(t, globalEnv, GetEnvironment())

	env, err := NewEnvironment(context.TODO(), "second", &Configuration{MongoDBURI: "mongodb://localhost:27017", NumWorkers: 2, DatabaseName: testDatabaseName})
	assert.NoError(t, err)
	SetEnvironment(env)
	second := GetEnvironment()
	assert.NotEqual(t, first, second)
}

func TestDatabaseSessionAccessor(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var env Environment

	_, _, err := GetSessionWithConfig(env)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "is nil")

	env, err = NewEnvironment(ctx, "test", &Configuration{MongoDBURI: "mongodb://localhost:27017"})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "amboy workers")
	_, _, err = GetSessionWithConfig(env)
	assert.Error(t, err)

	env, err = NewEnvironment(ctx, "test", &Configuration{MongoDBURI: "mongodb://localhost:27017", NumWorkers: 2, DatabaseName: testDatabaseName})
	assert.NoError(t, err)
	env.(*envState).client = nil
	_, _, err = GetSessionWithConfig(env)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "db session")

	env, err = NewEnvironment(ctx, "test", &Configuration{MongoDBURI: "mongodb://localhost:27017", NumWorkers: 2, DatabaseName: testDatabaseName})
	assert.NoError(t, err)
	conf, db, err := GetSessionWithConfig(env)
	assert.NoError(t, err)
	assert.NotNil(t, conf)
	assert.NotNil(t, db)
}

func TestEnvironmentConfiguration(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	const ename = "cedar-test-env"
	for name, test := range map[string]func(t *testing.T, conf *Configuration){
		"ErrorsForInvalidConfig": func(t *testing.T, conf *Configuration) {
			env, err := NewEnvironment(ctx, ename, &Configuration{})
			assert.Error(t, err)
			assert.Nil(t, env)
		},
		"PanicsWithNilConfig": func(t *testing.T, conf *Configuration) {
			assert.Panics(t, func() {
				NewEnvironment(ctx, ename, nil)
			})
		},
		"ErrorsWithMongoDBThatDoesNotExist": func(t *testing.T, conf *Configuration) {
			conf.MongoDBURI = "mongodb://localhost:27016"

			env, err := NewEnvironment(ctx, ename, conf)
			assert.Error(t, err)
			assert.Nil(t, env)
		},
		"VerifyFixtures": func(t *testing.T, conf *Configuration) {
			assert.NotNil(t, conf)
			assert.NoError(t, conf.Validate())
		},
		"ValidConfigUsesLocalConfig": func(t *testing.T, conf *Configuration) {
			conf.DisableRemoteQueue = true

			env, err := NewEnvironment(ctx, ename, conf)
			assert.NoError(t, err)
			q, err := env.GetLocalQueue()
			assert.NoError(t, err)
			assert.NotNil(t, q)
			assert.False(t, strings.Contains(fmt.Sprintf("%T", q), "remote"))
		},
		"DefaultsToRemoteQueueType": func(t *testing.T, conf *Configuration) {
			env, err := NewEnvironment(ctx, ename, conf)
			assert.NoError(t, err)
			q, err := env.GetRemoteQueue()
			assert.NoError(t, err)
			assert.NotNil(t, q)
			assert.True(t, strings.Contains(fmt.Sprintf("%T", q), "remote"))
		},
		// "": func(t *testing.T, conf *Configuration) {},
	} {
		t.Run(name, func(t *testing.T) {
			conf := &Configuration{
				MongoDBURI:         "mongodb://localhost:27017",
				NumWorkers:         2,
				MongoDBDialTimeout: 10 * time.Millisecond,
				SocketTimeout:      time.Second,
				DatabaseName:       testDatabaseName,
			}
			test(t, conf)
		})
	}
}
