Goffli
==========
[![Go Report Card](https://goreportcard.com/badge/github.com/spiral/goffli)](https://goreportcard.com/report/github.com/spiral/goffli)

Goffli is colorful and programmable ffmpeg wrapper with ability to share snippets using GitHub Gist.
> When you hate to google ffmpeg snippets.

# Installation
Make sure to [install Golang](https://golang.org/doc/install) at your machine.

```
go get "github.com/spiral/goffli"
```

You can also find banaries [here](https://github.com/spiral/goffli/releases).

# Usage
By default Goffli only able to display media information about given file:

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

* https://gist.github.com/wolfy-j/8009a8b3be1004d933e105494c64c372 - Copy media content from one container to another (by @wolfy-j)

> Feel free to share your own snippets.

# Snippet related operations
To get list of all available snippets

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

### Sample Snippet
```lua
--@Description: Copy media content from one container to another
--@Version: 1.0 <Apr 22, 2018>
--@Source: https://gist.github.com/wolfy-j/8009a8b3be1004d933e105494c64c372
--@Author: Wolfy-J <wolfy.jd@gmail.com>

local input = ask("Source file", "exists")
local output = ask("Output file")

if input == "" or output == "" then
        print("\n<red>Script error: </reset><red+hb>not enough arguments</reset>\n")
        return
end

require("ffmpeg").run({
    "-i", input,
    "-acodec", "copy",
    "-vcodec", "copy",
    "-y", output
})

print("<green+hb>Conversion complete!</reset>\n")
```

License:
--------
The MIT License (MIT). Please see [`LICENSE`](./LICENSE) for more information.
