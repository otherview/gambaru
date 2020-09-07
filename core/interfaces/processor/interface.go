package interface_processor

import "github.com/otherview/gambaru/core/sessions"

type ProcessorInterface interface {
	Execute(session *sessions.Session) error
}
