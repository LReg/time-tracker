package parse

import (
	"bufio"
	"os"
	"regexp"
	"strings"
)

var gitCommitRe = regexp.MustCompile(`^\[\w+ [a-f0-9]+\] (.+)$`)

type stdinParser func(input string) (parsed string, matched bool)

var stdinParsers = []stdinParser{
	parseGitCommit,
}

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

func parseGitCommit(input string) (string, bool) {
	firstLine := input
	if idx := strings.IndexByte(input, '\n'); idx != -1 {
		firstLine = input[:idx]
	}

	m := gitCommitRe.FindStringSubmatch(firstLine)
	if len(m) != 2 {
		return "", false
	}

	return strings.TrimSpace(m[1]), true
}

func parseFromStdin(data string) string {
	data = strings.TrimSpace(data)
	if data == "" {
		return ""
	}

	for _, p := range stdinParsers {
		if parsed, matched := p(data); matched {
			return parsed
		}
	}

	if idx := strings.IndexByte(data, '\n'); idx != -1 {
		return data[:idx]
	}

	return data
}