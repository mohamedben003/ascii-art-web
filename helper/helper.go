package helper

import (
	"fmt"
	"os"
	"strings"
)

func LoadBanner(path string) (map[rune][]string, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("can't open read err", err)
		return nil, err
	}

	lines := strings.Split(string(content), "\n")

	// skip the first empty line (same as loadBannerV1)
	if len(lines) > 0 && lines[0] == "" {
		lines = lines[1:]
	}

	banner := make(map[rune][]string)
	char := rune(32)

	var block []string
	for _, line := range lines {
		// empty line means block is complete
		if line == "" {
			if len(block) == 8 {
				banner[char] = block
				char++
			}
			block = []string{}
			continue
		}

		block = append(block, line)
	}

	// catch last block if file doesn't end with newline
	if len(block) == 8 {
		banner[char] = block
	}

	return banner, nil
}

func CheckNewLine(text string) bool {
	for _, char := range text {
		if char != '\n' {
			return false
		}
	}
	return true
}

func MainHelper(text, bannerTitle string) (string, error) {
	// Normalize newlines
	text = strings.ReplaceAll(text, "\r\n", `\n`)
	text = strings.ReplaceAll(text, "\n", `\n`)

	if text == "" {
		return "", nil
	}

	if CheckNewLine(text) {
		return text, nil
	}

	// Determine which banner file to use
	var bannerFile string
	switch bannerTitle {
	case "shadow":
		bannerFile = "helper/shadow.txt"
	case "standard":
		bannerFile = "helper/standard.txt"
	case "thinkertoy":
		bannerFile = "helper/thinkertoy.txt"
	default:
		return "", fmt.Errorf("invalid banner: %s", bannerTitle)
	}

	banner, err := LoadBanner(bannerFile)
	if err != nil {
		return "", fmt.Errorf("could not load banner file: %v", err)
	}

	result := ""
	lines := strings.Split(text, "\\n")

	for i, line := range lines {
		if line == "" {
			result += "\n"
			continue
		}

		for row := 0; row < 8; row++ {
			for _, char := range line {
				asciiArt, ok := banner[char]
				if ok && row < len(asciiArt) {
					result += asciiArt[row]
				}
			}
			result += "\n"
		}

		if i < len(lines)-1 && lines[i+1] != "" {
			result += "\n"
		}
	}

	return result, nil
}
