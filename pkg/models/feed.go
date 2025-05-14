package models

type Feed struct {
  Name string `"yaml:name"`
  Url string `"yaml:url"`
}

type FeedItem struct {
  Title string `"json:title"`
  Link string `"json:link"`
  Description string `"json:description"`
}