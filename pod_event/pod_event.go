package pod_event

import (
	"encoding/json"
	"io"
	"os"
	"sync"
	"time"
)

type Event interface {
	AddEvent(log []LogField)
}

type podEvent struct {
	FileHandle *os.File
	mutex      *sync.Mutex
	lastday    int64
}

type depEvent struct {
}

type LogField struct {
	Timestamp string
	Depid     int64
	Item      string
	Event     EVENT_ITEM
	Content   string
	SN        string
}

var PodEventObj = newPodEvent()

func init() {
	t := time.Now().Local()
	filename := LOG_BASE_PATH + "/" + FILENAME + t.Format(F_DATE)

	if checkFileExist(filename) {
		PodEventObj.FileHandle, _ = os.OpenFile(filename, os.O_APPEND, 0666)
	} else {
		PodEventObj.FileHandle, _ = os.Create(filename)
	}
	PodEventObj.mutex = new(sync.Mutex)
	PodEventObj.lastday = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()).Unix() //给其第一次赋值
}

func newPodEvent() *podEvent {
	return &podEvent{}
}

func (po *podEvent) AddEvent(data []LogField) {
	t := time.Now().Local()

	msgPack := make(map[string]interface{})
	var jsonStr []byte
	msg := ""

	for i := 0; i < len(data); i++ {
		msgPack["Timestamp"] = t.Format(F_DATETIME)
		msgPack["Depid"] = data[i].Depid
		msgPack["Item"] = data[i].Item
		msgPack["Event"] = data[i].Event
		msgPack["Content"] = data[i].Content
		msgPack["SN"] = data[i].SN

		jsonStr, _ = json.Marshal(msgPack)
		msg += string(jsonStr) + "\n"

	}

	filename := LOG_BASE_PATH + "/" + FILENAME + t.Format(F_DATE)

	po.mutex.Lock()
	writeLog(filename, msg)
	po.mutex.Unlock()
}

func (de *depEvent) AddEvent(log []LogField) {

}

func OperationLogAdd(log []LogField) {
	PodEventObj.AddEvent(log)
}

func writeLog(filename, message string) {
	t := time.Now().Local()
	todaySt := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()).Unix() //今日0点时间戳
	if PodEventObj.lastday != todaySt {                                                 //跨天了就更新FileHandle，之后这一天都不会再更新FileHandle了，不会重复打开
		if checkFileExist(filename) {
			PodEventObj.FileHandle, _ = os.OpenFile(filename, os.O_APPEND, 0666)
		} else {
			PodEventObj.FileHandle, _ = os.Create(filename)
		}
		PodEventObj.lastday = todaySt
	}

	io.WriteString(PodEventObj.FileHandle, message)
}

func checkFileExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}
