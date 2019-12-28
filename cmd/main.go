package main

import (
	"fmt"
	"log"

	"github.com/tgfjt/aozora"
)

func main() {
	// 作家別作品一覧(拡充版)
	listOfWorks, err := aozora.ListOfWorks()
	if err != nil {
		log.Fatal("error:", err)
	}

	fmt.Println(listOfWorks[138].OriginBook()[0].ParentBook.Title)
}
