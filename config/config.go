package config

import "os"

func DSN() string {
	return os.Getenv("DB_PATH")
}

func Port() string {
	return ":" + os.Getenv("PORT")
}

func Arrow() string {
	arrow := os.Getenv("ARROW")
	if arrow  == "" {
		return "http://127.0.0.1`"
	}
	return arrow
}
