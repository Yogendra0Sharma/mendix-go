package mendix_test

import (
	"context"
	"testing"

	. "github.com/yogendra0sharma/mendix-go-sdk"
)

const (
	testAPIToken = "7fC3atWGdK7u9g8oNF9GpokfyJRHSbauBk2xWBRZ754y9xF4UyeeMFawbVo7s9XKzDdcVELgSGY5hpPukNSyY9C69h61xr9CGafo"
	apiId        = "c5fc8942-813a-49c2-a222-3cbec0ae4ebd"
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
