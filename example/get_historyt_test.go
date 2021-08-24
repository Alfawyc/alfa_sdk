package example

import (
	"alfa_sdk/api/history"
	"context"
	"encoding/json"
	"log"
	"testing"
)

func TestGetHistory(t *testing.T) {
	ctx := context.Background()
	t.Run("getHistory", func(t *testing.T) {
		req := getSellerRequest(t, ScopeId)
		var adParents []history.AdParents
		adParents = append(adParents, history.AdParents{
			//CampaignId:             "",
			UseProfileIdAdvertiser: true,
			AdGroupId:              "14915352913755",
		})
		requestBody := history.HistoryRequestBody{
			FromDate: 1626095902,
			ToDate:   1628774302,
			Count:    50,
			EventType: history.EventType{
				AD: history.AdFilter{
					Filters:      []string{"DEFAULT_BID_AMOUNT"},
					EventTypeIds: []string{"14915352913755"},
					Parents:      adParents,
				},
			},
		}
		str, _ := json.Marshal(requestBody)
		log.Println(string(str))
		getHistory := history.NewGetHistory()

		res, err := req.Execute(ctx, getHistory, nil)
		if err != nil {
			log.Fatalln("err:", err.Error())
		}
		log.Println(res)
	})
}
