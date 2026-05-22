## Build

```sh
# Linux
GOOS=linux GOARCH=amd64 go build -o portfolio-cli-linux ./cmd

# macOS (Intel)
GOOS=darwin GOARCH=amd64 go build -o portfolio-cli-macos ./cmd

# macOS (Apple Silicon)
GOOS=darwin GOARCH=arm64 go build -o portfolio-cli-macos-arm64 ./cmd

# Windows
GOOS=windows GOARCH=amd64 go build -o portfolio-cli.exe ./cmd
```
