package main

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/same-ou/effective-go/internal/database"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	ApiKey    string    `json:"api_key"`
}

func databaseUserToUser(dbUser database.User) User {
	return User{
		ID:        dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Name:      dbUser.Name,
		ApiKey:    dbUser.ApiKey,
	}
}

type Feed struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	UserID    uuid.UUID `json:"user_id"`
}

func feedDAO(databaseFeed database.Feed) Feed {
	return Feed{
		ID:        databaseFeed.ID,
		CreatedAt: databaseFeed.CreatedAt,
		UpdatedAt: databaseFeed.UpdatedAt,
		Name:      databaseFeed.Name,
		Url:       databaseFeed.Url,
		UserID:    databaseFeed.UserID,
	}
}

func mapFeeds(dbfeeds []database.Feed) []Feed {
	feeds := []Feed{}
	for _, feed := range dbfeeds {
		feeds = append(feeds, feedDAO(feed))
	}
	return feeds
}

type FeedFollow struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    uuid.UUID `json:"user_id"`
	FeedID    uuid.UUID `json:"feed_id"`
}

func toFeedFollow(dbfeedFollow database.FeedFollow) FeedFollow {
	return FeedFollow{
		ID:        dbfeedFollow.ID,
		CreatedAt: dbfeedFollow.CreatedAt,
		UpdatedAt: dbfeedFollow.UpdatedAt,
		UserID:    dbfeedFollow.UserID,
		FeedID:    dbfeedFollow.FeedID,
	}
}

type Post struct {
	ID          uuid.UUID `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Title       string    `json:"title"`
	Description *string   `json:"description"`
	PublishedAt time.Time `json:"published_at"`
	Url         string    `json:"url"`
	FeedID      uuid.UUID `json:"feed_id"`
}

func toPost(dbPost database.Post) Post {
	return Post{
		ID:        dbPost.ID,
		CreatedAt: dbPost.CreatedAt,
		UpdatedAt: dbPost.UpdatedAt,
		Title:     dbPost.Title,
		Description: func(ns sql.NullString) *string {
			if ns.Valid {
				return &ns.String
			}
			return nil
		}(dbPost.Description),
		PublishedAt: dbPost.PublishedAt,
		Url:         dbPost.Url,
		FeedID:      dbPost.FeedID,
	}
}
