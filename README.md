# 24-Hour-Hackerathon-Timer
a terminal script that will print random time with timestamp every hour for 24 hour, used to keep track of time in a 24 hour hackerathon

# Install
1. You will need to install go on your machine: https://go.dev/doc/install
2. Setup GOPATH

Add the following to your shell config
```bash
export PATH=${PATH}:$HOME/go/bin
```
More information: https://go.dev/wiki/GOPATH#gopath-variable

3. Install the binary
```bash
go install github.com/epicseven-cup/24-Hour-Hackerathon-Timer@latest 
```

There could be delays between the Goproxy and GitHub binarys, you can use the direct setup
```bash
GOPROXY=direct go install github.com/epicseven-cup/24-Hour-Hackerathon-Timer@latest
```

# How to use


```bash
 24-Hour-Hackerathon-Timer -h
```

# Example

```bash
24-Hour-Hackerathon-Timer #Generate a random word every hour
```