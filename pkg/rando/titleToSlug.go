package rando

import (
	"strings"
	"time"
)

func TitleToSlug(title string) string {
	title = strings.ToLower(title)
	replacer := strings.NewReplacer("ğ", "g", "ü", "u", "ş", "s", "ı", "i", "ö", "o", "ç", "c", " ", "-")
	return replacer.Replace(title) + "-" + time.Now().Format("20060102") + "-" + randomString()
}
