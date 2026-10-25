# spootify

[![CI](https://github.com/FrankKair/spootify/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/FrankKair/spootify/actions/workflows/ci.yml)

`spootify` fetches information / lyrics about the album / song you're currently listening to.

Requires **macOS** and the **Spotify desktop app**.

## Install

With Go 1.25+:

```bash
go install github.com/FrankKair/spootify@latest
```

Ensure Go's binary directory (usually `~/go/bin`) is on your `PATH`.

Without Go, download the [latest binary](https://github.com/FrankKair/spootify/releases/latest): `darwin-arm64` for Apple Silicon or `darwin-x64` for Intel. Rename it to `spootify`, run `chmod +x spootify`, and place it in a directory on your `PATH`.

## Setup

For album information, get a [Last.fm API key](https://www.last.fm/api/account/create) and add it to your shell configuration:

```bash
export LASTFM_API_KEY="your-key-here"
```

Lyrics work without a key, via [LRCLIB](https://lrclib.net).

## Usage

Open Spotify and play a song, then run:

```bash
spootify            # show album info + lyrics
spootify -lyrics    # show only lyrics
spootify -info      # show only album info
```

## Build from source

```bash
git clone https://github.com/FrankKair/spootify.git
cd spootify
make build
./spootify
```

Use `make test` to run tests or `make install` to install your checkout.

### Motivation

I enjoy learning about the context and stories behind the music I listen to.

### Inspiration

[shpotify](https://github.com/hnarayanan/shpotify)'s usage of AppleScript to control the Spotify macOS client made me realize how useful AppleScript can be.
