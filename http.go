package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

var mux *http.ServeMux

func initHTTP() {
	mux = http.NewServeMux()
	mux.HandleFunc("/", coolqHandler)
}

func coolqHandler(response http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodPost {
		var postData postDataType
		body, _ := ioutil.ReadAll(request.Body)
		json.Unmarshal(body, &postData)
		switch {
		case postData.PostType == "notice":
			// 群员增减，重新获取列表
			getMemberList(postData.GroupID)
			responseJSON := map[string]interface{}{
				"reply": "",
				"block": false,
			}
			responseData, _ := json.Marshal(responseJSON)
			response.WriteHeader(http.StatusOK)
			response.Write(responseData)
		case postData.PostType == "group":
			// 处理群消息
			liver(response, request, postData)
		}
	} else {
		response.WriteHeader(http.StatusBadRequest)
	}
}
