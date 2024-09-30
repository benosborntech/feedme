package utils

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"

	"github.com/benosborntech/feedme/apigw/types"
)

func GetUserId(serviceType types.ServiceType, sub string) int {
	uniqueId := fmt.Sprintf("%s:%s", serviceType, sub)

	hash := sha256.Sum256([]byte(uniqueId))
	userId := binary.BigEndian.Uint64(hash[:8])

	return int(userId)
}
