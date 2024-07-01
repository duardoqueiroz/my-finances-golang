package logger

type Logger interface {
	InfoF(message string, args ...interface{})
	WarningF(message string, args ...interface{})
	ErrorF(message string, args ...interface{})
  Fatalf(message string, args ...interface{}) 
	Fatalln(args ...interface{})
	WithFields(keyValues Fields) Logger
	WithError(err error) Logger
}

type Fields map[string]interface{}
