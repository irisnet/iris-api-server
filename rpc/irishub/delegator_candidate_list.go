package irishub

import (
	irisProtoc "github.com/irisnet/irishub-rpc/codegen/server"
	"github.com/irisnet/irishub-server/rpc"
	"github.com/irisnet/irishub-server/rpc/vo"
	"golang.org/x/net/context"
)

type DelegatorCandidateListHandler struct {

}

func (h DelegatorCandidateListHandler) Handler(ctx context.Context, req *irisProtoc.DelegatorCandidateListRequest) (
	*irisProtoc.DelegatorCandidateListResponse, error) {
	
	reqVO := h.BuildRequest(req)
	
	resVO, err := candidateService.DelegatorCandidateList(reqVO)
	
	if err.IsNotNull() {
		return nil, rpc.ConvertIrisErrToGRPCErr(err)
	}
	return h.BuildResponse(resVO), nil
}

func (h DelegatorCandidateListHandler) BuildRequest(req *irisProtoc.DelegatorCandidateListRequest) vo.DelegatorCandidateListReqVO {
	
	reqVO := vo.DelegatorCandidateListReqVO{
		Address: req.GetAddress(),
		Page: req.GetPage(),
		PerPage: req.GetPerPage(),
		Q: req.GetQ(),
	}
	
	return reqVO
}

func (h DelegatorCandidateListHandler) BuildResponse(resVO vo.DelegatorCandidateListResVO) *irisProtoc.DelegatorCandidateListResponse {
	var (
		response irisProtoc.DelegatorCandidateListResponse
		resCandidateDescription irisProtoc.Candidate_Description
		resCandidateDelegator irisProtoc.Delegator
		resCandidate irisProtoc.Candidate
		resCandidates []*irisProtoc.Candidate
	)
	
	candidates := resVO.Candidates
	if len(candidates) > 0 {
		for _, v := range candidates {
			// description
			resCandidateDescription = irisProtoc.Candidate_Description{
				Details: v.Description.Details,
				Identity: v.Description.Identity,
				Moniker: v.Description.Moniker,
				Website: v.Description.Website,
			}
			
			// delegator
			var resCandidateDelegators []*irisProtoc.Delegator
			
			if len(v.Delegators) > 0 {
				delegator := v.Delegators[0]
				resCandidateDelegator = irisProtoc.Delegator{
					Address: delegator.Address,
					PubKey: delegator.PubKey,
					Shares: uint64(delegator.Shares),
				}
				resCandidateDelegators = append(resCandidateDelegators, &resCandidateDelegator)
			}
			
			
			resCandidate = irisProtoc.Candidate{
				Address: v.Address,
				PubKey: v.PubKey,
				Shares: uint64(v.Shares),
				VotingPower: v.VotingPower,
				Description: &resCandidateDescription,
				Delegators: resCandidateDelegators,
			}
			
			resCandidates = append(resCandidates, &resCandidate)
		}
	}
	
	response = irisProtoc.DelegatorCandidateListResponse{
		Candidate: resCandidates,
	}
	
	return &response
}
