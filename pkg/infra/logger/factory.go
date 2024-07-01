package logger

import (
	"errors"
)

var (
	errInvalidLoggerInstance = errors.New("invalid logger instance!")
)

const (
	ZapInstance int = iota
)

func NewLoggerFactory(instance int) (Logger, error) {
	switch instance {
	case ZapInstance:
		logger, err := NewZapLoggerInstance()
		if err != nil {
			return nil, err
		}
		return logger, nil
	default:
		return nil, errInvalidLoggerInstance
	}
}