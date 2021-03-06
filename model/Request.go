package model

import (
	h "baseadmin/helper"
	"fmt"
	"github.com/go-gorp/gorp"
	"reflect"
	"time"
)

type Request struct {
	Id         int64     `db:"id, primarykey, autoincrement"`
	RemoteAddr string    `db:"remote_address, size:100"`
	Header     string    `db:"header,size:10000"`
	Body       string    `db:"body,size:10000"`
	Time       time.Time `db:"time"`
}

func (_ Request) IsLanguageModel() bool {
	return false
}

func (_ Request) GetTable() string {
	return "request"
}

func (_ Request) GetPrimaryKey() []string {
	return []string{"id"}
}

func (r Request) BuildStructure(dbmap *gorp.DbMap) {
	Conf := h.GetConfig()
	if Conf.Mode.RebuildStructure {
		h.PrintlnIf(fmt.Sprintf("Drop %v table", r.GetTable()), Conf.Mode.RebuildStructure)
		dbmap.Exec(fmt.Sprintf("DROP TABLE IF EXISTS %s;", r.GetTable()))
	}

	h.PrintlnIf(fmt.Sprintf("Create %v table", r.GetTable()), Conf.Mode.RebuildStructure)
	dbmap.CreateTablesIfNotExists()
	var indexes map[int]map[string]interface{} = make(map[int]map[string]interface{})

	indexes = map[int]map[string]interface{}{
		0: {
			"name":   "IDX_REQUEST_REMOTE_ADDRESS",
			"type":   "hash",
			"field":  []string{"remote_address"},
			"unique": false,
		},
	}
	tablemap, err := dbmap.TableFor(reflect.TypeOf(Request{}), false)
	h.Error(err, "", h.ErrorLvlError)
	for _, index := range indexes {
		h.PrintlnIf(fmt.Sprintf("Create %s index", index["name"].(string)), Conf.Mode.RebuildStructure)
		tablemap.AddIndex(index["name"].(string), index["type"].(string), index["field"].([]string)).SetUnique(index["unique"].(bool))
	}
	dbmap.CreateIndex()
}

func (_ Request) IsAutoIncrement() bool {
	return true
}
