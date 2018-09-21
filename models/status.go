package models

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
