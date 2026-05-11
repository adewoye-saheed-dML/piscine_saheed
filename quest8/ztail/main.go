package main

import (
	"os"
)

func main() {
	if len(os.Args) < 4 {
		os.Exit(1)
	}

	if os.Args[1] != "-c" {
		os.Exit(1)
	}

	n := int64(0)
	for _, ch := range os.Args[2] {
		if ch >= '0' && ch <= '9' {
			n = n*10 + int64(ch-'0')
		} else {
			os.Exit(1)
		}
	}

	files := os.Args[3:]
	multipleFiles := len(files) > 1
	hasError := false

	// We now track the index 'i' to know exactly which file we are processing
	for i, file := range files {
		f, err := os.Open(file)
		if err != nil {
			os.Stdout.WriteString(err.Error() + "\n")
			hasError = true
			continue
		}

		info, err := f.Stat()
		if err != nil {
			os.Stdout.WriteString(err.Error() + "\n")
			hasError = true
			f.Close()
			continue
		}

		if multipleFiles {
			// If it is NOT the very first file argument, we prepend a newline
			if i > 0 {
				os.Stdout.WriteString("\n")
			}
			os.Stdout.WriteString("==> " + file + " <==\n")
		}

		size := info.Size()
		startIdx := int64(0)

		if size > n {
			startIdx = size - n
		}

		_, err = f.Seek(startIdx, 0)
		if err == nil {
			buf := make([]byte, 4096)
			for {
				bytesRead, rErr := f.Read(buf)
				if bytesRead > 0 {
					os.Stdout.Write(buf[:bytesRead])
				}
				if rErr != nil {
					break
				}
			}
		}
		f.Close()
	}

	if hasError {
		os.Exit(1)
	}
}
