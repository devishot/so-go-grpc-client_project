package _examples

import (
	"log"
)

func ExampleClientRepository_CreateAndDelete() {
	testRows := getClientTestRows()

	testRows.Get()
	defer testRows.Release()

	//Output:
}

func ExampleClientRepository_Get() {
	testRows := getClientTestRows()
	repo := testRows.repo

	cl := testRows.Get()
	defer testRows.Release()

	_, err := repo.Get(cl.ID)
	if err != nil {
		log.Fatal(err)
	}

	//Output:
}
