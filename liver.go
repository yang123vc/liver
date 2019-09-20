package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

func liver(response http.ResponseWriter, request *http.Request, postData postDataType) {
	switch {
	case strings.Contains(postData.RawMessage, "帮助") ||
		strings.Contains(postData.RawMessage, "规则") ||
		strings.Contains(postData.RawMessage, "命令") ||
		strings.Contains(postData.RawMessage, "格式"):
		help(response)
	case strings.Contains(postData.RawMessage, "积分"):
		grade(response, postData)
	case strings.Contains(postData.RawMessage, "打卡") ||
		strings.Contains(postData.RawMessage, "签到"):
		check(response, postData)
	default:
		responseJSON := map[string]interface{}{
			"reply": "",
			"block": false,
		}
		responseData, _ := json.Marshal(responseJSON)
		response.WriteHeader(http.StatusOK)
		response.Write(responseData)
	}
}

func getAt(message []messageType) []int64 {
	var qqs []int64
	for _, v := range message {
		if v.Type == "at" {
			qq, _ := strconv.ParseInt(v.Data["qq"], 10, 64)
			qqs = append(qqs, qq)
		}
	}
	return qqs
}

func help(response http.ResponseWriter) {
	reply := "凌晨0、1、2、3、4、5点内可以签到，每小时包括前三十分钟和后三十分钟两个时间段，每个时间段只可签到一次，每次根据签到时间（并非时间段）计算积分并累加。\n" +
		"命令：\n" +
		"！积分：查询自己的积分\n" +
		"！积分 @某人：查询某人的积分\n" +
		"！打卡/签到：打卡\n" +
		"！排名/排行：查看排名情况\n" +
		"！排名/排行 @某人：查询某人的排名\n" +
		"！帮助/规则/命令/格式：显示帮助信息"
	responseJSON := map[string]interface{}{
		"reply": reply,
		"block": true,
	}
	responseData, _ := json.Marshal(responseJSON)
	response.WriteHeader(http.StatusOK)
	response.Write(responseData)
}
