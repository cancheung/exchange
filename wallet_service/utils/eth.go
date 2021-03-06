package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/shopspring/decimal"
	"io/ioutil"
	"math/big"
	"net/http"
	"strings"
	"github.com/tidwall/gjson"
	"digicon/common/errors"
	"strconv"
)

func RpcGetValue(url string, address string, contract string, deci int) (string, error) {
	data := make(map[string]interface{})
	if contract == "" {
		data["id"] = 1
		data["jsonrpc"] = "2.0"
		data["method"] = "eth_getBalance"
		data["params"] = []string{address, "latest"}
	} else {
		data["id"] = 1
		data["jsonrpc"] = "2.0"
		data["method"] = "eth_call"
		param := make(map[string]string)
		param["data"] = "0x70a08231000000000000000000000000" + address[2:]
		param["to"] = contract
		data["params"] = []interface{}{param, "latest"}

	}
	//fmt.Println(data)
	rsp, err := RpcPost(url, data)
	if err != nil {
		return "", err
	}

	ret := make(map[string]interface{})
	err = json.Unmarshal(rsp, &ret)
	if err != nil {
		return "", err
	}
	//fmt.Println(ret)
	result, ok := ret["result"]
	if !ok {
		return "", nil
	}
	var balance string
	balance, ok = result.(string)
	temp, _ := new(big.Int).SetString(balance[2:], 16)
	amount := decimal.NewFromBigInt(temp, int32(8-deci)).IntPart()
	re := decimal.New(amount, -8)
	//fmt.Println("value", amount)
	return re.String(), nil
}
func RpcSendRawTx(url string, signtx string) (map[string]interface{}, error) {
	data := make(map[string]interface{})
	data["id"] = 1
	data["jsonrpc"] = "2.0"
	data["method"] = "eth_sendRawTransaction"
	data["params"] = []string{strings.Join([]string{"0x", signtx}, "")}
	rsp, err := RpcPost(url, data)
	if err != nil {
		return nil, err
	}
	ret := make(map[string]interface{})
	err = json.Unmarshal(rsp, &ret)
	return ret, err
}
func RpcPost(url string, send map[string]interface{}) ([]byte, error) {
	bytesData, err := json.Marshal(send)
	//fmt.Println(string(bytesData))
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	reader := bytes.NewReader(bytesData)
	request, err := http.NewRequest("POST", url, reader)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	//byte数组直接转成string，优化内存
	return respBytes, nil
	//str := (*string)(unsafe.Pointer(&respBytes))
	//fmt.Println(*str)
}

//获取随机数
func RpcGetNonce(url,address string) (int,error) {
	data := make(map[string]interface{})
	data["id"] = 1
	data["jsonrpc"] = "2.0"
	data["method"] = "eth_getTransactionCount"
	data["params"] = []string{address,"latest"}
	rsp, err := RpcPost(url, data)
	if err != nil {
		return 0, err
	}
	if gjson.Get(string(rsp),"error").String() != "" {
		return 0,errors.New("获取随机数失败")
	}
	num := gjson.Get(string(rsp),"result").String()
	temp, _ := new(big.Int).SetString(num[2:], 16)
	amount := decimal.NewFromBigInt(temp, 0).IntPart()
	return int(amount), nil
}

//获取gasprice
func RpcGetGasPrice(url string) (int64,error) {
	data := make(map[string]interface{})
	data["id"] = 1
	data["jsonrpc"] = "2.0"
	data["method"] = "eth_gasPrice"
	data["params"] = []string{}
	rsp, err := RpcPost(url, data)
	if err != nil {
		return 0, err
	}

	if gjson.Get(string(rsp),"error").String() != "" {
		return 0,errors.New("获取gasprice失败")
	}

	num := gjson.Get(string(rsp),"result").String()

	temp, _ := new(big.Int).SetString(num[2:], 16)
	amount := decimal.NewFromBigInt(temp, int32(8-8)).IntPart()
	//re := decimal.New(amount, -8)

	return int64(amount), nil
}

//十六进制转十进制
func ToHex(balance string) decimal.Decimal {
	temp, _ := new(big.Int).SetString(balance[2:], 16)
	amount := decimal.NewFromBigInt(temp, 0).IntPart()
	re := decimal.New(amount, -8)
	return re
}

//验证交易结果
func RpcGetTransactionReceipt(url string, signtx string) (error,string) {
	data := make(map[string]interface{})
	data["id"] = 1
	data["jsonrpc"] = "2.0"
	data["method"] = "eth_getTransactionReceipt"
	data["params"] = []string{signtx}
	rsp, err := RpcPost(url, data)
	if err != nil {
		return err,""
	}
	if gjson.Get(string(rsp),"error").String() != "" {
		return errors.New("error"),""
	}
	return nil,gjson.Get(string(rsp),"result").String()
}

//查询交易信息
func RpcGetTransaction(url string, signtx string) (error,string) {
	data := make(map[string]interface{})
	data["id"] = 1
	data["jsonrpc"] = "2.0"
	data["method"] = "eth_getTransactionByHash"
	data["params"] = []string{signtx}
	rsp, err := RpcPost(url, data)
	if err != nil {
		return err,""
	}
	if gjson.Get(string(rsp),"error").String() != "" {
		return errors.New("error"),""
	}
	return nil,gjson.Get(string(rsp),"result").String()
}

//获取区块高度
func GetBlockNumber(url string) (int64, error) {
	data := make(map[string]interface{})
	data["jsonrpc"] = "2.0"
	data["method"] = "eth_blockNumber"
	data["params"] = []string{}
	data["id"] = 1
	rsp, err := RpcPost(url, data)
	if err != nil {
		return 0, err
	}
	if gjson.Get(string(rsp),"error").String() != "" {
		return 0,errors.New("error")
	}
	result := gjson.Get(string(rsp),"result").String()
	number, err := strconv.ParseInt(result, 0, 64)
	if err != nil {
		return 0,errors.New("error")
	}
	return number,nil
}

//根据区块数查询区块
func GetblockBynumber(url string,num int) ([]byte, error) {
	data := make(map[string]interface{})
	data["jsonrpc"] = "2.0"
	data["method"] = "eth_getBlockByNumber"
	strconv.FormatInt(int64(num), 16)
	data["params"] = []interface{}{fmt.Sprintf("0x%s", strconv.FormatInt(int64(num), 16)), true}
	data["id"] = 1
	rsp, err := RpcPost(url,data)

	return rsp, err
}
