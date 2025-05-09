package kzstring

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/meth-suchatchai/kz-blog-api/lib/errors"
	"strconv"
)

func CombineAll(params ...string) string {
	var newString string
	if len(params) > 0 {
		for _, v := range params {
			newString += v
		}
	}

	return newString
}

func ReplaceMobileCountryCode(params ...int) (int, *fiber.Error) {
	var mobile int
	var newString string
	if len(params) > 0 {
		for _, v := range params {
			newString += fmt.Sprintf("%d", v)
		}
	}

	if newString == "" {
		return 0, errors.NewDefaultFiberMessageError("mobile number is empty")
	}

	mobile, err := strconv.Atoi(newString)
	if err != nil {
		return 0, errors.NewDefaultFiberError(err)
	}

	return mobile, nil
}
