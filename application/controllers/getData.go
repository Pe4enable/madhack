package controllers

import (
	"net/http"
	"encoding/json"
	. "../models"
	"../utils"
	"github.com/gorilla/mux"
	"fmt"
)

func GetData(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	employeeId := vars["employeeId"]
	fmt.Println(employeeId)
	data := make([]DataEmployee, 0)
	decoder := json.NewDecoder(r.Body)
	_ = decoder.Decode(&data)

	res, statusCode, err := utils.SendRequest("wildberries/hs/accountant/goods_from_doc/", "GET", nil)
	b, _ := json.Marshal(res)
	utils.SetHeadersAndWrite(w, b, statusCode, err)
}
