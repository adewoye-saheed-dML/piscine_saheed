package piscine

type food struct {
	preptime int
}

func FoodDeliveryTime(order string) int {
	var f food

	switch order {
	case "burger":
		f.preptime = 15
	case "chips":
		f.preptime = 10
	case "nuggets":
		f.preptime = 12
	default:
		return 404
	}
	return f.preptime
}
