package entities

import (
  "github.com/alphamystic/parkfinder/lib/utils"
)

type Newsletter struct {
  NewsID string
  Email string
  Subscribed bool
  utils.TimeStamps
}
