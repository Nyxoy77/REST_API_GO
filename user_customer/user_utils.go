package user_customer

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Nyxoy/restAPI/caching"
	"github.com/Nyxoy/restAPI/db"
	"github.com/Nyxoy/restAPI/models"
	"github.com/Nyxoy/restAPI/utils"
	"github.com/go-playground/validator"
	"github.com/spf13/viper"
)

var validate = utils.NewValidator()

func AddItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var item ItemToBeAdded
	var checkItem []models.Product
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		utils.WriteError(w, http.StatusBadRequest, "Error decoding")
		return
	}

	if err := validate.Struct(item); err != nil {
		errors := err.(validator.ValidationErrors)
		for _, value := range errors {
			log.Printf("Validation failed for field %s , conditon %s /n", value.Field(), value.Tag())
		}
		utils.WriteError(w, http.StatusBadRequest, "Invalid Input Fields")
		return
	}

	response, err1 := db.CreateRestyClient().R().SetQueryParam("id", fmt.Sprintf("eq.%d", item.Product_Id)).
		Get(viper.GetString("DB_BASE_URL") + "/rest/v1/products")
	if err1 != nil {
		utils.WriteError(w, http.StatusInternalServerError, "Internal Server errror")
		return
	}
	fmt.Println(string(response.Body()))
	fmt.Println(response.StatusCode())
	if response.StatusCode() == 200 {
		if err := json.Unmarshal(response.Body(), &checkItem); err != nil {
			utils.WriteError(w, http.StatusBadRequest, "Invalid data body")
			return
		}
		if len(checkItem) > 0 {
			if item.Quantity > checkItem[0].StockQuantity {
				utils.WriteError(w, http.StatusBadRequest, "Not enough stock")
				return
			}
		}

	}
	var cartItem = CartItem{

		User_Id:        item.User_Id,
		Product_Id:     item.Product_Id,
		Quantity:       item.Quantity,
		Price_Per_Unit: checkItem[0].Price,
	}
	resp, err := db.CreateRestyClient().R().
		SetBody(cartItem).
		Post(viper.GetString("DB_BASE_URL") + "/rest/v1/cart")
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "Error accessing the database")
		fmt.Println(err)
		return
	}
	fmt.Println(string(resp.Body()))
	if resp.StatusCode() == 201 {
		utils.WriteError(w, http.StatusOK, "Item added to the cart")
	}

	// Now update the products table

	res, err2 := db.CreateRestyClient().R().SetBody(map[string]interface{}{
		"stock_quantity": checkItem[0].StockQuantity - item.Quantity,
	}).Patch(viper.GetString("DB_BASE_URL") + fmt.Sprintf("/rest/v1/products?id=eq.%d", item.Product_Id))
	if err2 != nil {
		utils.WriteError(w, http.StatusInternalServerError, "Internal Server error")
		return
	}
	if res.StatusCode() == 201 {
		utils.WriteError(w, http.StatusOK, "Product quantity updated ")
	}
}

func RemoveItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var item DeleteItem
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		utils.WriteError(w, http.StatusBadRequest, "Error decoding")
		return
	}

	if err := validate.Struct(item); err != nil {
		errors := err.(validator.ValidationErrors)
		for _, value := range errors {
			log.Printf("Validation failed for field %s , conditon %s /n", value.Field(), value.Tag())
		}
		utils.WriteError(w, http.StatusBadRequest, "Invalid Input Fields")
		return
	}
	var existItem []CartItem
	resp, err := db.CreateRestyClient().R().
		SetQueryParam("user_id", fmt.Sprintf("eq.%d", item.User_Id)).
		SetQueryParam("product_id", fmt.Sprintf("eq.%d", item.Product_Id)).
		Get(viper.GetString("DB_BASE_URL") + "/rest/v1/cart")

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "Error accessing the database")
		fmt.Println(err)
		return
	}
	if resp.StatusCode() == 200 {
		if err := json.Unmarshal(resp.Body(), &existItem); err != nil {
			utils.WriteError(w, http.StatusInternalServerError, "Error parsing the data")
			fmt.Println(err)
			return
		}
		if len(existItem) == 0 {
			utils.WriteError(w, http.StatusNotFound, "No item exists in the cart")
			return
		}
	}
	// fmt.Println(len(existItem))
	count := existItem[0].Quantity - item.Quantity
	fmt.Println(existItem)
	if count == 0 {
		// Remove the product from the cart when the count is 0
		resp, err := db.CreateRestyClient().R().
			SetQueryParam("id", fmt.Sprintf("eq.%d", existItem[0].Id)).
			Delete(viper.GetString("DB_BASE_URL") + "/rest/v1/cart")
		if err != nil {
			utils.WriteError(w, http.StatusInternalServerError, "Error removing the product from the database")
			fmt.Println(err)
			return
		}
		if resp.StatusCode() == 204 {
			utils.WriteError(w, http.StatusOK, "Product removed successfully")
		} else {
			utils.WriteError(w, http.StatusInternalServerError, "Failed to remove the product")
		}
	} else {
		// Update the quantity of the product in the cart when the count is not 0
		if count < 1 {
			utils.WriteError(w, http.StatusBadRequest, "Quantity must be greater than 0")
			return
		}

		resp, err := db.CreateRestyClient().R().
			SetQueryParam("id", fmt.Sprintf("eq.%d", existItem[0].Id)).
			SetBody(map[string]interface{}{"quantity": count}).
			Patch(viper.GetString("DB_BASE_URL") + "/rest/v1/cart")

		if err != nil {
			utils.WriteError(w, http.StatusInternalServerError, "Error updating the product in the database")
			fmt.Println(err)
			return
		}
		fmt.Println(resp.StatusCode())

		if resp.StatusCode() == 204 {
			utils.WriteError(w, http.StatusOK, "Product quantity updated successfully")
		} else {
			utils.WriteError(w, http.StatusInternalServerError, "Failed to update the product quantity")
		}
	}

}

func GetPrice(Id int) float64 {
	// var price float64
	var price []struct {
		Price float64 `json:"price"`
	}

	if err := caching.GetCache(fmt.Sprintf("product:%d", Id), &price); err == nil {
		log.Println("The cache hit ")
		// fmt.Println(err)
		return price[0].Price
	}
	resp, err := db.CreateRestyClient().R().SetQueryParam("id", fmt.Sprintf("eq.%d", Id)).
		SetQueryParam("select", "price").
		Get(viper.GetString("DB_BASE_URL") + "/rest/v1/products")
	if err != nil {
		log.Println("Error fetching the price from the database")
		fmt.Println(err)
		return -1
	}

	if resp.StatusCode() == 200 {
		json.Unmarshal(resp.Body(), &price)
		caching.SetCache(fmt.Sprintf("product:%d", Id), price, time.Duration(time.Hour))
		return price[0].Price

	}
	return -1
	// price = strconv.At(string(resp.Body()))
}
