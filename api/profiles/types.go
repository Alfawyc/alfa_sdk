package profiles

import (
	"encoding/json"
	"strconv"
)

var GetProfilesURL string = "/v2/profiles"

// Profiles ...
type Profiles struct {
	ProfileId       int64   `json:"profileId"`
	ProfileStringId string  `json:"profile_string_id"`
	CountryCode     string  `json:"countryCode"`
	CurrencyCode    string  `json:"currencyCode"`
	DailyBudget     float64 `json:"dailyBudget"`
	TimeZone        string  `json:"timezone"`
	AccountInfo     AccountInfo
}

// AccountInfo ...
type AccountInfo struct {
	MarketplaceStringId string `json:"marketplaceStringId"`
	Id                  string `json:"id"`
	Type                string `json:"type"`
	Name                string `json:"name"`
	SubType             string `json:"subType"`
	ValidPaymentMethod  bool   `json:"validPaymentMethod"`
}

// RegisterSellerProfilesResponse create profiles request response
type RegisterSellerProfilesResponse struct {
	Status        string `json:"status"`
	StatusDetails string `json:"statusDetails"`
	ProfileId     int64  `json:"profileId"`
}

// CreateProfilesReq Create a seller profile for sandbox. SANDBOX ONLY
type CreateProfilesReq struct {
	CountryCode string `json:"countryCode"`
}

func (p *Profiles) UnmarshalJSON(data []byte) error {
	type xProfiles Profiles
	x := &xProfiles{}
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	x.ProfileStringId = strconv.Itoa(int(x.ProfileId))
	*p = Profiles(*x)

	return nil
}
