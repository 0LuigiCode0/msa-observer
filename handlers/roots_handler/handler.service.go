package roots_handler

import (
	"fmt"
	"net/http"

	"x-msa-core/grpc/msa_utils"
	core_helper "x-msa-core/helper"
	"x-msa-observer/store/mongo/model"

	goutill "github.com/0LuigiCode0/go-utill"
	"github.com/0LuigiCode0/logger"
	"github.com/gorilla/mux"
)

func (h *handler) AddServices(w http.ResponseWriter, r *http.Request) {
	_, err := h.Helper().AuthGuard(r)
	if err != nil {
		logger.Log.Warningf(core_helper.KeyErorrAccessDenied+": %v", err)
		h.respError(w, core_helper.ErorrAccessDeniedToken, core_helper.KeyErorrAccessDenied)
		return
	}

	req := &model.RequestServiceModel{}
	if err = goutill.JsonParse(r.Body, req); err != nil {
		logger.Log.Warningf(core_helper.KeyErrorParse+" json: %v", err)
		h.respError(w, core_helper.ErrorParse, core_helper.KeyErrorParse+" json")
		return
	}
	if err = goutill.Validator(true, map[string]interface{}{
		"title":    &req.Title,
		"peer_id":  req.PeerID,
		"group_id": req.GroupID,
	}); err != nil {
		logger.Log.Warningf(core_helper.KeyErrorInvalidParams+": %v", err)
		h.respError(w, core_helper.ErrorInvalidParams, fmt.Sprintf(core_helper.KeyErrorInvalidParams+": %v", err))
		return
	}
	if req.Role < msa_utils.RoleService_Observer || req.Role > msa_utils.RoleService_Interface {
		logger.Log.Warning(core_helper.KeyErrorInvalidParams + " role")
		h.respError(w, core_helper.ErrorInvalidParams, core_helper.KeyErrorInvalidParams+" role")
		return
	}

	_, err = h.MongoStore().ServiceStore().SelectByPeerAndGroup(req.GroupID, req.PeerID)
	if err == nil {
		logger.Log.Errorf(core_helper.KeyErrorExist+" service: %v", err)
		h.respError(w, core_helper.ErrorExist, core_helper.KeyErrorExist+" service")
		return
	}

	_, err = h.MongoStore().PeerStore().SelectByID(req.PeerID)
	if err != nil {
		logger.Log.Errorf(core_helper.KeyErrorNotFound+" peer: %v", err)
		h.respError(w, core_helper.ErrorNotFound, core_helper.KeyErrorNotFound+" peer")
		return
	}

	group, err := h.MongoStore().GroupStore().SelectByID(req.GroupID)
	if err != nil {
		logger.Log.Errorf(core_helper.KeyErrorNotFound+" group: %v", err)
		h.respError(w, core_helper.ErrorNotFound, core_helper.KeyErrorNotFound+" group")
		return
	}

	key := h.Helper().KeyGenerate()
	service := &model.ServiceModel{
		Title:   req.Title,
		Key:     key,
		Status:  msa_utils.StatusService_Created,
		PeerID:  req.PeerID,
		GroupID: req.GroupID,
		Role:    req.Role,
		Version: group.Version,
	}

	if err = h.MongoStore().ServiceStore().Save(service); err != nil {
		logger.Log.Errorf(core_helper.KeyErrorSave+" product: %v", err)
		h.respError(w, core_helper.ErrorSave, core_helper.KeyErrorSave+" product")
		return
	}

	h.respOk(w, "ok")
}

func (h *handler) DeleteServices(w http.ResponseWriter, r *http.Request) {
	_, err := h.Helper().AuthGuard(r)
	if err != nil {
		logger.Log.Warningf(core_helper.KeyErorrAccessDenied+": %v", err)
		h.respError(w, core_helper.ErorrAccessDeniedToken, core_helper.KeyErorrAccessDenied)
		return
	}

	vars := mux.Vars(r)
	if _, ok := vars["key"]; !ok {
		logger.Log.Warning(core_helper.KeyErrorInvalidParams + " key")
		h.respError(w, core_helper.ErrorInvalidParams, core_helper.KeyErrorInvalidParams+" key")
		return
	}
	key := vars["type"]
	if err = goutill.Validator(true, map[string]interface{}{
		"key": &key,
	}); err != nil {
		logger.Log.Warningf(core_helper.KeyErrorInvalidParams+": %v", err)
		h.respError(w, core_helper.ErrorInvalidParams, fmt.Sprintf(core_helper.KeyErrorInvalidParams+": %v", err))
		return
	}

	service, err := h.MongoStore().ServiceStore().SelectByKey(key)
	if err != nil {
		logger.Log.Errorf(core_helper.KeyErrorNotFound+" service: %v", err)
		h.respError(w, core_helper.ErrorNotFound, core_helper.KeyErrorNotFound+" service")
		return
	}

	if err = h.MongoStore().ServiceStore().DeleteByKey(key); err != nil {
		logger.Log.Errorf(core_helper.KeyErrorDelete+" service: %v", err)
		h.respError(w, core_helper.ErrorDelete, core_helper.KeyErrorDelete+" service")
		return
	}

	if service.Role == msa_utils.RoleService_Monitor {
		h.Grps().DeleteMonitor(service.Key)
	}

	h.respOk(w, "ok")
}

func (h *handler) SetStatusServices(w http.ResponseWriter, r *http.Request) {
	_, err := h.Helper().AuthGuard(r)
	if err != nil {
		logger.Log.Warningf(core_helper.KeyErorrAccessDenied+": %v", err)
		h.respError(w, core_helper.ErorrAccessDeniedToken, core_helper.KeyErorrAccessDenied)
		return
	}

}
func (h *handler) RebuildServices(w http.ResponseWriter, r *http.Request) {
	_, err := h.Helper().AuthGuard(r)
	if err != nil {
		logger.Log.Warningf(core_helper.KeyErorrAccessDenied+": %v", err)
		h.respError(w, core_helper.ErorrAccessDeniedToken, core_helper.KeyErorrAccessDenied)
		return
	}

}
func (h *handler) GetAllServices(w http.ResponseWriter, r *http.Request) {
	_, err := h.Helper().AuthGuard(r)
	if err != nil {
		logger.Log.Warningf(core_helper.KeyErorrAccessDenied+": %v", err)
		h.respError(w, core_helper.ErorrAccessDeniedToken, core_helper.KeyErorrAccessDenied)
		return
	}
	req := &model.FilterServiceModel{}
	if err = goutill.JsonParse(r.Body, req); err != nil {
		logger.Log.Warningf(core_helper.KeyErrorParse+" json: %v", err)
		h.respError(w, core_helper.ErrorParse, core_helper.KeyErrorParse+" json")
		return
	}

	resp := &model.ResponseServiceListModel{}
	resp.Services, resp.Count, err = h.MongoStore().ServiceStore().SelectFilter(req)
	if err != nil {
		logger.Log.Errorf(core_helper.KeyErrorNotFound+" services: %v", err)
		h.respError(w, core_helper.ErrorNotFound, core_helper.KeyErrorNotFound+" services")
		return
	}

	h.respOk(w, resp)
}
