package main
import "fmt"

func main() {
  checkout := NewCheckout()

  fmt.Println("Adding 3 products")
  checkout.AddProduct("Test Product 1", 100)
  checkout.AddProduct("Test Product 2", 200)
  checkout.AddProduct("Test Product 3", 300)

  printBreak()

  fmt.Println("Removing Test Product 3")
  checkout.RemoveProduct("Test Product 3")
  
  printBreak()

  checkout.RemoveProduct("Here is your basket")
  for k, v := range checkout.basket { 
    fmt.Println(k, v)
  }

  printBreak()
  
  fmt.Println("Your total before discounts is:", checkout.Total())

  // Create a simple promotion
  promoType := "buy_x_get_one_free"
  name := "If you spend over £60, then you get 10 percent off your purchase"

  rules := map[string]int {
    "condition": 60,
    "discount": 10.00,
  }

  promotion := NewPromotion(promoType, name, rules);

  // Hold our promotions in a slice
  promotions := []Promotion{}
  promotions = append(promotions, *promotion)

  // Create our promo calculator
  promoCalc := NewPromotionCalculator(promotions, checkout.basket, checkout.Total())
  promotionalDiscount := promoCalc.Calculate()
  finalTotal := checkout.Total() - promotionalDiscount

  printBreak()
  fmt.Println("Your total with promotions applied is:", finalTotal)
}

func printBreak() {
  fmt.Println("- - - - - - - -")
}