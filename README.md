# Hex Viewer CLI

A command-line tool written in Go to display the content of a file as a hex dump, including the offset, hexadecimal data, and their ASCII equivalent.

## Features
- Display the offset in hexadecimal (8 digits).
- Read the file in configurable block sizes (default is 16 bytes).
- Show data in hexadecimal format.
- Display readable ASCII characters; all others are replaced by a `.`.
- Supports both binary and text files.

## How to use

1. **Build the project**
   ```go build .```
2. **Run the program**
   ```./cli filename```
3. **Enter block size** (optional, default is 16):
- The program will ask for the block size (in bytes) at startup.

## Example output
```
00000000 7061636b616765206d61696e0a |package main.|
00000010 0a696d706f727420280a202022 | package main.|
```

## Project structure

- `main.go` : Program entry point.
- Function `PrintHexWithASCII` : Displays each block with offset, hex, and ASCII text.

## Requirements

- Go 1.19 or higher.
- A file to analyze.
