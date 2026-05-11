package main

import (
	"io"
	"os"
)

func catFile(name string) error {
	f, err := os.Open(name)
	if err != nil {
		return err
	}
	defer f.Close()
	io.Copy(os.Stdout, f)
	return nil
}

func main() {
	if len(os.Args) == 1 {
		io.Copy(os.Stdout, os.Stdin)
		return
	}
	for _, name := range os.Args[1:] {
		if err := catFile(name); err != nil {
			os.Stderr.Write([]byte("ERROR: " + err.Error() + "\n"))
			os.Exit(1)
		}
	}
}
