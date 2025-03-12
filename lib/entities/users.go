package entities

import (
  "github.com/alphamystic/parkfinder/lib/utils"
)

type UserData struct {
  ID int
  UserID string
  Role string
  Phone string
  Name string
  Email string
  Password string
  utils.TimeStamps
}
