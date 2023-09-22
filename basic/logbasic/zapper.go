package logbasic

import "go.uber.org/zap"

type ZapLogger struct {
	logger *zap.Logger
}

func (z ZapLogger) Debug(msg string) {
	z.logger.Debug(msg)
}

func (z ZapLogger) Info(msg string) {
	z.logger.Info(msg)
}
func (z ZapLogger) Warning(msg string) {
	z.logger.Warn(msg)
}
func (z ZapLogger) Error(msg string) {
	z.logger.Error(msg)
}
func (z ZapLogger) Fatal(msg string) {
	z.logger.Fatal(msg)
}

func NewZapLogger() *ZapLogger {
	z, _ := zap.NewDevelopment()
	return &ZapLogger{
		z,
	}
}
