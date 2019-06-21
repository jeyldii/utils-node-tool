package handlers

import (
	"github.com/button-tech/utils-node-tool/shared/responses"
	"github.com/imroc/req"
	"github.com/qiangxue/fasthttp-routing"
	"log"
	"strconv"
)

func GetBalance(c *routing.Context) error {

	address := c.Param("address")

	res, err := req.Get("https://nodes.wavesplatform.com/addresses/balance/" + address)
	if err != nil {
		log.Println(err)
		return err
	}

	var data responses.BalanceData

	err = res.ToJSON(&data)
	if err != nil {
		log.Println(err)
		return err
	}

	response := new(responses.BalanceResponse)

	response.Balance = strconv.FormatInt(data.Balance, 10)

	if err := responses.JsonResponse(c, response); err != nil {
		return err
	}

	return nil
}