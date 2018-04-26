package tpl

import "testing"

//go test -coverprofile=cover.out github.com/zzkkff/lib4go/db/tpl
// cover -func=cover.out

func TestAnalyzeTPL(t *testing.T) {

	input := make(map[string]interface{})
	input["name"] = "colin"
	input["name2"] = "colin2"
	input["name3_"] = "name3_"
	f := func() string {
		return ":"
	}

	//通用参数解析
	tpls := map[string][]interface{}{
		`1where dual`:               []interface{}{`1where dual`, 0},
		`2where name=@=`:            []interface{}{`2where name=@=`, 0},
		`3where name=@(2`:           []interface{}{`3where name=@(2`, 0},
		`4where name=@!`:            []interface{}{`4where name=@!`, 0},
		`5where name=@@`:            []interface{}{`5where name=@@`, 0},
		`6where name=@w`:            []interface{}{`6where name=:`, 1},
		`7where name=@w id=@id`:     []interface{}{`7where name=: id=:`, 2},
		`8where name=@w\r\n id=@id`: []interface{}{`8where name=:\r\n id=:`, 2},
		`9where id in(#ids)`:        []interface{}{`9where id in(NULL)`, 0},
		`10where name='#name'`:      []interface{}{`10where name='colin'`, 0},
		`11where id=0 &name`:        []interface{}{`11where id=0 and name=:`, 1},
		`12where id=0 &id`:          []interface{}{`12where id=0 `, 0},
		`13where id=0 |name`:        []interface{}{`13where id=0 or name=:`, 1},
		`14where id=0 |id`:          []interface{}{`14where id=0 `, 0},
		`15where id=0 !id`:          []interface{}{`15where id=0 !id`, 0},
		`16set name=colin~id`:       []interface{}{`16set name=colin`, 0},
		`17set id=0~name`:           []interface{}{`17set id=0,name=:`, 1},
		/*add by champly 2016年11月9日11:53:35*/
		`18where name=@name3_`: []interface{}{`18where name=:`, 1},
		/*end*/
	}

	for tpl, except := range tpls {
		actual, params, _ := AnalyzeTPL(tpl, input, f)
		if actual != except[0].(string) || len(params) != except[1].(int) {
			t.Errorf("AnalyzeTPL解析参数有误:except:%s actual:%s", except[0].(string), actual)
		}
	}

	//正确参数解析
	tpl := "select seq_wxaccountmenu_auto_id.nextval from where name=@name2"
	except := "select seq_wxaccountmenu_auto_id.nextval from where name=:"
	actual, params, _ := AnalyzeTPL(tpl, input, f)
	if actual != except || len(params) != 1 || params[0].(string) != input["name2"] {
		t.Error("AnalyzeTPL解析参数有误")
	}

	//值不存在
	tpl = "select seq_wxaccountmenu_auto_id.nextval from where name=@id"
	except = "select seq_wxaccountmenu_auto_id.nextval from where name=:"
	actual, params, _ = AnalyzeTPL(tpl, input, f)
	if actual != except || len(params) != 1 || params[0] != nil {
		t.Error("AnalyzeTPL解析参数有误")
	}

	//多个相同属性
	tpl = "select seq_wxaccountmenu_auto_id.nextval from where name=@id and id=@id"
	except = "select seq_wxaccountmenu_auto_id.nextval from where name=: and id=:"
	actual, params, _ = AnalyzeTPL(tpl, input, f)
	if actual != except || len(params) != 2 || params[0] != nil || params[1] != nil {
		t.Error("AnalyzeTPL解析参数有误")
	}

	/*add by champly 2016年11月9日11:54:52*/
	// 多个不同的参数
	tpl = "select seq_wxaccountmenu_auto_id.nextbal from where name=@name and name2='#name2' &name3_ |name ~name"
	except = "select seq_wxaccountmenu_auto_id.nextbal from where name=: and name2='colin2' and name3_=: or name=: ,name=:"
	actual, params, _ = AnalyzeTPL(tpl, input, f)
	if actual != except || len(params) != 4 || params[0].(string) != input["name"] || params[1].(string) != input["name3_"] || params[2].(string) != input["name"] || params[3].(string) != input["name"] {
		t.Error("AnalyzeTPL解析参数有误")
	}
	/*end*/

}
