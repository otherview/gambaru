package main

import (
	console "github.com/AsynkronIT/goconsole"
	"github.com/otherview/gambaru/core/flowchart"
	"github.com/otherview/gambaru/core/silo"
)

func main() {
	//simple_flow()
	simple_flow_json()
}

func simple_flow_json() {

	var flowJson = `
{
  "queues": [
    {
      "id": "a466f639-ac5a-4ccf-9d9d-a907faad47ae",
      "type": "SimpleQueue"
    },
	{
      "id": "55590ec0-e5af-4969-8599-66f3dbdc2c8a",
      "type": "SimpleQueue"
    }
  ],
  "processors": [
    {
      "id": "3dea85cf-a8b0-472e-8efc-e7ae7f5dd767",
      "type": "SimpleLogProcessor",
	  "config": {
	  	"executionTime": "10ms",
		"savePoint": true
	  },
      "inputQueues": [
		"a466f639-ac5a-4ccf-9d9d-a907faad47ae"
		]
    },
    {
      "id": "a204925e-4299-4bf6-b83c-f34649d90cf8",
      "type": "SimpleTextGeneratorProcessor",
	  "config": {
		"savePoint": true
	  },
      "outputQueues": [ 
		"55590ec0-e5af-4969-8599-66f3dbdc2c8a" 
		]
    },
	{
      "id": "7030f293-f9d3-469b-b30c-6139bbe97512",
      "type": "SimpleTextGeneratorProcessor",
	  "config": {
		"savePoint": true
	  },
      "outputQueues": [ 
		"55590ec0-e5af-4969-8599-66f3dbdc2c8a" 
	  ]
    },
	{
      "id": "369e46f0-1f94-4f5d-91ab-5ae8af1f7c12",
      "type": "SimpleTextFinderProcessor",
	  "config": {
	  	"executionTime": "10ms"
	  },
	  "inputQueues": [
		"55590ec0-e5af-4969-8599-66f3dbdc2c8a"
		],
      "outputQueues": [
		"a466f639-ac5a-4ccf-9d9d-a907faad47ae"
		]
    }
  ]
}
`

	silo := silo.NewSilo()
	newFlow := flowchart.NewFlow(silo)
	if err := newFlow.LoadJson(flowJson); err != nil {
		panic(err)
	}
	newFlow.CreateFlow()

	silo.Start()
	console.ReadLine()
	silo.Stop()
	console.ReadLine()

	return
}
