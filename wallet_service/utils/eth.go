package utils

import (
	"fmt"
	"bytes"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"math/big"
	"github.com/shopspring/decimal"
)
func RpcGetValue(url string,address string,contract string,deci int)(string,error){
	data := make(map[string]interface{})
	if contract==""{
		data["id"]=1
		data["jsonrpc"]="2.0"
		data["method"]="eth_getBalance"
		data["params"] = []string{address, "latest"}
	} else {
		data["id"]=1
		data["jsonrpc"]="2.0"
		data["method"]="eth_call"
		param := make(map[string]string)
		param["data"]="0x70a08231000000000000000000000000"+address[2:]
		param["to"] = contract
		data["params"] = []interface{}{param, "latest"}

	}
	fmt.Println(data)
	rsp,err:=RpcPost(url,data)
	if err != nil {
		return "",err
	}

	ret := make(map[string]interface{})
	err = json.Unmarshal(rsp,&ret)
	if err != nil {
		return "",err
	}
	fmt.Println(ret)
	result ,ok := ret["result"]
	if !ok {
		return "", nil
	}
	var balance string
	balance ,ok = result.(string)
	temp,_:=new(big.Int).SetString(balance[2:],16)
	amount := decimal.NewFromBigInt(temp,int32(8-deci)).IntPart()
	fmt.Println("value",amount)
	return string(amount),nil
}
func RpcPost(url string,send map[string]interface{}) ([]byte,error){
	bytesData, err := json.Marshal(send)
	fmt.Println(string(bytesData))
	if err != nil {
		fmt.Println(err.Error() )
		return nil,err
	}
	reader := bytes.NewReader(bytesData)
	request, err := http.NewRequest("POST", url, reader)
	if err != nil {
		fmt.Println(err.Error())
		return nil,err
	}
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
		return nil,err
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return nil,err
	}
	//byte数组直接转成string，优化内存
	return respBytes,nil
	//str := (*string)(unsafe.Pointer(&respBytes))
	//fmt.Println(*str)
}