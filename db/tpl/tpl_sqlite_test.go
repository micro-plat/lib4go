package tpl

import "testing"

func TestSqliteTPLGetContext(t *testing.T) {
	sqlite := MTPLContext{name: "sqlite", prefix: "?"}
	input := make(map[string]interface{})
	input["id"] = 1
	input["name"] = "colin"

	//正确参数解析
	tpl := "select seq_wxaccountmenu_auto_id.nextval from where id=@id and name=@name"
	except := "select seq_wxaccountmenu_auto_id.nextval from where id=? and name=?"
	actual, params := sqlite.GetSQLContext(tpl, input)
	if actual != except || len(params) != 2 || params[0] != input["id"] || params[1] != input["name"] {
		t.Error("GetSQLContext解析参数有误")
	}

	//正确参数解析o
	tpl = "select seq_wxaccountmenu_auto_id.nextval from where id=@id \r\nand name=@name"
	except = "select seq_wxaccountmenu_auto_id.nextval from where id=? \r\nand name=?"
	actual, params = sqlite.GetSQLContext(tpl, input)
	if actual != except || len(params) != 2 || params[0] != input["id"] || params[1] != input["name"] {
		t.Error("GetSQLContext解析参数有误")
	}

	/*add by champly 2016年11月9日15:02:44*/
	tpl = "select seq_wxaccountmenu_auto_id.nextval from where id=#id \r\nand name=@name"
	except = "select seq_wxaccountmenu_auto_id.nextval from where id=1 \r\nand name=?"
	actual, params = sqlite.GetSQLContext(tpl, input)
	if actual != except || len(params) != 1 || params[0] != input["name"] {
		t.Error("GetSQLContext解析参数有误")
	}

	tpl = "select seq_wxaccountmenu_auto_id.nextval from where |id \r\nand &name"
	except = "select seq_wxaccountmenu_auto_id.nextval from where or id=? \r\nand and name=?"
	actual, params = sqlite.GetSQLContext(tpl, input)
	if actual != except || len(params) != 2 || params[0] != input["id"] || params[1] != input["name"] {
		t.Error("GetSQLContext解析参数有误")
	}

	tpl = "select seq_wxaccountmenu_auto_id.nextval from where ~id~name~test"
	except = "select seq_wxaccountmenu_auto_id.nextval from where ,id=?,name=?"
	actual, params = sqlite.GetSQLContext(tpl, input)
	if actual != except || len(params) != 2 || params[0] != input["id"] || params[1] != input["name"] {
		t.Error("GetSQLContext解析参数有误")
	}

	tpl = "select seq_wxaccountmenu_auto_id.nextval from where ~!@#$|&"
	except = "select seq_wxaccountmenu_auto_id.nextval from where ~!@#$|&"
	actual, params = sqlite.GetSQLContext(tpl, input)
	if actual != except || len(params) != 0 {
		t.Error("GetSQLContext解析参数有误")
	}

	tpl = "select seq_wxaccountmenu_auto_id.nextval from where id=@id and id=@id and name=@name"
	except = "select seq_wxaccountmenu_auto_id.nextval from where id=? and id=? and name=?"
	actual, params = sqlite.GetSQLContext(tpl, input)
	if actual != except || len(params) != 3 || params[0] != input["id"] || params[1] != input["id"] || params[2] != input["name"] {
		t.Error("GetSQLContext解析参数有误")
	}

	tpl = "select seq_wxaccountmenu_auto_id.nextval from where id=@id and id=@id and name=@names"
	except = "select seq_wxaccountmenu_auto_id.nextval from where id=? and id=? and name=?"
	actual, params = sqlite.GetSQLContext(tpl, input)
	if actual != except || len(params) != 3 || params[0] != input["id"] || params[1] != input["id"] || params[2] != nil {
		t.Error("GetSQLContext解析参数有误")
	}
	/*end*/

}

func TestSqliteTPLGetSPContext(t *testing.T) {
	sqlite := MTPLContext{name: "sqlite", prefix: "?"}
	input := make(map[string]interface{})
	input["id"] = 1
	input["name"] = "colin"
	input["name_"] = "name_"

	//正确参数解析
	tpl := "order_create(@id,@name,@colin)"
	except := "order_create(?,?,?)"
	actual, params := sqlite.GetSPContext(tpl, input)
	if actual != except || len(params) != 3 || params[0] != input["id"] || params[1] != input["name"] || params[2] != nil {
		t.Error("GetSPContext解析参数有误")
	}

	/*add by champly 2016年11月10日09:25:45*/
	tpl = "order_create(&id&colin)"
	except = "order_create(and id=?)"
	actual, params = sqlite.GetSPContext(tpl, input)
	if actual != except || len(params) != 1 || params[0] != input["id"] {
		t.Error("GetSPContext解析参数有误")
	}

	tpl = "order_create(|id|colin)"
	except = "order_create(or id=?)"
	actual, params = sqlite.GetSPContext(tpl, input)
	if actual != except || len(params) != 1 || params[0] != input["id"] {
		t.Error("GetSPContext解析参数有误")
	}

	tpl = "order_create(#id#colin)"
	except = "order_create(1NULL)"
	actual, params = sqlite.GetSPContext(tpl, input)
	if actual != except || len(params) != 0 {
		t.Error("GetSPContext解析参数有误")
	}

	tpl = "order_create(~id~colin)"
	except = "order_create(,id=?)"
	actual, params = sqlite.GetSPContext(tpl, input)
	if actual != except || len(params) != 1 || params[0] != input["id"] {
		t.Error("GetSPContext解析参数有误")
	}

	tpl = "order_create(~@#|&!)"
	except = "order_create(~@#|&!)"
	actual, params = sqlite.GetSPContext(tpl, input)
	if actual != except || len(params) != 0 {
		t.Error("GetSPContext解析参数有误")
	}

	tpl = "order_create(name_=@name_,name_=@name__)"
	except = "order_create(name_=?,name_=?)"
	actual, params = sqlite.GetSPContext(tpl, input)
	if actual != except || len(params) != 2 || params[0] != input["name_"] || params[1] != nil {
		t.Error("GetSPContext解析参数有误")
	}

	tpl = "order_create(name_=@name_,name_=@name__"
	except = "order_create(name_=?,name_=?"
	actual, params = sqlite.GetSPContext(tpl, input)
	if actual != except || len(params) != 2 || params[0] != input["name_"] || params[1] != nil {
		t.Error("GetSPContext解析参数有误")
	}

	tpl = "order_create(@id,@name,@colin)"
	except = "order_create(?,?,?)"
	actual, params = sqlite.GetSPContext(tpl, nil)
	if actual != except || len(params) != 3 || params[0] != nil || params[1] != nil || params[2] != nil {
		t.Error("GetSPContext解析参数有误")
	}
	/*end*/
}

func TestSqliteTPLReplace(t *testing.T) {
	orcl := MTPLContext{name: "sqlite", prefix: "?"}
	input := make([]interface{}, 0, 2)

	tpl := "begin order_create(?,?,?);end;"
	except := "begin order_create(NULL,NULL,NULL);end;"
	actual := orcl.Replace(tpl, input)
	if actual != except {
		t.Error("Replace解析参数有误", actual)
	}

	tpl = ""
	except = ""
	actual = orcl.Replace(tpl, input)
	if actual != except {
		t.Error("Replace解析参数有误", actual)
	}

	input = append(input, 1)
	input = append(input, "colin")

	tpl = "begin order_create(?,?,?);end;"
	except = "begin order_create('1','colin',NULL);end;"
	actual = orcl.Replace(tpl, input)
	if actual != except {
		t.Error("Replace解析参数有误", actual)
	}

	tpl = "begin order_create(?);end;"
	except = "begin order_create('1');end;"
	actual = orcl.Replace(tpl, input)
	if actual != except {
		t.Error("Replace解析参数有误", actual)
	}

	/*change by champly 2016年11月10日10:06:51*/
	// tpl = "begin order_create(?,?,?);end;"
	// except = "begin order_create('1','colin',NULL);end;"
	// actual = orcl.Replace(tpl, input)
	// if actual != except {
	// 	t.Error("Replace解析参数有误", actual)
	// }
	/*end*/

	tpl = "begin order_create(?,'?234');end;"
	except = "begin order_create('1','?234');end;"
	actual = orcl.Replace(tpl, input)
	if actual != except {
		t.Error("Replace解析参数有误", actual)
	}

	/*add by champly 2016年11月10日10:08:52*/
	tpl = "begin order_create(?,?;end;"
	except = "begin order_create('1','colin';end;"
	actual = orcl.Replace(tpl, input)
	if actual != except {
		t.Error("Replace解析参数有误", actual)
	}

	tpl = "begin order_create(?,?）;end;"
	except = "begin order_create('1',?）;end;"
	actual = orcl.Replace(tpl, input)
	if actual != except {
		t.Error("Replace解析参数有误", actual)
	}

	tpl = "begin order_create(?,? );end;"
	except = "begin order_create('1','colin' );end;"
	actual = orcl.Replace(tpl, input)
	if actual != except {
		t.Error("Replace解析参数有误", actual)
	}

	tpl = "begin order_create(?,?"
	except = "begin order_create('1','colin'"
	actual = orcl.Replace(tpl, input)
	if actual != except {
		t.Error("Replace解析参数有误", actual)
	}

	tpl = "begin order_create(?,??"
	except = "begin order_create('1',?'colin'"
	actual = orcl.Replace(tpl, input)
	if actual != except {
		t.Error("Replace解析参数有误", actual)
	}

	tpl = "begin order_create(?,??@"
	except = "begin order_create('1',??@"
	actual = orcl.Replace(tpl, input)
	if actual != except {
		t.Error("Replace解析参数有误", actual)
	}

	tpl = "begin order_create(?,? "
	except = "begin order_create('1','colin' "
	actual = orcl.Replace(tpl, input)
	if actual != except {
		t.Error("Replace解析参数有误", actual)
	}
	/*end*/
}
