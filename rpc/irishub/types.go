package irishub

import (
	irisProtoc "github.com/irisnet/irishub-rpc/codegen/server"
	"github.com/irisnet/irishub-server/services"
	"golang.org/x/net/context"
)

var (
	shareHandler ShareHandler
	shareService services.ShareService
	
	candidateListHandler CandidateListHandler
	candidateService     services.CandidateService
	
	candidateDetailHandler CandidateDetailHandler
)


func Handler(ctx context.Context, req interface{}) (interface{}, error) {
	var (
		res interface{}
		err error
	)
	
	switch req.(type) {
	case *irisProtoc.TotalShareRequest:
		res, err = shareHandler.Handler(ctx, req.(*irisProtoc.TotalShareRequest))
		break
	case *irisProtoc.CandidateListRequest:
		res, err = candidateListHandler.Handler(ctx, req.(*irisProtoc.CandidateListRequest))
		break
	case *irisProtoc.CandidateDetailRequest:
		res, err = candidateDetailHandler.Handler(ctx, req.(*irisProtoc.CandidateDetailRequest))
		break
		
	}
	
	return res, err
}