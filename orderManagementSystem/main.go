package main

import (
	"errors"
	"fmt"
)

type Product struct {
	ID    int
	Name  string
	Price float64
	Stock int
}

type Order struct {
	ID       int
	Customer string
	Items    []Product
	Status   bool
}

type User struct {
	ID   int
	Name string
}

type Shop struct {
	ID       int
	Name     string
	Products []Product
	Orders   []Order
}

var ErrProductStock = errors.New("Product stock = 0")

func (s Shop) getAvailableProducts() []Product {
	products := []Product{}
	for i := range s.Products {
		curProd := s.Products[i]
		if curProd.Stock > 0 {
			products = append(products, curProd)
		}
	}

	return products
}

func (s *Shop) createProduct(name string, price float64, stock int) *Product {
	newProdID := len(s.Products)

	newProd := Product{
		ID:    newProdID,
		Name:  name,
		Price: price,
		Stock: stock,
	}

	s.Products = append(s.Products, newProd)

	return &s.Products[len(s.Products)-1]
}

func (o *Order) addProduct(product *Product) error {
	if product.Stock < 1 {
		return fmt.Errorf("addProduct error: %w", ErrProductStock)
	}
	o.Items = append(o.Items, *product)
	product.Stock -= 1

	return nil
}

func (s Shop) printProducts(printNotAvailable bool) {
	var products []Product

	if printNotAvailable {
		products = s.Products
	} else {
		products = s.getAvailableProducts()
	}

	for i := range products {
		curProd := products[i]
		fmt.Printf("#%v: %s\n- Price: %v\n- Stock: %v\n", curProd.ID, curProd.Name, curProd.Price, curProd.Stock)
	}
}

func (s Shop) printOrders() {
	for _, order := range s.Orders {
		fmt.Println(order)
	}
}

func (s *Shop) makeNewOrder(customer string, items []Product, status bool) int {
	orderID := len(s.Orders)

	order := Order{
		ID:       orderID,
		Customer: customer,
		Items:    items,
		Status:   status,
	}

	s.Orders = append(s.Orders, order)

	return orderID
}

func main() {
	wb := Shop{
		ID:   1,
		Name: "Wildberries",
	}

	mouse := wb.createProduct("Mouse", 100.50, 28)
	keyboard := wb.createProduct("Keyboard", 200.67, 93)

	wb.printProducts(true)

	user := User{
		ID:   1,
		Name: "utyara3",
	}

	wb.makeNewOrder(user.Name, []Product{*mouse, *keyboard}, false)
}
