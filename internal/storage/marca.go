package storage

import (
	"github.com/jrodolforojas/inside-goal-backend/internal/models"
	"github.com/mmcdole/gofeed"
)

type Marca struct {
	id      int64
	name    string
	feedURL string
}

func NewMarca() *Marca {
	return &Marca{
		id:      int64(MARCA_ID),
		name:    "Marca",
		feedURL: "https://e00-marca.uecdn.es/rss/futbol/futbol-internacional.xml",
	}
}

func (marca *Marca) GetNews(notices *[]models.Notice) error {
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL(marca.feedURL)

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

		var description string
		if len(item.Extensions["media"]["description"]) > 0 {
			description = item.Extensions["media"]["description"][0].Value
		}

		notice := models.Notice{
			Title:           item.Title,
			Author:          author,
			Description:     description,
			Link:            item.Link,
			Categories:      categories,
			PublicationDate: item.Published,
			Media:           media,
			ProviderID:      marca.id,
		}

		*notices = append(*notices, notice)
	}
	return nil
}
