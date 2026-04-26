package port

type Logger interface {
	Debug(msg string, keysAndValues ...interface{})
	Info(msg string, keysAndValues ...interface{})
	Warn(msg string, keysAndValues ...interface{})
	Error(msg string, err error, keysAndValues ...interface{})
	With(keysAndValues ...interface{}) Logger // コンテキストを付与した新しいロガーを生成する
}
