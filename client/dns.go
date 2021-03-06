package client

type DnsRecord struct {
	Address  string `json:"ip_addr,omitempty"`
	HostName string `json:"host,omitempty"`
}

func (d *DnsRecord) Query() string {
	return "ip_addr=" + d.Address + "&host=" + d.HostName
}

// Create creates a dns record
func (c *Client) DnsCreate(address, hostname string) (*DnsRecord, error) {
	dns := &DnsRecord{
		Address:  address,
		HostName: hostname,
	}
	if err := c.doRequest(dns.Query()); err != nil {
		return nil, err
	}
	return dns, nil
}

// Delete deletes a dns record
func (c *Client) DnsDelete(hostname string) (*DnsRecord, error) {
	dns := &DnsRecord{
		Address:  "delete",
		HostName: hostname,
	}
	if err := c.doRequest(dns.Query()); err != nil {
		return nil, err
	}
	return dns, nil
}
