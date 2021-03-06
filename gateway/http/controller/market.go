package controller

import (
	"digicon/gateway/rpc"
	. "digicon/proto/common"
	proto "digicon/proto/rpc"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type MarketGroup struct{}

func (s *MarketGroup) Router(r *gin.Engine) {
	action := r.Group("/market")
	{
		action.GET("/history/kline", s.HistoryKline)
		action.GET("/symbols", s.Symbols)
		action.GET("/entrust_quenes", s.EntrustQuene)

		action.GET("/trade_list", s.TradeList)

		action.GET("/quotation", s.Quotation)

		action.GET("/price", s.CurrentPrice)

		action.GET("/volume", s.Volume)

		action.GET("/symbols_title", s.SymbolsTitle)

		action.GET("/symbols_id", s.SymbolsById)

		action.POST("/cny_prices", s.CnyPrices)
	}
}

func (s *MarketGroup) HistoryKline(c *gin.Context) {
	ret := NewPublciError()
	defer func() {
		c.JSON(http.StatusOK, ret.GetResult())
	}()

	type KlineParam struct {
		Symbol string `form:"symbol" binding:"required"`
		Period string `form:"period" binding:"required"`
		Size   int32  `form:"size" binding:"required"`
	}

	var param KlineParam

	if err := c.ShouldBind(&param); err != nil {
		log.Errorf(err.Error())
		ret.SetErrCode(ERRCODE_PARAM, err.Error())
		return
	}

	rsp, err := rpc.InnerService.TokenService.CallHistoryKline(param.Symbol, param.Period, param.Size)
	if err != nil {
		ret.SetErrCode(ERRCODE_UNKNOWN, err.Error())
		return
	}
	ret.SetErrCode(ERRCODE_SUCCESS)
	ret.SetDataSection("list", rsp)
}

func (s *MarketGroup) Symbols(c *gin.Context) {
	ret := NewPublciError()
	defer func() {
		c.JSON(http.StatusOK, ret.GetResult())
	}()
	rsp, err := rpc.InnerService.PriceService.CallSymbols(&proto.NullRequest{})

	if err != nil {
		ret.SetErrCode(ERRCODE_UNKNOWN, err.Error())
		return
	}
	ret.SetErrCode(rsp.Err, rsp.Message)
	ret.SetDataSection("usdt", rsp.Usdt)
	ret.SetDataSection("btc", rsp.Btc)
	ret.SetDataSection("eth", rsp.Eth)
	ret.SetDataSection("sdc", rsp.Sdc)
}

func (s *MarketGroup) EntrustQuene(c *gin.Context) {
	ret := NewPublciError()
	defer func() {
		c.JSON(http.StatusOK, ret.GetResult())
	}()
	type EntrustQueneParam struct {
		Symbol string `form:"symbol" binding:"required"`
		Num    int64  `form:"num" `
	}

	var param EntrustQueneParam

	if err := c.ShouldBindQuery(&param); err != nil {
		log.Errorf(err.Error())
		ret.SetErrCode(ERRCODE_PARAM, err.Error())
		return
	}

	if param.Num <= 0 {
		param.Num = 5
	}

	param.Num -= 1
	rsp, err := rpc.InnerService.TokenService.CallEntrustQuene(&proto.EntrustQueneRequest{
		Symbol: param.Symbol,
		Num:    param.Num,
	})

	if err != nil {
		ret.SetErrCode(ERRCODE_UNKNOWN, err.Error())
		return
	}
	ret.SetErrCode(rsp.Err, rsp.Message)
	ret.SetDataSection("sell", rsp.Sell)
	ret.SetDataSection("buy", rsp.Buy)
}

func (s *MarketGroup) TradeList(c *gin.Context) {
	ret := NewPublciError()
	defer func() {
		c.JSON(http.StatusOK, ret.GetResult())
	}()
	type TradeListParam struct {
		Symbol string `form:"symbol" binding:"required"`
	}

	var param TradeListParam

	if err := c.ShouldBindQuery(&param); err != nil {
		log.Errorf(err.Error())
		ret.SetErrCode(ERRCODE_PARAM, err.Error())
		return
	}

	rsp, err := rpc.InnerService.TokenService.CallTrade(&proto.TradeRequest{
		Symbol: param.Symbol,
	})
	if err != nil {
		ret.SetErrCode(ERRCODE_UNKNOWN, err.Error())
		return
	}
	ret.SetErrCode(rsp.Err, rsp.Message)
	ret.SetDataSection("list", rsp.Data)
}

func (s *MarketGroup) Quotation(c *gin.Context) {
	ret := NewPublciError()
	defer func() {
		c.JSON(http.StatusOK, ret.GetResult())
	}()

	type QuotationParam struct {
		TokenId int32 `form:"token_id" binding:"required"`
	}

	var param QuotationParam

	if err := c.ShouldBindQuery(&param); err != nil {
		log.Errorf(err.Error())
		ret.SetErrCode(ERRCODE_PARAM, err.Error())
		return
	}

	rsp, err := rpc.InnerService.PriceService.CallQuotation(&proto.QuotationRequest{
		TokenId: param.TokenId,
	})
	if err != nil {
		ret.SetErrCode(ERRCODE_UNKNOWN, err.Error())
		return
	}
	ret.SetErrCode(ERRCODE_SUCCESS)
	ret.SetDataSection("list", rsp.Data)
}

func (s *MarketGroup) CurrentPrice(c *gin.Context) {
	ret := NewPublciError()
	defer func() {
		c.JSON(http.StatusOK, ret.GetResult())
	}()

	type SymbolPrice struct {
		Symbol string `form:"symbol" binding:"required"`
	}

	var param SymbolPrice

	if err := c.ShouldBindQuery(&param); err != nil {
		log.Errorf(err.Error())
		ret.SetErrCode(ERRCODE_PARAM, err.Error())
		return
	}

	rsp, err := rpc.InnerService.PriceService.CallCurrentPrice(&proto.CurrentPriceRequest{
		Symbol: param.Symbol,
	})
	if err != nil {
		ret.SetErrCode(ERRCODE_UNKNOWN, err.Error())
		return
	}
	ret.SetErrCode(ERRCODE_SUCCESS)
	ret.SetDataSection("list", rsp.Data)
}

func (s *MarketGroup) Volume(c *gin.Context) {
	ret := NewPublciError()
	defer func() {
		c.JSON(http.StatusOK, ret.GetResult())
	}()
	rsp, err := rpc.InnerService.PriceService.CallVolume(&proto.VolumeRequest{})
	if err != nil {
		ret.SetErrCode(ERRCODE_UNKNOWN, err.Error())
		return
	}
	ret.SetErrCode(ERRCODE_SUCCESS)
	ret.SetDataSection("day", rsp.DayVolume)
	ret.SetDataSection("week", rsp.WeekVolume)
	ret.SetDataSection("month", rsp.MonthVolume)
}



func (s *MarketGroup) SymbolsTitle(c *gin.Context) {
	ret := NewPublciError()
	defer func() {
		c.JSON(http.StatusOK, ret.GetResult())
	}()

	rsp, err := rpc.InnerService.PriceService.CallSymbolsTitle()
	if err != nil {
		ret.SetErrCode(ERRCODE_UNKNOWN, err.Error())
		return
	}
	ret.SetErrCode(ERRCODE_SUCCESS)
	ret.SetDataSection("list", rsp.Data)
}


func (s *MarketGroup) SymbolsById(c *gin.Context) {
	ret := NewPublciError()
	defer func() {
		c.JSON(http.StatusOK, ret.GetResult())
	}()
	type SymbolPrice struct {
		TokenId int32 `form:"token_id" binding:"required"`
	}

	var param SymbolPrice

	if err := c.ShouldBindQuery(&param); err != nil {
		log.Errorf(err.Error())
		ret.SetErrCode(ERRCODE_PARAM, err.Error())
		return
	}
	rsp, err := rpc.InnerService.PriceService.CallSymbolsById(&proto.SymbolsByIdRequest{
		TokenId:param.TokenId,
	})
	if err != nil {
		ret.SetErrCode(ERRCODE_UNKNOWN, err.Error())
		return
	}
	ret.SetErrCode(ERRCODE_SUCCESS)
	ret.SetDataSection("list", rsp.Data)
}



func (s *MarketGroup) CnyPrices(c *gin.Context) {
	ret := NewPublciError()
	defer func() {
		c.JSON(http.StatusOK, ret.GetResult())
	}()
	param:= &struct {
		TokenId []int32 `json:"token_id" binding:"required"`
	}{}


	if err := c.ShouldBindJSON(&param); err != nil {
		log.Errorf(err.Error())
		ret.SetErrCode(ERRCODE_PARAM, err.Error())
		return
	}
	rsp, err := rpc.InnerService.PriceService.CallCnyPrices(&proto.CnyPriceRequest{
		TokenTradeId:param.TokenId,
	})
	if err != nil {
		ret.SetErrCode(ERRCODE_UNKNOWN, err.Error())
		return
	}
	ret.SetErrCode(ERRCODE_SUCCESS)

	ret.SetDataSection("list", rsp.Data)
}