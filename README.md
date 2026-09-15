# spootify

[![CI](https://github.com/FrankKair/spootify/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/FrankKair/spootify/actions/workflows/ci.yml)

`spootify` fetches information / lyrics about the album / song you're currently listening to.

This utility is **macOS** only since it uses **AppleScript** to get information from Spotify.

### Setup

1. Get a free Last.fm API key at https://www.lastfm/api.account/create
2. Export it in your shell:
    ```bash
    export LASTFM_API_KEY="your-key-here"
    ```
    (add this to your `~/.zshrc` or `~/.bashrc` to persist it)

> **Note:** The Last.fm API is free for non-commercial use with no hard rate limit -- they just ask you to be reasonable (< 5 req/s). Lyrics come from [lrclib.net](https://www.lrclib.net), which requires no API key at all.

### Usage

Download the [binary release](https://github.com/FrankKair/spootify/releases) or clone the repo and install:

```bash
make install
```

Then call `spootify` on your terminal while Spotify is playing:

```bash
spootify            # show album info + lyrics
spootify -lyrics    # show only lyrics
spootify -info      # show only album info
spootify -version   # print version
```

### Motivation

I enjoy learning about the context and stories behind the music I listen to.

### Inspiration

[shpotify](https://github.com/hnarayanan/shpotify)'s usage of AppleScript to control the Spotify macOS client made me realize how useful AppleScript can be.
