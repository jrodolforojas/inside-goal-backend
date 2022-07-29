package storage

import (
	"github.com/jrodolforojas/inside-goal-backend/internal/models"
	"github.com/mmcdole/gofeed"
)

type Marca struct {
	id           int64
	name         string
	feedURL      string
	defaultImage string
}

func NewMarca() *Marca {
	return &Marca{
		id:           int64(MARCA_ID),
		name:         "Marca",
		feedURL:      "https://e00-marca.uecdn.es/rss/futbol/futbol-internacional.xml",
		defaultImage: "https://e00-marca.uecdn.es/assets/v27/img/destacadas/marca__logo-generica.jpg",
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
		if len(item.Extensions["media"]["content"]) > 0 {
			media = item.Extensions["media"]["content"][0].Attrs["url"]
		}

		var description string
		if len(item.Extensions["media"]["description"]) > 0 {
			description = item.Extensions["media"]["description"][0].Value
		}

		if media == "" {
			media = marca.defaultImage
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

func (marca *Marca) GetProvider() (*models.Provider, error) {
	return &models.Provider{
		ID:           marca.id,
		Name:         marca.name,
		FeedURL:      marca.feedURL,
		DefaultImage: marca.defaultImage,
	}, nil
}