package blockchain

import (
	commonProtoc "github.com/irisnet/blockchain-rpc/codegen/server"
	"github.com/irisnet/irishub-server/services"
	"golang.org/x/net/context"
)

var (
	buildTxHandler BuildTxHandler
	buildTxService services.BuildTxService
	
	postTxHandler PostTxHandler
	postTxService services.PostTxService
	
	sequenceHandler SequenceHandler
	sequenceService services.SequenceService
	
	balanceHandler BalanceHandler
	balanceService services.BalanceService
	
	txListHandler TxListHandler
	txService     services.TxService
	
	txDetailHandler TxDetailHandler
)

func Handler(ctx context.Context, req interface{}) (interface{}, error) {
	var (
		res interface{}
		err error
	)
	
	switch req.(type) {
	case *commonProtoc.BuildTxRequest:
		res, err = buildTxHandler.Handler(ctx, req.(*commonProtoc.BuildTxRequest))
		break
	case *commonProtoc.PostTxRequest:
		res, err = postTxHandler.Handler(ctx, req.(*commonProtoc.PostTxRequest))
		break
	case *commonProtoc.SequenceRequest:
		res, err = sequenceHandler.Handler(ctx, req.(*commonProtoc.SequenceRequest))
		break
	case *commonProtoc.BalanceRequest:
		res, err = balanceHandler.Handler(ctx, req.(*commonProtoc.BalanceRequest))
		break
	case *commonProtoc.TxListRequest:
		res, err = txListHandler.Handler(ctx, req.(*commonProtoc.TxListRequest))
		break
	case *commonProtoc.TxDetailRequest:
		res, err = txDetailHandler.Handler(ctx, req.(*commonProtoc.TxDetailRequest))
		break
	}
	
	return res, err
}
