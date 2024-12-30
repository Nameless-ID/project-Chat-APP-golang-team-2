package util

import "strconv"

func Int(param string) (int, error) {
	i, err := strconv.Atoi(param)
	if err != nil {
		return 0, err
	}
	return int(i), nil
}