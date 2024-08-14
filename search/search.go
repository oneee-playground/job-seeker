package search

import (
	"context"
	"errors"
	"fmt"
)

type JobPlatform string

const (
	PlatformWanted JobPlatform = "wanted"
)

type Options struct {
	Keywords []string
	// Limit softly limits the maximum count of found job positions.
	// If Limit < 0, there will be no limit.
	Limit    int
	ExpYears ExpYears
}

type Result struct {
	Platform JobPlatform
	Company  string
	Position string
	URL      string
}

func Search(ctx context.Context, platform Platform, opts Options) error {
	results, errchan := platform.Search(ctx, opts)

	if !opts.ExpYears.Valid() {
		return errors.New("experience years is not valid")
	}

	for {
		select {
		case <-ctx.Done():
			return nil
		case err := <-errchan:
			return err
		case result, ok := <-results:
			if !ok {
				return nil
			}
			fmt.Println(result)
		}
	}
}

type Platform interface {
	Search(ctx context.Context, opts Options) (<-chan Result, <-chan error)
}
