package utils

import (
	"math"

	"github.com/benosborntech/feedme/common/consts"
)

func HashPrecision(radius float64) int {
	hashDivideX := 8.0

	precision := 1
	currSize := consts.WORLD_CIRCUMFERENCE
	for currSize/hashDivideX > radius {
		precision += 1
		currSize = currSize / hashDivideX
	}

	return int(math.Min(float64(precision), consts.MAX_PRECISION))
}
