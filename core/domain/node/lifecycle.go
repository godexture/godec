package node

type Lifecycle interface {
	Init() error
	Flush() error
	Close() error
}
