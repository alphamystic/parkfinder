package entities

import (
  "ken/lib/utils"
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
