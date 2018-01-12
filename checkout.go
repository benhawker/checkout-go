// Handles the main checkout flow, add product, remove, get total.

package main

type Checkout struct {
  basket map[string]int
}

func NewCheckout() *Checkout {
  checkout := new(Checkout)
  checkout.basket = make(map[string]int)
  return checkout
}

func (checkout *Checkout) AddProduct(productName string, quantity int) {
  checkout.basket[productName] = quantity
}

func (checkout *Checkout) RemoveProduct(productName string) {
  _, ok := checkout.basket[productName];

  if ok {
    delete(checkout.basket, productName);
  }
}

func (checkout *Checkout) Total() int {
  total := 0
  // for _, _ := range checkout.basket { 
    total += 100
  // }

  return total
}
