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
	Feeds map[string][]models.FeedItem
}

func feedFetcher(w http.ResponseWriter, r *http.Request) {
	feeds, err := parser.LoadFeeds("config/feeds.yaml")
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to load Feeds: %v", err), http.StatusInternalServerError)
		return
	}

	results, err := fetcher.FetchAll(feeds)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to fetch Feeds: %v", err), http.StatusInternalServerError)
		return
	}

	if r.Header.Get("HX-Request") == "true" {
		renderFeedItems(w, results)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}

func renderFeedItems(w http.ResponseWriter, feeds map[string][]models.FeedItem) {
	funcMap := template.FuncMap(template.FuncMap{
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
	})
	tmpl := template.Must(template.New("feed-items").Funcs(funcMap).ParseFiles(`templates/feed_items.html`))
	w.Header().Set("Content-Type", "text/html")
	if err := tmpl.Execute(w, feeds); err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	w.Header().Set("Content-Type", "text/html")
	tmpl.Execute(w, nil)
}

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/api/feeds", feedFetcher)

	fmt.Printf("Server is running on port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}