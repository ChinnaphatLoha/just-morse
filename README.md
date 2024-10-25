# Morse CLI

A simple CLI application to convert text to Morse code and translate Morse code to text.

## Installation

1. **Download the latest binary** for your operating system from the [Releases](https://github.com/ChinnaphatLoha/just-morse/releases) page.
2. **Make the binary executable** (for Linux/macOS) and move it to a directory in your PATH.

   - **Linux/macOS**:

     ```bash
     chmod +x morse-linux   # or morse-mac
     sudo mv morse-linux /usr/local/bin/morse
     ```

   - **Windows**:
     - Rename the file to `morse.exe` if needed.
     - Add `morse.exe` to your system’s PATH, or run it directly.

3. **Run the Program**:

   - To convert text to Morse code:

     ```bash
     morse "Hello World"
     ```

   - To translate Morse code to text:

     ```bash
     morse --translate ".... . .-.. .-.. --- / .-- --- .-. .-.. -.."
     ```

## Building from Source

If you’d like to build the binaries yourself, clone this repository and run:

```bash
# Linux
GOOS=linux GOARCH=amd64 go build -o build/morse-linux ./cmd/morse

# macOS
GOOS=darwin GOARCH=amd64 go build -o build/morse-mac ./cmd/morse

# Windows
GOOS=windows GOARCH=amd64 go build -o build/morse-windows.exe ./cmd/morse
```
