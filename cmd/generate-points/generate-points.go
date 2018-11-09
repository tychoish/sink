package main

import (
	"math/rand"
	"os"
	"time"

	"github.com/evergreen-ci/sink/model"
	"github.com/mongodb/ftdc"
	"github.com/mongodb/grip"
	"github.com/mongodb/grip/message"
	"github.com/mongodb/mongo-go-driver/bson"
)

func main() {
	point := model.PerformancePoint{}
	startAt := time.Now()
	file, err := os.Create("metrics_mock.ftdc")
	grip.EmergencyFatal(err)

	collector := ftdc.NewStreamingCollector(1000, file)
	defer func() { grip.EmergencyFatal(file.Close()) }()
	defer func() { ftdc.FlushCollector(collector, file) }()

	totalOps := int64(44 * 1000 * 60 * 60 * 15)

	for i := int64(1); i <= totalOps; i++ {
		point.Timestamp = startAt.Add(time.Duration(i) * time.Second)
		point.Counters.Number = i
		point.Counters.Operations += i * rand.Int63n(10)
		point.Counters.Size += i * rand.Int63n(1000)

		point.Timers.Duration += time.Millisecond * time.Duration(i-1*rand.Int63n(100))
		point.Timers.Total += time.Millisecond * time.Duration(i*rand.Int63n(100))

		point.Guages.Workers = 1
		if i&1000 == 0 {
			point.Guages.State++
		}
		grip.InfoWhen(i%100000 == 0,
			message.Fields{
				"n":            i,
				"t":            totalOps,
				"info.metrics": collector.Info().MetricsCount,
				"info.samples": collector.Info().SampleCount,
			})

		payload, err := bson.Marshal(point)
		grip.EmergencyFatal(err)

		doc, err := bson.ReadDocument(payload)
		grip.EmergencyFatal(err)

		err = collector.Add(doc)
		grip.EmergencyFatal(err)
	}
}
