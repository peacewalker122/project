package model

import (
	"github.com/google/uuid"
	"github.com/peacewalker122/project/util/file"
	"time"
)

type Story struct {
	ID          string    `bson:"_id"`
	UserID      string    `bson:"user_id"`
	MediaLink   string    `bson:"media_link"`
	MentionUser []string  `bson:"mention_user,omitempty"`
	Content     file.MIME `bson:"content"`
	CreatedAt   time.Time `bson:"created_at"`
}

type StoryDB struct {
	ID          string    `bson:"_id"`
	UserID      string    `bson:"user_id"`
	MediaLink   string    `bson:"media_link"`
	MentionUser []string  `bson:"mention_user"`
	Content     file.MIME `bson:"content"`
	CreatedAt   time.Time `bson:"created_at"`
	ExpiredAt   time.Time `bson:"expired_at"`
}

func (s *Story) ToStoryDB() *StoryDB {
	id, _ := uuid.NewRandom()
	if s.ID == "" {
		s.ID = id.String()
	}
	return &StoryDB{
		ID:          s.ID,
		UserID:      s.UserID,
		MediaLink:   s.MediaLink,
		MentionUser: s.MentionUser,
		Content:     s.Content,
		CreatedAt:   time.Now(),
		ExpiredAt:   time.Now().Add(24 * time.Hour),
	}
}

func (s *StoryDB) ToStory() *Story {
	return &Story{
		ID:          s.ID,
		UserID:      s.UserID,
		MediaLink:   s.MediaLink,
		MentionUser: s.MentionUser,
		Content:     s.Content,
		CreatedAt:   s.CreatedAt,
	}
}
