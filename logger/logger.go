package logger

import (
	"os"

	"expense_tables/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var logLevelSeverity = map[zapcore.Level]string{
	zapcore.DebugLevel:  "DEBUG",
	zapcore.InfoLevel:   "INFO",
	zapcore.WarnLevel:   "WARNING",
	zapcore.ErrorLevel:  "ERROR",
	zapcore.DPanicLevel: "CRITICAL",
	zapcore.PanicLevel:  "PANIC",
	zapcore.FatalLevel:  "FATAL",
}

func NewLogger(cfg *config.Config) *zap.Logger {

	// если файл для логов необходимо переписывать
	// и если файл существует, то он удаляется  и создаётся заново;
	// параметр RewriteLog приравнивается false
	if cfg.RewriteLog {
		if _, err := os.Stat(cfg.LogFile); err == nil {
			os.Remove(cfg.LogFile)
			os.Create(cfg.LogFile)
			cfg.RewriteLog = false
		}
	}

	l := logger{
		LogLevel:        cfg.LogLevel,
		LogFileEnable:   cfg.LogFileEnable,
		LogStdoutEnable: cfg.LogStdoutEnable,
		LogFile:         cfg.LogFile,
		MaxSize:         cfg.MaxSize,
		MaxAge:          cfg.MaxAge,
		MaxBackups:      cfg.MaxBackups,
	}
	return l.initLogger(cfg)
}

func (l *logger) initLogger(cfg *config.Config) *zap.Logger {

	li := Logger(l)
	conf := li.newProductionEncoderConfig()
	conf.EncodeTime = zapcore.ISO8601TimeEncoder
	conf.EncodeLevel = li.customEncoderLevel
	conf.MessageKey = "message"
	conf.CallerKey = "caller"
	conf.TimeKey = "time"

	jsonEncoder := li.newJSONEncode(conf)
	textEncoder := li.newConsoleEncoder(conf)

	if !cfg.LogFileEnable && !cfg.LogStdoutEnable {
		return nil
	}

	var fileLogger zapcore.Core
	if cfg.LogFileEnable {
		fileLogger = li.newCore(
			jsonEncoder,
			li.addSync(&lumberjack.Logger{
				Filename:   l.LogFile,
				MaxSize:    l.MaxSize,
				MaxAge:     l.MaxAge,
				MaxBackups: l.MaxBackups,
			}),
			li.logLevel(cfg),
		)
	}

	var consoleLogger zapcore.Core
	if cfg.LogStdoutEnable {
		consoleLogger = li.newCore(
			textEncoder,
			zapcore.AddSync(os.Stdout),
			li.logLevel(cfg),
		)
	}

	if !cfg.LogFileEnable && cfg.LogStdoutEnable {
		return li.new(li.newTee(consoleLogger))
	} else if cfg.LogFileEnable && !cfg.LogStdoutEnable {
		return li.new(li.newTee(fileLogger))
	}
	return li.new(li.newTee(li.loggers(cfg, consoleLogger, fileLogger)...), li.zapOpts()...)
}
