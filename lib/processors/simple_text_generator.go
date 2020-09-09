package processors

import (
	"fmt"
	"math/rand"

	"github.com/otherview/gambaru/core/sessions"

	"github.com/otherview/gambaru/core/flowfiles"
)

type SimpleTextGeneratorProcessor struct{}

func NewSimpleTextGeneratorProcessor() *SimpleTextGeneratorProcessor {
	return &SimpleTextGeneratorProcessor{}
}

func (processor *SimpleTextGeneratorProcessor) Execute(session *sessions.Session) error {

	start := []string{
		"A river a thousand paces wide stands upon somebody else's legs,",
		"Lucky number slevin wants to go to hell,",
		"A late night does not make any sense,",
		"A sound you heard runs through everything,",
		"Whiskey on the table wants to set things right,",
		"Stupidity has its world rocked by trees (or rocks),",
	}
	middle := []string{
		" but ",
		" and ",
		" however ",
		" while doing so",
		" nevertheless",
	}
	phrase := start[rand.Intn(len(start))] + middle[rand.Intn(len(middle))] + start[rand.Intn(len(start))]

	newflowFile := flowfiles.NewFlowfile()

	_ = session.WriteFlowfileData(newflowFile, phrase)
	fmt.Println("Created -> ", phrase)

	_ = session.TransferFlowfile(newflowFile)
	_ = session.Commit()

	return nil
}
