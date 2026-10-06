package revkeen

import (
	"context"
	"errors"
	"iter"
)

var (
	// ErrPaginationDidNotAdvance protects callers from a server or adapter that
	// reports another page but repeats a state already visited by the iterator.
	ErrPaginationDidNotAdvance = errors.New("revkeen: pagination state did not advance")
	ErrNilPageFetcher          = errors.New("revkeen: page fetcher must not be nil")
)

// Page is the generator-independent result returned by a typed FetchPage.
// State may be a cursor string, an offset, a page number, or another comparable
// operation-specific value.
type Page[T any, State comparable] struct {
	Items     []T
	NextState State
	HasMore   bool
}

// FetchPage adapts one generated raw list call without replacing that call.
type FetchPage[T any, State comparable] func(context.Context, State) (Page[T, State], error)

// Iterator lazily fetches pages and can be ranged using All. It is immutable;
// each call to All starts again from InitialState.
type Iterator[T any, State comparable] struct {
	initialState State
	fetch        FetchPage[T, State]
}

func NewIterator[T any, State comparable](initialState State, fetch FetchPage[T, State]) (*Iterator[T, State], error) {
	if fetch == nil {
		return nil, ErrNilPageFetcher
	}
	return &Iterator[T, State]{initialState: initialState, fetch: fetch}, nil
}

// All returns a lazy, range-friendly sequence. Fetch and context errors are
// yielded once as the second value, after which iteration stops.
func (iterator *Iterator[T, State]) All(ctx context.Context) iter.Seq2[T, error] {
	return func(yield func(T, error) bool) {
		var zero T
		if iterator == nil || iterator.fetch == nil {
			yield(zero, ErrNilPageFetcher)
			return
		}
		if ctx == nil {
			ctx = context.Background()
		}

		state := iterator.initialState
		visited := map[State]struct{}{state: {}}
		for {
			if err := ctx.Err(); err != nil {
				yield(zero, err)
				return
			}
			page, err := iterator.fetch(ctx, state)
			if err != nil {
				yield(zero, err)
				return
			}
			for _, item := range page.Items {
				if err := ctx.Err(); err != nil {
					yield(zero, err)
					return
				}
				if !yield(item, nil) {
					return
				}
			}
			if len(page.Items) == 0 {
				return
			}
			if !page.HasMore {
				return
			}
			if _, seen := visited[page.NextState]; seen {
				yield(zero, ErrPaginationDidNotAdvance)
				return
			}
			visited[page.NextState] = struct{}{}
			state = page.NextState
		}
	}
}
