package api

import (
	"encoding/json"
	"hw/internal/models"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)


func GetData(n int) models.APIResponse {

	url := "https://fakerapi.it/api/v2/books?_quantity=" + strconv.Itoa(n)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Failed to fetch data from API: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("API returned non-OK status: %s", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

	var response models.APIResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	return response
}