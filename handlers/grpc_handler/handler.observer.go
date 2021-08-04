package grpc_handler

import (
	"context"
	"fmt"
	"x-msa-core/grpc/msa_observer"
	"x-msa-core/grpc/msa_utils"
	core_helper "x-msa-core/helper"

	"github.com/0LuigiCode0/logger"
)

func (h *handler) PushFirst(ctx context.Context, req *msa_observer.RequestPushFirst) (*msa_observer.ResponsePushFirst, error) {
	service, err := h.MongoStore().ServiceStore().SelectByKey(req.Key)
	if err != nil {
		logger.Log.Errorf(core_helper.KeyErrorNotFound+" service: %v", err)
		return nil, fmt.Errorf(core_helper.KeyErrorNotFound+" service: %v", err)
	}

	if h.MongoStore().ServiceStore().UpdateStatus(req.Key, msa_utils.StatusService_On); err != nil {
		logger.Log.Errorf(core_helper.KeyErrorUpdate+" service: %v", err)
		return nil, fmt.Errorf(core_helper.KeyErrorUpdate+" service: %v", err)
	}

	core_helper.Wg.Add(1)
	h.AddMonitor(req.Key, fmt.Sprintf("%v:%v", req.Host, req.Port))

	dependents, err := h.MongoStore().GroupStore().SelectForDependens(service.GroupID)
	if err != nil {
		logger.Log.Errorf(core_helper.KeyErrorNotFound+" dependents: %v", err)
		return nil, fmt.Errorf(core_helper.KeyErrorNotFound+" dependents: %v", err)
	}

	return &msa_observer.ResponsePushFirst{
		Dependents: dependents,
	}, nil
}
