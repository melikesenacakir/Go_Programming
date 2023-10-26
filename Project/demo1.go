package project

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Product struct {
	Id          int     `json:"id"`
	ProductName string  `json:"productName"`
	CategoryId  int     `json:"categoryId"`
	Price       float64 `json:"price"`
}
type Category struct {
	Id           int    `json:"id"`
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
	jsonProduct, _ := json.Marshal(Product{Id: 5, ProductName: "Telefon3", CategoryId: 1, Price: 3444.9})
	req, err := http.Post("http://localhost:3000/products",
		"application/json;charset=utf-8", bytes.NewBuffer(jsonProduct))
	if err != nil {
		fmt.Println(err)
	}
	defer req.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(req.Body)
	var productreq Product
	json.Unmarshal(bodyBytes, &productreq)
	return productreq, nil
}
