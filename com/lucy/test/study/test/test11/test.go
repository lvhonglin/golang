package main

import (
	"golang/com/lucy/test/study/test/test11/tof"
)

type Email struct {
	SummaryId          string
	EmailTitle         string
	AnalysisStartTime  string
	AnalysisEndTime    string
	TotalSceneRunTimes int
	SceneSuccessRate   float32
	SceneFailureTimes  int
	AlarmNum           int
	SceneTotal         int
	DailTaskList       []*DailTaskEmail
	Host               string
	GroupId            string
	ProjectId          string
}

type DailTaskEmail struct {
	DailName           string
	DailUuid           string
	Owner              string
	TotalSceneRunTimes int
	SceneSuccessRate   float32
	SceneSuccessTimes  int
	SceneFailureTimes  int
	SceneList          []*DailTaskSceneEmail
}
type DailTaskSceneEmail struct {
	SceneName    string
	SceneId      string
	Owner        string
	RunTimes     int
	SuccessRate  float32
	SuccessTimes int
	FailureTimes int
}

func main() {
	//file, _ := ioutil.ReadFile("F:\\垃圾箱\\123\\apimonitorOld.html")
	//dailTaskList := make([]*DailTaskEmail, 0)
	//dailTaskScene := &DailTaskSceneEmail{
	//	SceneName:    "场景1",
	//	SceneId:      "123",
	//	Owner:        "创建人1",
	//	RunTimes:     123,
	//	SuccessRate:  13.2,
	//	SuccessTimes: 55,
	//	FailureTimes: 66,
	//}
	//dailTaskScene2 := &DailTaskSceneEmail{
	//	SceneId:      "4567",
	//	SceneName:    "场景2",
	//	Owner:        "创建人3",
	//	RunTimes:     1223,
	//	SuccessRate:  123.2,
	//	SuccessTimes: 255,
	//	FailureTimes: 616,
	//}
	//dailTaskList = append(dailTaskList, &DailTaskEmail{
	//	DailName:           "拨测1",
	//	DailUuid:           "6ac558a1-d6e4-4860-924a-6add987384bf",
	//	Owner:              "lucy",
	//	TotalSceneRunTimes: 12,
	//	SceneSuccessRate:   55.6,
	//	SceneSuccessTimes:  67,
	//	SceneFailureTimes:  77,
	//	SceneList:          []*DailTaskSceneEmail{dailTaskScene},
	//})
	//dailTaskList = append(dailTaskList, &DailTaskEmail{
	//	DailName:           "拨测2",
	//	DailUuid:           "456",
	//	Owner:              "ryan",
	//	TotalSceneRunTimes: 42,
	//	SceneSuccessRate:   51.6,
	//	SceneSuccessTimes:  1,
	//	SceneFailureTimes:  727,
	//	SceneList:          []*DailTaskSceneEmail{dailTaskScene2},
	//})
	//email := Email{
	//	SummaryId:          "6b8b7fcc-2ff5-41e2-92b1-1142be0a587a",
	//	EmailTitle:         "QQ频道拨测日报",
	//	AnalysisStartTime:  "2023-02-11 00:00:00",
	//	AnalysisEndTime:    "2023-03-11 00:00:00",
	//	TotalSceneRunTimes: 123,
	//	SceneSuccessRate:   88.2,
	//	SceneFailureTimes:  44,
	//	AlarmNum:           66,
	//	SceneTotal:         123123,
	//	DailTaskList:       dailTaskList,
	//	Host:               "http://utest-biz-dev.21kunpeng.com",
	//	GroupId:            "utestGroup",
	//	ProjectId:          "b21e09dce88d11ec8588fa8e5b9d4a5d",
	//}
	//println(file[0:1])
	//outFile, _ := os.Create("F:\\垃圾箱\\123\\apimonitorNew.html")
	//tmpl, _ := template.New("test").Parse(string(file))
	//tmpl.Execute(outFile, email)
	//readFile, _ := ioutil.ReadFile("F:\\垃圾箱\\123\\apimonitorNew.html")
	content := string("率宏")
	m := tof.NewMail(
		tof.WithSender("itshelper@tencent.com"),
		tof.WithTitle("【快手集团招聘通知】"),
		tof.WithContent(content),
		tof.WithBodyFormat(1),
		tof.WithAddressee("lucyhllv@tencent.com"),
		tof.WithCC("lucyhllv@tencent.com"),
	)

	err := m.Send("dgvzdguzshcglzxnigldgfjagvu", "0pjCj3b9HH24P5eV2oUGVBhl51YkqPqg", "idc")
	if err != nil {
		println(err)
	}
}
