package zap

import (
	"go.uber.org/zap/zapcore"
	"time"
)

type Config struct {
	Name      string
	Json      bool
	Level     Level
	Metas     []zapcore.Field
	Selectors []string

	ToStderr    bool
	ToSyslog    bool
	ToFiles     bool
	ToEventLog  bool
	toObserver  bool
	toIODiscard bool
	Files       FileConfig
	addCaller   bool
	develop     bool
}

type FileConfig struct {
	Path            string
	Name            string
	MaxSize         uint
	MaxBackups      uint
	Permissions     uint32
	Interval        time.Duration
	RotateOnStartup bool
	RedirectStderr  bool
}

var defaultConfig = Config{
	Level:   InfoLevel,
	ToFiles: true,
	Files: FileConfig{
		MaxSize:         10 * 1024 * 1024,
		MaxBackups:      7,
		Permissions:     0600,
		Interval:        0,
		RotateOnStartup: true,
	},
	addCaller: true,
}

func DefaultConfig() Config {
	return defaultConfig
}
