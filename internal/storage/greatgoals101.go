package storage

import (
	"github.com/jrodolforojas/inside-goal-backend/internal/models"
	"github.com/mmcdole/gofeed"
)

type GreatGoals101 struct {
	id           int64
	name         string
	feedURL      string
	defaultImage string
}

func NewGreatGoals101() *GreatGoals101 {
	return &GreatGoals101{
		id:           int64(GREAT_GOALS_101_ID),
		name:         "101 Great Goals",
		feedURL:      "https://www.101greatgoals.com/feed/",
		defaultImage: "https://www.101greatgoals.com/wp-content/uploads/2022/02/101GG-logo.jpg",
	}
}

func (greatGoals101 *GreatGoals101) GetNews(notices *[]models.Notice) error {
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL(greatGoals101.feedURL)

	for _, item := range feed.Items {
		var author string
		if len(item.DublinCoreExt.Creator) > 0 {
			author = item.DublinCoreExt.Creator[0]
		}

		notice := models.Notice{
			Title:           item.Title,
			Link:            item.Link,
			PublicationDate: item.Published,
			Author:          author,
			Description:     item.Description,
			ProviderID:      greatGoals101.id,
			Media:           greatGoals101.defaultImage,
		}

		*notices = append(*notices, notice)
	}
	return nil
}

func (greatGoals101 *GreatGoals101) GetProvider() (*models.Provider, error) {
	return &models.Provider{
		ID:           greatGoals101.id,
		Name:         greatGoals101.name,
		FeedURL:      greatGoals101.feedURL,
		DefaultImage: greatGoals101.defaultImage,
	}, nil
}
