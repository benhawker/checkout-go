package promotioncalculator_test

import (
  "testing"
  "github.com/cheekybits/is"
  promotioncalculator "github.com/benhawker/checkout-go/promotion-calculator"
  checkout "github.com/benhawker/checkout-go/checkout"
  products "github.com/benhawker/checkout-go/products"
  promotions "github.com/benhawker/checkout-go/promotions"
)

func TestCheckout_AddValidProduct(t *testing.T) {
	is := is.New(t)
	pc := setup()

	// Applies 10 off on order over 60.
  is.Equal(pc.Calculate(), 10)
}

func setup() (*promotioncalculator.PromotionCalculator) {
	var prod products.Products
  productCatalogue := prod.GetProducts("../config/products.yml").Products
  co := checkout.NewCheckout(productCatalogue)

  co.AddProduct(001, 7)

  var promo promotions.Promotions
  promotions := promo.GetPromotions("../config/promotions.yml").Promotions
  	
  pc := promotioncalculator.NewPromotionCalculator(promotions, co.Basket, co.Total())
  return pc
}