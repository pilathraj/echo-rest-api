package models

import (
	uuid "github.com/satori/go.uuid"
)

type (
	ErrorNotice struct {
		Type    string `json:"type"`
		Message string `json:"message"`
	}
)

var EmptySku = uuid.UUID{}
