# GPM go package manager tool

## Like npm but in go

This tool is a simple and lightweight Go package manager, inspired by Node.js npm. It can find and install Go packages. Written entirely in Go with no external dependencies, it provides you with the ability to search pkg.go.dev without leaving your terminal. Additionally, it solves the issue of remembering the full path to a package. After the installation is finished, GMP will return the full path, allowing you to easily copy and paste it into your import block.

## Installation

Go to github repository [github repo](https://github.com/vladevelops/gpm) and clone the repo to install **gpm**.

```bash
cd gpm
go install
```

    This will install a gpm build in your GOPATH/bin directory

## Usage

```bash
gpm <command> value

```

The commands are:

|          Command          | Description                                                                                                                                                                                                                                                                                 |
| :-----------------------: | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
|     **`i <install>`**     | The process will start with the installation of the first package in the parsed list.                                                                                                                                                                                                       |
| **`is <search install>`** | This command enables you to search packages and grants you the flexibility to handpick which package you'd like to install. It empowers you with the ability to browse through the search results and make informed decisions about the installation of packages that best suit your needs. |
|    **`c <capacity>`**     | Max n for serch results, only to use with **`is`** command ex: `gpm is websocket  c 6`                                                                                                                                                                                                      |

Here some examples

simple [install]:

```bash
gpm i websocket

returns
Usage : import "github.com/gorilla/websocket"
```

gpm is [search install]

```bash
gpm is websocket
returns
[1] Name:  github.com/gorilla/websocket
Description: Package websocket implements the WebSocket protocol defined in RFC 6455.

[2] Name:  golang.org/x/net/websocket
Description: Package websocket implements a client and server for the WebSocket protocol as specified in RFC 6455.

[3] Name:  nhooyr.io/websocket
Description: Package websocket implements the RFC 6455 WebSocket protocol.

[4] Name:  v2ray.com/core/transport/internet/websocket
Description: Package websocket implements Websocket transport Websocket transport implements an HTTP(S) compliable, surveillance proof transport method with plausible deniability.
```

And after you provide which package to `go get`, it will be installed, and the full path will be returned.

---

## Development

Pull request are more then wellcome.

Please feel free to add fetures

## License

MIT License

Copyright (c) 2023 Vladyslav Topyekha

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
