package promotionCalculator

import (
  promotions "checkout-go/promotions"
)

type PromotionCalculator struct {
  promotions []promotions.Promotion
  basket map[string]int
  basketTotal int
}

func NewPromotionCalculator(promotions []promotions.Promotion, basket map[string]int, basketTotal int) *PromotionCalculator {
  promotionCalculator := new(PromotionCalculator)
  promotionCalculator.promotions = promotions
  promotionCalculator.basket = basket
  promotionCalculator.basketTotal = basketTotal 
  return promotionCalculator
}

func (promoCalc *PromotionCalculator) Calculate() int {
  discountTotal := 0 

  for _, promotion := range promoCalc.promotions {
    if promotion.PromoType == "basket_total_greater_than" {
      discountTotal += promoCalc.basket_total_greater_than(promotion)
    } else if promotion.PromoType == "buy_x_get_y_free" {
      discountTotal += promoCalc.buy_x_get_y_free(promotion)
    }
  }

  return discountTotal
}

func (promoCalc *PromotionCalculator) basket_total_greater_than(promotion promotions.Promotion) int {
  if (promoCalc.basketTotal > promotion.Condition) {
    return promotion.Discount
  }
  return 0
}

func (promoCalc *PromotionCalculator) buy_x_get_y_free(promotion promotions.Promotion) int {
  return 0
}

