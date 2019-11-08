package main

import (
	"flag"
	"github.com/hpcloud/tail"
	"julive.com/handle"
	"log"
	"strings"
	"time"
)

var configFile = flag.String("config", "./etc/log.toml", "log config file")
var currentDaysTime string
var currentHoursTime string
var hoursFile []string
var daysFile []string

func main() {
	currentDaysTime = time.Now().Format("2006-01-02")
	currentHoursTime = time.Now().Format("2006-01-02_01")
	// 创建一个接受匹配上的字符串管道
	var MonitorChan = make(chan handle.MonitorChan, 10)
	TailDaysPointer := make(chan *tail.Tail, 10)
	TailHoursPointer := make(chan *tail.Tail, 10)
	cfg, err := handle.NewConfigWithFile(*configFile)
	if err != nil {
		log.Println("error:", err)
		return
	}
	handle.LogStart(cfg)
	go handle.BuildDicTrie(cfg)
	for _, file := range cfg.LogList {
		if strings.Contains(file, "$todayhourstr") {
			hoursFile = append(hoursFile, file)
		} else {
			daysFile = append(daysFile, file)
		}
	}
	go handle.AddToDaysMonitor(daysFile, cfg, MonitorChan, TailDaysPointer)
	if len(hoursFile) != 0 {
		go handle.AddToHoursMonitor(hoursFile, cfg, MonitorChan, TailHoursPointer)
	}
	go func() {
		for {
			select {
			//读取sqldata
			case data := <-MonitorChan:
				go func() {
					handle.DingToInfo(&data, cfg)
				}()
			}
		}
	}()
	timerActive := time.NewTimer(600 * time.Second)
	for {
		select {
		case <-timerActive.C:
			if len(hoursFile) != 0 {
				timerCurrentHoursTime := time.Now().Format("2006-01-02_01")
				if timerCurrentHoursTime != currentHoursTime {
					if handle.CloseMonitor(TailHoursPointer) {
						currentHoursTime = timerCurrentHoursTime
						go handle.AddToHoursMonitor(hoursFile, cfg, MonitorChan, TailHoursPointer)
					}
				}
			}
			timerCurrentTime := time.Now().Format("2006-01-02")
			if timerCurrentTime != currentDaysTime {
				if handle.CloseMonitor(TailDaysPointer) {
					currentDaysTime = timerCurrentTime
					go handle.AddToDaysMonitor(daysFile, cfg, MonitorChan, TailDaysPointer)
				}
			}
			timerActive.Reset(600 * time.Second)
		}
	}
}
