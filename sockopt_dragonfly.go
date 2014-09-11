// Copyright 2014 Mikio Hara. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build go1.3

package tcp

import (
	"os"
	"syscall"
	"time"
)

func (c *Conn) setKeepAliveIdleInterval(d time.Duration) error {
	fd, err := c.sysfd()
	if err != nil {
		return err
	}
	d += (time.Millisecond - time.Nanosecond)
	msecs := int(d / time.Millisecond)
	return os.NewSyscallError("setsockopt", syscall.SetsockoptInt(fd, ianaProtocolTCP, sysTCP_KEEPIDLE, msecs))
}

func (c *Conn) setKeepAliveProbeInterval(d time.Duration) error {
	fd, err := c.sysfd()
	if err != nil {
		return err
	}
	d += (time.Millisecond - time.Nanosecond)
	msecs := int(d / time.Millisecond)
	return os.NewSyscallError("setsockopt", syscall.SetsockoptInt(fd, ianaProtocolTCP, sysTCP_KEEPINTVL, msecs))
}

func (c *Conn) setKeepAliveProbes(n int) error {
	fd, err := c.sysfd()
	if err != nil {
		return err
	}
	return os.NewSyscallError("setsockopt", syscall.SetsockoptInt(fd, ianaProtocolTCP, sysTCP_KEEPCNT, n))
}

func (c *Conn) setCork(on bool) error {
	fd, err := c.sysfd()
	if err != nil {
		return err
	}
	return os.NewSyscallError("setsockopt", syscall.SetsockoptInt(fd, ianaProtocolTCP, sysTCP_NOPUSH, boolint(on)))
}

func (c *Conn) info() (*Info, error) {
	return nil, errOpNoSupport
}
