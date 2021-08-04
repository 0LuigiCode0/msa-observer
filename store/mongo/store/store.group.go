package store

import (
	"x-msa-core/grpc/msa_observer"
	core_helper "x-msa-core/helper"
	"x-msa-observer/helper"
	"x-msa-observer/store/mongo/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//groupStore хранилище
type GroupStore interface {
	Save(group *model.GroupModel) error
	Update(group *model.GroupModel) error
	DeleteDependents(ids []primitive.ObjectID, groupID primitive.ObjectID) error
	SelectByID(id primitive.ObjectID) (*model.GroupModel, error)
	SelectByTitle(title string) (*model.GroupModel, error)
	SelectAll() ([]*model.GroupModel, error)
	SelectForDependens(groupID primitive.ObjectID) ([]*msa_observer.Info, error)
	DeleteByID(id primitive.ObjectID) error
}

//groupStore хранилище
type groupStore struct {
	db *mongo.Database
}

func InitGroupStore(db *mongo.Database) GroupStore {
	return &groupStore{db: db}
}

func (s *groupStore) Save(group *model.GroupModel) error {
	res, err := s.db.Collection(string(helper.CollGroups)).UpdateOne(core_helper.Ctx, primitive.M{
		"title": group.Title,
	}, primitive.M{
		"$setOnInsert": group,
	}, options.Update().SetUpsert(true))
	if res != nil {
		if id, ok := res.UpsertedID.(primitive.ObjectID); ok {
			group.ID = id
		}
	}
	return err
}

func (s *groupStore) Update(group *model.GroupModel) error {
	_, err := s.db.Collection(string(helper.CollGroups)).UpdateOne(core_helper.Ctx, primitive.M{
		"_id": group.ID,
	}, primitive.M{
		"$set": group,
	}, options.Update().SetUpsert(false))
	return err
}

func (s *groupStore) DeleteDependents(ids []primitive.ObjectID, groupID primitive.ObjectID) error {
	filter := primitive.M{}
	if len(ids) > 0 {
		filter["_id"] = primitive.M{"$in": ids}
	}
	_, err := s.db.Collection(string(helper.CollGroups)).UpdateMany(core_helper.Ctx, filter, primitive.M{
		"$pull": primitive.M{"dependents": groupID},
	})
	return err
}

func (s *groupStore) SelectByID(id primitive.ObjectID) (*model.GroupModel, error) {
	group := &model.GroupModel{}
	err := s.db.Collection(string(helper.CollGroups)).FindOne(core_helper.Ctx, primitive.M{
		"_id": id,
	}).Decode(group)
	return group, err
}

func (s *groupStore) SelectByTitle(title string) (*model.GroupModel, error) {
	group := &model.GroupModel{}
	err := s.db.Collection(string(helper.CollGroups)).FindOne(core_helper.Ctx, primitive.M{
		"title": title,
	}).Decode(group)
	return group, err
}

func (s *groupStore) SelectAll() ([]*model.GroupModel, error) {
	group := []*model.GroupModel{}
	res, err := s.db.Collection(string(helper.CollGroups)).Find(core_helper.Ctx, primitive.M{})
	if err != nil {
		return group, err
	}
	err = res.All(core_helper.Ctx, &group)
	return group, err
}

func (s *groupStore) SelectForDependens(groupID primitive.ObjectID) ([]*msa_observer.Info, error) {
	dep := []*msa_observer.Info{}

	pipeLimit := primitive.A{
		primitive.M{"$match": primitive.M{
			"_id": groupID,
		}},
		primitive.M{"$lookup": primitive.M{
			"from":         helper.CollServices,
			"localField":   "dependents",
			"foreignField": "group_id",
			"as":           "dependents",
		}},
		primitive.M{"$lookup": "$dependents"},
		primitive.M{"$addFields": primitive.M{
			"dependents.group_id": "$_id",
		}},
		primitive.M{"$replaceRoot": primitive.M{
			"newRoot": "$dependents",
		}},
	}

	res, err := s.db.Collection(string(helper.CollGroups)).Aggregate(core_helper.Ctx, pipeLimit)
	if err != nil {
		return dep, err
	}
	err = res.All(core_helper.Ctx, &dep)
	return dep, err
}

func (s *groupStore) DeleteByID(id primitive.ObjectID) error {
	_, err := s.db.Collection(string(helper.CollGroups)).DeleteMany(core_helper.Ctx, primitive.M{
		"_id": id,
	})
	return err
}
