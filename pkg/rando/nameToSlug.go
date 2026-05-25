package rando

import "strings"

func NameToSlug(firstName, lastName string) string {
	combined := strings.ToLower(firstName + "-" + lastName)
	replacer := strings.NewReplacer("ğ", "g", "ü", "u", "ş", "s", "ı", "i", "ö", "o", "ç", "c", " ", "-")
	return replacer.Replace(combined) + "-" + randomString()
}
