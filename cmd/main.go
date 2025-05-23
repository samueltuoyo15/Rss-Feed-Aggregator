package main 

import (
  "fmt"
  "log"
  "encoding/json"
  "net/http"
  "github.com/samueltuoyo15/Rss-Feed-Aggregator/pkg/fetcher"
  "github.com/samueltuoyo15/Rss-Feed-Aggregator/pkg/parser"
  )
  
  
func feedFetcher(w http.ResponseWriter, r *http.Request){
  feeds, err := parser.LoadFeeds("config/feeds.yaml")
  if err != nil {
    http.Error(w, "Failed to load Feeds: %v", http.StatusInternalServerError)
    return 
  }
  
  results, err := fetcher.FetchAll(feeds)
    if err != nil {
    http.Error(w, "Failed to fetch Feeds: %v", http.StatusInternalServerError)
    return 
  }
  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(results)
}


func main() {
    http.HandleFunc("/api/feeds", feedFetcher)
    fmt.Printf("Server is running on port 8080\n")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}
