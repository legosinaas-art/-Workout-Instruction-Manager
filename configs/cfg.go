package configs

import (
	"errors"
	"net"
	"net/url"
	"os"
)

func DNSFromEnv() (string, error) {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")
	ssl := os.Getenv("DB_SSLMODE")
	if ssl == "" {
		ssl = "disable"
	}
	if host == "" || port == "" || user == "" || name == "" {
		return "", errors.New("missing DB env: need DB_HOST, DB_PORT, DB_USER, DB_NAME (or DB_DSN)")
	}
	u := &url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(user, pass),
		Host:   net.JoinHostPort(host, port),
		Path:   "/" + name,
	}
	q := u.Query()
	q.Set("sslmode", ssl)
	u.RawQuery = q.Encode()
	return u.String(), nil
}
