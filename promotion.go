// For storing the logic for a specific promotion

package main

type Promotion struct {
	promoType string
	name string
	rules map[string]int
}

func NewPromotion(promoType string, name string, rules map[string]int) *Promotion {
	promotion := Promotion{
		promoType: promoType,
		name: name,
		rules: rules,
	}
	return &promotion
}