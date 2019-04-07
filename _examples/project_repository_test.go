package _examples

import (
	"fmt"

	"github.com/google/go-cmp/cmp"

	"github.com/devishot/so-go-grpc-client_project/domain"
)

func ExampleProjectRepository_CreateAndDelete() {
	t := getProjectTestRows()

	t.Get()
	defer t.Release()

	// Output:
}

func ExampleProjectRepository_Get() {
	t := getProjectTestRows()
	repo := t.repo
	p := t.Get()
	defer t.Release()

	p2, err := repo.Get(p.ID)
	if err != nil {
		fmt.Printf("error=%v", err)
		return
	}

	if !cmp.Equal(p, p2) {
		fmt.Printf("expected: %v, got: %v", p, p2)
	}

	//Output:
}

func ExampleProjectRepository_GetByClient() {
	t := getProjectTestRows()
	repo := t.repo
	p := t.Get()
	defer t.Release()

	arr, err := repo.GetByClient(p.ClientID)
	if err != nil {
		fmt.Printf("error=%v", err)
		return
	}

	expArr := []domain.ProjectEntity{p}

	for i, v := range expArr {
		if i >= len(arr) || arr[i].IsZero() || !cmp.Equal(v, arr[i]) {
			fmt.Printf("expected: %v, got: %v", expArr, arr)
			return
		}
	}

	//Output:
}

func ExampleProjectRepository_GetLastByClient() {
	t := getProjectTestRows()
	repo := t.repo
	p := t.GetLast()
	defer t.Release()

	p2, err := repo.GetLastByClient(p.ClientID)
	if err != nil {
		fmt.Printf("error=%v", err)
		return
	}

	if !cmp.Equal(p, p2) {
		fmt.Printf("expected: %v, got: %v", p, p2)
	}

	//Output:
}

func ExampleProjectRepository_PaginateForwardByClientByTimestamp() {
	t := getProjectTestRows()
	repo := t.repo
	p := t.Get()
	defer t.Release()

	first, after := 10, p.Timestamp

	arr, err := repo.PaginateForwardByClientByTimestamp(p.ClientID, first, after)
	if err != nil {
		fmt.Printf("error=%v", err)
		return
	}

	leftItems := t.items[1:]

	for i, p := range leftItems {
		if i >= len(arr) || arr[i].IsZero() || !cmp.Equal(p, arr[i]) {
			fmt.Printf("expected: %v, got: %v", leftItems, arr)
			return
		}
	}

	// Output:
}

func ExampleProjectRepository_PaginateBackwardByClientByTimestamp() {
	t := getProjectTestRows()
	repo := t.repo
	p := t.GetLast()
	defer t.Release()

	last, before := 10, p.Timestamp

	arr, err := repo.PaginateBackwardByClientByTimestamp(p.ClientID, last, before)
	if err != nil {
		fmt.Printf("error=%v", err)
		return
	}

	leftItems := t.items[:len(t.items)-1]

	for i, p := range leftItems {
		if i >= len(arr) || arr[i].IsZero() || !cmp.Equal(p, arr[i]) {
			fmt.Printf("expected: %v, got: %v", leftItems, arr)
			return
		}
	}

	// Output:
}
