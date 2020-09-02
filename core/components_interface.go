package core

type ProcessorInterface interface {
	Execute(payload interface{}) (interface{}, error)
}

type QueueInterface interface {
	Read() (interface{}, error)
	Write(interface{}) error
}
