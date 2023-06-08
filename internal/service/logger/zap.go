package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Zap struct {
	logger *zap.SugaredLogger
}

func NewLogger() Logger {
	//日志文件名称
	fileName := "go-admin.log"
	syncWriter := zapcore.AddSync(
		&lumberjack.Logger{
			Filename: fileName, //文件名称
			MaxSize:  2,        //MB
			//MaxAge:     0,
			MaxBackups: 0, //最大备份
			LocalTime:  true,
			Compress:   true, //是否启用压缩
		})
	//编码
	encoder := zap.NewProductionEncoderConfig()
	//时间格式
	encoder.EncodeTime = zapcore.ISO8601TimeEncoder
	core := zapcore.NewCore(
		// 编码器
		zapcore.NewJSONEncoder(encoder),
		syncWriter,
		zap.NewAtomicLevelAt(zap.DebugLevel))
	log := zap.New(
		core,
		zap.AddCaller(),
		zap.AddCallerSkip(1))
	var l Logger = log.Sugar()
	return l
}

func (z *Zap) Debug(args ...interface{}) {
	z.logger.Debug(args)
}

func (z *Zap) Debugf(template string, args ...interface{}) {
	z.logger.Debugf(template, args...)
}

func (z *Zap) Info(args ...interface{}) {
	z.logger.Info(args...)
}

func (z *Zap) Infof(template string, args ...interface{}) {
	z.logger.Infof(template, args...)
}

func (z *Zap) Warn(args ...interface{}) {
	z.logger.Warn(args...)
}

func (z *Zap) Warnf(template string, args ...interface{}) {
	z.logger.Warnf(template, args...)
}

func (z *Zap) Error(args ...interface{}) {
	z.logger.Error(args...)
}

func (z *Zap) Errorf(template string, args ...interface{}) {
	z.logger.Errorf(template, args...)
}

func (z *Zap) Fatal(args ...interface{}) {
	z.logger.Fatal(args...)
}

func (z *Zap) Fatalf(template string, args ...interface{}) {
	z.logger.Fatalf(template, args...)
}
