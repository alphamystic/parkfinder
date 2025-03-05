package entities

import (
  "ken/lib/utils"
)

type Newsletter struct {
  NewsID string
  Email string
  Subscribed bool
  utils.TimeStamps
}
