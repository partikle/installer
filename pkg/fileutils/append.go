package fileutils

import "os"

func Append(filename string, text []byte) error {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}

	defer f.Close()

	if _, err = f.Write(text); err != nil {
		return err
	}
	return nil
}
