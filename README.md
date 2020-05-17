# wgctrl [![builds.sr.ht status](https://builds.sr.ht/~mdlayher/wgctrl-go.svg)](https://builds.sr.ht/~mdlayher/wgctrl-go?) [![GoDoc](https://godoc.org/github.com/k-vpn/wgctrl-go?status.svg)](https://godoc.org/github.com/k-vpn/wgctrl-go) [![Go Report Card](https://goreportcard.com/badge/github.com/k-vpn/wgctrl-go)](https://goreportcard.com/report/github.com/k-vpn/wgctrl-go)

Package `wgctrl` enables control of WireGuard devices on multiple platforms.

For more information on WireGuard, please see <https://www.wireguard.com/>.

MIT Licensed.

```text
go get github.com/k-vpn/wgctrl-go
```

## Overview

`wgctrl` can control multiple types of WireGuard devices, including:

- Linux kernel module devices, via generic netlink
- userspace devices (e.g. wireguard-go), via the userspace configuration protocol
  - both UNIX-like and Windows operating systems are supported
- **Experimental:** OpenBSD kernel module devices, via ioctl interface
  - See <https://git.zx2c4.com/wireguard-openbsd/about/> for details. Specify
  environment variable `WGCTRL_OPENBSD_KERNEL=1` to enable this interface.

As new operating systems add support for in-kernel WireGuard implementations,
this package should also be extended to support those native implementations.

If you are aware of any efforts on this front, please
[file an issue](https://github.com/WireGuard/wgctrl-go/issues/new).

This package implements WireGuard configuration protocol operations, enabling
the configuration of existing WireGuard devices. Operations such as creating
WireGuard devices, or applying IP addresses to those devices, are out of scope
for this package.
