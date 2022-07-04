package MeanMedian

import (
	"sort"
	"strconv"
)

func CalculateMeanMedian(data []float64) map[string]string {
	result := make(map[string]string)

	Mean := mean(data)
	Median := median(data)

	result["Mean"] = strconv.FormatFloat(Mean, 'f', 6, 64)
	result["Median"] = strconv.FormatFloat(Median, 'f', 6, 64)

	return result
}

//calculate Mean
func mean(data []float64) float64 {
	var count float64 = 0
	len := len(data)
	decLen := float64(len)

	for _, num := range data {
		count = count + num
	}
	result := float64(count / decLen)

	return result
}

//calculate Median
func median(data []float64) float64 {
	var result float64 = 0
	sort.Float64s(data)
	lenData := len(data)
	if lenData%2 == 0 {
		data1 := data[:lenData]
		data2 := data[lenData:]

		float1 := data1[len(data1)]
		float2 := data2[0]

		result = float1 / float2
	} else {
		index := (lenData - 1) / 2

		result = float64(data[index+1])
	}

	return result
}
