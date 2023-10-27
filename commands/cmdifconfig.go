package commands

import (
	"fmt"
)

type cmdIfconfig struct{}

const result = `eth0: flags=4163<UP,BROADCAST,RUNNING,MULTICAST>  mtu 1500
	inet 124.33.11.22  netmask 255.255.255.255  broadcast 0.0.0.0
	inet6 fe80::9420:2ef:aea6:32a5  prefixlen 64  scopeid 0x20<link>
	inet6 2a02:5f9:c2c:f8f9::1  prefixlen 64  scopeid 0x0<global>
	ether 91:10:12:a5:32:a5  txqueuelen 1000  (Ethernet)
	RX packets 125604  bytes 36275519 (36.2 MB)
	RX errors 0  dropped 0  overruns 0  frame 0
	TX packets 69063  bytes 24399449 (24.3 MB)
	TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0
`

func (cmdIfconfig) execute(context CommandContext) (uint32, error) {
	_, err := fmt.Fprint(context.stdout, result)
	return 0, err
}
