package tool

import "testing"

var json1 = `{
    "userids" : [["userid1","userid2"],["userid3","userid4"]],
    "agentid" : 1,
    "task_id": "taskid122",
    "replace_name": "已收到"
}`

func TestFromJsonFile(t *testing.T) {
    t.Log(FromJson([]byte(json1)))
}
