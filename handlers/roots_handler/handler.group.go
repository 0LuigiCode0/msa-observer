package roots_handler

import (
	"fmt"
	"net/http"

	core_helper "x-msa-core/helper"
	"x-msa-observer/store/mongo/model"

	goutill "github.com/0LuigiCode0/go-utill"
	"github.com/0LuigiCode0/logger"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (h *handler) AddGroup(w http.ResponseWriter, r *http.Request) {
	_, err := h.Helper().AuthGuard(r)
	if err != nil {
		logger.Log.Warningf(core_helper.KeyErorrAccessDenied+": %v", err)
		h.respError(w, core_helper.ErorrAccessDeniedToken, core_helper.KeyErorrAccessDenied)
		return
	}

	req := &model.RequestGroupModel{}
	if err = goutill.JsonParse(r.Body, req); err != nil {
		logger.Log.Warningf(core_helper.KeyErrorParse+" json: %v", err)
		h.respError(w, core_helper.ErrorParse, core_helper.KeyErrorParse+" json")
		return
	}
	if err = goutill.Validator(true, map[string]interface{}{
		"title":      &req.Title,
		"rep_link":   &req.RepLink,
		"dependents": req.Dependents,
	}); err != nil {
		logger.Log.Warningf(core_helper.KeyErrorInvalidParams+": %v", err)
		h.respError(w, core_helper.ErrorInvalidParams, fmt.Sprintf(core_helper.KeyErrorInvalidParams+": %v", err))
		return
	}

	_, err = h.MongoStore().GroupStore().SelectByTitle(req.Title)
	if err == nil {
		logger.Log.Errorf(core_helper.KeyErrorExist + " group")
		h.respError(w, core_helper.ErrorExist, core_helper.KeyErrorExist+" group")
		return
	}

	group := &model.GroupModel{
		Title:      req.Title,
		RepLink:    req.RepLink,
		Dependents: req.Dependents,
	}

	if err = h.MongoStore().GroupStore().Save(group); err != nil {
		logger.Log.Errorf(core_helper.KeyErrorSave+" group: %v", err)
		h.respError(w, core_helper.ErrorSave, core_helper.KeyErrorSave+" group")
		return
	}

	h.respOk(w, "ok")
}

func (h *handler) DeleteGroup(w http.ResponseWriter, r *http.Request) {
	_, err := h.Helper().AuthGuard(r)
	if err != nil {
		logger.Log.Warningf(core_helper.KeyErorrAccessDenied+": %v", err)
		h.respError(w, core_helper.ErorrAccessDeniedToken, core_helper.KeyErorrAccessDenied)
		return
	}

	vars := mux.Vars(r)
	if _, ok := vars["group_id"]; !ok {
		logger.Log.Warning(core_helper.KeyErrorInvalidParams + " group_id")
		h.respError(w, core_helper.ErrorInvalidParams, core_helper.KeyErrorInvalidParams+" group_id")
		return
	}
	groupID, err := primitive.ObjectIDFromHex(vars["group_id"])
	if err != nil {
		logger.Log.Warningf(core_helper.KeyErrorParse+" group_id : %v", err)
		h.respError(w, core_helper.ErrorParse, core_helper.KeyErrorParse+" group_id")
		return
	}

	_, err = h.MongoStore().GroupStore().SelectByID(groupID)
	if err != nil {
		logger.Log.Errorf(core_helper.KeyErrorNotFound+" group: %v", err)
		h.respError(w, core_helper.ErrorNotFound, core_helper.KeyErrorNotFound+" group")
		return
	}

	_, err = h.MongoStore().ServiceStore().SelectByGroup(groupID)
	if err == nil {
		logger.Log.Errorf(core_helper.KeyErrorExist+" services: %v", err)
		h.respError(w, core_helper.ErrorExist, core_helper.KeyErrorExist+" services")
		return
	}

	if err = h.MongoStore().GroupStore().DeleteDependents(nil, groupID); err != nil {
		logger.Log.Errorf(core_helper.KeyErrorDelete+" group: %v", err)
		h.respError(w, core_helper.ErrorDelete, core_helper.KeyErrorDelete+" group")
		return
	}

	if err = h.MongoStore().GroupStore().DeleteByID(groupID); err != nil {
		logger.Log.Errorf(core_helper.KeyErrorDelete+" group: %v", err)
		h.respError(w, core_helper.ErrorDelete, core_helper.KeyErrorDelete+" group")
		return
	}

	h.respOk(w, "ok")
}

func (h *handler) SetGroup(w http.ResponseWriter, r *http.Request) {
	_, err := h.Helper().AuthGuard(r)
	if err != nil {
		logger.Log.Warningf(core_helper.KeyErorrAccessDenied+": %v", err)
		h.respError(w, core_helper.ErorrAccessDeniedToken, core_helper.KeyErorrAccessDenied)
		return
	}

	req := &model.RequestGroupModel{}
	if err = goutill.JsonParse(r.Body, req); err != nil {
		logger.Log.Warningf(core_helper.KeyErrorParse+" json: %v", err)
		h.respError(w, core_helper.ErrorParse, core_helper.KeyErrorParse+" json")
		return
	}
	if err = goutill.Validator(true, map[string]interface{}{
		"title":      &req.Title,
		"rep_link":   &req.RepLink,
		"dependents": req.Dependents,
	}); err != nil {
		logger.Log.Warningf(core_helper.KeyErrorInvalidParams+": %v", err)
		h.respError(w, core_helper.ErrorInvalidParams, fmt.Sprintf(core_helper.KeyErrorInvalidParams+": %v", err))
		return
	}

	group, err := h.MongoStore().GroupStore().SelectByID(req.ID)
	if err != nil {
		logger.Log.Errorf(core_helper.KeyErrorNotFound+" group: %v", err)
		h.respError(w, core_helper.ErrorNotFound, core_helper.KeyErrorNotFound+" group")
		return
	}

	group.Title = req.Title
	group.RepLink = req.RepLink
	group.Dependents = req.Dependents

	if err = h.MongoStore().GroupStore().Update(group); err != nil {
		logger.Log.Errorf(core_helper.KeyErrorUpdate+" group: %v", err)
		h.respError(w, core_helper.ErrorUpdate, core_helper.KeyErrorUpdate+" group")
		return
	}

	h.respOk(w, "ok")
}

func (h *handler) GetAllGroups(w http.ResponseWriter, r *http.Request) {
	_, err := h.Helper().AuthGuard(r)
	if err != nil {
		logger.Log.Warningf(core_helper.KeyErorrAccessDenied+": %v", err)
		h.respError(w, core_helper.ErorrAccessDeniedToken, core_helper.KeyErorrAccessDenied)
		return
	}

	resp, err := h.MongoStore().GroupStore().SelectAll()
	if err != nil {
		logger.Log.Errorf(core_helper.KeyErrorNotFound+" groups: %v", err)
		h.respError(w, core_helper.ErrorNotFound, core_helper.KeyErrorNotFound+" groups")
		return
	}

	h.respOk(w, resp)
}
