package utils

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetQueryIntGin(c *gin.Context, key string, defaultVal int) (int, error) {
	value := c.DefaultQuery(key, strconv.Itoa(defaultVal))
	if value == "" {
		return defaultVal, nil
	}
	intValue, err := strconv.Atoi(value)
	if err != nil {
		return defaultVal, err
	}
	return intValue, nil
}
