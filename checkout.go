package main

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

func (checkout *Checkout) Total() int {
  total := 0
  for _, v := range checkout.basket { 
    total += v
  }

  return total
}
