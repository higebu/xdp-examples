# XDP Example

## Setup

```
sudo apt install libbpfcc libbpfcc-dev bpfcc-tools
go get github.com/higebu/xdp-example/cmd/checkmod
```

## Verify module

```
sudo cmd/checkmod/checkmod -src src/xdp_example.c
```
