package shadowclient

import (
	"net/url"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/chesiver/snappy/core"
)

var config struct {
	Verbose    bool
	UDPTimeout time.Duration
	TCPCork    bool
}

func Start(client, cipher, password, socks string) error {
	config.Verbose = true
	var key []byte
	if client != "" { // client mode
		addr := client
		cipher := cipher
		password := password
		var err error
		if strings.HasPrefix(addr, "ss://") {
			addr, cipher, password, err = parseURL(addr)
			if err != nil {
				return err
			}
		}
		ciph, err := core.PickCipher(cipher, key, password)
		if err != nil {
			return err
		}
		if socks != "" {
			go socksLocal(socks, addr, ciph.StreamConn)
		}
	}
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh
	return nil
}

func parseURL(s string) (addr, cipher, password string, err error) {
	u, err := url.Parse(s)
	if err != nil {
		return
	}
	addr = u.Host
	if u.User != nil {
		cipher = u.User.Username()
		password, _ = u.User.Password()
	}
	return
}
