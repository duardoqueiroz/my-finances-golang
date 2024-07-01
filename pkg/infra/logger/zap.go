package logger

import (
	"go.uber.org/zap"
)

type zapLogger struct {
	logger *zap.SugaredLogger
}

func NewZapLoggerInstance() (Logger,error) {
	logger, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}

	sugar := logger.Sugar()
	defer logger.Sync()

	return &zapLogger{
		logger: sugar,
	}, nil
}

func (z *zapLogger) InfoF(message string, args ...interface{}) {
	z.logger.Infof(message, args)
}

func (z *zapLogger) ErrorF(message string, args ...interface{}) {
	z.logger.Errorf(message,args)
}

func (z *zapLogger) WarningF(message string, args ...interface{}) {
	z.logger.Warnf(message, args)
}

func (z *zapLogger) Fatalf(message string, args ...interface{}) {
	z.logger.Fatalf(message, args)
}

func (z *zapLogger) Fatalln( args ...interface{}) {
	z.logger.Fatalln(args)
}

func (l *zapLogger) WithFields(fields Fields) Logger {
	var f = make([]interface{}, 0)
	for index, field := range fields {
		f = append(f, index)
		f = append(f, field)
	}

	log := l.logger.With(f...)
	return &zapLogger{logger: log}
}

func (l *zapLogger) WithError(err error) Logger {
	var log = l.logger.With(err.Error())
	return &zapLogger{logger: log}
}


