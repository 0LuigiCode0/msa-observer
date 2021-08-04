package store

import (
	core_helper "x-msa-core/helper"
	"x-msa-observer/helper"
	"x-msa-observer/store/mongo/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//groupStore хранилище
type PeerStore interface {
	Save(peer *model.PeerModel) error
	Update(peer *model.PeerModel) error
	SelectByID(id primitive.ObjectID) (*model.PeerModel, error)
	SelectByAddr(host string, port int32) (*model.PeerModel, error)
	SelectAll() ([]*model.PeerModel, error)
	DeleteByID(id primitive.ObjectID) error
}

//groupStore хранилище
type peerStore struct {
	db *mongo.Database
}

func InitPeerStore(db *mongo.Database) PeerStore {
	return &peerStore{db: db}
}

func (s *peerStore) Save(peer *model.PeerModel) error {
	res, err := s.db.Collection(string(helper.CollPeers)).UpdateOne(core_helper.Ctx, primitive.M{
		"host": peer.Host,
		"port": peer.Port,
	}, primitive.M{
		"$setOnInsert": peer,
	}, options.Update().SetUpsert(true))
	if res != nil {
		if id, ok := res.UpsertedID.(primitive.ObjectID); ok {
			peer.ID = id
		}
	}
	return err
}

func (s *peerStore) Update(peer *model.PeerModel) error {
	_, err := s.db.Collection(string(helper.CollPeers)).UpdateOne(core_helper.Ctx, primitive.M{
		"_id": peer.ID,
	}, primitive.M{
		"$set": peer,
	}, options.Update().SetUpsert(false))
	return err
}

func (s *peerStore) SelectByID(id primitive.ObjectID) (*model.PeerModel, error) {
	peer := &model.PeerModel{}
	err := s.db.Collection(string(helper.CollPeers)).FindOne(core_helper.Ctx, primitive.M{
		"_id": id,
	}).Decode(peer)
	return peer, err
}

func (s *peerStore) SelectByAddr(host string, port int32) (*model.PeerModel, error) {
	peer := &model.PeerModel{}
	err := s.db.Collection(string(helper.CollPeers)).FindOne(core_helper.Ctx, primitive.M{
		"host": host,
		"port": port,
	}).Decode(peer)
	return peer, err
}

func (s *peerStore) SelectAll() ([]*model.PeerModel, error) {
	peer := []*model.PeerModel{}
	res, err := s.db.Collection(string(helper.CollPeers)).Find(core_helper.Ctx, primitive.M{})
	if err != nil {
		return peer, err
	}
	err = res.All(core_helper.Ctx, &peer)
	return peer, err
}

func (s *peerStore) DeleteByID(id primitive.ObjectID) error {
	_, err := s.db.Collection(string(helper.CollPeers)).DeleteMany(core_helper.Ctx, primitive.M{
		"_id": id,
	})
	return err
}
