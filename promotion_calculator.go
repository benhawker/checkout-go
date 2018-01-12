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
	discountTotal := 0 

	for _, promotion := range promoCalc.promotions {
		if promotion.promoType == "basket_total_greater_than" {
			discountTotal += promoCalc.basket_total_greater_than(promotion)
		} else if promotion.promoType == "buy_x_get_y_free" {
			discountTotal += promoCalc.buy_x_get_y_free(promotion)
		}
	}

	return discountTotal
}

func (promoCalc *PromotionCalculator) basket_total_greater_than(promotion Promotion) int {
	if (promoCalc.basketTotal > promoCalc.promotions[0].rules["condition"]) {
		return promoCalc.promotions[0].rules["discount"]
	}
	return 0
}

func (promoCalc *PromotionCalculator) buy_x_get_y_free(promotion Promotion) int {
	return 0
}

