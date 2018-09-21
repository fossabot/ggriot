package ggriot

import "github.com/soowan/ggriot/models"

// GetServerStatus returns the current service status for the requested region.
// This doesn't apply to rate limits, however we are still obeying the limit.
// This will change in the future possibly.
func GetServerStatus(region string) (ss *models.ServerStatus, err error) {
	err = apiRequest("https://"+region+"."+Base+BaseStatus, &ss)
	if err != nil {
		return nil, err
	}

	return ss, nil
}
