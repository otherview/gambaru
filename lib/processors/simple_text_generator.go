package processors

import (
	"math/rand"

	"fmt"

	"github.com/otherview/gambaru/core/repository"

	"github.com/otherview/gambaru/core/flowfile"
)

type SimpleTextGeneratorProcessor struct {
	repository *repository.Repository
}

func NewSimpleTextGeneratorProcessor(repository *repository.Repository) *SimpleTextGeneratorProcessor {
	return &SimpleTextGeneratorProcessor{repository: repository}
}

func (processor *SimpleTextGeneratorProcessor) Execute(flowfile *flowfiles.Flowfile) (*flowfiles.Flowfile, error) {

	phrases := []string{
		"A river a thousand paces wide stands upon somebody else's legs.",
		"Lucky number slevin wants to go to hell.",
		"A late night does not make any sense.",
		"A sound you heard runs through everything.",
		"Whiskey on the table wants to set things right.",
		"Stupidity has its world rocked by trees (or rocks).",
	}

	newflowFile := flowfiles.NewFlowfile()

	phrase := phrases[rand.Intn(len(phrases))]
	processor.repository.Write(newflowFile, phrase)
	fmt.Println("Created -> ", phrase)

	return newflowFile, nil
}
