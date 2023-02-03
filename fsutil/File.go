package fsutil

import (
	"bufio"
	"os"
)

func TruncWrite(d , v string) error {
	file, err := os.OpenFile(d, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	writer.WriteString(v)
	writer.Flush()
	return nil
}

func Write(d , v string)error {
	file, err := os.OpenFile(d, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	writer.WriteString(v)
	writer.Flush()
	return nil
}
