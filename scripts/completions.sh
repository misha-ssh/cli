#!/bin/sh
set -e
rm -rf completions
mkdir completions
for sh in bash zsh fish; do
	go run ./cmd/misha/main.go completion "$sh" >"completions/misha.$sh"
done