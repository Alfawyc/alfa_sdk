package example

import (
	"alfa_sdk/api/profiles"
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

//SANDBOX ONLY
func TestCreatProfiles(t *testing.T) {
	t.Run("CreateProfiles", func(t *testing.T) {
		req := getSellerRequest(t, "")
		res := new(profiles.RegisterSellerProfilesResponse)

		createProfiles := profiles.NewCreateProfiles().SetRequestBody(&profiles.CreateProfilesReq{CountryCode: "US"})
		resp, err := req.Execute(context.Background(), createProfiles, &res)
		assert.Nil(t, err)
		assert.NotNil(t, resp)
		t.Logf("%v", resp)
		t.Logf("%v", res)
	})
}
