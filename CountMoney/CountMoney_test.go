package CountMoney

import (
	"fmt"
	"testing"
)

func TestCountCash(t *testing.T) {
	result := CountCash(23000, 4)

	fmt.Println(result)
}
