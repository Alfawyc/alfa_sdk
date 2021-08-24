package example

import (
	"alfa_sdk/api/sp/suggestedkeywords"
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetAdGroupSuggestedKeywords(t *testing.T) {
	t.Run("getAdGroupSuggestedKeywords", func(t *testing.T) {
		req := getSellerRequest(t, "")
		getAdGroupSuggestedKeywords := suggestedkeywords.NewGetAdGroupSuggestedKeyword().SetAdGroupID("126286185975014").SetMaxNumSuggestions(1).EnabledSuggestBids()

		result := make([]suggestedkeywords.AdgroupsSuggestedKeywords, 0)

		resp, err := req.Execute(context.Background(), getAdGroupSuggestedKeywords, &result)

		assert.Nil(t, err)
		assert.NotNil(t, resp)

		t.Logf("%v", resp)
		t.Logf("%v", result)
	})
}
