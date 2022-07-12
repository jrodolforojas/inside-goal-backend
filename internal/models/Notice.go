package models

type Notice struct {
	Title           string
	Author          string
	ProviderID      int64
	Description     string
	PublicationDate string
	Categories      []string
	Media           string
	Link            string
}

