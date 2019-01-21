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

// accountPO 积分账户的持久化模型，持久化模型用于解决模型与持久化存储基础设施的关系。
//
// 持久化模型与领域模型并不一定是一对一映射，如果一个领域模型对应的是关系数据库里的多张表，
// 并且持久化模型使用的是 ORM 实现，那么一个领域模型就可能对应多个持久化模型。在这种情况下资源
// 库在进行持久化模型的读写的时候需要维护数据的事务一致性。
type accountPO struct {
	datas map[string]interface{}
	logs  []map[string]interface{}
}

func accountKey(id string) string {
	return "account." + id
}

func ownerKey(id string) string {
	return "owner." + id + ".accounts"
}
