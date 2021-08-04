package hub_helper

import (
	"net/http"
	"x-msa-observer/core/database"
	"x-msa-observer/handlers/grpc_handler/grpc_helper"
	"x-msa-observer/handlers/roots_handler/roots_helper"
	"x-msa-observer/helper"
	"x-msa-observer/store/mongo/model"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Helper interface {
	AuthGuard(r *http.Request) (*model.UserModel, error)
	Hash(data string) string
	GenerateJWT(id primitive.ObjectID) (string, error)
	KeyGenerate() string
}

type HelperForHandler interface {
	database.DBForHandler
	Helper() Helper
	Config() *helper.Config
	Router() *mux.Router
	SetHandler(hh http.Handler)
	Grps() grpc_helper.MSA
}

type HandlerForHelper interface {
	database.DBForHandler
	Roots() roots_helper.Handler
	Grps() grpc_helper.MSA
	Config() *helper.Config
}

type help struct {
	HandlerForHelper
}
