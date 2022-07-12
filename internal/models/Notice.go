package models

type Notice struct {
	Title           string
	Author          string
	Provider        Provider
	Description     string
	PublicationDate string
	Categories      []string
	Media           string
	Link            string
}

type Provider struct {
	Id   int64
	Name string
}
