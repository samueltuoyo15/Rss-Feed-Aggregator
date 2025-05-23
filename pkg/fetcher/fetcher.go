package fetcher

import (
	"log"
	"strings"
	"sync"
	"github.com/mmcdole/gofeed"
	"github.com/samueltuoyo15/Rss-Feed-Aggregator/pkg/models"
)

func FetchAll(feeds []models.Feed) (map[string][]models.FeedItem, error) {
	var (
		wg sync.WaitGroup
		mu sync.Mutex
		parser = gofeed.NewParser()
		results = make(map[string][]models.FeedItem)
	)

	for _, feed := range feeds {
		wg.Add(1)
		go func(f models.Feed) {
			defer wg.Done()

			parsed, err := parser.ParseURL(f.Url)
			if err != nil {
				log.Printf("Failed to fetch %s: %v", f.Name, err)
				return
			}

			var items []models.FeedItem
			for _, item := range parsed.Items {
				feedItem := models.FeedItem{
					Title: item.Title,
					Link: item.Link,
					PublishedAt: item.Published,
					Description: item.Description,
					Categories: item.Categories,
					Content: item.Content,
					Author: item.Author,
					Image: item.Image,
					GUID: item.GUID,
				}

				if item.Image != nil {
					feedItem.Image = item.Image
				} else if len(item.Enclosures) > 0 {
					for _, enc := range item.Enclosures {
						if strings.HasPrefix(enc.Type, "image/") {
							feedItem.Image = &gofeed.Image{
								URL:   enc.URL,
								Title: "",
							}
							break
						}
					}
				}

				items = append(items, feedItem)
			}

			mu.Lock()
			results[f.Name] = items
			mu.Unlock()
		}(feed)
	}

	wg.Wait()
	return results, nil
}