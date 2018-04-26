package db

import (
	"reflect"
	"testing"
)

type tqDbData struct {
	data []QueryRow
	cols []string
	err  error
}

type teDBData struct {
	row int64
	err error
}
type tpData struct {
	query  string
	input  map[string]interface{}
	args   []interface{}
	result tqDbData
	tp     int
	/*add by champly 2016年11月12日00:47:08*/
	repInput []interface{}
	/*end*/
}

/*add by champly 2016年11月11日13:44:19*/
type tIDBTrans struct {
	qdata  map[string]tqDbData
	edata  map[string]teDBData
	tpData map[string]tpData
}

/*end*/

type tDB struct {
	qdata  map[string]tqDbData
	edata  map[string]teDBData
	tpData map[string]tpData
}

type tTPL struct {
	data map[string]tpData
}

func (t *tTPL) GetSQLContext(tpl string, input map[string]interface{}) (query string, args []interface{}) {
	return t.data[tpl].query, t.data[tpl].args
}
func (t *tTPL) GetSPContext(tpl string, input map[string]interface{}) (query string, args []interface{}) {
	return t.data[tpl].query, t.data[tpl].args
}
func (t *tTPL) Replace(sql string, args []interface{}) string {
	return "REPLACE"
}

func (t *tDB) Query(q string, input ...interface{}) ([]QueryRow, []string, error) {
	return t.qdata[q].data, t.qdata[q].cols, t.qdata[q].err
}
func (t *tDB) Execute(q string, input ...interface{}) (int64, error) {
	return t.edata[q].row, t.edata[q].err
}

/*change by champly 2016年11月11日10:51:29*/
func (t *tDB) Begin() (IDBTrans, error) {
	// return &tIDBTrans{qdata: t.qdata, edata: t.edata}, nil
	return &tIDBTrans{qdata: t.qdata, edata: t.edata, tpData: t.tpData}, nil
}

/*end*/

/*add by champly 2016年11月11日14:27:40*/
func (t *tIDBTrans) GetSQLContext(tpl string, input map[string]interface{}) (query string, args []interface{}) {
	return t.tpData[tpl].query, t.tpData[tpl].args
}
func (t *tIDBTrans) GetSPContext(tpl string, input map[string]interface{}) (query string, args []interface{}) {
	return t.tpData[tpl].query, t.tpData[tpl].args
}
func (t *tIDBTrans) Replace(sql string, args []interface{}) (r string) {
	return "REPLACE"
}
func (t *tIDBTrans) Rollback() error {
	return nil
}
func (t *tIDBTrans) Commit() error {
	return nil
}

func (t *tIDBTrans) Query(q string, input ...interface{}) ([]QueryRow, []string, error) {
	return t.qdata[q].data, t.qdata[q].cols, t.qdata[q].err
}
func (t *tIDBTrans) Execute(q string, input ...interface{}) (int64, error) {
	return t.edata[q].row, t.edata[q].err
}

/*end*/

func makeDB(q map[string]tqDbData, e map[string]teDBData, p map[string]tpData) *DB {
	d := &DB{}
	tdb := &tDB{}
	tdb.edata = e
	tdb.qdata = q

	/*add by champly 2016年11月12日01:25:34*/
	tdb.tpData = p
	/*end*/

	d.db = tdb

	ttp := &tTPL{}
	ttp.data = p

	d.tpl = ttp
	return d
}

func TestDBQuery(t *testing.T) {
	queryMap := make(map[string]tqDbData)
	executeMap := make(map[string]teDBData)
	tplMap := make(map[string]tpData)

	tplMap = map[string]tpData{
		"select a from dual": tpData{
			tp:    1,
			query: "select 'a' from dual",
			args:  []interface{}{"a", 1},
			input: map[string]interface{}{
				"name": "colin",
			},
			repInput: []interface{}{"id", "age"},
		},
		"update order set t=1 where id=2": tpData{
			tp:    2,
			query: "update order set t=1 where id='2'",
			args:  []interface{}{"a", 1},
			input: map[string]interface{}{
				"name": "colin",
			},
			repInput: []interface{}{"id", "age"},
		},
	}

	queryMap["select 'a' from dual"] = tqDbData{
		err:  nil,
		cols: []string{"name", "id"},
		data: []QueryRow{
			map[string]interface{}{
				"name": "colin",
			},
			map[string]interface{}{
				"id": "id",
			},
		},
	}

	executeMap["update order set t=1 where id='2'"] = teDBData{
		err: nil,
		row: 2,
	}

	db := makeDB(queryMap, executeMap, tplMap)
	for k, v := range tplMap {
		switch v.tp {
		case 1:
			result, sql, input, err := db.Query(k, v.input)
			if !reflect.DeepEqual(result, queryMap[v.query].data) || sql != v.query || err != nil || !reflect.DeepEqual(input, v.args) {
				t.Error("Query返回参数有误", len(result), len(queryMap[v.query].data))
			}
			dt, sql, input, err := db.Scalar(k, v.input)
			if dt != queryMap[v.query].data[0]["name"] || sql != v.query || err != nil || !reflect.DeepEqual(input, v.args) {
				t.Error("Scalar", len(result), len(queryMap[v.query].data))
			}

		case 2:
			row, sql, input, err := db.Execute(k, v.input)
			if row != executeMap[v.query].row || sql != v.query || err != nil || !reflect.DeepEqual(input, v.args) {
				t.Error("execute返回参数有误")
			}

			row, sql, input, err = db.ExecuteSP(k, v.input)
			if row != executeMap[v.query].row || sql != v.query || err != nil || !reflect.DeepEqual(input, v.args) {
				t.Error("ExecuteSP返回参数有误")
			}

			/*add by champly 2016年11月11日11:06:07*/
			result := db.Replace(k, v.repInput)
			if result != "REPLACE" {
				t.Error("Replace返回参数错误")
			}
			/*end*/
		}
	}
}

/*add by champly 2016年11月11日11:07:45*/
func TestDBTrans(t *testing.T) {
	queryMap := make(map[string]tqDbData)
	executeMap := make(map[string]teDBData)
	tplMap := make(map[string]tpData)

	tplMap = map[string]tpData{
		"select a from dual": tpData{
			tp:    1,
			query: "select 'a' from dual",
			args:  []interface{}{"a", 1},
			input: map[string]interface{}{
				"name": "colin",
			},
		},
		"update order set t=1 where id=2": tpData{
			tp:    2,
			query: "update order set t=1 where id='2'",
			args:  []interface{}{"a", 1},
			input: map[string]interface{}{
				"name": "colin",
			},
		},
	}

	queryMap["select 'a' from dual"] = tqDbData{
		err:  nil,
		cols: []string{"name", "id"},
		data: []QueryRow{
			map[string]interface{}{
				"name": "colin",
			},
			map[string]interface{}{
				"id": "id",
			},
		},
	}

	executeMap["update order set t=1 where id='2'"] = teDBData{
		err: nil,
		row: 2,
	}

	db := makeDB(queryMap, executeMap, tplMap)
	tdb, err := db.Begin()
	if err != nil {
		t.Error("创建事务失败")
	}

	for k, v := range tplMap {
		switch v.tp {
		case 1:
			result, sql, input, err := tdb.Query(k, v.input)
			if !reflect.DeepEqual(result, queryMap[v.query].data) || sql != v.query || err != nil || !reflect.DeepEqual(input, v.args) {
				t.Error("Query返回参数有误", len(result), len(queryMap[v.query].data))
			}
			dt, sql, input, err := tdb.Scalar(k, v.input)
			if dt != queryMap[v.query].data[0]["name"] || sql != v.query || err != nil || !reflect.DeepEqual(input, v.args) {
				t.Error("Scalar", len(result), len(queryMap[v.query].data))
			}

		case 2:
			row, sql, input, err := tdb.Execute(k, v.input)
			if row != executeMap[v.query].row || sql != v.query || err != nil || !reflect.DeepEqual(input, v.args) {
				t.Error("execute返回参数有误")
			}

			row, sql, input, err = tdb.ExecuteSP(k, v.input)
			if row != executeMap[v.query].row || sql != v.query || err != nil || !reflect.DeepEqual(input, v.args) {
				t.Error("ExecuteSP返回参数有误")
			}
		}
	}

	err = tdb.Rollback()
	if err != nil {
		t.Error("Rollback失败")
	}
	err = tdb.Commit()
	if err != nil {
		t.Error("Commit失败")
	}
}

/*end*/
