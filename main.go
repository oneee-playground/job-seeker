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
		Keywords: []string{"Go", "Golang", "golang"},
		Limit:    10,
		ExpYears: search.ExpYears{0, 3},
	}); err != nil {
		panic(err)
	}
}
