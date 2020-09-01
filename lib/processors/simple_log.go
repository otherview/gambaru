package processors

type SimpleLogProcessor struct{}

func NewSimpleLogProcessor() *SimpleLogProcessor {
	return &SimpleLogProcessor{}
}

func (processor *SimpleLogProcessor) Execute(payload interface{}) (interface{}, error) {

	//	fmt.Println(payload)

	return nil, nil
}
