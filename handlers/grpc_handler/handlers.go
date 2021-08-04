package grpc_handler

import (
	"fmt"
	"x-msa-core/grpc/msa_monitor"
	"x-msa-core/grpc/msa_utils"
	core_helper "x-msa-core/helper"
	"x-msa-core/observer/server"
	"x-msa-observer/handlers/grpc_handler/grpc_helper"
	"x-msa-observer/helper"
	"x-msa-observer/hub/hub_helper"
	"x-msa-observer/store/mongo/model"

	"github.com/0LuigiCode0/logger"
)

type handler struct {
	hub_helper.HelperForHandler
	server.ObserverServer
}

func InitHandler(hub hub_helper.HelperForHandler, conf *helper.HandlerConfig) (H grpc_helper.Handler, err error) {
	h := &handler{
		HelperForHandler: hub,
		ObserverServer:   server.NewObserverServer(conf.Key, fmt.Sprintf("%v:%v", conf.Host, conf.Port)),
	}
	H = h

	if err = h.initDependents(); err != nil {
		logger.Log.Errorf("init dependents error: %v", err)
		return
	}

	h.SetPushFirst(h.PushFirst)

	logger.Log.Servicef("gserver started at address: %v", fmt.Sprintf("%v:%v", conf.Host, conf.Port))
	return
}

func (h *handler) initDependents() error {
	services, _, err := h.MongoStore().ServiceStore().SelectFilter(&model.FilterServiceModel{
		Statuses: []msa_utils.StatusService{
			msa_utils.StatusService_On,
		},
		Roles: []msa_utils.RoleService{
			msa_utils.RoleService_Monitor,
		},
	})
	if err != nil {
		logger.Log.Errorf(core_helper.KeyErrorNotFound+" services: %v", err)
		return err
	}

	for _, v := range services {
		core_helper.Wg.Add(1)
		go h.AddMonitor(v.Key, fmt.Sprintf("%v:%v", v.Host, v.Port))
	}
	return nil
}

func (h *handler) AddMonitor(key, addr string) {
	defer core_helper.Wg.Done()

	if err := h.Monitors().Add(key, addr); err != nil {
		logger.Log.Warningf("canot added monitors key %v: %v", key, err)
		return
	}
	monitor, err := h.Monitors().Get(key)
	if err != nil {
		logger.Log.Warningf("canot find monitors key %v: %v", key, err)
		return
	}
	monitor.AddService(&msa_monitor.RequestAddService{})
}

func (h *handler) DeleteMonitor(key string) error {
	if err := h.Monitors().Delete(key); err != nil {
		return err
	}
	return nil
}
