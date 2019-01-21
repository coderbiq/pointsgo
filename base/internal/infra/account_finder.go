package infra

import (
	"errors"
	"regexp"
	"strconv"

	"github.com/coderbiq/pointsgo/base/internal/model"
)

type accountFinder struct {
}

func (finder accountFinder) Detail(accountID int64) (model.AccountReader, error) {
	if accountData, has := db.get(accountKey(strconv.FormatInt(accountID, 10))); has {
		po := accountData.(accountPO)
		datas := po.datas
		datas["logs"] = po.logs
		return model.AccountReaderFromData(datas), nil
	}

	return model.AccountReader{}, nil
}

func (finder accountFinder) ByID(accountID int64, fields []string) (map[string]interface{}, error) {
	datas := make(map[string]interface{})
	if adatas, has := db.get(accountKey(strconv.FormatInt(accountID, 10))); has {
		po := adatas.(accountPO)
		logFields := []string{}
		for _, field := range fields {
			if matched, err := regexp.MatchString("logs.*", field); err == nil && matched {
				logFields = append(logFields, field[5:])
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
