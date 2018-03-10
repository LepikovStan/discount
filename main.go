package main

import (
	"github.com/LepikovStan/discount/crawler"
	"log"
	"github.com/LepikovStan/discount/parser"
	"fmt"
	"github.com/LepikovStan/discount/db"
	"github.com/LepikovStan/discount/data"
	"flag"
	"net/http"
	"encoding/json"
	"html/template"
	"time"
	"strings"
)

func stopCrawl(lastCrawlDate string) bool {
	shortForm := "2006-01-02"
	NOW := time.Now().Format(shortForm)

	lastCrawlDateTime := strings.Split(lastCrawlDate, " ")
	lastCrawlTime, _ := time.Parse(shortForm, lastCrawlDateTime[0])
	today, _ := time.Parse(shortForm, NOW)
	return !lastCrawlTime.Before(today)
}

func crawl() {
	database := db.New()
	database.SetCrawlDate()
	lastCrawlDate := database.GetLastCrawlDate()
	if stopCrawl(lastCrawlDate.Date) {
		return
	}

	cr := crawler.New()
	pr := parser.New()
	urlsList := data.Get("MoonyXLBoys")

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
			database.SetGood(result, mtype, lastCrawlDate.Id)
		}
	}
}

var templates *template.Template
func initTemplates() {
	templates = template.New("index") // Create a template.
	templates, _ = templates.ParseFiles("./server/static/index.html")
}

var database *db.Db
func server() {
	database = db.New()
	//initTemplates()

	http.HandleFunc("/", postHandler)
	log.Println("Listening...")
	http.ListenAndServe(":3001", nil)
}

type result struct {
	GoodsJSON string
}
func postHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	templates = template.New("index") // Create a template.
	templates, _ = templates.ParseFiles("./server/static/index.html")

	goods := database.GetGoods()
	goodsJSON, err := json.Marshal(goods)
	if err != nil {
		log.Fatal(err)
	}

	templates.Lookup("index.html").Execute(w, result{ string(goodsJSON) })
	end := time.Now()
	fmt.Println("/", end.Sub(start))
}

var serverUp bool
func main() {
	flag.BoolVar(&serverUp, "server", false, "")
	flag.Parse()

	if serverUp {
		server()
		return
	}
	crawl()
	fmt.Println("Done")
}
