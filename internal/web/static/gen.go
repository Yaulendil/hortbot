// Package static contains the static HTTP resources served at /static/.
package static

//go:generate go run github.com/mjibson/esc -o=static.go -pkg=static -ignore=^(gen|static)\.go$ -modtime=0 .
