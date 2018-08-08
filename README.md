## bear-cli

Interface with the [Bear](https://www.bear-writer.com) notetaking app from your command line. Unofficial - this tool is not endorsed or supported by Shiny Frog.

macOS only.

### Install


### Usage


### Developing
To build on OSX:

0) Install the XCode CLT

1) Install gcc-8

    brew install gcc

2) Build with these modes

    go build -o bear-cli -buildmode=c-shared -compiler gccgo main.go


### License
MIT

### Credits

This project contains the binary distribution of https://github.com/martinfinke/xcall, downloaded from https://github.com/martinfinke/xcall/releases/tag/v1.0.1.

Thanks to Martin Finke for this library, which enables `x-callback-url` communication with Bear.
