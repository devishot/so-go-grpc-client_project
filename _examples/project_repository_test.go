package _examples

import (
	"fmt"
	"log"

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
		log.Fatal(err)
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
		log.Fatal(err)
	}

	expArr := []domain.ProjectEntity{p}

	for i, v := range expArr {
		if i >= len(arr) || arr[i].IsZero() || !cmp.Equal(v, arr[i]) {
			fmt.Printf("expected: %v, got: %v", expArr, arr)
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
		log.Fatal(err)
	}

	if !cmp.Equal(p, p2) {
		fmt.Printf("expected: %v, got: %v", p, p2)
	}

	//Output:
}
