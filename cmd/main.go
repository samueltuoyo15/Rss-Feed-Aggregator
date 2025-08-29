package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
	"github.com/samueltuoyo15/Rss-Feed-Aggregator/pkg/fetcher"
	"github.com/samueltuoyo15/Rss-Feed-Aggregator/pkg/models"
	"github.com/samueltuoyo15/Rss-Feed-Aggregator/pkg/parser"
)

type TemplateData struct {
	Feeds []models.FeedWithItems
}

func feedFetcher(w http.ResponseWriter, r *http.Request) {
	feeds, err := parser.LoadFeeds("config/feeds.yaml")
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to load Feeds: %v", err), http.StatusInternalServerError)
		return
	}

	sortedResults, _, err := fetcher.FetchAll(feeds)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to fetch Feeds: %v", err), http.StatusInternalServerError)
		return
	}

	if r.Header.Get("HX-Request") == "true" {
		if err := renderFeedItems(w, sortedResults); err != nil {
			log.Printf("Template rendering error: %v", err)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(sortedResults); err != nil {
		log.Printf("JSON encoding error: %v", err)
	}
}

func renderFeedItems(w http.ResponseWriter, feeds []models.FeedWithItems) error {
	funcMap := template.FuncMap{
		"formatDate": func(t string) string {
			parsed, err := time.Parse(time.RFC1123Z, t)
			if err != nil {
				parsed, err = time.Parse(time.RFC1123, t)
				if err != nil {
					return t 
				}
			}
			return parsed.Format("Jan 2, 2006")
		},
		"truncate": func(s string, length int) string {
			if len(s) <= length {
				return s
			}
			return s[:length] + "..."
		},
	}

	tmpl, err := template.New("feed-items").Funcs(funcMap).ParseFiles("templates/feed_items.html")
	if err != nil {
		return fmt.Errorf("template parsing error: %w", err)
	}

	w.Header().Set("Content-Type", "text/html")
	if err := tmpl.Execute(w, TemplateData{Feeds: feeds}); err != nil {
		return fmt.Errorf("template execution error: %w", err)
	}
	return nil
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	w.Header().Set("Content-Type", "text/html")
	if err := tmpl.Execute(w, nil); err != nil {
		log.Printf("Home template error: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/api/feeds", feedFetcher)

	fmt.Printf("Server is running on port 5000\n")
	if err := http.ListenAndServe(":5000", nil); err != nil {
		log.Fatal(err)
	}
}