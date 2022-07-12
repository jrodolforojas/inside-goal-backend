package storage

import (
	"github.com/jrodolforojas/inside-goal-backend/internal/models"
	"github.com/mmcdole/gofeed"
)

type NYTimes struct {
	id      int64
	name    string
	feedURL string
}

func NewNYTimes() *NYTimes {
	return &NYTimes{
		id:      int64(NYTIMES_ID),
		name:    "New York Times",
		feedURL: "https://rss.nytimes.com/services/xml/rss/nyt/Soccer.xml",
	}
}

func (nytimes *NYTimes) GetNews(notices *[]models.Notice) error {
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL(nytimes.feedURL)

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

		notice := models.Notice{
			Title:           item.Title,
			Link:            item.Link,
			Description:     item.Description,
			Author:          author,
			PublicationDate: item.Published,
			Categories:      categories,
			Media:           media,
			ProviderID:      nytimes.id,
		}

		*notices = append(*notices, notice)
	}

	return nil
}
