package main
import "fmt"

type Checkout struct {
  basket map[string]int
}

func NewCheckout() *Checkout {
  checkout := new(Checkout)
  checkout.basket= make(map[string]int)
  return checkout
}

func (checkout *Checkout) AddProduct(productName string, price int) {
  checkout.basket[productName] = price
}

func (checkout *Checkout) RemoveProduct(productName string) {
  _, ok := checkout.basket[productName];

  if ok {
    delete(checkout.basket, productName);
  }
}

func (checkout *Checkout) Total() int{
  total := 0
  for _, v := range checkout.basket { 
    total += v
  }

  return total
}

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
} 