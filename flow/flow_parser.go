package flow

import "encoding/json"

func Parse(flowJson string) *FlowModel {

	var flow FlowModel
	if err := json.Unmarshal([]byte(flowJson), &flow); err != nil {
		// TODO yep
		panic(err)
	}

	return &flow
}
