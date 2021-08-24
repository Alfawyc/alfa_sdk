package report

var (
	RequestReportURL string = "/v2/sp/{recordType}/report"
	GetReportURL     string = "/v2/reports/{reportId}"
)

type RequestBodyReq struct {
	StateFilter string `json:"stateFilter,omitempty"`

	//The type of campaign. Only required for asins report - don't use with other report types. sponsoredProducts
	CampaignType string `json:"campaign_type,omitempty"`

	//Dimension on which the report is segmented
	Segment string `json:"segment,omitempty"`

	//The date for which to retrieve the performance report in YYYYMMDD format.
	ReportDate string `json:"reportDate,omitempty"`

	//A comma-separated list of the metrics to be included in the report.
	Metrics string `json:"metrics,omitempty"`
}

// RequestReportResponse request report response
type RequestReportResponse struct {
	ReportID      string `json:"reportId,omitempty"`
	RecordType    string `json:"recordType,omitempty"`
	Status        string `json:"status,omitempty"`
	StatusDetails string `json:"statusDetails,omitempty"`
}

// GetReportResponse check report info response
type GetReportResponse struct {
	ReportId     string `json:"reportId"`
	Status       string `json:"status"`
	StatusDetail string `json:"statusDetails"`
	Location     string `json:"location"`
	FileSize     int64  `json:"fileSize"`
}

var ReportMetrics map[string][]string = map[string][]string{
	"campaigns":  []string{"campaignName", "campaignId", "campaignStatus", "campaignBudget", "campaignRuleBasedBudget", "applicableBudgetRuleId", "applicableBudgetRuleName", "impressions", "clicks", "cost", "attributedConversions1d", "attributedConversions7d", "attributedConversions14d", "attributedConversions30d", "attributedConversions1dSameSKU", "attributedConversions7dSameSKU", "attributedConversions14dSameSKU", "attributedConversions30dSameSKU", "attributedUnitsOrdered1d", "attributedUnitsOrdered7d", "attributedUnitsOrdered14d", "attributedUnitsOrdered30d", "attributedSales1d", "attributedSales7d", "attributedSales14d", "attributedSales30d", "attributedSales1dSameSKU", "attributedSales7dSameSKU", "attributedSales14dSameSKU", "attributedSales30dSameSKU", "attributedUnitsOrdered1dSameSKU", "attributedUnitsOrdered7dSameSKU", "attributedUnitsOrdered14dSameSKU", "attributedUnitsOrdered30dSameSKU"},
	"adGroups":   []string{"campaignName", "campaignId", "adGroupName", "adGroupId", "impressions", "clicks", "cost", "attributedConversions1d", "attributedConversions7d", "attributedConversions14d", "attributedConversions30d", "attributedConversions1dSameSKU", "attributedConversions7dSameSKU", "attributedConversions14dSameSKU", "attributedConversions30dSameSKU", "attributedUnitsOrdered1d", "attributedUnitsOrdered7d", "attributedUnitsOrdered14d", "attributedUnitsOrdered30d", "attributedSales1d", "attributedSales7d", "attributedSales14d", "attributedSales30d", "attributedSales1dSameSKU", "attributedSales7dSameSKU", "attributedSales14dSameSKU", "attributedSales30dSameSKU", "attributedUnitsOrdered1dSameSKU", "attributedUnitsOrdered7dSameSKU", "attributedUnitsOrdered14dSameSKU", "attributedUnitsOrdered30dSameSKU"},
	"keywords":   []string{"campaignName", "campaignId", "adGroupName", "adGroupId", "keywordId", "keywordText", "matchType", "impressions", "clicks", "cost", "attributedConversions1d", "attributedConversions7d", "attributedConversions14d", "attributedConversions30d", "attributedConversions1dSameSKU", "attributedConversions7dSameSKU", "attributedConversions14dSameSKU", "attributedConversions30dSameSKU", "attributedUnitsOrdered1d", "attributedUnitsOrdered7d", "attributedUnitsOrdered14d", "attributedUnitsOrdered30d", "attributedSales1d", "attributedSales7d", "attributedSales14d", "attributedSales30d", "attributedSales1dSameSKU", "attributedSales7dSameSKU", "attributedSales14dSameSKU", "attributedSales30dSameSKU", "attributedUnitsOrdered1dSameSKU", "attributedUnitsOrdered7dSameSKU", "attributedUnitsOrdered14dSameSKU", "attributedUnitsOrdered30dSameSKU"},
	"productAds": []string{"campaignName", "campaignId", "adGroupName", "adGroupId", "impressions", "clicks", "cost", "currency", "asin", "sku", "attributedConversions1d", "attributedConversions7d", "attributedConversions14d", "attributedConversions30d", "attributedConversions1dSameSKU", "attributedConversions7dSameSKU", "attributedConversions14dSameSKU", "attributedConversions30dSameSKU", "attributedUnitsOrdered1d", "attributedUnitsOrdered7d", "attributedUnitsOrdered14d", "attributedUnitsOrdered30d", "attributedSales1d", "attributedSales7d", "attributedSales14d", "attributedSales30d", "attributedSales1dSameSKU", "attributedSales7dSameSKU", "attributedSales14dSameSKU", "attributedSales30dSameSKU", "attributedUnitsOrdered1dSameSKU", "attributedUnitsOrdered7dSameSKU", "attributedUnitsOrdered14dSameSKU", "attributedUnitsOrdered30dSameSKU"},
	"targets":    []string{"campaignId", "adGroupId", "targetId", "impressions", "clicks", "cost", "attributedConversions1d", "attributedConversions7d", "attributedConversions14d", "attributedConversions30d", "attributedConversions1dSameSKU", "attributedConversions7dSameSKU", "attributedConversions14dSameSKU", "attributedConversions30dSameSKU", "attributedUnitsOrdered1d", "attributedUnitsOrdered7d", "attributedUnitsOrdered14d", "attributedUnitsOrdered30d", "attributedSales1d", "attributedSales7d", "attributedSales14d", "attributedSales30d", "attributedSales1dSameSKU", "attributedSales7dSameSKU", "attributedSales14dSameSKU", "attributedSales30dSameSKU", "attributedUnitsOrdered1dSameSKU", "attributedUnitsOrdered7dSameSKU", "attributedUnitsOrdered14dSameSKU", "attributedUnitsOrdered30dSameSKU"},
}
