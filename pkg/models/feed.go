package models

import "github.com/mmcdole/gofeed"

type Feed struct {
  Name string `yaml:"name"`
  Url string `yaml:"url"`
}

type FeedItem struct {
  Title string `json:"title"`
  Link string `json:"link"`
  PublishedAt string `json:"published"`
  Description string `json:"description"`
  Categories []string `json:"categories"`
  Content string `json:"content"`
  Author *gofeed.Person `json:"author"`
  ThumbnailUrl string `json:"thumbnail_url"`
  GUID string `json:"guid"`
}

type FeedWithItems struct {
    Feed Feed
    Items []FeedItem
}