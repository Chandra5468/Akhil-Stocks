package models

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/Chandra5468/Akhil-Stocks/types"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func createConnection() *sql.DB {
	env := os.Getenv("APP_ENV")

	if env == "" { // if env is not specified then loading local env
		env = "local"
	}

	envFile := fmt.Sprintf("envs/.env.%s", env)
	// load env file
	err := godotenv.Load(envFile)

	if err != nil {
		log.Fatal("error loading godotenv .env file", err)
	}

	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()

	if err != nil {
		log.Fatal(err)
	}

	log.Println("successfully connected to postgresql")

	return db
}

func InsertStock(stock *types.Stock) int16 {
	db := createConnection()
	defer db.Close()

	statement := `insert into stock(name,price,number) values ($1,$2,$3) returning stockid`
	var id int16
	// db.Exec()
	err := db.QueryRow(statement, stock.Name, stock.Price, stock.Company).Scan(&id)

	if err != nil {
		log.Fatalf("Unable to execute the query %s", err.Error())
	}

	log.Println("Inserted a single record %d", id)
	return id
}

func GetStock(id int16) (*types.Stock, error) {
	db := createConnection()
	defer db.Close()

	var stock *types.Stock

	statement := `select * from stock where stockid =$1`

	row := db.QueryRow(statement, id)

	err := row.Scan(&stock.StockId, &stock.Name, &stock.Price, &stock.Company)

	switch err {
	case sql.ErrNoRows:
		log.Println("No rows are found")
		return nil, fmt.Errorf("no rows found")
	case nil:
		return stock, nil
	default:
		return nil, fmt.Errorf("unable to scan the rows")
	}

}

func GetAllStocks() (*[]types.Stock, error) {
	db := createConnection()
	defer db.Close()

	var stocks []types.Stock

	statement := `select * from stock`

	rows, err := db.Query(statement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var stock types.Stock
		err := rows.Scan(&stock.StockId, &stock.Name, &stock.Price, &stock.Company)

		if err != nil {
			log.Fatalf("unable to scan the row %s", err.Error())
		}

		stocks = append(stocks, stock)
	}

	return &stocks, nil
}

func UpdateStock(id int16, stock *types.Stock) int16 {
	db := createConnection()
	defer db.Close()
	// update table_name set field1='abc' where field2='xyz';
	statement := `update stock set name = $2, price = $3, company = $4 where stockid = $1`

	res, err := db.Exec(statement, id, stock.Name, stock.Price, stock.Company)

	if err != nil {
		log.Println("Unable to execute query, some issue ", err)
		return 0
	}

	rA, err := res.RowsAffected()

	if err != nil {
		log.Println("Error while checking rows affected during update command", err)
		return 0
	}

	return int16(rA)
}

func DeleteStock(id int16) int16 {
	db := createConnection()
	defer db.Close()

	statement := `delete from stock where stockid=$1`

	res, err := db.Exec(statement, id)

	if err != nil {
		log.Println("unable to execute the query ", err)
		return 0
	}
	rA, err := res.RowsAffected()

	if err != nil {
		log.Println("unable to check rows affected ", err)
		return 0
	}

	return int16(rA)
}
