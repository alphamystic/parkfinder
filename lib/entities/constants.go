package entities

import (
  "errors"
)

var UserNotLoggedIn = errors.New("User Not Logged in.")

var NoCLaims = errors.New("No Claims")
