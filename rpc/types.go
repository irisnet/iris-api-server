package rpc

import (
	
	"github.com/irisnet/irishub-server/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ConvertIrisErrToGRPCErr(err errors.IrisError) error {
	return status.Error(codes.Code(err.ErrCode), err.ErrMsg)
}
