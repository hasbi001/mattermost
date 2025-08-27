package middleware

import (
  "time"
)

var lastTrigger time.Time
var maxPerDay = 20
var todayCount = 0

func AllowTrigger() bool {
  now := time.Now()

  // reset harian
  if now.Day() != lastTrigger.Day() {
    todayCount = 0
  }

  if todayCount >= maxPerDay {
    return false
  }

  // enforce minimal jeda antar trigger
  if now.Sub(lastTrigger) < 30*time.Minute {
    return false
  }

  todayCount++
  lastTrigger = now
  return true
}