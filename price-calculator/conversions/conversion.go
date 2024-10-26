package conversions

import (
	"errors"
	"fmt"
	"strconv"
)

func StringsToFloat(strings []string) ([]float64, error) {
	var floats []float64

	for _, line := range strings {
		floatPrices, err := strconv.ParseFloat(line, 64)

		if err != nil {
			fmt.Println(err)
			return nil, errors.New("Conversion Failed")
		}

		floats = append(floats, floatPrices)
	}
	return floats, nil
}
