package main

import (
	"github.com/posthog/posthog-go"
)

func main() {
	client, _ := posthog.NewWithConfig("phc_5x6Jl9AYh65uUievqzRNdj1eoGXqQqJPlHNpcSMFHSs", posthog.Config{Endpoint: "https://us.i.posthog.com"})
	defer client.Close()

	client.Enqueue(posthog.Capture{
		DistinctId: "test-user",
		Event:      "test-snippet",
	})
}
