package middleware

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Chandra5468/Akhil-Stocks/models"
	"github.com/Chandra5468/Akhil-Stocks/types"
	"github.com/gorilla/mux"
)

func CreateStock(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	stock := &types.Stock{}

	err := json.NewDecoder(r.Body).Decode(stock)

	if err != nil {
		log.Fatalf("unable to decode req.body in create stock handler, %s", err.Error())
		return
	}

	insertId := models.InsertStock(stock)

	res := &types.Response{
		ID:      insertId,
		Message: "stock created successfully",
	}

	json.NewEncoder(w).Encode(res)
}

func GetStock(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatal("unable to convert string id to int id", err)
		return
	}

	stock, err := models.GetStock(int16(id))

	if err != nil {
		log.Fatalf("unable to get stock, %s", err.Error())
	}

	json.NewEncoder(w).Encode(stock)
}

func GetStocks(w http.ResponseWriter, r *http.Request) {
	stocks, err := models.GetAllStocks()

	if err != nil {
		log.Fatal("unable to get all stocks", err)
	}

	json.NewEncoder(w).Encode(stocks)
}

func UpdateStock(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("unable to convert the string to id %s", err.Error())
		return
	}

	stock := &types.Stock{}
	err = json.NewDecoder(r.Body).Decode(stock)

	if err != nil {
		log.Fatalf("unable to decode the request %s", err.Error())
		return
	}

	updatedRows := models.UpdateStock(int16(id), stock)

	msg := fmt.Sprintf("stock updated successfully. Total rows/records affected %v", updatedRows)

	res := &types.Response{
		ID:      int16(id),
		Message: msg,
	}

	json.NewEncoder(w).Encode(res)
}

func DeleteStock(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("unable to convert string to int %v", err)
		return
	}

	deletedRows := models.DeleteStock(int16(id))

	msg := fmt.Sprintf("stock deleted successfully, total rows affected/deleted %v", deletedRows)

	res := &types.Response{
		ID:      int16(id),
		Message: msg,
	}

	json.NewEncoder(w).Encode(res)
}
