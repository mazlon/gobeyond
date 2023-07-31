package auth

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAuth(t *testing.T) {
	testCases := []struct {
		userID string
	}{
		{
			userID: `maz`,
		},
	}

	for _, test := range testCases {
		t.Run(test.userID, func(t *testing.T) {
			token, err := GenerateToken(test.userID)
			assert.NoError(t, err)
			tokenClaim, err := TokenVerifier(token)
			assert.NoError(t, err)
			assert.Equal(t, test.userID, tokenClaim)

		})
	}
}
