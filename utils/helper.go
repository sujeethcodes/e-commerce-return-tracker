package utils

import "time"

func FormateDate() (string, error) {
	now := time.Now()
	formatted := now.Format("02-01-2006")
	return formatted, nil
}
