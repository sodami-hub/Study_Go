package main

import (
	"time"
)

type Courier struct {
	Name string
}

type Product struct {
	Name  string
	Price int
	ID    int
}

type Parcel struct {
	Pdt           *Product
	ShippedTime   time.Time
	DeliveredTime time.Time
}

func (c *Courier) SendProduct(product *Product) *Parcel {
	parcel := &Parcel{}
	parcel.Pdt = product
	parcel.ShippedTime = time.Now()

	return parcel
}

func (p *Parcel) Delivered() *Product {
	p.DeliveredTime = time.Now()
	return p.Pdt
}
