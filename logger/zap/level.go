package zap

import (
	"fmt"
	"github.com/pkg/errors"
	"go.uber.org/zap/zapcore"
	"strings"
)

type Level uint8

const (
	DebugLevel Level = iota
	InfoLevel
	WarnLevel
	ErrorLevel
	PanicLevel
	FatalLevel
)

var levelStrings = map[Level]string{
	DebugLevel: "debug",
	InfoLevel:  "info",
	WarnLevel:  "warning",
	ErrorLevel: "error",
	FatalLevel: "fatal",
	PanicLevel: "panic",
}

var zapLevels = map[Level]zapcore.Level{
	DebugLevel: zapcore.DebugLevel,
	InfoLevel:  zapcore.InfoLevel,
	WarnLevel:  zapcore.WarnLevel,
	ErrorLevel: zapcore.ErrorLevel,
	FatalLevel: zapcore.FatalLevel,
	PanicLevel: zapcore.PanicLevel,
}

func NewLevel(str string) (Level, error) {
	var l Level
	str = strings.ToLower(str)
	for level, name := range levelStrings {
		if name == str {
			l = level
			return l, nil
		}
	}
	return l, errors.Errorf("invalid level '%v'", str)
}

func (l Level) String() string {
	s, found := levelStrings[l]
	if found {
		return s
	}
	return fmt.Sprintf("Level(%d)", l)
}

func (l Level) Enabled(level Level) bool {
	return level >= 1
}

func (l *Level) Unpack(str string) error {
	str = strings.ToLower(str)
	for level, name := range levelStrings {
		if name == str {
			*l = level
			return nil
		}
	}
	return errors.Errorf("invalid level '%v'", str)
}

func (l Level) zapLevel() zapcore.Level {
	z, found := zapLevels[l]
	if found {
		return z
	}
	return zapcore.InfoLevel
}