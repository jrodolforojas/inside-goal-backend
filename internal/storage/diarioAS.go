package storage

import (
	"sync"

	"github.com/jrodolforojas/inside-goal-backend/internal/models"
	"github.com/mmcdole/gofeed"
)

type DiarioASStorage struct {
	id           int64
	name         string
	feedURL      string
	defaultImage string
	mu           *sync.Mutex
}

func NewDiarioAS(mu *sync.Mutex) *DiarioASStorage {
	return &DiarioASStorage{
		id:           int64(DIARIOAS_ID),
		name:         "Diario AS",
		feedURL:      "https://feeds.as.com/mrss-s/pages/as/site/as.com/section/futbol/portada/",
		defaultImage: "https://upload.wikimedia.org/wikipedia/commons/thumb/d/d8/Diario_AS.svg/1200px-Diario_AS.svg.png",
		mu:           mu,
	}
}

func (diarioAS *DiarioASStorage) GetNews(notices *[]models.Notice) error {
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL(diarioAS.feedURL)

	diarioASNotices := []models.Notice{}

	for _, item := range feed.Items {
		var author string

		if len(item.DublinCoreExt.Creator) > 0 {
			author = item.DublinCoreExt.Creator[0]
		}

		var categories []string
		categories = append(categories, item.Categories...)

		var media string
		if len(item.Enclosures) > 0 {
			media = item.Enclosures[0].URL
		}

		if media == "" {
			media = diarioAS.defaultImage
		}

		notice := models.Notice{
			Title:           item.Title,
			Link:            item.Link,
			Author:          author,
			Description:     item.Description,
			PublicationDate: *item.PublishedParsed,
			Categories:      categories,
			Media:           media,
			ProviderID:      diarioAS.id,
		}

		diarioASNotices = append(diarioASNotices, notice)
	}

	diarioAS.mu.Lock()
	*notices = append(*notices, diarioASNotices...)
	diarioAS.mu.Unlock()
	return nil
}

func (diarioAS *DiarioASStorage) GetProvider() (*models.Provider, error) {
	return &models.Provider{
		ID:           diarioAS.id,
		Name:         diarioAS.name,
		FeedURL:      diarioAS.feedURL,
		DefaultImage: diarioAS.defaultImage,
	}, nil
}
