# Inside Goal Backend

## Demo
[Inside Goal API](https://inside-goal.up.railway.app/)

## About
Inside Goal collects data from soccer rss feed.

RSS Feeds:

* [ESPN](https://www.espn.com/espn/rss/soccer/news)
* [DiarioAS](https://as.com/rss/futbol/mundial.xml)
* [Marca](https://e00-marca.uecdn.es/rss/futbol/futbol-internacional.xml)
* [New York Times](https://rss.nytimes.com/services/xml/rss/nyt/Soccer.xml)
* [Fox Sports](https://api.foxsports.com/v2/content/optimized-rss?partnerKey=MB0Wehpmuj2lUhuRhQaafhBjAJqaPU244mlTDK1i&size=30&tags=fs/soccer,soccer/epl/league/1,soccer/mls/league/5,soccer/ucl/league/7,soccer/europa/league/8,soccer/wc/league/12,soccer/euro/league/13,soccer/wwc/league/14,soccer/nwsl/league/20,soccer/cwc/league/26,soccer/gold_cup/league/32,soccer/unl/league/67)
* [Yahoo Sports](https://sports.yahoo.com/soccer/rss/)
* [90 min](https://www.90min.com/posts.rss)
* [101 Great Goals](https://www.101greatgoals.com/feed/) (not supported)

## Technologies
[![Go](https://img.shields.io/badge/Go-57b9d3?style=for-the-badge&logo=go&logoColor=white&labelColor=101010)]()

## Endpoints
### GET `{url}/news`
Returns the notices of all providers.

	{
	"Title": "England Delivers Its Dream, Beating Germany to Win Euro 2022",
	"Author": "The New York Times",
	"ProviderID": 4,
	"Description": "England won its first European Championship on Chloe Kelly’s goal in extra time, thrilling a record home crowd at Wembley Stadium.",
	"PublicationDate": "Mon, 01 Aug 2022 12:59:42 +0000",
	"Categories": null,
	"Media": "https://static01.nyt.com/images/2022/07/31/sports/31england-celebration/31england-celebration-moth.jpg",
	"Link": "https://www.nytimes.com/live/2022/07/31/sports/uefa-england-germany-euro"
    }
    ...

### GET `{url}/providers`
Returns providers information.

	{
	    "ID": 1,
	    "Name": "ESPN",
	    "FeedURL": "https://www.espn.com/espn/rss/soccer/news",
	    "DefaultImage": "https://media.wired.com/photos/5927404ccfe0d93c47432c13/master/pass/espn-logo.png"
	}
	...


## How to run?

In the project directory run:

`go run main.go`

Run locally the server.
