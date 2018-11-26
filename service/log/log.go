package log

import (
	"github.com/sirupsen/logrus"
)

type Fields = logrus.Fields

var log = logrus.New()

// Adds a field to the log entry, note that it doesn't log until you call
// Debug, Print, Info, Warn, Fatal or Panic. It only creates a log entry.
// If you want multiple fields, use `WithFields`.
func WithField(key string, value interface{}) *logrus.Entry {
	return log.WithField(key, value)
}

// Adds a struct of fields to the log entry. All it does is call `WithField` for
// each `Field`.
func WithFields(fields Fields) *logrus.Entry {
	return log.WithFields(fields)
}

// Add an error as single field to the log entry.  All it does is call
// `WithError` for the given `error`.
func WithError(err error) *logrus.Entry {
	return log.WithError(err)
}

func Debugf(format string, args ...interface{}) {
	log.Debugf(format, args...)
}

func Infof(format string, args ...interface{}) {
	log.Infof(format, args...)
}

func Printf(format string, args ...interface{}) {
	log.Printf(format, args...)
}

func Warnf(format string, args ...interface{}) {
	log.Warnf(format, args...)
}

func Warningf(format string, args ...interface{}) {
	log.Warningf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	log.Errorf(format, args...)
}

func Fatalf(format string, args ...interface{}) {
	log.Fatalf(format, args...)
}

func Panicf(format string, args ...interface{}) {
	log.Panicf(format, args...)
}

func Debug(args ...interface{}) {
	log.Debug(args...)
}

func Info(args ...interface{}) {
	log.Info(args...)
}

func Print(args ...interface{}) {
	log.Print(args...)
}

func Warn(args ...interface{}) {
	log.Warn(args...)
}

func Warning(args ...interface{}) {
	log.Warning(args...)
}

func Error(args ...interface{}) {
	log.Error(args...)
}

func Fatal(args ...interface{}) {
	log.Fatal(args...)
}

func Panic(args ...interface{}) {
	log.Panic(args...)
}

func Debugln(args ...interface{}) {
	log.Debugln(args...)
}

func Infoln(args ...interface{}) {
	log.Infoln(args...)
}

func Println(args ...interface{}) {
	log.Println(args...)
}

func Warnln(args ...interface{}) {
	log.Warnln(args...)
}

func Warningln(args ...interface{}) {
	log.Warningln(args...)
}

func Errorln(args ...interface{}) {
	log.Errorln(args...)
}

func Fatalln(args ...interface{}) {
	log.Fatalln(args...)
}

func Panicln(args ...interface{}) {
	log.Panicln(args...)
}

//When file is opened with appending mode, it's safe to
//write concurrently to a file (within 4k message on Linux).
//In these cases user can choose to disable the lock.
func SetNoLock() {
	log.SetNoLock()
}

func Level() logrus.Level {
	return log.Level
}

func SetLevel(level logrus.Level) {
	log.SetLevel(level)
}

func AddHook(hook logrus.Hook) {
	log.AddHook(hook)
}
