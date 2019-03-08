# spt-album-info

`spt-album-info` uses [Last.fm](https://www.last.fm/) wiki section to fetch information about the album you're currently listening to.

This utility is **macOS** only since it uses **AppleScript** to get information about Spotify.

```
go build
./spt-album-info
```

```
https://www.last.fm/music/Led+Zeppelin/Led+Zeppelin+II/+wiki
Led Zeppelin II is the second studio album by English rock band Led Zeppelin...
```

### Motivation

I enjoy learning about the context and stories behind the music I listen to.

### Inspiration

[shpotify](https://github.com/hnarayanan/shpotify)'s usage of AppleScript to control the Spotify macOS client made me realize how useful AppleScript can be.
