package storage

import (
	"sync"

	"github.com/jrodolforojas/inside-goal-backend/internal/models"
	"github.com/mmcdole/gofeed"
)

type NinetyMin struct {
	id           int64
	name         string
	feedURL      string
	defaultImage string
	mu           *sync.Mutex
}

func NewNinetyMin(mu *sync.Mutex) *NinetyMin {
	return &NinetyMin{
		id:           int64(NINETY_MIN_ID),
		name:         "90 min",
		feedURL:      "https://www.90min.com/posts.rss",
		defaultImage: "https://images2.minutemediacdn.com/image/upload/c_fill,w_1440,ar_1:1,f_auto,q_auto,g_auto/frontier/sites/logos/90min_New_Logo_90min_Horizontal_Orange_White.png",
		mu:           mu,
	}
}

func (ninetyMin *NinetyMin) GetNews(notices *[]models.Notice) error {
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL(ninetyMin.feedURL)

	ninentyMinNotices := []models.Notice{}

	for _, item := range feed.Items {
		var media string
		if len(item.Extensions["media"]["thumbnail"]) > 0 {
			media = item.Extensions["media"]["thumbnail"][0].Attrs["url"]
		}

		if media == "" {
			media = ninetyMin.defaultImage
		}

		author := "90 min"
		if item.Author != nil {
			author = item.Author.Name
		}

		notice := models.Notice{
			Title:           item.Title,
			Link:            item.Link,
			Author:          author,
			Description:     item.Description,
			PublicationDate: *item.PublishedParsed,
			Media:           media,
			ProviderID:      ninetyMin.id,
		}

		ninentyMinNotices = append(ninentyMinNotices, notice)
	}

	ninetyMin.mu.Lock()
	*notices = append(*notices, ninentyMinNotices...)
	ninetyMin.mu.Unlock()

	return nil
}

func (ninetyMin *NinetyMin) GetProvider() (*models.Provider, error) {
	return &models.Provider{
		ID:           ninetyMin.id,
		Name:         ninetyMin.name,
		FeedURL:      ninetyMin.feedURL,
		DefaultImage: ninetyMin.defaultImage,
	}, nil
}
