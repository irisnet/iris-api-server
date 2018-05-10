package rests

import (
	"github.com/gin-gonic/gin"
	"github.com/irisnet/iris-api-server/rests/errors"
	"github.com/irisnet/iris-api-server/rests/vo"
)

type DelegatorRoute struct {
}


func RegisterRoutesDelegator(r *gin.Engine) {

	delegatorRoute := DelegatorRoute{}

	rg := r.Group("/delegators")
	{
		rg.GET("/:address/candidates", delegatorRoute.DelegatorCandidateList)
		rg.GET("/:address/candidates/:pub_key", delegatorRoute.DelegatorCandidateDetail)

	}
}

func (r DelegatorRoute) DelegatorCandidateList(c *gin.Context)  {
	address := c.Param("address")
	var (
		listVo vo.DelegatorCandidateListVo
	)
	err := c.ShouldBindQuery(&listVo)
	if err != nil {
		irisErr = irisErr.New(errors.EC40001, errors.EM40001)
		c.JSON(HttpStatusOk, BuildExpResponse(irisErr))
	}
	listVo.Address = address

	response, iriErr := candidateService.DelegatorCandidateList(listVo)
	if iriErr.IsNotNull() {
		c.JSON(HttpStatusOk, BuildExpResponse(irisErr))
	}
	c.JSON(HttpStatusOk, response)

}

func (r DelegatorRoute) DelegatorCandidateDetail(c *gin.Context)  {
	pubKey := c.Param("pub_key")
	address := c.Param("address")

	response, iriErr := candidateService.Detail(pubKey, address)
	if iriErr.IsNotNull() {
		c.JSON(HttpStatusOk, BuildExpResponse(irisErr))
	}
	c.JSON(HttpStatusOk, response)

}