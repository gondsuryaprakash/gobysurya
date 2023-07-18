package regularexpression

import (
	"regexp"
	"strings"
)

func IsMobile(ua string) bool {
	isValidMobile := regexp.MustCompile(`(ios|ipod|ipad|iphone|android)`)
	return isValidMobile.MatchString(strings.ToLower(ua))
}
