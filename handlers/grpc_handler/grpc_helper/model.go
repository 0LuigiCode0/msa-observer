package grpc_helper

type Handler interface {
	Close()

	AddMonitor(key, addr string)
	DeleteMonitor(key string) error
}

type MSA interface {
	AddMonitor(key, addr string)
	DeleteMonitor(key string) error
}
