package util

import (
	"bufio"
	"os"
	"strings"
)

var GoogleClientID, GoogleClientSecret, GithubClientID, GithubClientSecret string

func LoadEnvVariables(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		os.Setenv(key, value)
	}

	GithubClientID = os.Getenv("GITHUB_CLIENT_ID")
	GithubClientSecret = os.Getenv("GITHUB_CLIENT_SECRET")
	GoogleClientID = os.Getenv("GOOGLE_CLIENT_ID")
	GoogleClientSecret = os.Getenv("GOOGLE_CLIENT_SECRET")

	return scanner.Err()
}
