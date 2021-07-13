package Implemention

import "os"

func GetHost() string {
	host := os.Getenv("JIZHANG_TEST_API_HOST")
	if len(host) == 0 {
		return "http://localhost"
	}
	return host
}
