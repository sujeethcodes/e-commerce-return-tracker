package utils

import "time"

func FormateDate() (time.Time, error) {
	now := time.Now()
	// Format to string and parse again to remove time part
	formatted := now.Format("02-01-2006")
	dateOnly, err := time.Parse("02-01-2006", formatted)
	if err != nil {
		return time.Time{}, err
	}
	return dateOnly, nil
}
