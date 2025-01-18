package log

import (
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	l  *zap.Logger
	s  *zap.SugaredLogger
	mu sync.Mutex
)

func generateLogFileName() string {
	t := time.Now()
	tf := t.Format("2006-01-02")

	return tf + ".log"
}

func createLogFile(filePath string) (*os.File, error) {
	return os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
}

func setupLogger(f *os.File) (*zap.Logger, error) {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000")

	core := zapcore.NewTee(
		zapcore.NewCore(
			zapcore.NewConsoleEncoder(encoderConfig),
			zapcore.Lock(os.Stdout),
			zap.NewAtomicLevelAt(zap.InfoLevel),
		),
		zapcore.NewCore(
			zapcore.NewConsoleEncoder(encoderConfig),
			zapcore.AddSync(f),
			zap.NewAtomicLevelAt(zap.InfoLevel),
		),
	)

	return zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1)), nil
}

func LogInit() error {
	mu.Lock()
	defer mu.Unlock()

	fileName := generateLogFileName()
	filePath := "../log"
	fullPath := filepath.Join(filePath, fileName)

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		if err := os.MkdirAll(filePath, 0755); err != nil {
			return err
		}
	}

	file, err := createLogFile(fullPath)
	if err != nil {
		return err
	}

	l, err = setupLogger(file)
	if err != nil {
		return err
	}

	s = l.Sugar()

	return nil
}

func LogRotation() {
	crontab := cron.New(cron.WithSeconds())
	task := func() {
		LogInit()
	}
	spec := "0 0 0 * * *"
	crontab.AddFunc(spec, task)
	crontab.Start()
	select {}
}

func Gorm() *zap.Logger {
	return l
}

func Info(a ...interface{}) {
	s.Info(a...)
}

func Warn(a ...interface{}) {
	s.Warn(a...)
}

func Debug(a ...interface{}) {
	s.Debug(a...)
}

func Error(a ...interface{}) {
	s.Error(a...)
}

func Panic(a ...interface{}) {
	s.Panic(a...)
}

func Fatal(a ...interface{}) {
	s.Fatal(a...)
}
