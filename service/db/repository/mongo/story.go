package mongo

import (
	"context"
	"errors"
	"fmt"
	contract "github.com/peacewalker122/project/service/db/contract/mongo"
	"github.com/peacewalker122/project/service/db/model/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

const cOLLECTION = "stories"

type StoryRepository struct {
	*mongo.Database
}

func NewStoryRepository(db *mongo.Database) *StoryRepository {
	return &StoryRepository{db}
}

func (s *StoryRepository) Get(ctx context.Context, id string) (*model.Story, error) {
	var StoryData model.Story

	err := s.Collection(cOLLECTION).FindOne(ctx, bson.M{
		"_id": id,
	}).Decode(&StoryData)
	if err != nil {
		return nil, err
	}

	return &StoryData, nil
}

func (s *StoryRepository) GetAll(ctx context.Context, usersID []string) ([]*model.Story, error) {
	var (
		storyData []*model.Story
		err       error
		filter    bson.M
	)

	if len(usersID) == 0 {
		return nil, errors.New("no user id provided")
	}

	filter = bson.M{
		"user_id": bson.M{
			"$in": usersID,
		},
	}

	cursor, err := s.Collection(cOLLECTION).Find(ctx, filter)
	if err != nil {
		return nil, errors.New("error while getting stories")
	}

	for cursor.Next(ctx) {
		var story model.Story
		if err := cursor.Decode(&story); err != nil {
			return nil, fmt.Errorf("error decoding document: %s", err)
		}
		storyData = append(storyData, &story)
	}

	cursor.Close(ctx)
	return storyData, nil
}

func (s *StoryRepository) Create(ctx context.Context, story *model.Story) error {

	_, err := s.Collection(cOLLECTION).InsertOne(ctx, story.ToStoryDB())
	if err != nil {
		log.Println(err)
		return errors.New("error while creating story")
	}

	return nil
}

func (s *StoryRepository) Delete(ctx context.Context, id string) error {
	err := s.Collection(cOLLECTION).FindOneAndDelete(ctx, bson.M{
		"_id": id,
	}).Err()
	if err != nil {
		return errors.New("error while deleting story")
	}

	return nil
}

var _ contract.StoriesRepoitory = (*StoryRepository)(nil)
