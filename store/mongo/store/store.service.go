package store

import (
	"fmt"
	"x-msa-core/grpc/msa_utils"
	core_helper "x-msa-core/helper"
	"x-msa-observer/helper"
	"x-msa-observer/store/mongo/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Store хранилище
type ServiceStore interface {
	Save(srv *model.ServiceModel) error
	UpdateStatus(key string, status msa_utils.StatusService) error
	SelectByID(id primitive.ObjectID) (*model.ServiceModel, error)
	SelectByKey(key string) (*model.ServiceModel, error)
	SelectByGroup(groupID primitive.ObjectID) ([]*model.ServiceModel, error)
	SelectByPeer(peerID primitive.ObjectID) ([]*model.ServiceModel, error)
	SelectByPeerAndGroup(groupID, peerID primitive.ObjectID) (*model.ServiceModel, error)
	SelectFilter(filter *model.FilterServiceModel) ([]*model.ServiceModel, int64, error)
	DeleteByKey(key string) error
}

//Store хранилище
type serviceStore struct {
	db *mongo.Database
}

func InitServiceStore(db *mongo.Database) ServiceStore {
	return &serviceStore{db: db}
}

func (s *serviceStore) Save(srv *model.ServiceModel) error {
	res, err := s.db.Collection(string(helper.CollServices)).UpdateOne(core_helper.Ctx, primitive.M{
		"peer_id":  srv.PeerID,
		"group_id": srv.GroupID,
	}, primitive.M{
		"$setOnInsert": srv,
	}, options.Update().SetUpsert(true))
	if res != nil {
		if id, ok := res.UpsertedID.(primitive.ObjectID); ok {
			srv.ID = id
		}
	}
	return err
}

func (s *serviceStore) UpdateStatus(key string, status msa_utils.StatusService) error {
	_, err := s.db.Collection(string(helper.CollServices)).UpdateOne(core_helper.Ctx, primitive.M{
		"key": key,
	}, primitive.M{
		"$set": primitive.M{"status": status},
	}, options.Update().SetUpsert(false))
	return err
}

func (s *serviceStore) SelectByID(id primitive.ObjectID) (*model.ServiceModel, error) {
	srv := &model.ServiceModel{}
	err := s.db.Collection(string(helper.CollServices)).FindOne(core_helper.Ctx, primitive.M{
		"_id": id,
	}).Decode(srv)
	return srv, err
}

func (s *serviceStore) SelectByKey(key string) (*model.ServiceModel, error) {
	srv := &model.ServiceModel{}
	err := s.db.Collection(string(helper.CollServices)).FindOne(core_helper.Ctx, primitive.M{
		"key": key,
	}).Decode(srv)
	return srv, err
}

func (s *serviceStore) SelectByGroup(groupID primitive.ObjectID) ([]*model.ServiceModel, error) {
	srv := []*model.ServiceModel{}
	res, err := s.db.Collection(string(helper.CollServices)).Find(core_helper.Ctx, primitive.M{
		"group_id": groupID,
	})
	if err != nil {
		return srv, err
	}
	err = res.All(core_helper.Ctx, &srv)
	return srv, err
}

func (s *serviceStore) SelectByPeer(peerID primitive.ObjectID) ([]*model.ServiceModel, error) {
	srv := []*model.ServiceModel{}
	res, err := s.db.Collection(string(helper.CollServices)).Find(core_helper.Ctx, primitive.M{
		"peer_id": peerID,
	})
	if err != nil {
		return srv, err
	}
	err = res.All(core_helper.Ctx, &srv)
	return srv, err
}

func (s *serviceStore) SelectByPeerAndGroup(groupID, peerID primitive.ObjectID) (*model.ServiceModel, error) {
	srv := &model.ServiceModel{}
	err := s.db.Collection(string(helper.CollServices)).FindOne(core_helper.Ctx, primitive.M{
		"peer_id":  peerID,
		"group_id": groupID,
	}).Decode(srv)
	return srv, err
}

func (s *serviceStore) SelectFilter(filter *model.FilterServiceModel) ([]*model.ServiceModel, int64, error) {
	srv := []*model.ServiceModel{}

	ff := primitive.M{}
	if len(filter.Statuses) > 0 {
		ff["status"] = primitive.M{"$in": filter.Statuses}
	}
	if len(filter.Roles) > 0 {
		ff["role"] = primitive.M{"$in": filter.Roles}
	}
	if len(filter.Peers) > 0 {
		ff["peer_id"] = primitive.M{"$in": filter.Peers}
	}
	if len(filter.Groups) > 0 {
		ff["group_id"] = primitive.M{"$in": filter.Groups}
	}
	if filter.Host != "" {
		ff["host"] = primitive.M{"$regex": fmt.Sprintf("^(%v)", filter.Host), "$options": "i"}
	}
	if filter.Port > 0 {
		ff["$expr"] = primitive.M{"$regexMatch": primitive.M{
			"input": primitive.M{"$convert": primitive.M{
				"input": "$port",
				"to":    "string",
			}},
			"regex":   fmt.Sprintf("^(%v)", filter.Port),
			"options": "i",
		}}
	}
	if filter.Title != "" {
		ff["title"] = primitive.M{"$regex": fmt.Sprintf("(%v)", filter.Title), "$options": "i"}
	}

	pipeLimit := primitive.A{
		primitive.M{"$match": ff},
	}

	pipeCount := primitive.A{
		primitive.M{"$match": ff},
		primitive.M{"$count": "count"},
	}

	if filter.Offset > 0 {
		pipeLimit = append(pipeLimit, primitive.M{"$skip": filter.Offset})
	}
	if filter.Limit > 0 {
		pipeLimit = append(pipeLimit, primitive.M{"$limit": filter.Limit})
	}

	res, err := s.db.Collection(string(helper.CollServices)).Aggregate(core_helper.Ctx, pipeLimit)
	if err != nil {
		return srv, 0, err
	}
	if err = res.All(core_helper.Ctx, &srv); err != nil {
		return srv, 0, err
	}
	res, err = s.db.Collection(string(helper.CollServices)).Aggregate(core_helper.Ctx, pipeCount)
	if err != nil {
		return srv, 0, err
	}
	defer res.Close(core_helper.Ctx)
	count := &struct {
		Count int64 `bson:"count"`
	}{}
	for res.Next(core_helper.Ctx) {
		err = res.Decode(count)
	}
	return srv, count.Count, err
}

func (s *serviceStore) DeleteByKey(key string) error {
	_, err := s.db.Collection(string(helper.CollServices)).DeleteMany(core_helper.Ctx, primitive.M{
		"key": key,
	})
	return err
}
