package apperror

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func ExtractValidationError(fe validator.FieldError) (string) {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "gte":
		return "Should be greater than or equal with " + fe.Param()
	case "lte":
		return "Should be less than or equal with " + fe.Param()
	case "min":
		return "Should be greater than or equal with " + fe.Param()
	case "max":
		return "Should be less than or equal with " + fe.Param()
	case "email":
		return fe.Field() + " must be in email format"
	case "datetime":
		return "date is not valid"
	}

	return "unknown error"
}


func IsPasswordValid(password string) bool {
    // At least one letter
    hasLetter := regexp.MustCompile(`[A-Za-z]`).MatchString(password)
    // At least one digit
    hasDigit := regexp.MustCompile(`[0-9]`).MatchString(password)
    // At least one symbol (you can customize the symbol set)
    hasSymbol := regexp.MustCompile(`[!@#\$%\^&\*\(\)_\+\-=\[\]\{\};:'",<>\.\?/\\|]`).MatchString(password)

    return hasLetter && hasDigit && hasSymbol
}

func IsAlphanumeric(input string) bool {
    match, _ := regexp.MatchString(`^[a-zA-Z0-9]+$`, input)
    return match
}

