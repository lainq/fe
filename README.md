<div align="center">
  <h1><code>Fe</code></h1>
</div>

- ![What is it?](https://github.com/pranavbaburaj/fe/blob/main/README.md#what-is-fe-)
- ![Video demo](https://github.com/pranavbaburaj/fe/blob/main/README.md#video-demo)
- ![Installation](https://github.com/pranavbaburaj/fe/blob/main/README.md#installation)

## What is `Fe` ?
Fe is a simple tool which you can use to move through your directories from the command line

## Video demo

<!-- [https://user-images.githubusercontent.com/70764593/147730564-f0db7b9a-c91f-4a71-86df-60954e3cebc5.mp4 -->
[![asciicast]()](https://user-images.githubusercontent.com/70764593/147730564-f0db7b9a-c91f-4a71-86df-60954e3cebc5.mp4)


## Installation

### Normal installation
Binaries of the package can be collected from the [releases](https://github.com/pranavbaburaj/fe/releases/) page

### Build from source

Inorder to build from source, make sure to have go and git installed.

```sh
sudo apt update
sudo apt install go git
```

Now, clone the repository and run `go build`

```sh
# clone the repository
git clone https://github.com/pranavbaburaj/fe.git
cd fe

# Install the dependencies
go get

# Build it into a binary
go build
```
