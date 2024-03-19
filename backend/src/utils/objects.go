package utils

import "github.com/gin-gonic/gin"

func ToStructOrSendError[T any](c *gin.Context) (*T, error) {
	res := new(T)
	err := c.ShouldBindJSON(&res)
	if err != nil {
		c.JSON(422, ApiError{-1, "Json parsing error"})
		return nil, err
	}
	return res, nil
}
