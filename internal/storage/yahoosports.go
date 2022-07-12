package storage

import (
	"github.com/jrodolforojas/inside-goal-backend/internal/models"
	"github.com/mmcdole/gofeed"
)

type YahooSports struct {
	id      int64
	name    string
	feedURL string
}

func NewYahooSports() *YahooSports {
	return &YahooSports{
		id:      int64(YAHOO_SPORTS_ID),
		name:    "Yahoo Sports",
		feedURL: "https://sports.yahoo.com/soccer/rss/",
	}
}

func (yahooSports *YahooSports) GetNews(notices *[]models.Notice) error {
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL(yahooSports.feedURL)

	for _, item := range feed.Items {
		var author string
		if len(item.DublinCoreExt.Creator) > 0 {
			author = item.DublinCoreExt.Creator[0]
		}

		var categories []string
		categories = append(categories, item.Categories...)

		notice := models.Notice{
			Title:           item.Title,
			Author:          author,
			Description:     item.Description,
			PublicationDate: item.Published,
			Link:            item.Link,
			Categories:      categories,
			ProviderID:      yahooSports.id,
		}

		*notices = append(*notices, notice)
	}
	return nil
}