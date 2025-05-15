package main 

import (
  "fmt"
  "log"
  "github.com/samueltuoyo15/Rss-Feed-Aggregator/pkg/fetcher"
  "github.com/samueltuoyo15/Rss-Feed-Aggregator/pkg/parser"
  )
  
func main(){
  feeds, err := parser.LoadFeeds("config/feeds.yaml")
  if err != nil {
    log.Fatalf("Failed to load Feeds: %v", err)
  }
  
  results, err := fetcher.FetchAll(feeds)
    if err != nil {
    log.Fatalf("Failed to fetch Feeds: %v", err)
  }
  
  for name, items := range results {
    fmt.Printf("=== %s ===\n", name)
    for _, item := range items {
     	fmt.Printf("📰 %s\n🔗 %s\n\n", item.Title, item.Link)
    }
  }
}


