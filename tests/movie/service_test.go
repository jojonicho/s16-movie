package movie

import (
	"context"
	"os"
	"testing"

	movie "github.com/jojonicho/s16-movie/movie"
)

func TestDetail(t *testing.T) {
	apiKey := os.Getenv("OMDB_KEY")
	svc := movie.NewService(apiKey)

	ctx := context.Context(nil)
	opts := movie.DetailOpts{
		Title: "pokemon",
	}

	data, err := svc.Detail(ctx, opts)

	// Response = true

	if err != nil {
		t.Errorf("Error is not nil")
	}

	if data.Response != "True" {
		t.Errorf("Response is not true")
	}
}

func TestSearch(t *testing.T) {
	apiKey := os.Getenv("OMDB_KEY")
	svc := movie.NewService(apiKey)

	ctx := context.Context(nil)
	opts := movie.SearchOpts{
		Title: "pokemon",
	}

	_, err := svc.Search(ctx, opts)

	if err != nil {
		t.Errorf("Error is not nil")
	}
}

func TestDetail_ById(t *testing.T) {
	apiKey := os.Getenv("OMDB_KEY")
	svc := movie.NewService(apiKey)

	ctx := context.Context(nil)
	opts := movie.DetailOpts{
		ID: "tt5884052",
	}

	data, err := svc.Detail(ctx, opts)

	if err != nil {
		t.Errorf("Error is not nil")
	}

	if data.Title != "Pokémon: Detective Pikachu" {
		t.Errorf("Title is not 'Pokémon: Detective Pikachu'")
	}
}
