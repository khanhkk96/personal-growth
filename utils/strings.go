package utils

import (
	"strconv"
	"strings"
	"time"
)

func IsEmpty(s *string) bool {
	return s == nil || *s == ""
}

func ParseDurationFromEnv(raw string) (time.Duration, error) {
	if strings.HasSuffix(raw, "d") {
		daysStr := strings.TrimSuffix(raw, "d")
		days, err := strconv.Atoi(daysStr)
		if err != nil {
			return 0, err
		}
		return time.Hour * 24 * time.Duration(days), nil
	}
	// fallback: parse bình thường (ví dụ: "10h", "30m")
	return time.ParseDuration(raw)
}
