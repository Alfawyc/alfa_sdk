package example

import (
	"alfa_sdk/api/sp/keywords"
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateKeyword(t *testing.T) {
	t.Run("createKeyword", func(t *testing.T) {
		req := getSellerRequest(t, SandBxoScopeId)
		list := make([]*keywords.CreateKeywordsReq, 0)
		list = append(list, &keywords.CreateKeywordsReq{
			CampaignId:  109238869710315,
			AdGroupId:   213079305122647,
			State:       "enabled",
			KeywordText: "create test",
			MatchType:   "exact",
			Bid:         0.02,
		})
		createKeywords := keywords.NewCreateKeywords().SetBody(list)
		res := make([]keywords.KeywordResponse, 0)
		resp, err := req.Execute(context.Background(), createKeywords, &res)
		assert.Nil(t, err)
		assert.NotNil(t, resp)
		t.Logf("%v", res)
		t.Logf("%v", resp)
	})
}
