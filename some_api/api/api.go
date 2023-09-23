package api

import (
  "encoding/json"
  "net/http"
)

type TurboCoinsBalance struct {
  Username string
  Amount int
}

type TurboCoinsBalanceResponse struct {
  Code int
  Balance int64
}

type TurboError struct {
  Code int
  Message string
}