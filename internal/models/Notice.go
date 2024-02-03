package models

import "time"

type Notice struct {
	Title           string
	Author          string
	ProviderID      int64
	Description     string
	PublicationDate time.Time
	Categories      []string
	Media           string
	Link            string
}
