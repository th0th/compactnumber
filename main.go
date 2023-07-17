package compactnumber

import (
	"fmt"
	"math"
	"strings"
)

type Formattable interface {
	int | int16 | int32 | int64 | uint | uint16 | uint32 | uint64
}

func Format[T Formattable](v T) string {
	u := uint64(v)
	f := float64(v)

	var factor float64
	var letter string

	if u < 1_000 {
		return fmt.Sprintf("%d", u)
	} else if u < 1_000_000 {
		factor = 1000
		letter = "K"
	} else if u < 1_000_000_000 {
		factor = 1_000_000
		letter = "M"
	} else if u < 1_000_000_000_000 {
		factor = 1_000_000_000
		letter = "B"
	}

	return strings.Replace(
		fmt.Sprintf(
			"%.1f%s",
			math.Round((factor/10)*f/factor)/(factor/10),
			letter,
		),
		".0",
		"",
		1,
	)
}
