package dict

import (
	"bufio"
	"os"
	"strings"
)

// FormatKaikki reads a wiktionary dump from `kaikki.org` and formats it to be
// valid JSON.
func FormatKaikki(fp string) (b []byte, err error) {
	f, err := os.Open(fp)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	sb := strings.Builder{}
	sb.WriteString("[")

	if scanner.Scan() {
		sb.WriteString(scanner.Text())
	}

	for scanner.Scan() {
		sb.WriteByte(',')
		sb.WriteString(scanner.Text())
	}

	sb.WriteString("]")

	return []byte(sb.String()), scanner.Err()
}
