// Handles the main checkout flow, add product, remove, get total.

package checkout

import (
  products "github.com/benhawker/checkout-go/products"
  "fmt"
)

type Checkout struct {
  Basket map[int]int
  productCatalogue []products.Product
}

func NewCheckout(productCatalogue []products.Product) *Checkout {
  checkout := new(Checkout)
  checkout.Basket = make(map[int]int)
  checkout.productCatalogue = productCatalogue
  return checkout
}

func (checkout *Checkout) AddProduct(productCode int, quantity int) {
  // Refactor of products to map will allow easier 
  // check that product is in our catalogue
  checkout.Basket[productCode] = quantity
}

// Removes
func (checkout *Checkout) RemoveProduct(productCode int) {
  _, ok := checkout.Basket[productCode];

  if ok {
    delete(checkout.Basket, productCode);
  } else {
    fmt.Println("We could find this Product in your basket.")
  }
}

func (checkout *Checkout) Total() int {
  total := 0

  for code, quantity := range checkout.Basket {
    // TODO: Store products in a map with productCode as key 
    // to remove the need to iterate all products.
    for _, product := range checkout.productCatalogue {
      if code == product.Code {
        total += (quantity * product.Price)
      }
    }
  }
  return total
}
