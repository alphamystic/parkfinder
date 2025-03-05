package entities

import (
  "ken/lib/utils"
)

type ContactUs struct{
  CuID string
  Name string
  Email string
  Subject string
  Message string
  Handled bool
  utils.TimeStamps
}
