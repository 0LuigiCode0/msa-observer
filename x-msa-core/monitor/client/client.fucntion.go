package client

import (
	"fmt"
	"x-msa-core/grpc/msa_monitor"
	"x-msa-core/helper"
)

func NewMonitorClient(addr string) (MonitorClient, error) {
	var err error
	client := &monitorClient{}

	client.conn, client.ctx, client.close, err = helper.CreateConn(addr)
	if err != nil {
		return nil, fmt.Errorf("cannot create conn: %v", err)
	}
	client.client = msa_monitor.NewMonitorClient(client.conn)

	return client, nil
}

func (m *monitorClient) Close() { m.close(); m.conn.Close() }

func (m *monitorClient) AddService(req *msa_monitor.RequestAddService) (*msa_monitor.Response, error) {
	return m.client.AddService(helper.Ctx, req)
}
func (m *monitorClient) DeleteService(req *msa_monitor.RequestDelService) (*msa_monitor.Response, error) {
	return m.client.DeleteService(helper.Ctx, req)
}
func (m *monitorClient) RebuildService(req *msa_monitor.RequestRebuildService) (*msa_monitor.Response, error) {
	return m.client.RebuildService(helper.Ctx, req)
}
func (m *monitorClient) SetService(req *msa_monitor.RequestSetService) (*msa_monitor.Response, error) {
	return m.client.SetService(helper.Ctx, req)
}
