// Copyright 2016 Mikio Hara. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package tcpopt

/*
#include <sys/socket.h>

#include <netinet/tcp.h>
*/
import "C"

const (
	sysSOL_SOCKET = C.SOL_SOCKET

	sysSO_KEEPALIVE = C.SO_KEEPALIVE
	sysSO_SNDBUF    = C.SO_SNDBUF
	sysSO_RCVBUF    = C.SO_RCVBUF

	sysTCP_NODELAY         = C.TCP_NODELAY
	sysTCP_KEEPALIVE       = C.TCP_KEEPALIVE
	sysTCP_KEEPINTVL       = C.TCP_KEEPINTVL
	sysTCP_KEEPCNT         = C.TCP_KEEPCNT
	sysTCP_CONNECTION_INFO = C.TCP_CONNECTION_INFO
	sysTCP_NOTSENT_LOWAT   = C.TCP_NOTSENT_LOWAT

	sysTCPCI_OPT_TIMESTAMPS           = C.TCPCI_OPT_TIMESTAMPS
	sysTCPCI_OPT_SACK                 = C.TCPCI_OPT_SACK
	sysTCPCI_OPT_WSCALE               = C.TCPCI_OPT_WSCALE
	sysTCPCI_OPT_ECN                  = C.TCPCI_OPT_ECN
	sysTCPCI_FLAG_LOSSRECOVERY        = C.TCPCI_FLAG_LOSSRECOVERY
	sysTCPCI_FLAG_REORDERING_DETECTED = C.TCPCI_FLAG_REORDERING_DETECTED

	sizeofTCPConnInfo = C.sizeof_struct_tcp_connection_info
)

type sysTCPConnInfo C.struct_tcp_connection_info