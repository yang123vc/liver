package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type allowTime []int

func (allow allowTime) has(time int) bool {
	result := false
	for _, v := range allow {
		if v == time {
			result = true
			break
		}
	}
	return result
}

func check(response http.ResponseWriter, postData postDataType) {
	now := time.Now()
	var u user
	db.Where(&user{Group: postData.GroupID, QQ: postData.UserID}).First(&u)
	var reply string
	if u.Ban {
		reply = "您被封禁，不能打卡！"
	} else {
		var allow allowTime
		allow = []int{0, 1, 2, 3, 4, 5}
		if allow.has(now.Hour()) {
			if now.Unix() >= u.Next.Unix() {
				gradeGetted := float64(now.Hour()) + (0.01 * float64(now.Minute()))
				if gradeGetted == 0.0 {
					gradeGetted = 0.01
				}
				newGrade := u.Grade + gradeGetted
				next := nextTime(now.Unix())
				if u.Grade == 0.0 {
					newUser := user{
						QQ:    postData.UserID,
						Group: postData.GroupID,
						Grade: newGrade,
						Next:  next,
						Ban:   false,
					}
					db.Create(&newUser)
				} else {
					db.Model(&u).UpdateColumns(user{Grade: newGrade, Next: next})
				}
				if newGrade >= 1000.0 {
					reply = "打卡成功，本次打卡您获得" + fmt.Sprintf("%.2f", gradeGetted) + "积分，总积分为" + fmt.Sprintf("%.2f", newGrade) + "，" + next.Format("2006/01/02 15:04:05") + "后可继续打卡。\n亲，您的肝活跃积分较高，这边建议是早睡觉呢。"
				} else {
					reply = "打卡成功，本次打卡您获得" + fmt.Sprintf("%.2f", gradeGetted) + "积分，总积分为" + fmt.Sprintf("%.2f", newGrade) + "，" + next.Format("2006/01/02 15:04:05") + "后可继续打卡。"
				}
			} else {
				reply = "本时间段内您已经打过卡了，下一个打卡时间段将于" + u.Next.Format("2006/01/02 15:04:05") + "开启。"
			}
		} else {
			reply = "凌晨0、1、2、3、4、5点内才可以打卡！"
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
