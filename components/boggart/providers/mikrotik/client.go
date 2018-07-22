package mikrotik

import (
	"errors"
	"sync"
	"time"

	"gopkg.in/routeros.v2"
)

type Client struct {
	mutex  sync.Mutex
	client *routeros.Client
}

func NewClient(address, username, password string, timeout time.Duration) (*Client, error) {
	client, err := routeros.DialTimeout(address, username, password, timeout)
	if err != nil {
		return nil, err
	}

	return &Client{
		client: client,
	}, nil
}

func (c *Client) runArgs(sentence []string) (*routeros.Reply, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	return c.client.RunArgs(sentence)
}

func (c *Client) SystemRouterboard() (map[string]string, error) {
	reply, err := c.runArgs([]string{"/system/routerboard/print"})
	if err != nil {
		return nil, nil
	}

	if len(reply.Re) == 0 {
		return nil, errors.New("Empty reply from device")
	}

	return reply.Re[0].Map, nil
}

func (c *Client) SystemResource() (map[string]string, error) {
	reply, err := c.runArgs([]string{"/system/resource/print"})
	if err != nil {
		return nil, nil
	}

	if len(reply.Re) == 0 {
		return nil, errors.New("Empty reply from device")
	}

	return reply.Re[0].Map, nil
}

func (c *Client) WifiClients() ([]map[string]string, error) {
	reply, err := c.runArgs([]string{"/interface/wireless/registration-table/print"})
	if err != nil {
		return nil, err
	}

	if len(reply.Re) == 0 {
		return nil, errors.New("Empty reply from device")
	}

	clients := make([]map[string]string, 0, len(reply.Re))

	for _, re := range reply.Re {
		clients = append(clients, re.Map)
	}

	return clients, nil
}

func (c *Client) PPPActiveConnections() ([]map[string]string, error) {
	reply, err := c.runArgs([]string{"/ppp/active/print"})
	if err != nil {
		return nil, err
	}

	if len(reply.Re) == 0 {
		return nil, errors.New("Empty reply from device")
	}

	clients := make([]map[string]string, 0, len(reply.Re))

	for _, re := range reply.Re {
		clients = append(clients, re.Map)
	}

	return clients, nil
}

func (c *Client) EthernetStats() ([]map[string]string, error) {
	reply, err := c.runArgs([]string{"/interface/print", "stats"})
	if err != nil {
		return nil, err
	}

	if len(reply.Re) == 0 {
		return nil, errors.New("Empty reply from device")
	}

	stats := make([]map[string]string, 0, len(reply.Re))

	for _, re := range reply.Re {
		stats = append(stats, re.Map)
	}

	return stats, nil
}

func (c *Client) DNSStatic() (map[string]string, error) {
	reply, err := c.runArgs([]string{"/ip/dns/static/print"})
	if err != nil {
		return nil, err
	}

	if len(reply.Re) == 0 {
		return nil, errors.New("Empty reply from device")
	}

	table := make(map[string]string, len(reply.Re))
	for _, re := range reply.Re {
		if re.Map["disabled"] == "true" {
			continue
		}

		table[re.Map["address"]] = re.Map["name"]
	}

	return table, nil
}

func (c *Client) ARPTable() (map[string]map[string]string, error) {
	reply, err := c.runArgs([]string{"/ip/arp/print"})
	if err != nil {
		return nil, err
	}

	if len(reply.Re) == 0 {
		return nil, errors.New("Empty reply from device")
	}

	table := make(map[string]map[string]string, len(reply.Re))
	for _, re := range reply.Re {
		table[re.Map["mac-address"]] = re.Map
	}

	return table, nil
}

func (c *Client) DHCPLease() (map[string]string, error) {
	reply, err := c.runArgs([]string{"/ip/dhcp-server/lease/print"})
	if err != nil {
		return nil, err
	}

	if len(reply.Re) == 0 {
		return nil, errors.New("Empty reply from device")
	}

	table := make(map[string]string, len(reply.Re))
	for _, re := range reply.Re {
		table[re.Map["mac-address"]] = re.Map["host-name"]
	}

	return table, nil
}
