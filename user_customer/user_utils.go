package user_customer

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Nyxoy/restAPI/db"
	"github.com/Nyxoy/restAPI/utils"
	"github.com/go-playground/validator"
	"github.com/spf13/viper"
)

var validate = utils.NewValidator()

func AddItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var item Item

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

	resp, err := db.CreateRestyClient().R().
		SetBody(item).
		Post(viper.GetString("DB_BASE_URL") + "/rest/v1/cart")
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "Error accessing the database")
		fmt.Println(err)
		return
	}

	if resp.StatusCode() == 201 {
		utils.WriteError(w, http.StatusOK, "Item added to the cart")
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
	var existItem []Item
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
