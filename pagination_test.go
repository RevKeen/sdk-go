package revkeen

import (
	"context"
	"errors"
	"reflect"
	"testing"
)

func TestIteratorLazilyRangesMultiplePages(t *testing.T) {
	fetches := 0
	iterator, err := NewIterator(0, func(_ context.Context, state int) (Page[string, int], error) {
		fetches++
		switch state {
		case 0:
			return Page[string, int]{Items: []string{"a", "b"}, NextState: 1, HasMore: true}, nil
		case 1:
			return Page[string, int]{Items: []string{"c"}}, nil
		default:
			t.Fatalf("unexpected pagination state %d", state)
			return Page[string, int]{}, nil
		}
	})
	if err != nil {
		t.Fatal(err)
	}
	if fetches != 0 {
		t.Fatal("iterator fetched before it was ranged")
	}

	var items []string
	for item, rangeErr := range iterator.All(context.Background()) {
		if rangeErr != nil {
			t.Fatal(rangeErr)
		}
		items = append(items, item)
	}
	if !reflect.DeepEqual(items, []string{"a", "b", "c"}) || fetches != 2 {
		t.Fatalf("items/fetches = %v/%d", items, fetches)
	}
}

func TestIteratorDetectsRepeatedOrCyclicState(t *testing.T) {
	for _, test := range []struct {
		name string
		next func(int) int
	}{
		{name: "same state", next: func(state int) int { return state }},
		{name: "cycle", next: func(state int) int { return 1 - state }},
	} {
		t.Run(test.name, func(t *testing.T) {
			iterator, err := NewIterator(0, func(_ context.Context, state int) (Page[int, int], error) {
				return Page[int, int]{Items: []int{state}, NextState: test.next(state), HasMore: true}, nil
			})
			if err != nil {
				t.Fatal(err)
			}
			var gotErr error
			for _, rangeErr := range iterator.All(context.Background()) {
				if rangeErr != nil {
					gotErr = rangeErr
				}
			}
			if !errors.Is(gotErr, ErrPaginationDidNotAdvance) {
				t.Fatalf("error = %v; want ErrPaginationDidNotAdvance", gotErr)
			}
		})
	}
}

func TestIteratorPropagatesFetchAndContextErrors(t *testing.T) {
	fetchErr := errors.New("page failed")
	iterator, err := NewIterator("", func(_ context.Context, state string) (Page[int, string], error) {
		if state == "" {
			return Page[int, string]{Items: []int{1}, NextState: "next", HasMore: true}, nil
		}
		return Page[int, string]{}, fetchErr
	})
	if err != nil {
		t.Fatal(err)
	}
	var items []int
	var gotErr error
	for item, rangeErr := range iterator.All(context.Background()) {
		if rangeErr != nil {
			gotErr = rangeErr
			continue
		}
		items = append(items, item)
	}
	if !reflect.DeepEqual(items, []int{1}) || !errors.Is(gotErr, fetchErr) {
		t.Fatalf("items/error = %v/%v", items, gotErr)
	}

	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	fetches := 0
	cancelled, err := NewIterator(0, func(context.Context, int) (Page[int, int], error) {
		fetches++
		return Page[int, int]{}, nil
	})
	if err != nil {
		t.Fatal(err)
	}
	gotErr = nil
	for _, rangeErr := range cancelled.All(ctx) {
		gotErr = rangeErr
	}
	if !errors.Is(gotErr, context.Canceled) || fetches != 0 {
		t.Fatalf("cancelled iterator error/fetches = %v/%d", gotErr, fetches)
	}
}

func TestIteratorSupportsEarlyBreakAndRejectsNilFetcher(t *testing.T) {
	fetches := 0
	iterator, err := NewIterator(0, func(_ context.Context, state int) (Page[int, int], error) {
		fetches++
		return Page[int, int]{Items: []int{1, 2}, NextState: state + 1, HasMore: true}, nil
	})
	if err != nil {
		t.Fatal(err)
	}
	for range iterator.All(context.Background()) {
		break
	}
	if fetches != 1 {
		t.Fatalf("early break fetched %d pages; want 1", fetches)
	}
	if _, err := NewIterator[int, string]("", nil); !errors.Is(err, ErrNilPageFetcher) {
		t.Fatalf("nil fetcher error = %v", err)
	}
}
