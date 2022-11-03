package time

import "time"

func GetTimeFromString(input string) (time.Time, error) {
	result, err := time.Parse(time.RFC3339, input)
	if err != nil {
		return time.Time{}, err
	}

	return result, nil
}
