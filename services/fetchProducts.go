package services

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Nyxoy/restAPI/db"
	"github.com/Nyxoy/restAPI/models"
	"github.com/Nyxoy/restAPI/utils"
	"github.com/spf13/viper"
)

// var token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InNoaXZhbW1Ab3AuY29tIiwiZXhwIjoxNzM0NzYwMzc4LCJpc3N1ZWRfYXQiOjE3MzQ2NzM5NzgsInVzZXJfaWQiOjExfQ.7Rbn2reHwF5_Z7kuVdw-e5RkGArxDRFk59AP-QfL-S8"

func FetchProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("This is a protected route"))
	var prod []models.Product
	resp, err := db.CreateRestyClient().R().SetQueryParam("select", "*").Get(viper.GetString("DB_BASE_URL") + "/rest/v1/products")
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "An error occured while fetching the product data")
		fmt.Println("The error while fetching prodcuts ", err)
		return
	}
	fmt.Println(string(resp.Body()))
	if er := json.Unmarshal(resp.Body(), &prod); er != nil {
		utils.WriteError(w, http.StatusBadRequest, "An error occured while parsing the prod")
		return
	}
	for _, value := range prod {
		fmt.Println(value)
	}

}
