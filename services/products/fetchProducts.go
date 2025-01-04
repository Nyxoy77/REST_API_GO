package productservices

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
	"github.com/spf13/viper"
)

func FetchProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var prod []models.Product

	// Try fetching from cache
	if err := caching.GetCache("products", &prod); err == nil && len(prod) > 0 {
		json.NewEncoder(w).Encode(prod)
		log.Println("The cache was hit")
		return
	} else {
		log.Println("Cache miss or empty cache, querying the database.")
	}

	// Fetch from Supabase
	resp, err := db.CreateRestyClient().R().
		SetQueryParam("select", "*").
		Get(viper.GetString("DB_BASE_URL") + "/rest/v1/products")

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "Failed to fetch product data from Supabase.")
		log.Printf("Error fetching products: %v\n", err)
		return
	}

	// Unmarshal response body
	if er := json.Unmarshal(resp.Body(), &prod); er != nil {
		utils.WriteError(w, http.StatusBadRequest, "Failed to parse product data from database response.")
		log.Printf("Error unmarshaling products: %v\n", er)
		return
	}

	// Set data in cache
	if err := caching.SetCache("products", prod, time.Hour); err != nil {
		log.Println("Error occurred while setting data in cache.")
		fmt.Println(err)
	}

	// Send response
	if err := json.NewEncoder(w).Encode(prod); err != nil {
		utils.WriteError(w, http.StatusBadRequest, "Failed to encode product data to JSON.")
		log.Printf("Error encoding products: %v\n", err)
	}
}
