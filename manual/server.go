package manual

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/tmvrus/grpc-validation/manual/api"
)

type Server struct {
	api.UnimplementedCarShopServer
}

func (Server) GetCarsForSale(_ context.Context, req *api.GetCarsForSaleRequest) (*api.GetCarsForSaleResponse, error) {
	if err := validate(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return nil, nil
}

func validate(r *api.GetCarsForSaleRequest) error {
	if r.YearFrom < 1980 {
		return fmt.Errorf("invalid YearFrom")
	}
	if r.YearTo > uint32(time.Now().Year()) {
		return fmt.Errorf("YearTo in future")
	}
	if r.YearFrom > r.YearTo {
		return fmt.Errorf("YearFrom cant be more than YearTo")
	}
	if len(r.Color) == 0 || len(r.Color) > 4 {
		return fmt.Errorf("invalid Color set")
	}

	validColor := map[string]struct{}{
		"red":   {},
		"green": {},
		"black": {},
		"white": {},
	}
	uniqueColor := map[string]struct{}{}
	for _, c := range r.Color {
		_, ok := uniqueColor[c]
		if ok {
			return fmt.Errorf("not uniq color: %s", c)
		}
		_, ok = validColor[c]
		if !ok {
			return fmt.Errorf("invalid color value: %s", c)
		}
		uniqueColor[c] = struct{}{}
	}

	if r.MileageFrom > r.MileageTo {
		return fmt.Errorf("invalid MileageTo: %d", r.MileageTo)
	}
	if r.MileageTo > 1000000 {
		return fmt.Errorf("invalid MileageTo: %d", r.MileageTo)
	}
	return nil
}
