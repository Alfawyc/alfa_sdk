package token

import (
	"alfa_sdk/types"
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClient(t *testing.T) {
	const (
		clientID     = "amzn1.application-oa2-client.632a4d1db3f44fb0bb445f5cb86faa61"
		clientSecret = "c93e275bc3c8686e584b6532d9a529fefa3743221bda3229e71c7de9ca0bc2e1"
		RefreshToken = "Atzr|IwEBIF17gJ3UIgRhbk3yOEw5IMVZEPRVfyM9Wg87MpdhRM1u1egkEuEWJ2N2IxOCwXctAb4kMtzLsK_xgi_o1Geq8J8pi6-uURuKqud2WH8aZn-QdBHhxcjaD_JiGq1EC4XFrttztmDGDT9ErB8jeSopXYA_1BwJZSsv4VydBv4lHqM_1dazQ6Ae8iQ65u2mRkeQRMKBf1g31DEbfr6TvkrQQsFhVaolU74DQmtBSxWJRf2BHLQWM3rUsT4vEc6Pt0IXiBK63PQShmhfwu-uHY4WCl-teGD_yslxsDLUb-ghEL7j6R6nF-cx57hCaNw38k0G8I9CWS2S0ThmPRugT5Ya_t_GLAdnv6s_vCoocg06o50If-vjSLyDIornUWZsRcBVg68ltPSvX5QoleHZrYsRbx9mCx354EyAuUPIvvJXcplYQ8vPW1g3TCqMD5F1LBHKPi70JuUxpQY-51T7Oi_X5qOfTKJ5iHrRNWyHBeR4omNDZw"
		authCode     = "SRzSdCohWgykyDQciVyI"
	)
	ctx := context.Background()

	t.Run("TestNewClient", func(t *testing.T) {
		client := NewClient(&Config{
			ClientID:     clientID,
			ClientSecret: clientSecret,
		})
		assert.NotNil(t, client)

		t.Run("TestExchange", func(t *testing.T) {
			token, err := client.Exchange(ctx, authCode, "https://api.oovp.com/v1/authorize/ad/callback")
			assert.Nil(t, err)
			assert.NotNil(t, token)
			printToken(t, token)
			t.Run("TestGetAccessToken", func(t *testing.T) {
				token2, err := client.GetAccessToken(ctx, token.RefreshToken)
				assert.Nil(t, err)
				assert.NotNil(t, token2)
				printToken(t, token2)
				t.Run("TestGetAccessToken from cache", func(t *testing.T) {
					token3, err := client.GetAccessToken(ctx, token2.RefreshToken)
					assert.Nil(t, err)
					assert.NotNil(t, token3)
					printToken(t, token3)
					assert.Equal(t, token3.AccessToken, token2.AccessToken)
					assert.Equal(t, token3.RefreshToken, token2.RefreshToken)
					assert.Equal(t, token3.Expiry, token2.Expiry)
				})
			})
		})
	})

}

func printToken(t *testing.T, token *types.Token) {
	t.Logf("refresh_token = %s", token.RefreshToken)
	t.Logf("access_token = %s", token.AccessToken)
	t.Logf("token_type = %s", token.TokenType)
	t.Logf("expiry = %s", token.Expiry)
}
