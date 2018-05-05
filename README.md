Goffli
==========
[![Go Report Card](https://goreportcard.com/badge/github.com/spiral/goffli)](https://goreportcard.com/report/github.com/spiral/goffli)

Goffli is colorful and programmable (LUA) **FFmpeg CLI wrapper** with ability to share snippets over GitHub Gist.
> When you hate to google ffmpeg bash scripts.

# Installation
Make sure to [install Golang](https://golang.org/doc/install) at your machine.

```
go get "github.com/spiral/goffli"
```

You can also find binaries [here](https://github.com/spiral/goffli/releases).

# Usage
By default Goffli only able to display media information about given file:

```
goffli info video.mp4
```

In order to extend Goffli functionality load snippet using GitHub Gist url:

```
goffli get https://gist.github.com/wolfy-j/d4ece481eb8c9bd8a438967d77603ce7 video2gif
```

You can use this snippet immediatelly:

```
goffli video2gif input.mp4 result.gif
```

# Available Snippets

Snippet         | URL
----            | ---
copy            | https://gist.github.com/wolfy-j/8009a8b3be1004d933e105494c64c372
video2gif       | https://gist.github.com/wolfy-j/d4ece481eb8c9bd8a438967d77603ce7

> Feel free to share your own snippets.

# Snippet related operations
In order to extend Goffli capabilities install *lua* script with desired ffmpeg options. 

```
goffli get [gist-url] [snippet-name]
```

Once installed you can evaluate snippet using it's name

```
goffli [snippet-name] [args]
```

To get list of all installed snippets

```
goffli list
```

To remove snippet from Goffli

```
goffli remove [snippet-name]
```

To display content of the snippet

```
goffli snow [snippet-name]
```

### Running local snippets
You can also evaluate local lua snippet without downloading it from GitHub Gists.

```
goffli run snippet.lua [args]
```

# Coding the Snippet
Coding the snippet is easy, you can utilize set of functions embedded to Lua machine in order to make usage more user friendly.

### Metadata description
@TODO

### Input functions
@TODO

### Temp files and directories
@TODO

### FFmpeg functions
@TODO

License:
--------
The MIT License (MIT). Please see [`LICENSE`](./LICENSE) for more information.
