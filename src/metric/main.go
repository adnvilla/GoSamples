package main

import (
	"log"
	"time"

	"github.com/DataDog/datadog-go/statsd"
)

func main() {
	statsd, err := statsd.New("127.0.0.1:8125")
	if err != nil {
		log.Fatal(err)
	}

	for {

		statsd.Tags = append(statsd.Tags, "service:qwerty-poiuytre")
		err := statsd.Count("example_metric.histogram", int64(0), statsd.Tags, 1)
		//statsd.Tags = []string{}
		if err != nil {
			log.Println(err)
		}
		log.Println("Done...")
		time.Sleep(time.Duration(0) * time.Second)
	}
}
