package renderer

import (
	"io"
	"regexp"
)

type findStringWriter struct {
	channel chan string
	found   bool
	writer  io.Writer
	regex   *regexp.Regexp
}

func (w *findStringWriter) Write(p []byte) (int, error) {
	if w.writer != nil {
		n, err := w.writer.Write(p)
		if w.found {
			return n, err
		}
	}

	if !w.found {
		matches := w.regex.FindAllString(string(p), -1)
		for _, match := range matches {
			select {
			case w.channel <- match:
				break
			default:
				break
			}
			w.found = true
		}

		if w.found {
			close(w.channel)
		}
	}
	return len(p), nil
}

// NewFindStringWriter creates a writer which writes matches to a channel
func NewFindStringWriter(channel chan string, r *regexp.Regexp, writer io.Writer) io.Writer {
	return &findStringWriter{
		channel: channel,
		found:   false,
		regex:   r,
		writer:  writer,
	}
}
