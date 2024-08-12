package main

import (
	"context"
	"net/http"

	"github.com/oneee-playground/job-seeker/platform/wanted"
	"github.com/oneee-playground/job-seeker/search"
)

func main() {
	wanted := wanted.NewPlatform(http.DefaultClient)

	if err := search.Search(context.Background(), wanted, search.Options{
		Keywords: []string{"Go", "Golang"},
		Limit:    20,
	}); err != nil {
		panic(err)
	}
}
