package rando

import "strings"

func NameToSlug(fullName string) string {
	combined := strings.ToLower(fullName + "-")
	replacer := strings.NewReplacer("ğ", "g", "ü", "u", "ş", "s", "ı", "i", "ö", "o", "ç", "c", " ", "-")
	return replacer.Replace(combined) + "-" + randomString()
}
