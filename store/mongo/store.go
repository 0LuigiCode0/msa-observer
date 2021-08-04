package mongoStore

import (
	"x-msa-observer/store/mongo/store"

	"go.mongodb.org/mongo-driver/mongo"
)

type Store interface {
	UserStore() store.UserStore
	ServiceStore() store.ServiceStore
	PeerStore() store.PeerStore
	GroupStore() store.GroupStore
}

type s struct {
	user    store.UserStore
	service store.ServiceStore
	peer    store.PeerStore
	group   store.GroupStore
}

func InitStore(db *mongo.Database) Store {
	return &s{
		user:    store.InitUserStore(db),
		service: store.InitServiceStore(db),
		peer:    store.InitPeerStore(db),
		group:   store.InitGroupStore(db),
	}
}

func (s *s) UserStore() store.UserStore       { return s.user }
func (s *s) ServiceStore() store.ServiceStore { return s.service }
func (s *s) PeerStore() store.PeerStore       { return s.peer }
func (s *s) GroupStore() store.GroupStore     { return s.group }
