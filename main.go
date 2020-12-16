package main

import (
	"BcRPCCode/entity"
	"BcRPCCode/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

func main() {

	fmt.Println("hello world")
	/**
	 * 1、准备一个json格式的数据（rpc通信协议规范）
	 * 2、发送一个post请求，发送http链接到rpc服务节点（比特币节点）
	 */

	//1、准备一个json数据（string）
	//json数据：序列化，反序列化
	//go: json.Marshal

	//获取节点的区块的总数：getblockcount
	rpcReq := entity.RPCRequest{
		Id:      time.Now().Unix(),
		Method:  GETBLOCK,
		Jsonrpc: RPCVERSION,
		Params: params("000000000019d6689c085ae165831e934ff763ae46a2a6c172b3f1b60a8ce26f"),
	}
	reqBytes, err := json.Marshal(&rpcReq)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("准备好的json数据：",string(reqBytes))

	//2、发送一个post请求
	client := &http.Client{}
	request, err := http.NewRequest("POST", RPCURL, strings.NewReader(string(reqBytes)))
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	//请求头设置
	request.Header.Add("Encoding", "UTF-8")
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", "Basic "+utils.Base64Str(RPCUSER + ":" + RPCPASSWORD))

	//java: HttpResponsr reponse = client.execute(post)
	//java返回响应类：HttpResponse
	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	code := response.StatusCode
	if code == 200 {
		fmt.Println("请求成功")
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		defer response.Body.Close()

		result := entity.RPCResult{}
		err = json.Unmarshal(body, &result)
		if err != nil {
			log.Fatal(err.Error())
			return
		}
		fmt.Println(result)

	}else {
		fmt.Println("状态码：",code)
		fmt.Println("请求失败")
	}
}

func params(para interface{}) []interface{} {
	return []interface{}{para}
}




