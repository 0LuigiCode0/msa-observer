package roots_handler

import (
	"fmt"
	"net/http"
	core_helper "x-msa-core/helper"

	goutill "github.com/0LuigiCode0/go-utill"
	"github.com/0LuigiCode0/logger"
	"github.com/gorilla/mux"
)

func (h *handler) Auth(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if _, ok := vars["login"]; !ok {
		logger.Log.Warningf(core_helper.KeyErrorInvalidParams+": %v", "login")
		h.respError(w, core_helper.ErrorInvalidParams, fmt.Sprintf(core_helper.KeyErrorInvalidParams+": %v", "login"))
		return
	}
	login := vars["login"]
	if _, ok := vars["pwd"]; !ok {
		logger.Log.Warningf(core_helper.KeyErrorInvalidParams+": %v", "pwd")
		h.respError(w, core_helper.ErrorInvalidParams, fmt.Sprintf(core_helper.KeyErrorInvalidParams+": %v", "pwd"))
		return
	}
	pwd := vars["pwd"]
	if err := goutill.Validator(true, map[string]interface{}{"login": &login, "pwd": &pwd}); err != nil {
		logger.Log.Warningf(core_helper.KeyErrorInvalidParams+": %v", err)
		h.respError(w, core_helper.ErrorInvalidParams, fmt.Sprintf(core_helper.KeyErrorInvalidParams+": %v", err))
		return
	}

	user, err := h.MongoStore().UserStore().SelectByLogin(login)
	if err != nil {
		logger.Log.Errorf(core_helper.KeyErrorNotFound+" user : %v", err)
		h.respError(w, core_helper.ErrorNotFound, core_helper.KeyErrorNotFound+" user")
		return
	}
	if user.Password != h.Helper().Hash(pwd) {
		logger.Log.Warningf(core_helper.KeyErorrAccessDenied+": %v", "invalid password")
		h.respError(w, core_helper.ErorrAccessDeniedParams, core_helper.KeyErorrAccessDenied)
		return
	}

	token, err := h.Helper().GenerateJWT(user.ID)
	if err != nil {
		logger.Log.Warningf(core_helper.KeyErrorGenerate+" jwt: %v", err)
		h.respError(w, core_helper.ErrorGenerate, core_helper.KeyErrorGenerate+" jwt")
		return
	}

	h.respOk(w, token)
}
