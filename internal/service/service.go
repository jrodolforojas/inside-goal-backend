package service

import (
	"context"

	"github.com/Rviewer-Challenges/4TOWna2EWttcHFagNbUw/api/internal/models"
	"github.com/Rviewer-Challenges/4TOWna2EWttcHFagNbUw/api/internal/storage"
)

type Feed struct {
	News *storage.Storage
}

func New(storage *storage.Storage) *Feed {
	return &Feed{
		News: storage,
	}
}

func (feed *Feed) GetNews(ctx context.Context) ([]models.Notice, error) {
	notices := []models.Notice{}

	errc := make(chan error, 8)

	go func() {
		err := feed.News.GetNewsESPN(&notices)
		if err != nil {
			errc <- err
			return
		}

		errc <- nil
	}()

	go func() {
		err := feed.News.GetNewsDiarioAS(&notices)
		if err != nil {
			errc <- err
			return
		}

		errc <- nil
	}()

	go func() {
		err := feed.News.GetNewsMarca(&notices)
		if err != nil {
			errc <- err
			return
		}

		errc <- nil
	}()

	go func() {
		err := feed.News.GetNewsNYTimes(&notices)
		if err != nil {
			errc <- err
			return
		}

		errc <- nil
	}()

	go func() {
		err := feed.News.GetNewsFoxSports(&notices)
		if err != nil {
			errc <- err
			return
		}

		errc <- nil
	}()

	go func() {
		err := feed.News.GetNewsYahoo(&notices)
		if err != nil {
			errc <- err
			return
		}

		errc <- nil
	}()

	go func() {
		err := feed.News.GetNews101GreatGoals(&notices)
		if err != nil {
			errc <- err
			return
		}

		errc <- nil
	}()

	go func() {
		err := feed.News.Get90Min(&notices)
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
