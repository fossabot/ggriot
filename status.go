package ggriot

import "github.com/soowan/ggriot/models"

// ServerStatus returns the current service status for the requested region.
// This doesn't apply to rate limits, however we are still obeying the limit.
// This will change in the future possibly.
func ServerStatus(region string) (ss *models.ServerStatus, err error) {
	err = apiRequest("https://"+region+"."+Base+BaseStatus, &ss)
	if err != nil {
		return ss, err
	}

	return ss, nil
}
