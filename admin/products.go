package admin

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Nyxoy/restAPI/db"
	"github.com/Nyxoy/restAPI/models"
	"github.com/Nyxoy/restAPI/utils"
	"github.com/spf13/viper"
)

// Add a product

func AddProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var product = &models.Product{}
	if err := json.NewDecoder(r.Body).Decode(product); err != nil {
		utils.WriteError(w, http.StatusBadRequest, "Invalid request body")
		log.Printf("Invalid request body")
		return
	}
	fmt.Println(product)
	//Enter into the database
	resp, err2 := db.CreateRestyClient().R().SetBody(product).Post(viper.GetString("DB_BASE_URL") + "/rest/v1/products")
	if err2 != nil {
		utils.WriteError(w, http.StatusInternalServerError, "Error occured entering the databse")
		log.Println("Error occured entering the databse")
		return
	}
	fmt.Println(resp.StatusCode())
	fmt.Println(string(resp.Body()))

	if resp.StatusCode() == 201 {
		utils.WriteError(w, http.StatusOK, "Product Added Successfully")
	}
}

func RemoveProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Url me hoga ID
	var id = r.URL.Query().Get("id")
	fmt.Printf("%T", id)
	//Enter into the database
	if id == "" {
		utils.WriteError(w, http.StatusBadRequest, "Product ID is required")
		log.Println("Product ID is missing")
		return
	}
	actId, _ := strconv.Atoi(id)
	// fmt.Println(actId)

	resp, err := db.CreateRestyClient().R().Get(viper.GetString("DB_BASE_URL") + "/rest/v1/products?id=eq." + fmt.Sprintf("%d", actId))
	if err != nil || resp.StatusCode() != 201 {
		utils.WriteError(w, http.StatusNotFound, "Product not found")
		log.Printf("Error finding product with ID %s", id)
		return
	}

	resp1, err2 := db.CreateRestyClient().R().Delete(viper.GetString("DB_BASE_URL") + "/rest/v1/products?id=eq." + fmt.Sprintf("%d", actId))
	if err2 != nil {
		utils.WriteError(w, http.StatusInternalServerError, "Error occured entering the databse")
		log.Println("Error occured entering the databse")
		return
	}
	if resp1.StatusCode() == 204 {
		utils.WriteError(w, http.StatusOK, "Deleted the product ")
	}
}

func UpdatePrice(w http.ResponseWriter, r *http.Request) {
	type Price struct {
		Id           int `json:"id" validate:"required"`
		UpdatedPrice int `json:"price" validate:"required,min=1"`
	}
	var price Price
	if err := json.NewDecoder(r.Body).Decode(&price); err != nil {
		utils.WriteError(w, http.StatusBadRequest, "Invalid body")
		return
	}
	// Check if it exists withing the database
	res, er := db.CreateRestyClient().R().SetQueryParam("select", "count").Get(viper.GetString("DB_BASE_URL") + "/rest/v1/products?id=eq." + fmt.Sprintf("%d", price.Id))
	if er != nil {
		utils.WriteError(w, http.StatusInternalServerError, "Database Query Error")
		return
	}
	if res.StatusCode() == 200 {
		var prod []struct {
			Count int `json:"count"`
		}
		if err := json.Unmarshal(res.Body(), &prod); err != nil {
			utils.WriteError(w, http.StatusBadRequest, "Error parsing the data")
			return
		}
		if len(prod) == 0 || prod[0].Count == 0 {
			utils.WriteError(w, http.StatusNotFound, "Product does not exist")
			return
		}
	}

	resp, err := db.CreateRestyClient().R().SetBody(price).Patch(viper.GetString("DB_BASE_URL") + "/rest/v1/products?id=eq." + fmt.Sprintf("%d", price.Id))
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "Database query error")
		return
	}

	if resp.StatusCode() == 204 {
		utils.WriteError(w, http.StatusOK, "Price Updated Successfully")
	}
}
