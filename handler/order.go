package handler

import (
	"fmt" 
	"net/http"
)

type Order struct{

}

func (o *Order) Create(w http.ResponseWriter, r *http.Request){
	fmt.Println("Order has been created")
}

func (o *Order) List(w http.ResponseWriter, r *http.Request){
	fmt.Println("List all orders")
}

func (o *Order) GetByID(w http.ResponseWriter, r *http.Request){
	fmt.Println("Get by id method has been called")
}

func (o *Order) UpdateByID(w http.ResponseWriter, r *http.Request){
	fmt.Println("Update method has been called")
}

func (o *Order) Delete(w http.ResponseWriter, r *http.Request){
	fmt.Println("Delete method has been called")
}