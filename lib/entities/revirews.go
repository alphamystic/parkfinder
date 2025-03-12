package entities

import (
  "github.com/alphamystic/parkfinder/lib/utils"
)

type Review struct {
  ID int
  ReviewID string
  ParkID string
  UserID string
  UserName string
  Comment string
  utils.TimeStamps
}
