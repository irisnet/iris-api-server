package irishub

import (
	irisProtoc "github.com/irisnet/irisnet-rpc/irishub/codegen/server/model"
	"golang.org/x/net/context"
)

type IRISHubRPCServices struct {
}

func (s IRISHubRPCServices) GetCandidateList(ctx context.Context, req *irisProtoc.CandidateListRequest) (
	[]*irisProtoc.Candidate, error) {

	res, err := Handler(ctx, req)
	if err != nil {
		return []*irisProtoc.Candidate{}, err
	}
	return res.([]*irisProtoc.Candidate), err
}

func (s IRISHubRPCServices) GetCandidateDetail(ctx context.Context, req *irisProtoc.CandidateDetailRequest) (
	*irisProtoc.Candidate, error) {

	res, err := Handler(ctx, req)
	if err != nil {
		return &irisProtoc.Candidate{}, err
	}
	return res.(*irisProtoc.Candidate), err
}

func (s IRISHubRPCServices) GetDelegatorCandidateList(ctx context.Context, req *irisProtoc.DelegatorCandidateListRequest) (
	[]*irisProtoc.Candidate, error) {

	res, err := Handler(ctx, req)
	if err != nil {
		return []*irisProtoc.Candidate{}, err
	}
	return res.([]*irisProtoc.Candidate), err
}

func (s IRISHubRPCServices) GetDelegatorTotalShares(ctx context.Context, req *irisProtoc.TotalShareRequest) (
	*irisProtoc.TotalShareResponse, error) {

	res, err := Handler(ctx, req)
	if err != nil {
		return &irisProtoc.TotalShareResponse{}, err
	}
	return res.(*irisProtoc.TotalShareResponse), err
}

func (s IRISHubRPCServices) GetValidatorExRate(ctx context.Context, req *irisProtoc.ValidatorExRateRequest) (
	r *irisProtoc.ValidatorExRateResponse, err error) {

	res, err := Handler(ctx, req)
	if err != nil {
		return &irisProtoc.ValidatorExRateResponse{}, err
	}
	return res.(*irisProtoc.ValidatorExRateResponse), err
}

func (s IRISHubRPCServices) GetWithdrawInfo(ctx context.Context, req *irisProtoc.WithdrawAddrRequest) (
	r *irisProtoc.WithdrawAddrResponse, err error) {

	res, err := Handler(ctx, req)
	if err != nil {
		return &irisProtoc.WithdrawAddrResponse{}, err
	}
	return res.(*irisProtoc.WithdrawAddrResponse), err
}
