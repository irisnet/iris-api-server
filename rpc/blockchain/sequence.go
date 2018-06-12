package blockchain

import (
	commonProtoc "github.com/irisnet/blockchain-rpc/codegen/server"
	"github.com/irisnet/irishub-server/rpc"
	"github.com/irisnet/irishub-server/rpc/vo"
	"golang.org/x/net/context"
)

type SequenceHandler struct {

}

func (c SequenceHandler) Handler(ctx context.Context, request *commonProtoc.SequenceRequest) (
	*commonProtoc.SequenceResponse, error) {
	
	reqVO := c.buildRequest(request)
	resVO, err := sequenceService.GetSequence(reqVO)
	
	if err.IsNotNull() {
		return nil, rpc.ConvertIrisErrToGRPCErr(err)
	}
	
	return c.buildResponse(resVO), nil
}

func (c SequenceHandler) buildRequest(req *commonProtoc.SequenceRequest) vo.SequenceReqVO {
	reqVO := vo.SequenceReqVO{
		Address: req.GetAddress(),
	}
	
	return reqVO
}

func (c SequenceHandler) buildResponse(resVO vo.SequenceResVO) *commonProtoc.SequenceResponse {
	response := commonProtoc.SequenceResponse{
		Sequence: resVO.Sequence,
		Height: resVO.Height,
	}
	
	return &response
}