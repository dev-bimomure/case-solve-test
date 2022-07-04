package CountMoney

import (
	"log"
)

func CountCash(price int, bills int) []int {
	if price < 0 || price > 100000 {
		log.Println("Bills exceed condition")
		return nil
	}
	if bills < 0 || bills > 10 {
		log.Println("Bills exceed condition")
		return nil
	}
	notes := []int{100000, 50000, 20000, 10000, 5000, 1000}

	currentPrice := price
	var result []int
	for i := 0; i < bills; i++ {

		for j := 0; j < len(notes); j++ {
			if currentPrice >= notes[j] {
				result = append(result, notes[j])
				currentPrice = currentPrice - notes[j]
				break
			}

		}

	}

	return result
}
