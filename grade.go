package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func grade(response http.ResponseWriter, postData postDataType) {
	qqs := getAt(postData.Message)
	if len(qqs) > 0 {
		// at了某些人
		var users []user
		db.Where("qq IN (?) AND group = ?", qqs, postData.GroupID).Find(&users)
		var reply []messageType
		reply = append(reply, messageType{
			Type: "text",
			Data: map[string]string{
				"text": "您查询的积分如下：",
			},
		})
		for _, qq := range qqs {
			var name string
			if members[postData.GroupID][postData.UserID].Card != "" {
				name = members[postData.GroupID][postData.UserID].Card
			} else {
				name = members[postData.GroupID][postData.UserID].Nickname
			}
			var u user
			for _, item := range users {
				if item.QQ == qq {
					u = item
					break
				}
			}
			reply = append(reply, messageType{
				Type: "text",
				Data: map[string]string{
					"text": "\n" + name + " 的肝活跃积分：" + fmt.Sprintf("%.2f", u.Grade),
				},
			})
		}
		responseJSON := map[string]interface{}{
			"reply": reply,
			"block": true,
		}
		responseData, _ := json.Marshal(responseJSON)
		response.WriteHeader(http.StatusOK)
		response.Write(responseData)
	} else {
		// 没有at
		var u user
		db.Where(&user{Group: postData.GroupID, QQ: postData.UserID}).First(&u)
		reply := "您的肝活跃积分：" + fmt.Sprintf("%.2f", u.Grade)
		responseJSON := map[string]interface{}{
			"reply": reply,
			"block": true,
		}
		responseData, _ := json.Marshal(responseJSON)
		response.WriteHeader(http.StatusOK)
		response.Write(responseData)
	}
}
