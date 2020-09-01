package processors

import (
	"math/rand"
)

type SimpleTextGeneratorProcessor struct {
	Input  chan []byte
	Output chan []byte
}

func NewSimpleTextGeneratorProcessor() *SimpleTextGeneratorProcessor {
	return &SimpleTextGeneratorProcessor{}
}

func (processor *SimpleTextGeneratorProcessor) Execute(payload interface{}) (interface{}, error) {

	phrases := []string{
		"A river a thousand paces wide stands upon somebody else's legs.",
		"Lucky number slevin wants to go to hell.",
		"A late night does not make any sense.",
		"A sound you heard runs through everything.",
		"Whiskey on the table wants to set things right.",
		"Stupidity has its world rocked by trees (or rocks).",
	}

	return phrases[rand.Intn(len(phrases))], nil
}
