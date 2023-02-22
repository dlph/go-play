package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/dustin/go-humanize"
	"github.com/spf13/afero"
)

func printProgress(n uint64) {
	// Clear the line by using a character return to go back to the start and remove
	// the remaining characters by filling it with spaces
	fmt.Printf("\r%s", strings.Repeat(" ", 35))

	// Return again and print current status of download
	// We use the humanize package to print the bytes in a meaningful way (e.g. 10 MB)
	fmt.Printf("\rDownloading... %s complete\n", humanize.Bytes(n))
}

type WriterFunc func([]byte) (int, error)

func (wr WriterFunc) Write(p []byte) (int, error) {
	return wr(p)
}

func countWriter() WriterFunc {
	var total int

	return func(p []byte) (int, error) {
		n := len(p)
		total += n
		printProgress(uint64(total))
		return n, nil
	}
}

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}

func run() error {
	fs := afero.NewOsFs()
	name := "goheadscale.mp4"
	out, err := fs.Create(name)
	if err != nil {
		return err
	}
	defer out.Close()

	resp, err := http.Get("https://ftp.osuosl.org/pub/fosdem/2023/UD2.218A/goheadscale.mp4")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if _, err := io.Copy(out, io.TeeReader(resp.Body, countWriter())); err != nil {
		return err
	}

	return nil
}
