// https://stackoverflow.com/questions/39147446/return-reference-to-struct-in-go-lang

package main

import (
  "fmt"
  promotions "checkout-go/promotions"
  products "checkout-go/products"
  co "checkout-go/checkout"
  pc "checkout-go/promotioncalculator"
)

func main() {
  checkout := co.NewCheckout()

  fmt.Println("Adding 3 products")
  checkout.AddProduct("001", 2)
  checkout.AddProduct("002", 2)
  checkout.AddProduct("003", 2)

  printBreak()

  fmt.Println("Removing Test Product 3")
  checkout.RemoveProduct("003")
  
  printBreak()

  checkout.RemoveProduct("Here is your basket")
  for k, v := range checkout.Basket { 
    fmt.Println(k, v)
  }

  printBreak()
  
  fmt.Println("Your total before discounts is:", checkout.Total())

  var prod products.Products
  products := prod.GetProducts().Products
  fmt.Println("Here are your products:", products)

  var promo promotions.Promotions
  promotions := promo.GetPromotions().Promotions
  fmt.Println("Here are your promotions:", promotions)

  // // Create our promo calculator
  promoCalc := pc.NewPromotionCalculator(promotions, checkout.Basket, checkout.Total())
  promotionalDiscount := promoCalc.Calculate()

  fmt.Println("Discount to be applied is:", promotionalDiscount)

  finalTotal := checkout.Total() - promotionalDiscount

  printBreak()
  fmt.Println("Your total with promotions applied is:", finalTotal)
}

func printBreak() {
  fmt.Println("- - - - - - - -")
}