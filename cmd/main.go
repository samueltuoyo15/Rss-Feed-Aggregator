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
	tmpl := template.Must(template.New("feed-items").Funcs(funcMap).Parse(`
		{{ range $source, $items := . }}
			{{ range $items }}
				<article class="w-full p-4 bg-white rounded-lg border border-gray-200 hover:shadow-md transition-shadow">
					<div class="flex items-center gap-2 mb-2">
						{{ if .Categories }}
						<span class="px-2 py-1 text-xs font-medium rounded-full bg-amber-100 text-amber-800">
							{{ index .Categories 0 }}
						</span>
						{{ end }}
						<time class="text-xs italic text-gray-400">
							{{ if .PublishedAt }}
								{{ formatDate .PublishedAt }}
							{{ else }}
								Unknown date
							{{ end }}
						</time>
						<span class="text-xs text-gray-500 ml-auto">{{ $source }}</span>
					</div>
					<h2 class="text-lg font-bold text-gray-900">
						<a href="{{ .Link }}" target="_blank" rel="noopener noreferrer">{{ .Title }}</a>
					</h2>
					<p class="mt-1 text-gray-600 line-clamp-2">
						{{ if .Description }}
							{{ .Description }}
						{{ else if .Content }}
							{{ truncate .Content 150 }}
						{{ end }}
					</p>
					
					<div class="flex items-center justify-between mt-3 pt-3 border-t border-gray-100">
						<div class="flex space-x-4">
							<button class="flex items-center text-gray-500 hover:text-gray-700 text-sm">
								<svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4 mr-1" fill="none" viewBox="0 0 24 24" stroke="currentColor">
									<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 12h.01M12 12h.01M16 12h.01M21 12c0 4.418-4.03 8-9 8a9.863 9.863 0 01-4.255-.949L3 20l1.395-3.72C3.512 15.042 3 13.574 3 12c0-4.418 4.03-8 9-8s9 3.582 9 8z" />
								</svg>
								0
							</button>
						</div>
						
						<div class="flex space-x-2">
							<button class="text-gray-500 hover:text-gray-700 p-1 rounded-full hover:bg-gray-100">
								<svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
									<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 5a2 2 0 012-2h10a2 2 0 012 2v16l-7-3.5L5 21V5z" />
								</svg>
							</button>
							<button class="text-gray-500 hover:text-gray-700 p-1 rounded-full hover:bg-gray-100">
								<svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
									<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8.684 13.342C8.886 12.938 9 12.482 9 12c0-.482-.114-.938-.316-1.342m0 2.684a3 3 0 110-2.684m0 2.684l6.632 3.316m-6.632-6l6.632-3.316m0 0a3 3 0 105.367-2.684 3 3 0 00-5.367 2.684zm0 9.316a3 3 0 105.368 2.684 3 3 0 00-5.368-2.684z" />
								</svg>
							</button>
						</div>
					</div>
				</article>
			{{ end }}
		{{ end }}
	`))


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