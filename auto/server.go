package manual

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/tmvrus/grpc-validation/auto/api"
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
	if err := r.Validate(); err != nil {
		return fmt.Errorf("auto validation error: %s", err.Error())
	}
	if r.YearFrom > r.YearTo {
		return fmt.Errorf("YearFrom cant be more than YearTo")
	}
	if r.YearTo > uint32(time.Now().Year()) {
		return fmt.Errorf("YearTo in future")
	}
	if r.MileageFrom > r.MileageTo {
		return fmt.Errorf("invalid MileageTo: %d", r.MileageTo)
	}
	return nil
}
