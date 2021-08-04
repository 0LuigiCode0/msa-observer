package dependents

import (
	"fmt"
	"x-msa-core/monitor/client"
)

func NewMDM() MonitorDependentsManager {
	return &monitorDependents{
		dep: map[string]client.MonitorClient{},
	}
}

func (m *monitorDependents) Add(key, addr string) error {
	client, err := client.NewMonitorClient(addr)
	if err != nil {
		return fmt.Errorf("cannot craete monitor client: %v", err)
	}

	m.rw.Lock()
	defer m.rw.Unlock()

	if _, ok := m.dep[key]; ok {
		return fmt.Errorf("monitor exist by key: %v", key)
	}
	m.dep[key] = client

	return nil
}

func (m *monitorDependents) Get(key string) (client.MonitorClient, error) {
	m.rw.Lock()
	defer m.rw.Unlock()

	if c, ok := m.dep[key]; ok {
		return c, nil
	} else {
		return nil, fmt.Errorf("monitor not found by key: %v", key)
	}
}

func (m *monitorDependents) Delete(key string) error {
	m.rw.Lock()
	defer m.rw.Unlock()

	if c, ok := m.dep[key]; ok {
		c.Close()
		delete(m.dep, key)
		return nil
	} else {
		return fmt.Errorf("monitor not found by key: %v", key)
	}
}
