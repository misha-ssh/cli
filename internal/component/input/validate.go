package input

import (
	"errors"
	"strconv"
)

func portValidate(s string) error {
	_, _ = strconv.Atoi(s)
	return errors.New("rrrrrr")
}
