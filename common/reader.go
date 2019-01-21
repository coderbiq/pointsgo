package common

import "time"

// AccountReader 积分账户数据读取器
type AccountReader struct {
	datas map[string]interface{}
}

// AccountReaderFromData 根据路由创建账户读取器
func AccountReaderFromData(datas map[string]interface{}) AccountReader {
	return AccountReader{datas: datas}
}

// ID 返回账户标识
func (reader AccountReader) ID() int64 {
	return reader.datas["id"].(int64)
}

// OwnerID 返回账户所属会员标识
func (reader AccountReader) OwnerID() string {
	return reader.datas["ownerId"].(string)
}

// Points 返回账户当前可用积分
func (reader AccountReader) Points() uint {
	return reader.datas["points"].(uint)
}

// Deposited 返回账户总充值积分
func (reader AccountReader) Deposited() uint {
	return reader.datas["deposited"].(uint)
}

// Consumed 返回账户总消费积分
func (reader AccountReader) Consumed() uint {
	return reader.datas["consumed"].(uint)
}

// Logs 返回日志列表
func (reader AccountReader) Logs() []AccountLogReader {
	logs := []AccountLogReader{}
	if datas, has := reader.datas["logs"]; has {
		for _, data := range datas.([]map[string]interface{}) {
			logs = append(logs, AccountLogReader{data: data})
		}
	}
	return logs
}

// CreatedAt 返回账户创建时间
func (reader AccountReader) CreatedAt() time.Time {
	return time.Unix(reader.datas["created"].(int64), 0)
}

// AccountLogReader 账户日志数据读取器
type AccountLogReader struct {
	data map[string]interface{}
}

// Action 返回操作名称
func (reader AccountLogReader) Action() string {
	return reader.data["action"].(string)
}

// Desc 返回详细描述
func (reader AccountLogReader) Desc() string {
	return reader.data["desc"].(string)
}

// CreatedAt 返回创建时间
func (reader AccountLogReader) CreatedAt() time.Time {
	return time.Unix(reader.data["created"].(int64), 0)
}
