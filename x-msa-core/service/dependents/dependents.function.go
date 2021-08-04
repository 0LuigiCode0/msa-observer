package dependents

import (
	"fmt"
	"x-msa-core/helper"
	"x-msa-core/service/client"
)

func NewSDM() ServiceDependentsManager {
	return &serviceDependents{
		dep: map[string]client.ServiceClient{},
	}
}

func (s *serviceDependents) Add(key, addr, group string) error {
	client, err := client.NewServiceClient(addr, helper.GroupsType(group))
	if err != nil {
		return fmt.Errorf("cannot craete service client: %v", err)
	}

	s.rw.Lock()
	defer s.rw.Unlock()

	if _, ok := s.dep[key]; ok {
		return fmt.Errorf("service exist by key: %v", key)
	}
	s.dep[key] = client

	return nil
}

func (s *serviceDependents) Get(key string) (client.ServiceClient, error) {
	s.rw.Lock()
	defer s.rw.Unlock()

	if c, ok := s.dep[key]; ok {
		return c, nil
	} else {
		return nil, fmt.Errorf("service not found by key: %v", key)
	}
}

func (s *serviceDependents) Delete(key string) error {
	s.rw.Lock()
	defer s.rw.Unlock()

	if c, ok := s.dep[key]; ok {
		c.Close()
		delete(s.dep, key)
		return nil
	} else {
		return fmt.Errorf("service not found by key: %v", key)
	}
}
