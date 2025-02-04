package storage

import (
	"sync"

	"github.com/jrodolforojas/inside-goal-backend/internal/models"
	"github.com/mmcdole/gofeed"
)

type ESPNStorage struct {
	id           int64
	name         string
	feedURL      string
	defaultImage string
	mu           *sync.Mutex
}

func NewESPN(mu *sync.Mutex) *ESPNStorage {
	return &ESPNStorage{
		id:           int64(ESPN_ID),
		name:         "ESPN",
		feedURL:      "https://www.espn.com/espn/rss/soccer/news",
		defaultImage: "https://media.wired.com/photos/5927404ccfe0d93c47432c13/master/pass/espn-logo.png",
		mu:           mu,
	}
}

func (espn *ESPNStorage) GetNews(notices *[]models.Notice) error {
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL(espn.feedURL)

	espnNotices := []models.Notice{}
	for _, item := range feed.Items {

		var media string

		if len(item.Enclosures) > 0 {
			media = item.Enclosures[0].URL
		}

		if media == "" {
			media = espn.defaultImage
		}

		notice := models.Notice{
			ProviderID:      espn.id,
			Title:           item.Title,
			Link:            item.Link,
			PublicationDate: *item.PublishedParsed,
			Author:          "ESPN",
			Description:     item.Description,
			Media:           media,
		}

		espnNotices = append(espnNotices, notice)
	}

	espn.mu.Lock()
	*notices = append(*notices, espnNotices...)
	espn.mu.Unlock()

	return nil
}

func (espn *ESPNStorage) GetProvider() (*models.Provider, error) {
	return &models.Provider{
		ID:           espn.id,
		Name:         espn.name,
		FeedURL:      espn.feedURL,
		DefaultImage: espn.defaultImage,
	}, nil
}
