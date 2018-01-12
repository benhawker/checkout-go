package main
import "fmt"
// import "strconv"

func main() {
  checkout := NewCheckout()
  checkout.AddProduct("Test Product 1", 100)
  checkout.AddProduct("Test Product 2", 200)
  checkout.AddProduct("Test Product 3", 300)

  checkout.RemoveProduct("Test Product 3")
  
  for k, v := range checkout.basket { 
    fmt.Println(k, v)
  }

  fmt.Println("Your total is:")
  fmt.Println(checkout.Total())

  promoType := "buy_x_get_one_free"
  name := "If you spend over £60, then you get 10 percent off your purchase"

  rules := map[string]int {
    "condition": 60,
    "discount": 10.00,
  }

  promotion := NewPromotion(promoType, name, rules);

  promotions := []Promotion{}
  promotions = append(promotions, *promotion)

  //We now have promotions.
  fmt.Println(promotions)

  promoCalc := NewPromotionCalculator(promotions, checkout.basket, checkout.Total())

  fmt.Println("Apply this discount:", promoCalc.Calculate())

  promotionalDiscount := promoCalc.Calculate()
  finalTotal := checkout.Total() - promotionalDiscount
  fmt.Println(finalTotal)
} 