package movie

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type Service interface {
	Search(ctx context.Context, opts SearchOpts) ([]Movie, error)
	Detail(ctx context.Context, opts DetailOpts) (MovieDetail, error)
}

type service struct {
	apiKey string
}

func NewService(
	apiKey string,
) Service {
	return &service{apiKey}
}

func (s *service) Search(ctx context.Context, opts SearchOpts) ([]Movie, error) {
	url := fmt.Sprintf("http://www.omdbapi.com/?apikey=%s&s=%s&type=%s&y=%s&page=%s", s.apiKey, opts.Title, opts.Type, opts.Year, opts.Page)

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	var result MovieSearchResponse
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result.Movies, nil
}

func (s *service) Detail(ctx context.Context, opts DetailOpts) (MovieDetail, error) {
	url := fmt.Sprintf("http://www.omdbapi.com/?apikey=%s&i=%s&t=%s&type=%s&y=%s&plot=%s", s.apiKey, opts.ID, opts.Title, opts.Type, opts.Year, opts.Plot)

	res, err := http.Get(url)
	if err != nil {
		return MovieDetail{}, err
	}

	defer res.Body.Close()

	var result MovieDetail
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return MovieDetail{}, err
	}

	return result, nil
}
