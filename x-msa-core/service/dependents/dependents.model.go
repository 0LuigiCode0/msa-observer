package dependents

import (
	"sync"
	"x-msa-core/service/client"
)

type ServiceDependentsManager interface {
	Add(key, addr, group string) error
	Get(key string) (client.ServiceClient, error)
	Delete(key string) error
}

type serviceDependents struct {
	dep map[string]client.ServiceClient
	rw  sync.Mutex
}
