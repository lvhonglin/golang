package main

import "encoding/json"

func main() {
	a := "{\"TaskId\":7429056073065205514,\"RecordId\":7855417377670595743,\"State\":{\"State\":6,\"Timestamp\":1708511352},\"Type\":5,\"Active\":false,\"StartTime\":\"\",\"EndTime\":\"\"}"
	taskInfo := TaskInfo{}
	err := json.Unmarshal([]byte(a), &taskInfo)
	println(err)
}

type TaskInfo struct {
	TaskId    int64
	RecordId  int64
	State     State
	Type      int32
	Active    bool
	StartTime string
	EndTime   string
}

// State
type State struct {
	State     int32
	Timestamp int64
}
