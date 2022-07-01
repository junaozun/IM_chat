package utils

import (
	"fmt"
	"time"

	"github.com/junaozun/IM_chat/utils/encrypt"
)

func GenerateToken() string {
	str := fmt.Sprintf("%d", time.Now().Unix())
	return encrypt.Md5Encode(str)
}
