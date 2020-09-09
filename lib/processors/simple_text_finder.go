package processors

import (
	"fmt"
	"strings"

	"github.com/otherview/gambaru/core/sessions"
)

type SimpleTextFinderProcessor struct{}

func NewSimpleTextFinderProcessor() *SimpleTextFinderProcessor {
	return &SimpleTextFinderProcessor{}
}

func (processor *SimpleTextFinderProcessor) Execute(session *sessions.Session) error {

	flowfile := session.GetFlowfile()
	if flowfile == nil {
		return nil
	}

	data, err := session.ReadFlowfileData(flowfile)
	if err != nil {
		// TODO yep
		panic(err)
	}

	if strings.Contains(data.(string), "Whiskey") {
		err = session.Remove(flowfile)
		if err != nil {
			// TODO yep
			panic(err)
		}
		fmt.Println("Removed -> ", data)

		err = session.Commit()
		if err != nil {
			// TODO yep
			panic(err)
		}
		return nil
	}

	err = session.TransferFlowfile(flowfile)
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
