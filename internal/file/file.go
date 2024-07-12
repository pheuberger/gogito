package file

import (
	"os"
)

func Write(filepath string, text string) error {
	f, err := os.Create(filepath)
	if err != nil {
		return err
	} else {
		defer f.Close()
	}

	if _, err := f.WriteString(text); err != nil {
		return err
	}
	return nil
}
