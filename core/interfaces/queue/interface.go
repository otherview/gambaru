package interface_queue

import "github.com/otherview/gambaru/core/flowfiles"

type QueueInterface interface {
	Read() (*flowfiles.Flowfile, error)
	Write(*flowfiles.Flowfile) error
}
