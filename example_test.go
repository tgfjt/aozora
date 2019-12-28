package aozora_test

import (
	"fmt"
	"log"

	"github.com/tgfjt/aozora"
)

// Thanks: https://mattn.kaoriya.net/software/lang/go/20160328114637.htm

func ExampleListOfWorks() {
	list, err := aozora.ListOfWorks()

	if err != nil {
		log.Fatal(err)
	}
	if len(list) == 0 {
		log.Fatal("the list is empty.")
	}

	fmt.Println(list[88].OriginBook()[0].ParentBook.Title)

	// Output:
	// 筑摩全集類聚版芥川龍之介全集
}
