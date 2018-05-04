# Goffli
RoadRunner
==========
[![GoDoc](https://godoc.org/github.com/spiral/goffli?status.svg)](https://godoc.org/github.com/spiral/goffli)
[![Go Report Card](https://goreportcard.com/badge/github.com/spiral/goffli)](https://goreportcard.com/report/github.com/spiral/goffli)

Goffli is programmable and colorful ffmpeg wrapper with ability to share snippets using GitHub Gist.
> When you hate to google ffmpeg snippets.

# Installation
Make sure to [install Golang](https://golang.org/doc/install) at your machine.

```
go get github.com/spiral/goffli
```

You can also find banaries [here](https://github.com/spiral/goffli/releases).

# Usage
By default Goffli only able to display media informarmation about given file:

```
goffli info video.mp4
```

# Install/Update the snippet
In order to extend Goffli capabilities install *lua* script with desired ffmpeg options. 

```
goffli get [gist-url] [snippet-name]
```

Once installed you can evaluate snippet using it's name

```
goffli [snippet-name] [args]
```

# Running local scripts
You can also evaluate local lua snippet without downloading it from GitHub Gists.

```
goffli run snippet.lua [args]
```

# Available Snippets

LIST IS HERE

> Feel free to suggest your own snippet for the list.

# Snippet related operations
To get list of all available snippets run

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

License:
--------
The MIT License (MIT). Please see [`LICENSE`](./LICENSE) for more information.
