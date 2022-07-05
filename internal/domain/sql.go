package domain

import (
	"bufio"
	"errors"
	"io"
)

var EmptyInputError = errors.New("empty input value")

// DetergeSql TODO: Using this method, we are currently really only understanding the input one line at a time
func DetergeSql(in io.Reader, out io.Writer) error {

	scanner := bufio.NewScanner(in)
	lines := 0

	for scanner.Scan() {
		line := scanner.Text()

		// TODO: Run substitutions

		_, err := out.Write([]byte(line))

		if err != nil {
			return err
		}

		lines++
	}

	if lines == 0 {
		return EmptyInputError
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
