package e

import "fmt"

func Wrap(errMsg string, err error) error {
	if err == nil {
		return nil
	}

	return fmt.Errorf("%s: %w", errMsg, err)
}
