package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Product struct {
	ProductID string
	Category  string
	Market    *Marketplace
}
type Marketplace struct {
	Name    string
	Address string
	Website string
}

var products []Product

// Checks if there is any product
func (p *Product) IsEmpty() bool {
	return p.Category == ""
}

func serveRoot(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello I am a ROOT page</h1>"))
}

func getAllProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func getAProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, val := range products {
		if val.ProductID == params["id"] {
			json.NewEncoder(w).Encode(val)
			return
		} else {
			json.NewEncoder(w).Encode("No Product Found!")
			return
		}
	}
}

func addProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some Data!")
	}
	var product Product
	_ = json.NewDecoder(r.Body).Decode(&product)
	if product.IsEmpty() {
		json.NewEncoder(w).Encode("Please send some Data!")
		return
	}
	product.ProductID = strconv.Itoa(rand.Intn(100))
	// this returns the new product list
	products = append(products, product)
	json.NewEncoder(w).Encode(product)
}

func deleteProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, val := range products {
		if val.ProductID == params["id"] {
			products = append(products[:index], products[index+1:]...)
			json.NewEncoder(w).Encode(val)
			return
		} else {
			json.NewEncoder(w).Encode("No Product Found!")
			return
		}
	}
}

func updateCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, val := range products {
		if val.ProductID == params["id"] {
			products = append(products[:index], products[index+1:]...)
			var product Product
			_ = json.NewDecoder(r.Body).Decode(&product)
			product.ProductID = params["id"]
			products = append(products, product)
			json.NewEncoder(w).Encode(product)
			return
		} else {
			continue
		}
	}
}

func main() {
	fmt.Println("Let's build API")
	products = append(products, Product{
		ProductID: "1",
		Category:  "Kitchen",
		Market: &Marketplace{
			Name:    "Flipkart",
			Address: "RameshNagar",
			Website: "flipkart.com",
		},
	})
	products = append(products, Product{
		ProductID: "2",
		Category:  "Gardening",
		Market: &Marketplace{
			Name:    "Flipkart",
			Address: "RameshNagar",
			Website: "flipkart.com",
		},
	})
	products = append(products, Product{
		ProductID: "3",
		Category:  "Sleeping",
		Market: &Marketplace{
			Name:    "Flipkart",
			Address: "RameshNagar",
			Website: "flipkart.com",
		},
	})
	r := mux.NewRouter()
	r.HandleFunc("/", serveRoot).Methods("GET")
	r.HandleFunc("/getall", getAllProducts).Methods("GET")
	r.HandleFunc("/getproduct/{id}", getAProduct).Methods("GET")
	r.HandleFunc("/add", addProduct).Methods("POST")
	r.HandleFunc("/update/{id}", updateCourse).Methods("PUT")
	r.HandleFunc("/delete/{id}", deleteProduct).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":4000", r))
}
