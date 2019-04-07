package _examples

import (
	"fmt"
	"time"
)

func ExampleClientRepository_CreateAndDelete() {
	t := getClientTestRows()

	items := t.Get(1)
	defer t.Delete(items)
	defer t.Release()

	//Output:
}

func ExampleClientRepository_Get() {
	t := getClientTestRows()
	repo := t.repo

	items := t.Get(1)
	cl := items[0]
	defer t.Delete(items)
	defer t.Release()

	_, err := repo.Get(cl.ID)
	if err != nil {
		fmt.Printf("error=%v", err)
		return
	}

	//Output:
}

func ExampleClientRepository_GetFirst() {
	t := getClientTestRows()
	repo := t.repo

	items := t.Get(2)
	expCl := items[0]
	defer t.Delete(items)
	defer t.Release()

	resCl, err := repo.GetFirst()
	if err != nil {
		fmt.Printf("error=%v", err)
		return
	}

	if expCl.ID != resCl.ID {
		fmt.Printf("expected: %v, got: %v", expCl, resCl)
	}

	//Output:
}

func ExampleClientRepository_GetLast() {
	t := getClientTestRows()
	repo := t.repo

	items := t.Get(2)
	expCl := items[1]
	defer t.Delete(items)
	defer t.Release()

	resCl, err := repo.GetLast()
	if err != nil {
		fmt.Printf("error=%v", err)
		return
	}

	if expCl.ID != resCl.ID {
		fmt.Printf("expected: %v, got: %v", expCl, resCl)
	}

	//Output:
}

func ExampleClientRepository_PaginateForwardByTimestamp() {
	t := getClientTestRows()
	repo := t.repo

	items := t.Get(2)
	afterItem := items[1]
	defer t.Release()
	defer t.Delete(items)

	time.Sleep(time.Second)

	items2 := t.Get(2)
	defer t.Delete(items2)

	resItems, err := repo.PaginateForwardByTimestamp(2, afterItem.Timestamp)
	if err != nil {
		fmt.Printf("error=%v", err)
		return
	}

	if len(resItems) != 2 || resItems[0].ID != items2[0].ID {
		fmt.Printf("expected: %v, got: %v", items2, resItems)
	}

	//Output:
}
