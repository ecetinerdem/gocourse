package fakebackend

import (
	"bytes"
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

func GetAllProducts() ([]Product, error) {

	//Get request sent
	response, errOne := http.Get("http://localhost:3000/products")

	//if not nill there is an error because err is nil at the request
	if errOne != nil {
		fmt.Println(errOne)
	}

	//No matter what, close th request's response body
	defer response.Body.Close()
	//I am not sure of this part
	body, errTwo := io.ReadAll(response.Body)

	if errTwo != nil {
		fmt.Println(errTwo)
	}

	//Create products array
	var products []Product
	//Unmarshall takes the response body and assigns it to products array
	json.Unmarshal(body, &products)

	return products, nil
}

func AddProduct() (Product, error) {

	product := Product{Id: 4, ProductName: "Workstation", CategoryId: 4, UnitPrice: 6000.00}

	jsonProduct, errOne := json.Marshal(product)

	if errOne != nil {
		fmt.Println(errOne)
	}

	response, errtwo := http.Post("http://localhost:3000/products", "application/json:charset=utg-8", bytes.NewBuffer(jsonProduct))

	if errtwo != nil {
		fmt.Println(errtwo)
	}

	defer response.Body.Close()

	body, errThree := io.ReadAll(response.Body)

	if errThree != nil {
		fmt.Println(errThree)
	}

	var productResponse Product

	json.Unmarshal(body, &productResponse)

	return productResponse, nil
}
