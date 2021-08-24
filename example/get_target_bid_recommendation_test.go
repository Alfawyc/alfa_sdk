package example

import (
	"alfa_sdk/api/sp/binrecommendation"
	"context"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestGetTargetBidRecommendation(t *testing.T) {
	t.Run("targetBidRecommendation", func(t *testing.T) {
		req := getSellerRequest(t, ScopeId)

		targetExpression := &binrecommendation.TargetExpression{
			Value: "",
			Type:  "queryHighRelMatches",
		}
		targetExpressions := make([]*binrecommendation.TargetExpression, 0)
		targetExpressions = append(targetExpressions, targetExpression)

		expressions := make([][]*binrecommendation.TargetExpression, 0)
		expressions = append(expressions, targetExpressions)

		body := &binrecommendation.TargetBindRecommendationReq{
			AdGroupID:   "32654599095432",
			Expressions: expressions,
		}
		str, _ := json.Marshal(body)
		log.Println(string(str))

		targetBidRecommendation := binrecommendation.NewGetTargetBidRecommendation().SetBody(body)

		result := binrecommendation.TargetBidRecommendationResp{}

		res, err := req.Execute(context.Background(), targetBidRecommendation, &result)

		assert.Nil(t, err)
		assert.NotNil(t, res)

		t.Logf("%v", res)
		t.Logf("%v", result)
	})
}
