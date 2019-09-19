package main

import (
	"encoding/json"
	"net/http"
)

var member map[int64]map[int64]memberType

func initMember() {
	member = make(map[int64]map[int64]memberType)
	getMemberListCron()
}

func getMemberListCron() {
	for _, group := range cfg.Liver.Group {
		data := map[string]interface{}{
			"group_id": group,
		}
		bodyByte, err := fetch(http.MethodPost, "get_group_member_list", data)
		if err == nil {
			memberMap := make(map[int64]memberType)
			var body struct {
				Data []memberType
			}
			json.Unmarshal(bodyByte, &body)
			for _, m := range body.Data {
				memberMap[m.QQ] = m
			}
			member[group] = memberMap
		}
	}
}
