package service

import (
	"context"
	"sort"

	"github.com/jrodolforojas/inside-goal-backend/internal/models"
	"github.com/jrodolforojas/inside-goal-backend/internal/storage"
)

type Feed struct {
}

func New() *Feed {
	return &Feed{}
}

const PROVIDERS = 7

func (feed *Feed) GetNews(ctx context.Context) ([]models.Notice, error) {
	notices := []models.Notice{}

	errc := make(chan error, PROVIDERS)

	go func() {
		espn := storage.NewESPN()
		err := espn.GetNews(&notices)
		if err != nil {
			errc <- err
			return
		}

		errc <- nil
	}()

	go func() {
		diarioAS := storage.NewDiarioAS()
		err := diarioAS.GetNews(&notices)
		if err != nil {
			errc <- err
			return
		}

		errc <- nil
	}()

	go func() {
		marca := storage.NewMarca()
		err := marca.GetNews(&notices)
		if err != nil {
			errc <- err
			return
		}

		errc <- nil
	}()

	go func() {
		nyTimes := storage.NewNYTimes()
		err := nyTimes.GetNews(&notices)
		if err != nil {
			errc <- err
			return
		}

		errc <- nil
	}()

	go func() {
		foxSports := storage.NewFoxSports()
		err := foxSports.GetNews(&notices)
		if err != nil {
			errc <- err
			return
		}

		errc <- nil
	}()

	go func() {
		yahooSports := storage.NewYahooSports()
		err := yahooSports.GetNews(&notices)
		if err != nil {
			errc <- err
			return
		}

		errc <- nil
	}()

	go func() {
		ninetyMin := storage.NewNinetyMin()
		err := ninetyMin.GetNews(&notices)
		if err != nil {
			errc <- err
			return
		}

		errc <- nil
	}()

	var err error
	for i := 0; i < PROVIDERS; i++ {
		e := <-errc
		if e != nil {
			err = e
		}
	}

	if err != nil {
		return nil, err
	}

	// order the notices by date
	sort.Slice(notices, func(i, j int) bool {
		return notices[i].PublicationDate.After(notices[j].PublicationDate)
	})

	return notices, nil
}

func (feed *Feed) GetProviders(ctx context.Context) ([]models.Provider, error) {
	providers := []models.Provider{}

	espn := storage.NewESPN()
	espnProvider, _ := espn.GetProvider()
	providers = append(providers, *espnProvider)

	diarioAS := storage.NewDiarioAS()
	diarioASProvider, _ := diarioAS.GetProvider()
	providers = append(providers, *diarioASProvider)

	marca := storage.NewMarca()
	marcaProvider, _ := marca.GetProvider()
	providers = append(providers, *marcaProvider)

	nyTimes := storage.NewNYTimes()
	nyTimesProvider, _ := nyTimes.GetProvider()
	providers = append(providers, *nyTimesProvider)

	foxsports := storage.NewFoxSports()
	foxsportsProvider, _ := foxsports.GetProvider()
	providers = append(providers, *foxsportsProvider)

	yahooSports := storage.NewYahooSports()
	yahooSportsProvider, _ := yahooSports.GetProvider()
	providers = append(providers, *yahooSportsProvider)

	ninetyMin := storage.NewNinetyMin()
	ninetyMinProvider, _ := ninetyMin.GetProvider()
	providers = append(providers, *ninetyMinProvider)

	return providers, nil
}
