package queues

type QueueInterface interface {
	Read() (interface{}, error)
	Write(interface{}) error
}
