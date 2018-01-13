// Handles the main checkout flow, add product, remove, get total.

package checkout

type Checkout struct {
  Basket map[string]int
}

func NewCheckout() *Checkout {
  checkout := new(Checkout)
  checkout.Basket = make(map[string]int)
  return checkout
}

func (checkout *Checkout) AddProduct(productCode string, quantity int) {
  checkout.Basket[productCode] = quantity
}

// Removes
func (checkout *Checkout) RemoveProduct(productCode string) {
  _, ok := checkout.Basket[productCode];

  if ok {
    delete(checkout.Basket, productCode);
  }
}

func (checkout *Checkout) Total() int {
  total := 0
  // for _, _ := range checkout.basket { 
    total += 100
  // }

  return total
}
