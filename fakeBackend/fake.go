package fakebackend

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Product struct {
	Id          int     `json:"id"`
	ProductName string  `json:"productName"`
	CategoryId  int     `json:"categoryId"`
	UnitPrice   float64 `json:"unitPrice"`
}

type Category struct {
	Id           int    `json:"id"`
	CategoryName string `json:"categoryName"`
}

func GetAllProducts() {

	//Get request sent
	response, err := http.Get("http://localhost:3000/products")

	//if not nill there is an error because err is nil at the request
	if err != nil {
		fmt.Println(err)
	}

	//No matter what, close th request's response body
	defer response.Body.Close()
	//I am not sure of this part
	body, _ := io.ReadAll(response.Body)

	//Create products array
	var products []Product
	//Unmarshall takes the response body and assigns it to products array
	json.Unmarshal(body, &products)
}
