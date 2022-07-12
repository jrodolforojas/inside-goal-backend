package service

import (
	"context"

	"github.com/jrodolforojas/inside-goal-backend/internal/models"
	"github.com/jrodolforojas/inside-goal-backend/internal/storage"
)

type Feed struct {
}

func New() *Feed {
	return &Feed{}
}

func (feed *Feed) GetNews(ctx context.Context) ([]models.Notice, error) {
	notices := []models.Notice{}

	errc := make(chan error, 8)

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
		greatGoals101 := storage.NewGreatGoals101()
		err := greatGoals101.GetNews(&notices)
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
	for i := 0; i < 8; i++ {
		e := <-errc
		if e != nil {
			err = e
		}
	}

	if err != nil {
		return nil, err
	}

	return notices, nil
}
