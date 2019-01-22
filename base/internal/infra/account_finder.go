package infra

import (
	"errors"
	"regexp"
	"strconv"
)

type accountFinder struct {
}

func (finder accountFinder) ByID(accountID int64, fields []string) (map[string]interface{}, error) {
	return finder.byKey(accountKey(strconv.FormatInt(accountID, 10)), fields)
}

func (finder accountFinder) ByOwnerID(ownerID string, fields []string) ([]map[string]interface{}, error) {
	accounts := []map[string]interface{}{}
	if data, has := db.get(ownerKey(ownerID)); has {
		keys := data.([]string)
		for _, key := range keys {
			if account, err := finder.byKey(key, fields); err == nil {
				accounts = append(accounts, account)
			}
		}
	}
	return accounts, nil
}

func (finder accountFinder) byKey(key string, fields []string) (map[string]interface{}, error) {
	datas := make(map[string]interface{})
	if adatas, has := db.get(key); has {
		po := adatas.(accountPO)
		logFields := []string{}
		for _, field := range fields {
			if matched, err := regexp.MatchString("log.*", field); err == nil && matched {
				logFields = append(logFields, field[4:])
			} else {
				if v, has := po.datas[field]; has {
					datas[field] = v
				}
			}
		}
		if len(logFields) > 0 {
			logs := []map[string]interface{}{}
			for _, log := range po.logs {
				l := make(map[string]interface{})
				for _, field := range logFields {
					l[field] = log[field]
				}
				logs = append(logs, l)
			}
			datas["logs"] = logs
		}
	} else {
		return datas, errors.New("没有找到指定账户")
	}
	return datas, nil
}
