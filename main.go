// https://stackoverflow.com/questions/39147446/return-reference-to-struct-in-go-lang

package main
import (
  "fmt"
  "gopkg.in/yaml.v2"
  "log"
  "io/ioutil"
  "path/filepath"
)

type Config struct {
  Products []Product  `yaml:"products"`
}

type Product struct {
  Code  int `yaml:"code"`
  Name  string `yaml:"name"`
  Price int `yaml:"price"`
}

func main() {
  filename, _ := filepath.Abs("./config/products.yml")
  yamlFile, err := ioutil.ReadFile(filename)
  if err != nil {
    log.Fatalf("error: %v", err)
  }

  var config Config

  err = yaml.Unmarshal(yamlFile, &config)
  if err != nil {
      panic(err)
  }

  fmt.Printf("Value: %#v\n", config.Products)
  fmt.Printf("Value: %#v\n", config.Products[0].Name)


  checkout := NewCheckout()

  fmt.Println("Adding 3 products")
  checkout.AddProduct("Test Product 2", 2)
  checkout.AddProduct("Test Product 2", 2)
  checkout.AddProduct("Test Product 3", 2)

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
  promoType := "basket_total_greater_than"
  name := "If you spend over £60, then you get 10 percent off your purchase"

  rules := map[string]int {
    "condition": 60,
    "discount": 10.00,
  }

  promotion := NewPromotion(promoType, name, rules);

  // A second promotion
  rules = map[string]int {
    "product_code": 001,
    "condition": 2,
    "get_free": 1,
  }

  promotion_two := NewPromotion("buy_x_get_y_free", "Buy Two T-shirts and Get One free", rules);

  // Hold our promotions in a slice (+ add them to the slice)
  promotions := []Promotion{}
  promotions = append(promotions, *promotion)
  promotions = append(promotions, *promotion_two)

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