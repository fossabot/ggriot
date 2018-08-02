package ggriot

// ServerStatus is returned with status for the region requested.
type ServerStatus struct {
	Name      string `json:"name"`
	RegionTag string `json:"region_tag"`
	Hostname  string `json:"hostname"`
	Services  []struct {
		Status    string        `json:"status"`
		Incidents []interface{} `json:"incidents"`
		Name      string        `json:"name"`
		Slug      string        `json:"slug"`
	} `json:"services"`
	Slug    string   `json:"slug"`
	Locales []string `json:"locales"`
}

// GetServerStatus returns the current service status for the requested region.
// This doesn't apply to rate limits, however we are still obeying the limit.
// This will change in the future possibly.
func GetServerStatus(region string) (ss *ServerStatus, err error) {
	err = apiRequest("https://"+region+"."+Base+BaseStatus, &ss)
	if err != nil {
		return nil, err
	}

	return ss, nil
}
