package mendix_test

import (
	"context"
	"testing"

	. "github.com/yogendra0sharma/mendix-go-sdk"
)

const (
	testAPIToken = ""
	apiId        = ""
)

func TestAPI(t *testing.T) {
	apiToken := testAPIToken
	var err error
	c := NewClientWithPATKey(apiToken)
	ctx := context.Background()
	_, err = c.GetRepoInfo(ctx, apiId)
	if err != nil {
		t.Fatalf("error: %v", err)
	}

}
