// https://stackoverflow.com/questions/39147446/return-reference-to-struct-in-go-lang

package main

import (
  "fmt"
  promotions "github.com/benhawker/checkout-go/promotions"
  products "github.com/benhawker/checkout-go/products"
  co "github.com/benhawker/checkout-go/checkout"
  pc "github.com/benhawker/checkout-go/promotion-calculator"
)

func main() {
  var prod products.Products
  products := prod.GetProducts("./config/products.yml").Products

  var promo promotions.Promotions
  promotions := promo.GetPromotions("./config/promotions.yml").Promotions

  // Create a new checkout
  checkout := co.NewCheckout(products)
  printBreak()

  fmt.Println("Adding 3 products")
  checkout.AddProduct(001, 1)
  checkout.AddProduct(002, 3)
  checkout.AddProduct(003, 2)

  printBreak()

  fmt.Println("Removing Test Product 3.")
  checkout.RemoveProduct(003)

  printBreak()
  
  fmt.Println("Removing product not in the basket.")
  checkout.RemoveProduct(0055)
  
  printBreak()

  fmt.Println("Current Basket contents:")
  for k, v := range checkout.Basket { 
    fmt.Println("Product Code", k, "with quantity", v)
  }

  printBreak()
  
  fmt.Println("Your total before discounts is:", checkout.Total())

  // // Create our promo calculator
  promoCalc := pc.NewPromotionCalculator(promotions, checkout.Basket, checkout.Total())
  promotionalDiscount := promoCalc.Calculate()

  fmt.Println("Discount to be applied is:", promotionalDiscount)

  finalTotal := checkout.Total() - promotionalDiscount

  printBreak()
  fmt.Println("Your total with promotions applied is:", finalTotal)
}

func printBreak() {
  fmt.Println("- - - - - - - - - - - - - - - -")
}