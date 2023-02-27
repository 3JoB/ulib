package doh

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"time"

	"github.com/miekg/dns"
	ucn "github.com/3JoB/ulib/net/client"
)

type DialFunc func(c context.Context, network string, address string) (net.Conn, error)

type DOHServer string

const (
	Alicloud DOHServer = "https://dns.alidns.com/dns-query"
	Cloudflare DOHServer = "https://cloudflare-dns.com/dns-query"
	Google     DOHServer = "https://dns.google/dns-query"
	Quad9      DOHServer = "https://dns.quad9.net/dns-query"
	OpenDNS    DOHServer = "https://doh.opendns.com/dns-query"
	AdGuard    DOHServer = "https://dns.adguard.com/dns-query"
)


func Dialer(addr DOHServer) (*net.Dialer, error) {
	addrs := string(addr)
	fmt.Println("init")
	dohURL, err := url.Parse(addrs)
	if err != nil {
		return nil, err
	}
	dnsClient := http.DefaultClient
	//dnsClient.Timeout = 10 * 1000000000
	s := &net.Dialer{
		//Timeout:   10 * 1000000000,
		KeepAlive: 0,
		Resolver: &net.Resolver{
			PreferGo: true,
			Dial:     parse(addrs, dohURL, dnsClient),
		},
	}
	return s, nil
}

func parse(addr string, ns *url.URL, dnsClient *http.Client) DialFunc {
	return func(c context.Context, network, address string) (net.Conn, error) {
		//server := ns.Hostname()
		query := url.Values{}
		query.Set("dns", address)
		query.Set("type", "A")
		ns.RawQuery = query.Encode()

		fmt.Printf("do: %v\n", ns.RawQuery)

		req, err := http.NewRequestWithContext(c, http.MethodGet, ns.String(), nil)
		if err != nil {
			return nil, err
		}

		resp, err := dnsClient.Do(req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		var msg dns.Msg
		ds := ucn.UnPackData(resp)
		fmt.Println(ds.String())
		if err := msg.Unpack(ds.Bytes()); err != nil {
			return nil, err
		}
		if len(msg.Answer) > 0 {
			aRecord, ok := msg.Answer[0].(*dns.A)
			if ok {
				ipAddress := aRecord.A.String()
				conn, err := net.DialTimeout(network, ipAddress+":53", 5*time.Second)
				if err != nil {
					return nil, err
				}
				return conn, nil
			}
		}

		return nil, fmt.Errorf("no valid response found for %s", address)
	}
}