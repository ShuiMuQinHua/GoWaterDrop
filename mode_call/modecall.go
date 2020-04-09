package mode_call

import (
	"io/ioutil"
	"math"
	"net"
	"os/exec"
	"strconv"
	"strings"
	"sync"
	"time"
)

type modeCall struct {
	ReportData map[string]interface{}
	socket     *net.UDPConn
	mutex      *sync.Mutex
}

var modCallObj = NewModelCall()

func init() {
	modCallObj.socket, _ = net.DialUDP("udp4", nil, &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: UDP_PORT,
	})

	modCallObj.mutex = new(sync.Mutex)
}

func (this *modeCall) StartModCall() {
	this.ReportData["start_time"] = time.Now().UnixNano()
}

func (this *modeCall) EndModCall(data map[string]string) {

	costTime := float64(time.Now().UnixNano()-this.ReportData["start_time"].(int64)) / math.Pow10(9)
	this.ReportData["response_time"] = strconv.FormatFloat(costTime, 'f', 6, 64)

	this.ReportData["datetime"] = time.Now().Format("2006-01-02 15:04:05")

	if _, ok := data["app_id"]; !ok {
		this.ReportData["app_id"] = APP_ID
	} else {
		this.ReportData["app_id"] = data["app_id"]
	}

	if _, ok := data["module"]; ok {
		this.ReportData["module_id"] = MODULES[data["module"]]
	} else {
		this.ReportData["module_id"] = DEFAULT_MODULE_ID
	}

	if _, ok := data["call_ip"]; !ok {
		this.ReportData["call_ip"] = this.getHostName()
	} else {
		this.ReportData["call_ip"] = data["call_ip"]
	}

	if _, ok := data["receive_ip"]; !ok {
		this.ReportData["receive_ip"] = "-"
	} else {
		this.ReportData["receive_ip"] = data["receive_ip"]
	}

	if _, ok := data["status"]; !ok {
		this.ReportData["status"] = "-"
	} else {
		this.ReportData["status"] = data["status"]
	}

	if _, ok := data["status_code"]; !ok {
		this.ReportData["status_code"] = "-"
	} else {
		this.ReportData["status_code"] = data["status_code"]
	}

	split := string(rune(1))

	msg := TOPIC + split +
		this.ReportData["datetime"].(string) + split +
		this.ReportData["app_id"].(string) + split +
		this.ReportData["module_id"].(string) + split +
		this.ReportData["call_ip"].(string) + split +
		this.ReportData["receive_ip"].(string) + split +
		this.ReportData["status"].(string) + split +
		this.ReportData["status_code"].(string) + split +
		this.ReportData["response_time"].(string)

		//	socket, _ := net.DialUDP("udp4", nil, &net.UDPAddr{
		//		IP:   net.IPv4(127, 0, 0, 1),
		//		Port: UDP_PORT,
		//	})

		//	defer socket.Close()

	//发送数据
	senddata := []byte(msg)

	modCallObj.mutex.Lock()

	modCallObj.socket.Write(senddata)

	modCallObj.mutex.Unlock()
}

func (this *modeCall) getHostName() string {

	cmd := exec.Command("/bin/bash", "-c", `uname -a`)
	stdout, _ := cmd.StdoutPipe()

	if err := cmd.Start(); err != nil {
		return ""
	}

	bytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		return ""
	}

	if err := cmd.Wait(); err != nil {
		return ""
	}

	output := string(bytes[:])
	hostname := strings.Split(output, " ")[1]
	return hostname
}

func NewModelCall() *modeCall {
	return &modeCall{ReportData: make(map[string]interface{})}
}
