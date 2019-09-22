package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func rank(response http.ResponseWriter, postData postDataType) {
	qqs := getAt(postData.Message)
	var reply []messageType
	var all []rankType
	allMap := make(map[int64]rankType)
	allNum := 0
	rows, _ := db.Raw("SELECT qq, grade, rank FROM " +
		"(SELECT qq, grade, " +
		"@curRank := IF(@prevRank = grade, @curRank, @incRank) AS rank, " +
		"@incRank := @incRank + 1, " +
		"@prevRank := grade " +
		"FROM users u, " +
		"(SELECT @curRank :=0, @prevRank := NULL, @incRank := 1) r " +
		"WHERE `group` = " + strconv.FormatInt(postData.GroupID, 10) + " " +
		"ORDER BY grade DESC) s").Rows()
	for rows.Next() {
		var rank rankType
		db.ScanRows(rows, &rank)
		all = append(all, rank)
		allMap[rank.QQ] = rank
		allNum = rank.Rank
	}
	if len(qqs) > 0 {
		// at了某些人
		reply = append(reply, messageType{
			Type: "text",
			Data: map[string]string{
				"text": "您查询的排名如下：",
			},
		})
		for _, qq := range qqs {
			var name string
			if members[postData.GroupID][qq].Card != "" {
				name = members[postData.GroupID][qq].Card
			} else {
				name = members[postData.GroupID][qq].Nickname
			}
			var text string
			if allMap[qq].Grade == 0.0 {
				text = "\n" + name + "：尚未打卡，没有名次"
			} else {
				text = "\n" + name + "：积分" + fmt.Sprintf("%.2f", allMap[qq].Grade) + "，排名" + strconv.Itoa(allMap[qq].Rank) + "/" + strconv.Itoa(allNum)
			}
			reply = append(reply, messageType{
				Type: "text",
				Data: map[string]string{
					"text": text,
				},
			})
		}
	} else {
		// 没有at
		if len(all) > 0 {
			reply = append(reply, messageType{
				Type: "text",
				Data: map[string]string{
					"text": "目前肝活跃积分榜前三：",
				},
			})
			n := 3
			if len(all) < 3 {
				n = len(all)
			}
			for i := 0; i < n; i++ {
				var name string
				if members[postData.GroupID][all[i].QQ].Card != "" {
					name = members[postData.GroupID][all[i].QQ].Card
				} else {
					name = members[postData.GroupID][all[i].QQ].Nickname
				}
				reply = append(reply, messageType{
					Type: "text",
					Data: map[string]string{
						"text": "\n" + name + "：积分" + fmt.Sprintf("%.2f", all[i].Grade) + "，排名" + strconv.Itoa(i+1) + "/" + strconv.Itoa(allNum),
					},
				})
			}
			self := all[postData.UserID]
			var selfText string
			if self.Rank == 0 {
				selfText = "\n您尚未打卡，没有名次"
			} else {
				selfText = "\n您积分" + fmt.Sprintf("%.2f", self.Grade) + "，排名" + strconv.Itoa(self.Rank) + "/" + strconv.Itoa(allNum)
			}
			reply = append(reply, messageType{
				Type: "text",
				Data: map[string]string{
					"text": selfText,
				},
			})
		} else {
			reply = append(reply, messageType{
				Type: "text",
				Data: map[string]string{
					"text": "目前无人上榜",
				},
			})
		}
	}
	responseJSON := map[string]interface{}{
		"reply": reply,
		"block": true,
	}
	responseData, _ := json.Marshal(responseJSON)
	response.Header().Add("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	response.Write(responseData)
}

func joinQQs(qqs []int64) string {
	var s string
	for i, qq := range qqs {
		if i == len(qqs)-1 {
			s += strconv.FormatInt(qq, 10)
		} else {
			s += strconv.FormatInt(qq, 10) + ", "
		}
	}
	return s
}
