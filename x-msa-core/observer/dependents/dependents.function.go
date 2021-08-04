package dependents

import (
	"fmt"
	"x-msa-core/observer/client"
)

func NewODM() ObserverDependentsManager {
	return &observerDependents{
		dep: map[string]client.ObserverClient{},
	}
}

func (o *observerDependents) Add(key, addr string) error {
	client, err := client.NewObserverClient(addr)
	if err != nil {
		return fmt.Errorf("cannot craete observer client: %v", err)
	}

	o.rw.Lock()
	defer o.rw.Unlock()

	if _, ok := o.dep[key]; ok {
		return fmt.Errorf("observer exist by key: %v", key)
	}
	o.dep[key] = client

	return nil
}

func (o *observerDependents) Get(key string) (client.ObserverClient, error) {
	o.rw.Lock()
	defer o.rw.Unlock()

	if c, ok := o.dep[key]; ok {
		return c, nil
	} else {
		return nil, fmt.Errorf("observer not found by key: %v", key)
	}
}

func (o *observerDependents) Delete(key string) error {
	o.rw.Lock()
	defer o.rw.Unlock()

	if c, ok := o.dep[key]; ok {
		c.Close()
		delete(o.dep, key)
		return nil
	} else {
		return fmt.Errorf("observer not found by key: %v", key)
	}
}
