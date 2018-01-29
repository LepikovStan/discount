package db

import (
	"database/sql"
	"log"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
)

type db struct {
	conn *sql.DB
}

func (db *db) init() {
	conn, err := sql.Open("mysql", "root:1@/discount")
	if err != nil {
		log.Fatal(err)
	}
	db.conn = conn
}

func (db *db) getMarketId(marketName string) int {
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

func (db *db) SetGood(good [3]string, marketName string) {
	name := good[0]
	price, err := strconv.Atoi(good[1])
	if err != nil {
		log.Fatal(err)
	}
	oldPrice, err := strconv.Atoi(good[2])
	if err != nil {
		log.Fatal(err)
	}
	marketId := db.getMarketId(marketName)
	db.conn.Query("insert into goods (name, price, oldPrice, marketId) values (?,?,?,?)", name, price, oldPrice, marketId)
}


type Good struct {
	id int
	marketId int
	name string
	price int
	oldPrice int
}
func (db *db) GetGood(id int) *Good {
	rows, err := db.conn.Query("select id, name, marketId, price, oldPrice from goods where id=?", id)
	defer rows.Close()

	if err != nil {
		log.Fatal(err)
	}

	var good Good
	for rows.Next() {
		if err := rows.Scan(&good.id, &good.name, &good.marketId, &good.price, &good.oldPrice); err != nil {
			log.Fatal(err)
		}
	}
	return &good
}

func New() *db {
	db := new(db)
	db.init()
	return db
}
