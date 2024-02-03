package storage

import (
	"github.com/jrodolforojas/inside-goal-backend/internal/models"
	"github.com/mmcdole/gofeed"
)

type FoxSports struct {
	id           int64
	name         string
	feedURL      string
	defaultImage string
}

func NewFoxSports() *FoxSports {
	return &FoxSports{
		id:           int64(FOX_SPORTS_ID),
		name:         "Fox Sports",
		feedURL:      "https://api.foxsports.com/v2/content/optimized-rss?partnerKey=MB0Wehpmuj2lUhuRhQaafhBjAJqaPU244mlTDK1i&size=30&tags=fs/soccer,soccer/epl/league/1,soccer/mls/league/5,soccer/ucl/league/7,soccer/europa/league/8,soccer/wc/league/12,soccer/euro/league/13,soccer/wwc/league/14,soccer/nwsl/league/20,soccer/cwc/league/26,soccer/gold_cup/league/32,soccer/unl/league/67",
		defaultImage: "https://upload.wikimedia.org/wikipedia/commons/thumb/5/56/Fox_Sports_logo1.svg/2560px-Fox_Sports_logo1.svg.png",
	}
}

func (foxSports *FoxSports) GetNews(notices *[]models.Notice) error {
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL(foxSports.feedURL)

	for _, item := range feed.Items {

		var categories []string
		categories = append(categories, item.Categories...)

		var media string
		if len(item.Extensions["media"]["content"]) > 0 {
			media = item.Extensions["media"]["content"][0].Attrs["url"]
		}

		if media == "" {
			media = foxSports.defaultImage
		}

		notice := models.Notice{
			Title:           item.Title,
			Link:            item.Link,
			Categories:      categories,
			Description:     item.Description,
			PublicationDate: *item.PublishedParsed,
			Media:           media,
			Author:          "Fox Sports",
			ProviderID:      foxSports.id,
		}

		*notices = append(*notices, notice)
	}
	return nil
}

func (foxSports *FoxSports) GetProvider() (*models.Provider, error) {
	return &models.Provider{
		ID:           foxSports.id,
		Name:         foxSports.name,
		FeedURL:      foxSports.feedURL,
		DefaultImage: foxSports.defaultImage,
	}, nil
}