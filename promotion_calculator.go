package main

import "fmt"

type PromotionCalculator struct {
	promotions []Promotion
	basket map[string]int
	basketTotal int
}

func NewPromotionCalculator(promotions []Promotion, basket map[string]int, basketTotal int) *PromotionCalculator {
	promotionCalculator := new(PromotionCalculator)
	promotionCalculator.promotions = promotions
	promotionCalculator.basket = basket
	promotionCalculator.basketTotal = basketTotal 

	// return promotion
	fmt.Println("It did not blow up.")
	return promotionCalculator
}


func (promoCalc *PromotionCalculator) Calculate() int {
	// Simply return the discount amount for now.
	return promoCalc.promotions[0].rules["discount"]
}

