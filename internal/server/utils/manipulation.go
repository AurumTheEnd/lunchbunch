package utils

import (
	"strconv"
	"strings"
)

func TrimPathToUint(path string, prefix string) (uint, error) {
	var stringId = strings.TrimSuffix(strings.TrimPrefix(path, prefix), "/")
	var id, err = strconv.ParseUint(stringId, 10, 0)
	if err != nil {
		return 0, err
	}

	return uint(id), nil
}
