package manual

import (
	"testing"

	"github.com/tmvrus/grpc-validation/auto/api"
)

func Test_ServerValidate(t *testing.T) {
	req := &api.GetCarsForSaleRequest{
		YearFrom:    2000,
		YearTo:      2020,
		Color:       []string{"green", "green"},
		MileageFrom: 150,
		MileageTo:   300,
	}

	err := validate(req)
	if err == nil {
		t.Error("expected error")
	} else {
		t.Logf("got expected error: %s", err.Error())
	}
}
