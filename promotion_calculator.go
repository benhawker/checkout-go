package main


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
	return promotionCalculator
}

func (promoCalc *PromotionCalculator) Calculate() int {
	// Simply return the discount amount for now.
	return promoCalc.promotions[0].rules["discount"]
}

