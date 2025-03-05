package entities

import (
  "ken/lib/utils"
)


type Park struct {
  ID int
  ParkID string
  Name string
  Location string
  Description string
  utils.TimeStamps
}
