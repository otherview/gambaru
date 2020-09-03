package processors

import (
	"fmt"

	"github.com/otherview/gambaru/core/flowfile"
	"github.com/otherview/gambaru/core/repository"
)

type SimpleLogProcessor struct {
	repository *repository.Repository
}

func NewSimpleLogProcessor(repository *repository.Repository) *SimpleLogProcessor {
	return &SimpleLogProcessor{repository: repository}
}

func (processor *SimpleLogProcessor) Execute(flowfile *flowfiles.Flowfile) (*flowfiles.Flowfile, error) {

	if flowfile == nil {
		return nil, nil
	}

	data, err := processor.repository.Read(flowfile)
	if err != nil {
		// TODO yep
		panic(err)
	}
	fmt.Println("Logging -> ", data)

	return nil, nil
}
