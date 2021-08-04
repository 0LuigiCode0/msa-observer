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

func (h *handler) AddPeer(w http.ResponseWriter, r *http.Request) {
	_, err := h.Helper().AuthGuard(r)
	if err != nil {
		logger.Log.Warningf(core_helper.KeyErorrAccessDenied+": %v", err)
		h.respError(w, core_helper.ErorrAccessDeniedToken, core_helper.KeyErorrAccessDenied)
		return
	}

	req := &model.RequestPeerModel{}
	if err = goutill.JsonParse(r.Body, req); err != nil {
		logger.Log.Warningf(core_helper.KeyErrorParse+" json: %v", err)
		h.respError(w, core_helper.ErrorParse, core_helper.KeyErrorParse+" json")
		return
	}
	if err = goutill.Validator(true, map[string]interface{}{
		"title": &req.Title,
		"host":  &req.Host,
		"port":  req.Port,
		"login": &req.Login,
		"pwd":   &req.Password,
	}); err != nil {
		logger.Log.Warningf(core_helper.KeyErrorInvalidParams+": %v", err)
		h.respError(w, core_helper.ErrorInvalidParams, fmt.Sprintf(core_helper.KeyErrorInvalidParams+": %v", err))
		return
	}

	_, err = h.MongoStore().PeerStore().SelectByAddr(req.Host, req.Port)
	if err == nil {
		logger.Log.Errorf(core_helper.KeyErrorExist + " peer")
		h.respError(w, core_helper.ErrorExist, core_helper.KeyErrorExist+" peer")
		return
	}

	peer := &model.PeerModel{
		Title:    req.Title,
		Host:     req.Host,
		Port:     req.Port,
		Login:    req.Login,
		Password: req.Password,
	}

	if err = h.MongoStore().PeerStore().Save(peer); err != nil {
		logger.Log.Errorf(core_helper.KeyErrorSave+" peer: %v", err)
		h.respError(w, core_helper.ErrorSave, core_helper.KeyErrorSave+" peer")
		return
	}

	h.respOk(w, "ok")
}

func (h *handler) DeletePeer(w http.ResponseWriter, r *http.Request) {
	_, err := h.Helper().AuthGuard(r)
	if err != nil {
		logger.Log.Warningf(core_helper.KeyErorrAccessDenied+": %v", err)
		h.respError(w, core_helper.ErorrAccessDeniedToken, core_helper.KeyErorrAccessDenied)
		return
	}

	vars := mux.Vars(r)
	if _, ok := vars["peer_id"]; !ok {
		logger.Log.Warning(core_helper.KeyErrorInvalidParams + " peer_id")
		h.respError(w, core_helper.ErrorInvalidParams, core_helper.KeyErrorInvalidParams+" peer_id")
		return
	}
	peerID, err := primitive.ObjectIDFromHex(vars["peer_id"])
	if err != nil {
		logger.Log.Warningf(core_helper.KeyErrorParse+" peer_id : %v", err)
		h.respError(w, core_helper.ErrorParse, core_helper.KeyErrorParse+" peer_id")
		return
	}

	_, err = h.MongoStore().PeerStore().SelectByID(peerID)
	if err != nil {
		logger.Log.Errorf(core_helper.KeyErrorNotFound+" peer: %v", err)
		h.respError(w, core_helper.ErrorNotFound, core_helper.KeyErrorNotFound+" peer")
		return
	}

	_, err = h.MongoStore().ServiceStore().SelectByPeer(peerID)
	if err == nil {
		logger.Log.Errorf(core_helper.KeyErrorExist+" services: %v", err)
		h.respError(w, core_helper.ErrorExist, core_helper.KeyErrorExist+" services")
		return
	}

	if err = h.MongoStore().PeerStore().DeleteByID(peerID); err != nil {
		logger.Log.Errorf(core_helper.KeyErrorDelete+" peer: %v", err)
		h.respError(w, core_helper.ErrorDelete, core_helper.KeyErrorDelete+" peer")
		return
	}

	h.respOk(w, "ok")
}

func (h *handler) SetPeer(w http.ResponseWriter, r *http.Request) {
	_, err := h.Helper().AuthGuard(r)
	if err != nil {
		logger.Log.Warningf(core_helper.KeyErorrAccessDenied+": %v", err)
		h.respError(w, core_helper.ErorrAccessDeniedToken, core_helper.KeyErorrAccessDenied)
		return
	}

	req := &model.RequestPeerModel{}
	if err = goutill.JsonParse(r.Body, req); err != nil {
		logger.Log.Warningf(core_helper.KeyErrorParse+" json: %v", err)
		h.respError(w, core_helper.ErrorParse, core_helper.KeyErrorParse+" json")
		return
	}
	if err = goutill.Validator(true, map[string]interface{}{
		"title": &req.Title,
		"host":  &req.Host,
		"port":  req.Port,
		"login": &req.Login,
	}); err != nil {
		logger.Log.Warningf(core_helper.KeyErrorInvalidParams+": %v", err)
		h.respError(w, core_helper.ErrorInvalidParams, fmt.Sprintf(core_helper.KeyErrorInvalidParams+": %v", err))
		return
	}

	peer, err := h.MongoStore().PeerStore().SelectByAddr(req.Host, req.Port)
	if err == nil {
		logger.Log.Errorf(core_helper.KeyErrorExist + " peer")
		h.respError(w, core_helper.ErrorExist, core_helper.KeyErrorExist+" peer")
		return
	}

	peer.Title = req.Title
	peer.Host = req.Host
	peer.Port = req.Port
	peer.Login = req.Login

	if err = h.MongoStore().PeerStore().Update(peer); err != nil {
		logger.Log.Errorf(core_helper.KeyErrorUpdate+" peer: %v", err)
		h.respError(w, core_helper.ErrorUpdate, core_helper.KeyErrorUpdate+" peer")
		return
	}

	h.respOk(w, "ok")
}

func (h *handler) GetAllPeers(w http.ResponseWriter, r *http.Request) {
	_, err := h.Helper().AuthGuard(r)
	if err != nil {
		logger.Log.Warningf(core_helper.KeyErorrAccessDenied+": %v", err)
		h.respError(w, core_helper.ErorrAccessDeniedToken, core_helper.KeyErorrAccessDenied)
		return
	}

	resp, err := h.MongoStore().PeerStore().SelectAll()
	if err != nil {
		logger.Log.Errorf(core_helper.KeyErrorNotFound+" peers: %v", err)
		h.respError(w, core_helper.ErrorNotFound, core_helper.KeyErrorNotFound+" peers")
		return
	}

	h.respOk(w, resp)
}
