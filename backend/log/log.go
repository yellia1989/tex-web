package log

import (
    "io"
    "fmt"
    "encoding/json"
    llog "github.com/yellia1989/tex-go/tools/log"
    "github.com/labstack/echo"
    "github.com/labstack/gommon/log"
)

var mylog *myLogger

func init() {
    mylog = &myLogger{logger: llog.GetDefaultLogger()}
    mylog.logger.SetLogName("web")
    mylog.logger.SetFileRoller("./", 10, 50)
}

type myLogger struct {
    logger *llog.Logger
}

func GetLogger() echo.Logger {
    return mylog
}

func (l *myLogger) Output() io.Writer {
    return l.logger.GetWriter()
}

type myWriter struct {
    io.Writer
}
func (w myWriter) Write(v []byte) (n int, err error) {
    return w.Write(v)
}
func (w myWriter) NeedPrefix() bool {
    return true
}

func (l *myLogger) SetOutput(w io.Writer) {
    l.logger.SetWriter(myWriter{w})
}

func (l *myLogger) Prefix() string {
    return ""
}

func (l *myLogger) SetPrefix(p string) {
}

func (l *myLogger) Level() (lv log.Lvl) {
    level := llog.GetLevel()
    switch level {
    case llog.DEBUG: lv = log.DEBUG
    case llog.INFO: lv = log.INFO
    case llog.WARN: lv = log.WARN
    case llog.ERROR: lv = log.ERROR
    default: panic("unknown log level")
    }

    return
}

func (l *myLogger) SetLevel(lv log.Lvl) {
    level := llog.DEBUG
    switch lv {
    case log.DEBUG: level = llog.DEBUG
    case log.INFO: level = llog.INFO
    case log.WARN: level = llog.INFO
    case log.ERROR: level = llog.ERROR
    default: panic("unsupport log level")
    }
    llog.SetLevel(level)
}

func (l *myLogger) SetHeader(h string) {
}

func (l *myLogger) Print(i ...interface{}) {
    llog.Debug(i...)
}

func (l *myLogger) Printf(format string, args ...interface{}) {
    llog.Debugf(format, args...)
}

func (l *myLogger) Printj(j log.JSON) {
    json, _ := json.Marshal(j)
    llog.Debug(json)
}

func (l *myLogger) Debug(i ...interface{}) {
    l.logger.Debug(i...)
}

func (l *myLogger) Debugf(format string, args ...interface{}) {
    l.logger.Debugf(format, args...)
}

func (l *myLogger) Debugj(j log.JSON) {
    json, _ := json.Marshal(j)
    l.logger.Debug(json)
}

func (l *myLogger) Info(i ...interface{}) {
    l.logger.Info(i...)
}

func (l *myLogger) Infof(format string, args ...interface{}) {
    l.logger.Infof(format, args...)
}

func (l *myLogger) Infoj(j log.JSON) {
    json, _ := json.Marshal(j)
    l.logger.Info(json)
}

func (l *myLogger) Warn(i ...interface{}) {
    l.logger.Warn(i...)
}

func (l *myLogger) Warnf(format string, args ...interface{}) {
    l.logger.Warnf(format, args...)
}

func (l *myLogger) Warnj(j log.JSON) {
    json, _ := json.Marshal(j)
    l.logger.Warn(json)
}

func (l *myLogger) Error(i ...interface{}) {
    l.logger.Error(i...)
}

func (l *myLogger) Errorf(format string, args ...interface{}) {
    l.logger.Errorf(format, args...)
}

func (l *myLogger) Errorj(j log.JSON) {
    json, _ := json.Marshal(j)
    l.logger.Error(json)
}

func (l *myLogger) Fatal(i ...interface{}) {
    l.Error(i...)
}

func (l *myLogger) Fatalj(j log.JSON) {
    l.Errorj(j)
}

func (l *myLogger) Fatalf(format string, args ...interface{}) {
    l.Errorf(format, args...)
}

func (l *myLogger) Panic(i ...interface{}) {
    l.logger.Error(i...)
    panic(fmt.Sprint(i...))
}

func (l *myLogger) Panicj(j log.JSON) {
    json,_ := json.Marshal(j)
    l.logger.Error(json)
    panic(fmt.Sprint(json))
}

func (l *myLogger) Panicf(format string, args ...interface{}) {
    l.logger.Errorf(format, args...)
    panic(fmt.Sprintf(format, args...))
}
