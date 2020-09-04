package core

import (
	"github.com/otherview/gambaru/core/flowfiles"
	"github.com/otherview/gambaru/core/sessions"
)

type ProcessorInterface interface {
	Execute(session *sessions.Session) error
}

type QueueInterface interface {
	Read() (*flowfiles.Flowfile, error)
	Write(*flowfiles.Flowfile) error
}
