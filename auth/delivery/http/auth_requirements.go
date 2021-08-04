package http

import (
	"fmt"
	"regexp"

	"github.com/spf13/viper"
)

var (
	userNameReq []func(*string) bool = []func(*string) bool{
		usernameMinLength_NonWhiteSpace,
	}

	passwordReq []func(*string) bool = []func(*string) bool{
		passMinLength_NonWhiteSpace,
		passUpperCase,
		passLowerCase,
		passDigits,
	}
)

func usernameMinLength_NonWhiteSpace(str *string) bool {
	min_length := viper.GetInt("app.auth.requirements.min_length.username")
	format := fmt.Sprintf(`[\S]{%d,}`, min_length)
	re := regexp.MustCompile(format)
	return re.MatchString(*str)
}

func passMinLength_NonWhiteSpace(str *string) bool {
	min_length := viper.GetInt("app.auth.requirements.min_length.password")
	format := fmt.Sprintf(`[\S]{%d,}`, min_length)
	re := regexp.MustCompile(format)
	return re.MatchString(*str)
}

func passUpperCase(str *string) bool {
	count := viper.GetInt("app.auth.requirements.number_of.upper_case")
	format := fmt.Sprintf(`[A-Z]{%d,}`, count)
	re := regexp.MustCompile(format)
	return re.MatchString(*str)
}

func passLowerCase(str *string) bool {
	count := viper.GetInt("app.auth.requirements.number_of.lower_case")
	format := fmt.Sprintf(`[a-z]{%d,}`, count)
	re := regexp.MustCompile(format)
	return re.MatchString(*str)
}

func passDigits(str *string) bool {
	count := viper.GetInt("app.auth.requirements.number_of.digits")
	format := fmt.Sprintf(`[\d]{%d,}`, count)
	re := regexp.MustCompile(format)
	return re.MatchString(*str)
}
