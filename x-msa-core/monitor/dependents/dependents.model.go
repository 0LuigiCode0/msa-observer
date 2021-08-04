package dependents

import (
	"sync"
	"x-msa-core/monitor/client"
)

type MonitorDependentsManager interface {
	Add(key, addr string) error
	Get(key string) (client.MonitorClient, error)
	Delete(key string) error
}

type monitorDependents struct {
	dep map[string]client.MonitorClient
	rw  sync.Mutex
}
