package model

import (
	"x-msa-core/grpc/msa_utils"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//UserModel модель пользователя в БД
type ServiceModel struct {
	ID      primitive.ObjectID      `bson:"_id,omitempty"`
	Status  msa_utils.StatusService `bson:"status"`
	Role    msa_utils.RoleService   `bson:"role"`
	PeerID  primitive.ObjectID      `bson:"peer_id"`
	GroupID primitive.ObjectID      `bson:"group_id"`
	Version string                  `bson:"version"`
	Title   string                  `bson:"title"`
	Key     string                  `bson:"key"`
	Host    string                  `bson:"host"`
	Port    int32                   `bson:"port"`
}

type FilterServiceModel struct {
	Limit    int64                     `json:"limit"`
	Offset   int64                     `json:"offset"`
	Groups   []primitive.ObjectID      `json:"groups"`
	Roles    []msa_utils.RoleService   `json:"roles"`
	Peers    []primitive.ObjectID      `json:"peers"`
	Statuses []msa_utils.StatusService `json:"statuses"`
	Title    string                    `json:"title"`
	Host     string                    `json:"host"`
	Port     int32                     `json:"port"`
}

//UserModel модель пользователя в БД
type RequestServiceModel struct {
	ID      primitive.ObjectID    `json:"id"`
	Role    msa_utils.RoleService `json:"role"`
	PeerID  primitive.ObjectID    `json:"peer_id"`
	GroupID primitive.ObjectID    `json:"group_id"`
	Title   string                `json:"title"`
	Key     string                `json:"key"`
	Host    string                `json:"host"`
	Port    int32                 `json:"port"`
}

type ResponseServiceListModel struct {
	Services []*ServiceModel `json:"services"`
	Count    int64           `json:"count"`
}
