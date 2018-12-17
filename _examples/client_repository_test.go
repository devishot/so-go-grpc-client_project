package _examples

import (
	"fmt"
)

func ExampleClientRepository_CreateAndDelete() {
	t := getClientTestRows()

	t.Get()
	defer t.Release()

	//Output:
}

func ExampleClientRepository_Get() {
	t := getClientTestRows()
	repo := t.repo

	cl := t.Get()
	defer t.Release()

	_, err := repo.Get(cl.ID)
	if err != nil {
		fmt.Print(err)
	}

	//Output:
}
