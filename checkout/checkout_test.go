package checkout_test

import (
  "testing"
  "github.com/cheekybits/is"
  checkout "github.com/benhawker/checkout-go/checkout"
  products "github.com/benhawker/checkout-go/products"
)

func TestCheckout_AddValidProduct(t *testing.T) {
	is := is.New(t)
	co, _ := setup()

  co.AddProduct(001, 2)
  is.Equal(co.Basket[001], 2)
}

func TestCheckout_AddProductMultipleQuantities(t *testing.T) {
	is := is.New(t)
  co, _ := setup()

  co.AddProduct(001, 2)
  is.Equal(co.Basket[001], 2)
}	

func TestCheckout_RemoveProduct(t *testing.T) {
	is := is.New(t)
  co, _ := setup()

  co.AddProduct(001, 2)
  is.Equal(co.Basket[001], 2)

  co.RemoveProduct(001)
  is.Equal(co.Basket, make(map[int]int))
}

func TestCheckout_SingleProductTotal(t *testing.T) {
	is := is.New(t)
  co, _ := setup()

  co.AddProduct(001, 2)

  is.Equal(co.Total(), 20) 
}

func TestCheckout_MultipleProductsTotal(t *testing.T) {
	is := is.New(t)
  co, _ := setup()

  co.AddProduct(001, 2)
  co.AddProduct(002, 3)
  co.AddProduct(003, 1)

  is.Equal(co.Total(), 110) 
}

func setup() (*checkout.Checkout, []products.Product) {
	var prod products.Products
  productCatalogue := prod.GetProducts("../config/products.yml").Products
  co := checkout.NewCheckout(productCatalogue)
  return co, productCatalogue
}

