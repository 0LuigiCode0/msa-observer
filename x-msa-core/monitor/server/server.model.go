package server

import (
	"x-msa-core/grpc/msa_monitor"
	dep_observer "x-msa-core/observer/dependents"
	dep_service "x-msa-core/service/dependents"

	"google.golang.org/grpc"
)

type MonitorServer interface {
	Start() error

	Observers() dep_observer.ObserverDependentsManager
	Services() dep_service.ServiceDependentsManager

	Close()
}

type monitorServer struct {
	msa_monitor.MonitorServer

	server *grpc.Server
	key    string
	addr   string

	observers dep_observer.ObserverDependentsManager
	services  dep_service.ServiceDependentsManager
}
