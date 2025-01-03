package user_customer

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Nyxoy/restAPI/db"
	"github.com/Nyxoy/restAPI/utils"
	"github.com/go-playground/validator"
	"github.com/spf13/viper"
)

func AddToOrders(w http.ResponseWriter, r *http.Request) {
	var comb CombinedOrder
	if err := json.NewDecoder(r.Body).Decode(&comb); err != nil {
		utils.WriteError(w, http.StatusBadRequest, "Invalid Body ")
		return
	}

	if err := validate.Struct(comb); err != nil {
		errors := err.(validator.ValidationErrors)
		for _, value := range errors {
			log.Printf("Validation failed for field %s , condition %s\n", value.Field(), value.Tag())
		}
		utils.WriteError(w, http.StatusBadRequest, "Validation failed")
		return
	}

	// Create order struct
	var order = Orders{
		User_Id:    comb.User_Id,
		Order_Date: time.Now(),
		Status:     "shipped",
	}

	// Parse the frontend order items
	var frontend []FrontEndOrder = comb.Items
	var OrderItem []OrderItems

	for _, val := range frontend {
		price := GetPrice(val.Product_Id)
		if price == -1 {
			utils.WriteError(w, http.StatusInternalServerError, "Error fetching the prices")
			return
		}
		obj1 := OrderItems{
			Product_Id:    val.Product_Id,
			Quantity:      val.Quantity,
			Price_At_Time: price,
		}
		OrderItem = append(OrderItem, obj1)
	}

	// Add the order to the database
	resp1, err1 := db.CreateRestyClient().R().
		SetBody(order).
		Post(viper.GetString("DB_BASE_URL") + "/rest/v1/orders")

	if err1 != nil {
		utils.WriteError(w, http.StatusInternalServerError, "Error storing the order overall information")
		return
	}
	// fmt.Println("Response from Orders Table:", string(resp1.Body()))

	if resp1.StatusCode() != 201 {
		utils.WriteError(w, http.StatusInternalServerError, "Failed to create the order")
		return
	}

	resp3, err3 := db.CreateRestyClient().R().
		SetQueryParam("select", "id").
		Get(viper.GetString("DB_BASE_URL") + "/rest/v1/orders")
	// fmt.Println("Response from Orders Table:", string(resp3.Body()))

	if err3 != nil {
		utils.WriteError(w, http.StatusInternalServerError, "Error retrieving the order overall information")
		return
	}
	// Extract the order ID from the response body
	var orderResponse []struct {
		Id int `json:"id"`
	}
	err := json.Unmarshal(resp3.Body(), &orderResponse)
	if err != nil || len(orderResponse) == 0 {
		utils.WriteError(w, http.StatusInternalServerError, "Error parsing the order response")
		return
	}

	orderId := orderResponse[0].Id
	// fmt.Println("Inserted Order ID:", orderId)

	// Insert the order items
	for _, val := range OrderItem {
		val.Order_Id = orderId
		resp2, err2 := db.CreateRestyClient().R().SetBody(val).
			Post(viper.GetString("DB_BASE_URL") + "/rest/v1/order_items")

		if err2 != nil {
			utils.WriteError(w, http.StatusInternalServerError, "Error storing the order items detailed information")
			return
		}

		if resp2.StatusCode() != 201 {
			utils.WriteError(w, http.StatusInternalServerError, "Failed to insert order items")
			return
		}

		fmt.Println("Inserted Order Item:", string(resp2.Body()))
	}

	utils.WriteError(w, http.StatusOK, "Order and items added successfully")
}
