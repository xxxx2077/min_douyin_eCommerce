package service

import (
	"context"
	"testing"
	product "github.com/xxxx2077/min_douyin_eCommerce/rpc_gen/kitex_gen/product"
)

func TestGetProduct_Run(t *testing.T) {
	ctx := context.Background()
	s := NewGetProductService(ctx)
	// init req and assert value

	req := &product.GetProductReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
