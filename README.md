# GPM go package manager helper
## Like npm but in go
    
This tool is make to help you with downloading and instaling go packages more easily. Espired by the Nodejs npm it can instal and find go backages from the main repository.

## Installation

Go to gitub repository [git repo](https://github.com/bytesociety/gpm) and clone the repo to install **gpm**.

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

| Command  | Description                |
|:--------:|----------------------------|
|  **`i`** | install first finded package |
|  **`id`** | serch n packages and let you choose which to install |
|  **`mx`** | max length list for serch results, only to use with **`id`** command|

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