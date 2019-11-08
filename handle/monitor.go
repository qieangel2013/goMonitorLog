package handle

import (
	"fmt"
	"github.com/hpcloud/tail"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type MonitorChan struct {
	file string
	key  string
	text string
}

func IsExist(f string) bool {
	_, err := os.Stat(f)
	return err == nil || os.IsExist(err)
}

func AddToDaysMonitor(days []string, cfg *Config, cline chan MonitorChan, tails chan *tail.Tail) bool {
	currentTime := time.Now()
	for _, file := range days {
		file = strings.Replace(file, "$todaystr", currentTime.Format("2006-01-02"), -1)
		if !IsExist(file) {
			Error.Println("%s文件不存在", file)
		} else {
			go ExcuteTail(file, cfg, cline, tails)
		}
	}
	return true
}

func ExcuteTail(file string, cfg *Config, cline chan MonitorChan, tails chan *tail.Tail) bool {
	tail_line, err := strconv.ParseInt(cfg.TailLine, 10, 64)
	if err != nil {
		Error.Println("error:", err)
	}
	seekInfo := tail.SeekInfo{Offset: -int64(tail_line), Whence: os.SEEK_END}
	t, err := tail.TailFile(file, tail.Config{Follow: true, Poll: true, Location: &seekInfo})
	if err != nil {
		Error.Println("error:", err)
	}
	for line := range t.Lines {
		strTmp := []rune(line.Text)
		findListData := MatchText("find_list", strTmp, []rune{}, '*')
		fillerListData := MatchText("fillter_list", strTmp, []rune{}, '*')
		if findListData != "" && fillerListData == "" {
			MonitorData := MonitorChan{file, findListData, line.Text}
			cline <- MonitorData
		}
	}
	tails <- t
	err = t.Wait()
	if err != nil {
		Error.Println("error:", err)
	}
	return true
}

func AddToHoursMonitor(hours []string, cfg *Config, cline chan MonitorChan, tails chan *tail.Tail) bool {
	currentTime := time.Now()
	for _, file := range hours {
		file = strings.Replace(file, "$todayhourstr", currentTime.Format("2006-01-02_01"), -1)
		// Info.Println(file)
		if !IsExist(file) {
			Error.Println("%s文件不存在", file)
		} else {
			go ExcuteTail(file, cfg, cline, tails)

		}
	}
	return true
}

func CloseMonitor(tail chan *tail.Tail) bool {
	for {
		select {
		case t := <-tail:
			t.Stop()
			t.Cleanup()
		default:
			goto Loop
		}
	}
Loop:
	return true
}

func DingToInfo(s *MonitorChan, cfg *Config) bool {
	if s.text != "" {
		// s.text = "3336666"
		s.text = strings.Replace(s.text, "\\", "\\\\", -1)
		ip := GetIp()
		formt := `#### ip:%s \n\n  #### category::php_error \n\n  ### **file**:<font color=#228B22 size=4>%s</font> \n\n #### key:%s \n\n #### error count:1 \n\n #### **错误**: \n\n > <font color=#0000FF size=4>%s</font> \n `
		text := fmt.Sprintf(formt, ip, s.file, s.key, s.text)
		content := `{"msgtype": "markdown",
					"markdown": {
            			"title":"服务端日志监控",
            			"text": "` + text + `"
        			}
			}`
		req, err := http.NewRequest("POST", cfg.DingWebhookUrl, strings.NewReader(content))
		if err != nil {
			// handle error
		}
		client := &http.Client{}
		req.Header.Set("Content-Type", "application/json; charset=utf-8")
		resp, err := client.Do(req)
		defer resp.Body.Close()
	}
	return true
}

func GetIp() string {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		Error.Println("error:", err)
	}

	for _, address := range addrs {

		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}
