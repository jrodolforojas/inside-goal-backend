package service

import (
	"context"
	"fmt"
	"sort"
	"sync"

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

	wg := sync.WaitGroup{}
	mu := sync.Mutex{}

	errMutex := sync.Mutex{}
	var errors []error

	wg.Add(1)
	go func() {
		defer wg.Done()
		espn := storage.NewESPN(&mu)
		if err := espn.GetNews(&notices); err != nil {
			errMutex.Lock()
			errors = append(errors, err)
			errMutex.Unlock()
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		diarioAS := storage.NewDiarioAS(&mu)
		if err := diarioAS.GetNews(&notices); err != nil {
			errMutex.Lock()
			errors = append(errors, err)
			errMutex.Unlock()
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		marca := storage.NewMarca(&mu)
		if err := marca.GetNews(&notices); err != nil {
			errMutex.Lock()
			errors = append(errors, err)
			errMutex.Unlock()
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		nyTimes := storage.NewNYTimes(&mu)
		if err := nyTimes.GetNews(&notices); err != nil {
			errMutex.Lock()
			errors = append(errors, err)
			errMutex.Unlock()
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		foxSports := storage.NewFoxSports(&mu)
		if err := foxSports.GetNews(&notices); err != nil {
			errMutex.Lock()
			errors = append(errors, err)
			errMutex.Unlock()
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		yahooSports := storage.NewYahooSports(&mu)
		if err := yahooSports.GetNews(&notices); err != nil {
			errMutex.Lock()
			errors = append(errors, err)
			errMutex.Unlock()
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		ninetyMin := storage.NewNinetyMin(&mu)
		if err := ninetyMin.GetNews(&notices); err != nil {
			errMutex.Lock()
			errors = append(errors, err)
			errMutex.Unlock()
		}
	}()

	wg.Wait()

	if len(errors) > 0 {
		fmt.Println("Errors occurred:")
		for _, err := range errors {
			fmt.Println(err)
		}
	}

	// order the notices by date
	sort.Slice(notices, func(i, j int) bool {
		return notices[i].PublicationDate.After(notices[j].PublicationDate)
	})

	fmt.Printf("len notices: %v\n", len(notices))

	return notices, nil
}

func (feed *Feed) GetProviders(ctx context.Context) ([]models.Provider, error) {
	providers := []models.Provider{}
	return providers, nil
}
