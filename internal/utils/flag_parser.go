package utils

import (
	"regexp"
	"strings"
)

func ParseFlags(args []string) map[string]string {
	flags := make(map[string]string)
	for _, arg := range args {
		if isFlag(arg) {
			key, value := extractFlag(arg)
			flags[key] = value
		}
	}

	return flags
}

func isFlag(value string) bool {
	// The value must follow the pattern '--key=value'.
	re := regexp.MustCompile(`^--([A-Za-z]+)=(.*)$`)
	return re.MatchString(value)
}

func extractFlag(value string) (string, string) {
	parts := strings.Split(value, "=")
	return strings.TrimPrefix(parts[0], "--"), parts[1]
}
