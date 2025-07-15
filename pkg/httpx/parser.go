package httpx

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func ParseIDFromParam(c *gin.Context, key string) (uint, error) {
	id, err := strconv.ParseUint(c.Param(key), 10, 32)
	if err != nil {
		return 0, err
	}

	return uint(id), nil
}
