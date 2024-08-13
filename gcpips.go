package gcpips

import (
	"encoding/json"
	"errors"
	"net/http"
)

// GCP_RANGES_URL is the URL used to pull the IP range information.
const GCP_RANGES_URL = `https://www.gstatic.com/ipranges/cloud.json`

type GCPRanges struct {
	SyncToken    string     `json:"syncToken"`
	CreationTime string     `json:"creationTime"`
	Prefixes     []GCPPrefix `json:"prefixes"`
}

type GCPPrefix struct {
	IPv4Prefix string `json:"ipv4Prefix,omitempty"`
	IPv6Prefix string `json:"ipv6Prefix,omitempty"`
	Service    string `json:"service"`
	Scope      string `json:"scope"`
}

func download() (*GCPRanges, error) {
	request, err := http.Get(GCP_RANGES_URL)
	if err != nil {
		return nil, err
	}
	defer request.Body.Close()

	if request.StatusCode != http.StatusOK {
		return nil, errors.New("Expected 200 OK")
	}
	var ranges GCPRanges
	decoder := json.NewDecoder(request.Body)
	err = decoder.Decode(&ranges)
	return &ranges, err
}

// Get will download and return a GCPRanges struct representing the
// IP ranges.
func Get() (*GCPRanges, error) {
	return download()
}
