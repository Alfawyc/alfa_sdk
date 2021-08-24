package example

import (
	amz_ad_sdk "alfa_sdk"
	"alfa_sdk/api/profiles"
	"context"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

var _req *amz_ad_sdk.Request

func TestProfiles(t *testing.T) {
	t.Run("getReq", func(t *testing.T) {
		req := getSellerRequest(t, "")
		assert.NotNil(t, req)
		t.Run("getProfiles", func(t *testing.T) {
			getProfiles := profiles.NewGetProfiles()
			result := make([]profiles.Profiles, 0)
			resp, err := req.Execute(context.Background(), getProfiles, &result)
			assert.NoError(t, err)
			assert.NotNil(t, resp)
			t.Logf("%v", resp)
			log.Println("profiles result", result)
		})
	})
}
