# HeidiSQL Password Decoder

A simple tool to decode HeidiSQL stored passwords. Built with Go + WebAssembly.

**Live:** https://anudeepd.github.io/heidisql-password-decoder/

## Usage

1. Open `index.html` in a browser (or visit the deployed site)
2. Enter an encoded HeidiSQL password (hex string)
3. Click **Decode** to reveal the password
4. Click **Copy** to copy the result

## Features

- Dark/Light theme toggle
- Clear error messages for invalid input
- Copy to clipboard
- Keyboard support (Enter to decode)

## Development

### Build

```bash
GOOS=js GOARCH=wasm go build -o decoder.wasm main.go
```

### Run locally

```bash
python3 -m http.server 8080
```

Then open `http://localhost:8080`

## Files

| File | Description |
|------|-------------|
| `main.go` | Go source code |
| `decoder.wasm` | Compiled WebAssembly module |
| `index.html` | User interface |
| `wasm_exec.js` | Go runtime for WebAssembly |

### Getting wasm_exec.js

This file comes bundled with Go. Copy it from your Go installation:

```bash
# Common paths (varies by OS/install)
cp "$(go env GOROOT)/lib/wasm/wasm_exec.js" .
# or
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .
```

## Deploy to GitHub Pages

1. Push all files to a GitHub repository
2. Go to **Settings → Pages**
3. Select **main** branch as source
4. Save and wait for deployment

## Credits

HeidiSQL password decoding algorithm based on:
- **Original:** https://gist.github.com/jpatters/4553139
- **Archived:** https://archive.softwareheritage.org/browse/content/sha1_git:037292fd4c6fe097c6c75246541fe647988775a0/?origin_url=https://gist.github.com/jpatters/4553139&path=HeidiDecode.js

**Vibe-coded** using the Big Pickle model via [OpenCode](https://opencode.ai)
