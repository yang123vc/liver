package main

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/robfig/cron"
)

func initCron() {
	c := cron.New()
	c.AddFunc("0 0,30 0,1,2,3,4,5 * * *", sendLiverMsgCron)
	c.AddFunc("0 15,45 * * * *", getMemberListCron)
	c.Start()
}

func sendLiverMsgCron() {
	now := time.Now()
	var t int
	if now.Minute() >= 30 {
		t = now.Hour()*2 + 2
	} else {
		t = now.Hour()*2 + 1
	}
	msg := "肝活跃检查（" + strconv.Itoa(t) + "/12）新的打卡时间段开始~"
	help := "\n发送“！帮助”查看帮助信息"
	message := msg + help
	if cfg.Liver.Special[now.Format("01-02")] != "" {
		message = strings.ReplaceAll(cfg.Liver.Special[now.Format("01-02")], "{msg}", msg)
		message = strings.ReplaceAll(cfg.Liver.Special[now.Format("01-02")], "{help}", help)
	}
	if cfg.Liver.Special[now.Format("01-02 15:04:05")] != "" {
		message = strings.ReplaceAll(cfg.Liver.Special[now.Format("01-02 15:04")], "{msg}", msg)
		message = strings.ReplaceAll(cfg.Liver.Special[now.Format("01-02 15:04")], "{help}", help)
	}
	for _, group := range cfg.Liver.Group {
		data := map[string]interface{}{
			"group_id": group,
			"message":  message,
		}
		fetch(http.MethodPost, "send_group_msg", data)
	}
}
