package processors

import (
	"fmt"

	"github.com/otherview/gambaru/core/sessions"
)

type SimpleLogProcessor struct{}

func NewSimpleLogProcessor() *SimpleLogProcessor {
	return &SimpleLogProcessor{}
}

func (processor *SimpleLogProcessor) Execute(session *sessions.Session) error {

	flowfile := session.GetFlowfile()
	if flowfile == nil {
		return nil
	}

	data, err := session.ReadFlowfileData(flowfile)
	if err != nil {
		// TODO yep
		panic(err)
	}
	fmt.Println("Logging ->      ", flowfile.ID, data)

	err = session.Remove(flowfile)
	if err != nil {
		// TODO yep
		panic(err)
	}

	err = session.Commit()
	if err != nil {
		// TODO yep
		panic(err)
	}

	return nil
}
