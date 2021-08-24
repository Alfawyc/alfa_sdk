package productad

import (
	"encoding/json"
	"strconv"
)

var (
	ProductAdURL         string = "/v2/sp/productAds"
	ProductAdURLExtended string = "/v2/sp/productAds/extended"
)

type ProductAd struct {
	AdId                 int64                  `json:"adId"`
	AdStringId           string                 `json:"ad_string_id"`
	CampaignId           int64                  `json:"campaignId"`
	CampaignStringId     string                 `json:"campaign_string_id"`
	AdGroupId            int64                  `json:"adGroupId"`
	AdGroupStringId      string                 `json:"ad_group_string_id"`
	Sku                  string                 `json:"sku"`
	Asin                 string                 `json:"asin"`
	State                string                 `json:"state"`
	CreationDate         int64                  `json:"creationDate"`
	LastUpdateDate       int64                  `json:"lastUpdatedDate"`
	ServingStatus        string                 `json:"servingStatus"`
	ServingStatusDetails []ServingStatusDetails `json:"servingStatusDetails"`
}

type ServingStatusDetails struct {
	Name     string `json:"name"`
	Severity string `json:"severity"`
	Message  string `json:"message"`
	HelpUrl  string `json:"helpUrl"`
}

// CreateProductAdReq create product ad req
type CreateProductAdReq struct {
	CampaignId int64 `json:"campaignId"`
	AdGroupId  int64 `json:"adGroupId"`
	//The SKU associated with the product. Defined for seller accounts only.
	Sku string `json:"sku"`
	//The ASIN associated with the product. Defined for vendors only.
	//Asin       string `json:"asin"`
	State string `json:"state"`
}

// UpdateProductAdReq update product ad resp
type UpdateProductAdReq struct {
	AdId  int64  `json:"adId"`
	State string `json:"state"`
}

// ProductAdResponse edit product ad response
type ProductAdResponse struct {
	AdId        int64  `json:"adId"`
	Code        string `json:"code"`
	Details     string `json:"details"`
	Description string `json:"description"`
}

func (p *ProductAd) UnmarshalJSON(data []byte) error {
	type xProductAd ProductAd
	x := &xProductAd{}
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	x.CampaignStringId = strconv.Itoa(int(x.CampaignId))
	x.AdGroupStringId = strconv.Itoa(int(x.AdGroupId))
	x.AdStringId = strconv.Itoa(int(x.AdId))
	*p = ProductAd(*x)

	return nil
}
