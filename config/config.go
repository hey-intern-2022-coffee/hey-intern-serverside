package config

import "os"

func DSN() string {
	return os.Getenv("DB_PATH")
}

func Port() string {
	port := os.Getenv("PORT")
	if port == "" {
		return ":8080"
	}
	return ":" + port
}

func Arrow() string {
	arrow := os.Getenv("ARROW")
	if arrow == "" {
		return "http://127.0.0.1`"
	}
	return arrow
}
