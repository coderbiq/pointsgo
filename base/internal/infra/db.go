package infra

// 模拟一个KV数据库
var db = database{datas: make(map[string]interface{})}

type database struct {
	datas map[string]interface{}
}

func (db *database) get(key string) (interface{}, bool) {
	data, has := db.datas[key]
	return data, has
}

func (db *database) set(key string, data interface{}) {
	db.datas[key] = data
}
