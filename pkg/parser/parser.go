package parser

import (
  "os"
 	"gopkg.in/yaml.v3"
  "github.com/samueltuoyo15/Rss-Feed-Aggregator/pkg/models"
  )

type Config struct {
  Feeds []models.Feed `yaml:"feeds"`
}

func LoadFeeds(path string) ([]models.Feed, error) {
  data, err := os.ReadFile(path)
  if err != nil {
    return nil, err
  }
  
  var config Config
  if err := yaml.Unmarshal(data, &config); err != nil {
    return nil, err
  }
  return config.Feeds, nil
}