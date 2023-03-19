package project

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Product struct {
	Id          int     `json: "id"`
	ProductName string  `json: "productName"`
	CategoryId  string  `json: "categoryId"`
	UnitPrice   float64 `json:"unitPrice"`
}

type Category struct {
	Id           int    `json: "id"`
	CategoryName string `json:"categoryName"`
}

func GetAllProducts() ([]Product, error) {

	response, err := http.Get("http://localhost:3000/products")
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(response.Body)

	var products []Product
	json.Unmarshal(bodyBytes, &products)
	return products, nil

}

func AddProduct() (Product, error) {
	product1 := Product{Id: 4, ProductName: "keyboard", CategoryId: "1", UnitPrice: 45.99}
	jsonProduct, err := json.Marshal(product1)
	if err != nil {
		return Product{}, err
	}

	response, _ := http.Post("http://localhost:3000/products", "application/json;charset =utf-8", bytes.NewBuffer(jsonProduct))
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(response.Body)

	var productResponse Product
	json.Unmarshal(bodyBytes, &productResponse)
	fmt.Println("kaydedildi", productResponse)
	return product1, nil

}
