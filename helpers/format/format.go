package format

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func ToNumberString(f int64) string {
	return strconv.FormatInt(f, 10)
}

func ToFloatString(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}

func ToBooleanString(b bool) string {
	if b == true {
		return "oui"
	}

	return "non"
}

func ToErrorMessage(s string) map[string]any {
	return gin.H{"message": s}
}
