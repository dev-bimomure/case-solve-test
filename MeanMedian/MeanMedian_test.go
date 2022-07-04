package MeanMedian

import (
	"fmt"
	"testing"
)

func TestMeanMEdian(t *testing.T) {
	data := []float64{3, 4, 6, 17, 25, 21, 23}
	result := CalculateMeanMedian(data)

	fmt.Println(result)
}
