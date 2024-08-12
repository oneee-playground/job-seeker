package search

import (
	"context"
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
	Limit int
}

type Result struct {
	Platform JobPlatform
	Position string
	URL      string
}

func Search(ctx context.Context, platform Platform, opts Options) error {
	results, errchan := platform.Search(ctx, opts)

	for {
		select {
		case <-ctx.Done():
			return nil
		case err := <-errchan:
			return err
		case result := <-results:
			fmt.Println(result)
		}
	}
}

type Platform interface {
	Search(ctx context.Context, opts Options) (<-chan Result, <-chan error)
}
