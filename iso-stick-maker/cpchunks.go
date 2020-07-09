package main

import (
	"bufio"
	"io"
	"os"
)

func cpChunks(src, dst string, percent chan<- int) error {

	data := make([]byte, 4*1024*1024)

	in, err := os.Open(src)
	if err != nil {
		return err
	}

	reader := bufio.NewReader(in)
	defer in.Close()

	fi, err := in.Stat()
	if err != nil {
		return err
	}

	out, err := os.OpenFile(dst, os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	writer := bufio.NewWriter(out)
	defer out.Close()

	total := 0

	for {
		count, err := reader.Read(data)
		total += count
		data = data[:count]

		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		_, err = writer.Write(data)
		if err != nil {
			return err
		}

		percent <- int(int64(total) * int64(100) / fi.Size())

	}
	return nil

}
