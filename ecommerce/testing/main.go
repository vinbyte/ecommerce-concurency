package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"
)

var wg sync.WaitGroup
var mutex sync.Mutex
var successCheckout []string

func main() {

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go puchase(i)
	}

	wg.Wait()

	fmt.Println("===========================")
	if len(successCheckout) > 0 {
		fmt.Println("success checkout : ", strings.Join(successCheckout[:], ","))
	}
}

func puchase(id int) {
	userID := strconv.Itoa(id)
	addToCart(userID)
	wg.Done()
}

func addToCart(userID string) {
	time.Sleep(time.Second * 1)
	type addCartResponse struct {
		Data struct {
			CartID int `json:"cart_id"`
		} `json:"data"`
	}
	var response addCartResponse

	data := url.Values{}
	data.Set("user_id", userID)
	data.Set("product_code", "P1")
	data.Set("qty", "1")

	client := &http.Client{}
	r, _ := http.NewRequest(http.MethodPost, "http://localhost:5050/v1/cart/add", strings.NewReader(data.Encode()))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	resp, _ := client.Do(r)
	fmt.Println("user "+userID+" add to cart", resp.Status)
	if resp.StatusCode == 200 {
		defer resp.Body.Close()
		json.NewDecoder(resp.Body).Decode(&response)
		checkout(userID, response.Data.CartID)
	}
}

func checkout(userID string, cartID int) {
	data := url.Values{}
	data.Set("cart_id", strconv.Itoa(cartID))

	client := &http.Client{}
	r, _ := http.NewRequest(http.MethodPost, "http://localhost:5050/v1/checkout", strings.NewReader(data.Encode()))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	resp, _ := client.Do(r)
	fmt.Println("user "+userID+" checkout", resp.Status)
	if resp.StatusCode == 200 {
		successCheckout = append(successCheckout, "user "+userID)
	}
}
