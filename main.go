package main

import (
	"github.com/LepikovStan/discount/crawler"
	"log"
	"github.com/LepikovStan/discount/parser"
	"fmt"
	"github.com/LepikovStan/discount/db"
	"github.com/LepikovStan/discount/data"
)

func main() {
	cr := crawler.New()
	pr := parser.New()
	database := db.New()
	urlsList := data.Get()

	for mtype, urls := range urlsList {
		for _, url := range urls {
			fmt.Println("Crawl url: ", url)
			body, err := cr.Crawl(url)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("Parse url: ", url)
			result, err := pr.Parse(body, mtype)

			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(result)
			database.SetGood(result, mtype)
		}
	}
	fmt.Println(database.GetGood(1))
}
