package processors

type ProcessorInterface interface {
	Execute(payload interface{}) (interface{}, error)
}

//type ProcessorInputOutputInterface interface {
//	Input
//}
