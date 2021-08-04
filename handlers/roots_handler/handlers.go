package roots_handler

import (
	"encoding/json"
	"net/http"
	core_helper "x-msa-core/helper"
	"x-msa-observer/handlers/roots_handler/roots_helper"
	"x-msa-observer/helper"
	"x-msa-observer/hub/hub_helper"

	"github.com/0LuigiCode0/logger"
)

type handler struct {
	hub_helper.HelperForHandler
}

func InitHandler(hub hub_helper.HelperForHandler, conf *helper.HandlerConfig) (H roots_helper.Handler, err error) {
	h := &handler{HelperForHandler: hub}
	H = h

	rCore := h.Router().PathPrefix("/core").Subrouter()
	rCore.HandleFunc("/auth", h.Auth).Queries("login", "{login}", "pwd", "{pwd}").Methods(http.MethodGet)

	rService := h.Router().PathPrefix("/service").Subrouter()
	rService.HandleFunc("/add", h.AddServices).Methods(http.MethodPost)
	rService.HandleFunc("/delete", h.DeleteServices).Queries("key", "{key}").Methods(http.MethodDelete)
	rService.HandleFunc("/set/status", h.SetStatusServices).Methods(http.MethodPost)
	rService.HandleFunc("/rebuild", h.RebuildServices).Methods(http.MethodPost)
	rService.HandleFunc("/get/all", h.GetAllServices).Methods(http.MethodGet)

	rGroup := h.Router().PathPrefix("/group").Subrouter()
	rGroup.HandleFunc("/add", h.AddGroup).Methods(http.MethodPost)
	rGroup.HandleFunc("/delete", h.DeleteGroup).Queries("group_id", "{group_id}").Methods(http.MethodDelete)
	rGroup.HandleFunc("/set", h.SetGroup).Methods(http.MethodPost)
	rGroup.HandleFunc("/get/all", h.GetAllGroups).Methods(http.MethodGet)

	rPeer := h.Router().PathPrefix("/peer").Subrouter()
	rPeer.HandleFunc("/add", h.AddPeer).Methods(http.MethodPost)
	rPeer.HandleFunc("/delete", h.DeletePeer).Queries("peer_id", "{peer_id}").Methods(http.MethodDelete)
	rPeer.HandleFunc("/set", h.SetPeer).Methods(http.MethodPost)
	rPeer.HandleFunc("/get/all", h.GetAllPeers).Methods(http.MethodGet)

	h.Router().Use(h.middleware)
	h.SetHandler(applyCORS(h.Router()))
	return
}

func (h *handler) respOk(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	resp := &core_helper.ResponseModel{
		Success: true,
		Result:  data,
	}
	buf, err := json.Marshal(resp)
	if err != nil {
		logger.Log.Warningf(core_helper.KeyErrorParse+": josn: %v", err)
		h.respError(w, core_helper.ErrorParse, core_helper.KeyErrorParse+": josn")
		return
	}
	_, err = w.Write(buf)
	if err != nil {
		logger.Log.Warningf(core_helper.KeyErrorWrite+": response: %v", err)
		h.respError(w, core_helper.ErrorWrite, core_helper.KeyErrorWrite+": response")
		return
	}
}

func (h *handler) respError(w http.ResponseWriter, code core_helper.ErrCode, msg string) {
	w.Header().Set("Content-Type", "application/json")
	resp := &core_helper.ResponseModel{
		Success: false,
		Result: &core_helper.ResponseError{
			Code: code,
			Msg:  msg,
		},
	}
	buf, _ := json.Marshal(resp)
	_, err := w.Write(buf)
	if err != nil {
		logger.Log.Warningf(core_helper.KeyErrorWrite+": response: %v", err)
	}
}
