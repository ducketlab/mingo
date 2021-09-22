package logger

type Logger interface {
	StandardLogger
	FormatLogger
	CompatibleLogger
	RecoveryLogger
	WithMetaLogger
}

type StandardLogger interface {
	Debug(msg ...interface{})
	Info(msg ...interface{})
	Warn(msg ...interface{})
	Error(msg ...interface{})
	Fatal(msg ...interface{})
	Panic(msg ...interface{})
}

type FormatLogger interface {
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Panicf(format string, args ...interface{})
}

type CompatibleLogger interface {
	Print(v ...interface{})
	Printf(format string, v ...interface{})
	Println(v ...interface{})
}

type RecoveryLogger interface {
	Recover(msg string)
}

type WithMetaLogger interface {
	Debugw(msg string, fields ...Field)
	Infow(msg string, fields ...Field)
	Warnw(msg string, fields ...Field)
	Errorw(msg string, fields ...Field)
	Fatalw(msg string, fields ...Field)
	Panicw(msg string, fields ...Field)
}
