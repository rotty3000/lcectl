package spinner

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"

	"github.com/briandowns/spinner"
	"github.com/docker/docker/pkg/stdcopy"
)

func SpinnerPipe(s *spinner.Spinner, prefix string) func(io.ReadCloser, bool) {
	return func(out io.ReadCloser, verbose bool) {
		if verbose {
			stdcopy.StdCopy(os.Stdout, os.Stderr, out)
		} else if s != nil {
			c := make(chan (string))
			go func() {
				for {
					msg := <-c
					s.Suffix = fmt.Sprintf(prefix, msg)
				}
			}()

			reader := bufio.NewReader(out)
			for {
				str, err := reader.ReadString('\n')
				if err != nil {
					close(c)
					break
				} else {
					c <- removeInvisibleChars(truncateText(strings.TrimSpace(str), 80))
				}
			}
		}
	}
}

func truncateText(s string, max int) string {
	if max > len(s) {
		return s
	}
	return s[:strings.LastIndex(s[:max], " ")]
}

func removeInvisibleChars(s string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsGraphic(r) || unicode.IsPrint(r) {
			return r
		}
		return -1
	}, s)
}