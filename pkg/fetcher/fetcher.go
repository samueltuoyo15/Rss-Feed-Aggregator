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
				thumbnailUrl := ""
				if mediaExt, ok := item.Extensions["media"]; ok {
					if thumbs, ok := mediaExt["thumbnail"]; ok && len(thumbs) > 0 {
						thumbnailUrl = thumbs[0].Attrs["url"]
					}
				}

				feedItem := models.FeedItem{
					Title: item.Title,
					Link: item.Link,
					PublishedAt: item.Published,
					Description: item.Description,
					Categories: item.Categories,
					Content: item.Content,
					Author: item.Author,
					ThumbnailUrl: thumbnailUrl, 
					GUID: item.GUID,
				}

				if thumbnailUrl == "" && len(item.Enclosures) > 0 {
					for _, enc := range item.Enclosures {
						if strings.HasPrefix(enc.Type, "image/") {
							feedItem.ThumbnailUrl = enc.URL
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