# Promo/Checkout application in Go.

A Go appliction that builds a checkout system with a variety of promotional rules.

Variation on a theme based on a prior [Ruby repo](https://github.com/benhawker/benhawker/promo).

===================

### Usage:

```
make run
```


Run the specs:
```
make test
```

===================

Example output:

```
- - - - - - - - - - - - - - - -
Adding 3 products
- - - - - - - - - - - - - - - -
Removing Test Product 3.
- - - - - - - - - - - - - - - -
Removing product not in the basket.
We could find this Product in your basket.
- - - - - - - - - - - - - - - -
Current Basket contents:
Product Code 1 with quantity 1
Product Code 2 with quantity 3
- - - - - - - - - - - - - - - -
Your total before discounts is: 70
Discount to be applied is: 10
- - - - - - - - - - - - - - - -
Your total with promotions applied is: 60
```