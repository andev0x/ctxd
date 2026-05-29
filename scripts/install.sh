#!/usr/bin/env bash
set -euo pipefail

if ! command -v go >/dev/null 2>&1; then
  echo "Go is required to install lea. Visit https://go.dev/dl/ to install Go." >&2
  exit 1
fi

echo "Installing lea with go install..."
go install github.com/PizenLabs/lea/cmd/lea@latest

LEA_BIN="$(go env GOPATH)/bin/lea"
if [ -x "$LEA_BIN" ]; then
  echo "lea installed: $($LEA_BIN version)"
  echo "Ensure $(go env GOPATH)/bin is in your PATH."
else
  echo "lea installed to $(go env GOPATH)/bin. Ensure it is in your PATH."
fi
