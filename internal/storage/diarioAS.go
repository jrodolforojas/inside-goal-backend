package storage

import (
	"github.com/jrodolforojas/inside-goal-backend/internal/models"
	"github.com/mmcdole/gofeed"
)

type DiarioASStorage struct {
	id      int64
	name    string
	feedURL string
}

func NewDiarioAS() *DiarioASStorage {
	return &DiarioASStorage{
		id:      int64(DIARIOAS_ID),
		name:    "Diario AS",
		feedURL: "https://as.com/rss/futbol/mundial.xml",
	}
}

func (diarioAS *DiarioASStorage) GetNews(notices *[]models.Notice) error {
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL(diarioAS.feedURL)

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

		notice := models.Notice{
			Title:           item.Title,
			Link:            item.Link,
			Author:          author,
			Description:     item.Description,
			PublicationDate: item.Published,
			Categories:      categories,
			Media:           media,
			ProviderID:      diarioAS.id,
		}

		*notices = append(*notices, notice)
	}
	return nil
}
