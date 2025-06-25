package utils

import (
	"regexp"
	"strings"
)

// IsValidEmail validates email format using regex
func IsValidEmail(email string) bool {
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	match, _ := regexp.MatchString(emailRegex, email)
	return match
}

// IsValidPhoneNumber validates phone number format
func IsValidPhoneNumber(phone string) bool {
	// Remove all non-numeric characters
	phoneRegex := `^\+?[1-9]\d{1,14}$`
	match, _ := regexp.MatchString(phoneRegex, phone)
	return match
}

// SanitizeString removes unwanted characters and trims whitespace
func SanitizeString(input string) string {
	// Remove leading and trailing whitespace
	sanitized := strings.TrimSpace(input)

	// Remove HTML tags (basic sanitization)
	htmlRegex := `<[^>]*>`
	re := regexp.MustCompile(htmlRegex)
	sanitized = re.ReplaceAllString(sanitized, "")

	return sanitized
}

// IsValidUsername validates username format
func IsValidUsername(username string) bool {
	// Username should be 3-30 characters, alphanumeric with underscores and hyphens
	usernameRegex := `^[a-zA-Z0-9_-]{3,30}$`
	match, _ := regexp.MatchString(usernameRegex, username)
	return match
}
