package db

import (
	"database/sql"
	"log"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
)

type Db struct {
	conn *sql.DB
}

func (db *Db) init() {
	conn, err := sql.Open("mysql", "root:1@/discount")
	if err != nil {
		log.Fatal(err)
	}
	db.conn = conn
}

func (db *Db) getMarketId(marketName string) int {
	rows, err := db.conn.Query("select id from markets where name=?", marketName)
	defer rows.Close()

	if err != nil {
		log.Fatal(err)
	}

	var marketId int
	for rows.Next() {
		if err := rows.Scan(&marketId); err != nil {
			log.Fatal(err)
		}
	}
	return marketId
}

func (db *Db) SetGood(good [3]string, marketName string, lastCrawlDateId int) {
	name := good[0]
	price, err := strconv.Atoi(good[1])
	if err != nil {
		price = 0
	}
	oldPrice, err := strconv.Atoi(good[2])
	if err != nil {
		oldPrice = 0
	}
	marketId := db.getMarketId(marketName)
	db.conn.Query("insert into goods (name, price, oldPrice, marketId, crawlId) values (?,?,?,?,?)", name, price, oldPrice, marketId, lastCrawlDateId)
}

func (db *Db) SetCrawlDate() {
	db.conn.Query("insert into crawl_dates (date) values (NOW())")
}

type CrawlDate struct {
	Id int
	Date string
}
func (db *Db) GetLastCrawlDate() *CrawlDate {
	rows, err := db.conn.Query("select * from crawl_dates order by date limit 1")
	defer rows.Close()

	if err != nil {
		log.Fatal(err)
	}

	var date CrawlDate
	for rows.Next() {
		if err := rows.Scan(&date.Id, &date.Date); err != nil {
			log.Fatal(err)
		}
	}
	return &date
}


type Good struct {
	Id int
	MarketId int `json:"-"`
	Name string
	Price int
	OldPrice int
	MarketName string
	CrawlId int
}
func (db *Db) GetGood(id int) *Good {
	rows, err := db.conn.Query("select id, name, marketId, price, oldPrice from goods where id=?", id)
	defer rows.Close()

	if err != nil {
		log.Fatal(err)
	}

	var good Good
	for rows.Next() {
		if err := rows.Scan(&good.Id, &good.Name, &good.MarketId, &good.Price, &good.OldPrice); err != nil {
			log.Fatal(err)
		}
	}
	return &good
}

func (db *Db) GetGoods() []*Good {
	res := []*Good{}
	rows, err := db.conn.Query("select goods.id, goods.name, goods.price, goods.oldPrice, markets.name, crawl_dates.Id from goods join markets on goods.marketId = markets.id join crawl_dates on goods.crawlId = crawl_dates.id")
	defer rows.Close()

	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var good Good
		if err := rows.Scan(&good.Id, &good.Name, &good.Price, &good.OldPrice, &good.MarketName, &good.CrawlId); err != nil {
			log.Fatal(err)
		}
		res = append(res, &good)
	}
	return res
}

func New() *Db {
	db := new(Db)
	db.init()
	return db
}
