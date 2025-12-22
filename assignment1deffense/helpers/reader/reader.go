package reader

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func NewReader() *Reader {
	return &Reader{reader: bufio.NewReader(os.Stdin)}
}

type Reader struct {
	reader *bufio.Reader
}

func (r *Reader) ReadValue(variable interface{}) {
	for {
		_, err := fmt.Scan(variable)
		if err != nil {
			fmt.Println("Wrong format")
			continue
		}

		break
	}
}

func (r *Reader) ReadString(value *string, until byte) error {
	var err error
	*value, err = r.reader.ReadString(until)
	if err != nil {
		return err
	}

	*value = strings.TrimSpace(*value)
	return nil
}
