package parse

import (
	"bufio"
	"os"
	"regexp"
	"strings"
)

var gitCommitRe = regexp.MustCompile(`^\[\w+ [a-f0-9]+\] (.+)$`)

func tryReadStdin() (string, bool) {
	stat, err := os.Stdin.Stat()
	if err != nil {
		return "", false
	}

	if (stat.Mode() & os.ModeCharDevice) != 0 {
		return "", false
	}

	scanner := bufio.NewScanner(os.Stdin)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return strings.Join(lines, "\n"), true
}

func parseFromStdin(data string) string {
	data = strings.TrimSpace(data)
	if data == "" {
		return ""
	}

	firstLine := data
	if idx := strings.IndexByte(data, '\n'); idx != -1 {
		firstLine = data[:idx]
	}

	if m := gitCommitRe.FindStringSubmatch(firstLine); len(m) == 2 {
		return strings.TrimSpace(m[1])
	}

	return firstLine
}