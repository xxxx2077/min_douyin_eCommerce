package service

import (
	"context"
	"testing"
	order "github.com/xxxx2077/min_douyin_eCommerce/rpc_gen/kitex_gen/order"
)

func TestListOrder_Run(t *testing.T) {
	ctx := context.Background()
	s := NewListOrderService(ctx)
	// init req and assert value

	req := &order.ListOrderReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
