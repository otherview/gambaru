package core

import "github.com/otherview/gambaru/core/flowfile"

type ProcessorInterface interface {
	Execute(payload *flowfiles.Flowfile) (*flowfiles.Flowfile, error)
}

type QueueInterface interface {
	Read() (*flowfiles.Flowfile, error)
	Write(*flowfiles.Flowfile) error
}
