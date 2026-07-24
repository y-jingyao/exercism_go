package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	//panic("Please implement the ParseCard function")
    	switch card {
	case "ace":
		return 11
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	case "ten", "jack", "queen", "king":
		return 10
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	//panic("Please implement the FirstTurn function")
    v1 := ParseCard(card1)
	v2 := ParseCard(card2)
	dv := ParseCard(dealerCard)
	total := v1 + v2

	if card1 == "ace" && card2 == "ace" {
		return "P"
	}
    
	if total == 21 {
		if dv == 10 || dv == 11 {
			return "S"
		}
		return "W"
	}
    
	if total >= 17 && total <= 20 {
		return "S"
	}
    
	if total >= 12 && total <= 16 {
		if dv >= 7 {
			return "H"
		}
		return "S"
	}
    
	return "H"
}
