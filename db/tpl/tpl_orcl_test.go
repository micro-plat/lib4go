package tpl

import "testing"

func TestORCTPLgetSPName(t *testing.T) {
	orcl := ATTPLContext{name: "oracle", prefix: ":"}
	input := map[string]string{
		"order_create(:1,:2,:3)":   "begin order_create(:1,:2,:3);end;",
		"order_create(:1,:2,:3);":  "begin order_create(:1,:2,:3);end;",
		"order_create(:1,:2,:3),":  "begin order_create(:1,:2,:3);end;",
		";order_create(:1,:2,:3)":  "begin order_create(:1,:2,:3);end;",
		",order_create(:1,:2,:3)":  "begin order_create(:1,:2,:3);end;",
		";order_create(:1,:2,:3),": "begin order_create(:1,:2,:3);end;",
		"#order_create(:1,:2,:3)":  "begin #order_create(:1,:2,:3);end;",
		"": "begin ;end;",
	}
	for i, except := range input {
		if orcl.getSPName(i) != except {
			t.Error("getSPName 解析参数有误")
		}
	}
}
func TestORCTPLGetContext(t *testing.T) {
	orcl := ATTPLContext{name: "oracle", prefix: ":"}
	input := make(map[string]interface{})
	input["id"] = 1
	input["name"] = "colin"
	input["condtion"] = "(id>1)"
	input["ids"] = `"1","2","3"`

	//正确参数解析
	tpl := "where id=@id and name=@name"
	except := "where id=:1 and name=:2"
	actual, params := orcl.GetSQLContext(tpl, input)
	if actual != except || len(params) != 2 || params[0] != input["id"] || params[1] != input["name"] {
		t.Error("GetSQLContext解析参数有误")
	}

	//正确参数解析
	tpl = "where id=@id \r\nand name=@name"
	except = "where id=:1 \r\nand name=:2"
	actual, params = orcl.GetSQLContext(tpl, input)
	if actual != except || len(params) != 2 || params[0] != input["id"] || params[1] != input["name"] {
		t.Error("GetSQLContext解析参数有误")
	}

	/*change by champly 2016年11月9日11:55:31*/
	// //正确参数解析【重复代码】
	// tpl = "where id=@id \r\nand name=@name"
	// except = "where id=:1 \r\nand name=:2"
	// actual, params = orcl.GetSQLContext(tpl, input)
	// if actual != except || len(params) != 2 || params[0] != input["id"] || params[1] != input["name"] {
	// 	t.Error("GetSQLContext解析参数有误")
	// }
	/*end*/

	//正确参数解析
	tpl = "where id=@id \r\nand name=@name &name"
	except = "where id=:1 \r\nand name=:2 and name=:3"
	actual, params = orcl.GetSQLContext(tpl, input)
	if actual != except || len(params) != 3 || params[0] != input["id"] || params[1] != input["name"] {
		t.Error("GetSQLContext解析参数有误")
	}

	//正确参数解析
	tpl = "where id=@id \r\nand name=@name |name &name"
	except = "where id=:1 \r\nand name=:2 or name=:3 and name=:4"
	actual, params = orcl.GetSQLContext(tpl, input)
	if actual != except || len(params) != 4 || params[0] != input["id"] || params[1] != input["name"] {
		t.Error("GetSQLContext解析参数有误")
	}

	//正确参数解析
	tpl = "where id=@id \r\nand name=@name |name &name #condtion"
	except = "where id=:1 \r\nand name=:2 or name=:3 and name=:4 (id>1)"
	actual, params = orcl.GetSQLContext(tpl, input)
	if actual != except || len(params) != 4 || params[0] != input["id"] || params[1] != input["name"] {
		t.Error("GetSQLContext解析参数有误")
	}

	//正确参数解析
	tpl = "where id=@id and name=@name |name &name id in (#ids)"
	except = `where id=:1 and name=:2 or name=:3 and name=:4 id in ("1","2","3")`
	actual, params = orcl.GetSQLContext(tpl, input)
	if actual != except || len(params) != 4 || params[0] != input["id"] || params[1] != input["name"] {
		t.Errorf("GetSQLContext解析参数有误:actual:%s,except:%s", actual, except)
	}

	//正确参数解析
	tpl = "update order_main set date=sysdate ~name ~status where id=@id"
	except = `update order_main set date=sysdate ,name=:1  where id=:2`
	actual, params = orcl.GetSQLContext(tpl, input)
	if actual != except || len(params) != 2 || params[1] != input["id"] || params[0] != input["name"] {
		t.Errorf("GetSQLContext解析参数有误:actual:%s,except:%s", actual, except)
	}

	/*add by champly 2016年11月9日11:55:57*/
	// 没有参数的解析
	tpl = "where status=@status"
	except = `where status=:1`
	actual, params = orcl.GetSQLContext(tpl, input)
	if actual != except || len(params) != 1 || params[0] != nil {
		t.Errorf("GetSQLContext解析参数有误:actual:%s,except:%s", actual, except)
	}

	// 没有参数的解析
	tpl = "udpate order_main set status=#status"
	except = `udpate order_main set status=NULL`
	actual, params = orcl.GetSQLContext(tpl, input)
	if actual != except || len(params) != 0 {
		t.Errorf("GetSQLContext解析参数有误:actual:%s,except:%s", actual, except)
	}

	// 没有参数的解析
	tpl = "udpate order_main set status=|status &status ~status"
	except = `udpate order_main set status=  `
	actual, params = orcl.GetSQLContext(tpl, input)
	if actual != except || len(params) != 0 {
		t.Errorf("GetSQLContext解析参数有误:actual:%s,except:%s", actual, except)
	}

	// 特殊字符的解析
	tpl = "udpate order_main set status=@!^&*|s"
	except = `udpate order_main set status=@!^&*`
	actual, params = orcl.GetSQLContext(tpl, input)
	if actual != except || len(params) != 0 {
		t.Errorf("GetSQLContext解析参数有误:actual:%s,except:%s", actual, except)
	}

	// 解析空字符串
	tpl = ""
	except = ``
	actual, params = orcl.GetSQLContext(tpl, input)
	if actual != except || len(params) != 0 {
		t.Errorf("GetSQLContext解析参数有误:actual:%s,except:%s", actual, except)
	}
	/*end*/
}

/*
func TestORCLNICEName(t *testing.T) {
	orcl := ATTPLContext{name: "oracle", prefix: ":"}
	input := make(map[string]interface{})

	input["id"] = 1
	input["name"] = "colin"
	tpl := "&t.id&t.name),"
	except := "t.id=:1 and t.name=:2"
	actual, params := orcl.GetSPContext(tpl, input)
	if actual != except || len(params) != 2 || params[0] != input["id"] || params[1] != input["name"] {
		t.Error("GetSPContext解析参数有误", actual)
	}
}
*/
func TestORCTPLGetSPContext(t *testing.T) {
	orcl := ATTPLContext{name: "oracle", prefix: ":"}
	input := make(map[string]interface{})

	input["id"] = 1
	input["name"] = "colin"
	//正确参数解析
	tpl := "order_create(@id,@name,@colin)"
	except := "begin order_create(:1,:2,:3);end;"
	actual, params := orcl.GetSPContext(tpl, input)
	if actual != except || len(params) != 3 || params[0] != input["id"] || params[1] != input["name"] || params[2] != nil {
		t.Error("GetSPContext解析参数有误")
	}

	//正确参数解析
	tpl = "order_create(@id,@name),"
	except = "begin order_create(:1,:2);end;"
	actual, params = orcl.GetSPContext(tpl, input)
	if actual != except || len(params) != 2 || params[0] != input["id"] || params[1] != input["name"] {
		t.Error("GetSPContext解析参数有误", actual)
	}

	/*add by champly 2016年11月9日11:56:24*/
	// sql为空解析
	tpl = ""
	except = "begin ;end;"
	actual, params = orcl.GetSPContext(tpl, input)
	if actual != except || len(params) != 0 {
		t.Error("GetSPContext解析参数有误", actual)
	}
	/*end*/
}
func TestORCTPLReplace(t *testing.T) {
	orcl := ATTPLContext{name: "oracle", prefix: ":"}
	input := make([]interface{}, 0, 2)

	tpl := "begin order_create(:1,:2,:3);end;"
	except := "begin order_create(NULL,NULL,NULL);end;"
	actual := orcl.Replace(tpl, input)
	if actual != except {
		t.Error("Replace解析参数有误:actual:", actual)
	}

	tpl = ""
	except = ""
	actual = orcl.Replace(tpl, input)
	if actual != except {
		t.Error("Replace解析参数有误", actual)
	}

	input = append(input, 1)
	input = append(input, "colin")

	tpl = "begin order_create(:1,:2,:3);end;"
	except = "begin order_create('1','colin',NULL);end;"
	actual = orcl.Replace(tpl, input)
	if actual != except {
		t.Error("Replace解析参数有误", actual)
	}

	tpl = "begin order_create(:11,:2,:3);end;"
	except = "begin order_create(NULL,'colin',NULL);end;"
	actual = orcl.Replace(tpl, input)
	if actual != except {
		t.Error("Replace解析参数有误", actual)
	}

	/*add by champly 2016年11月9日14:23:20*/
	tpl = "begin name=:1  where id=:2;end;"
	except = "begin name='1'  where id='colin';end;"
	actual = orcl.Replace(tpl, input)
	if actual != except {
		t.Error("Replace解析参数有误", actual)
	}

	tpl = ""
	except = ""
	actual = orcl.Replace(tpl, input)
	if actual != except {
		t.Error("Replace解析参数有误", actual)
	}

	tpl = "begin name=:1  where id=:2;end;"
	except = "begin name=:1  where id=:2;end;"
	actual = orcl.Replace(tpl, nil)
	if actual != except {
		t.Error("Replace解析参数有误", actual)
	}

	tpl = "begin name=:1  where id=:2|;end;"
	except = "begin name='1'  where id='colin'|;end;"
	actual = orcl.Replace(tpl, input)
	if actual != except {
		t.Error("Replace解析参数有误", actual)
	}

	tpl = "begin name=:1a  where id=:2a|;end;"
	except = "begin name=:1a  where id=:2a|;end;"
	actual = orcl.Replace(tpl, input)
	if actual != except {
		t.Error("Replace解析参数有误", actual)
	}

	// 匹配结尾
	tpl = "begin name=:1a  where id=:2"
	except = "begin name=:1a  where id='colin'"
	actual = orcl.Replace(tpl, input)
	if actual != except {
		t.Error("Replace解析参数有误", actual)
	}

	tpl = "begin name=:1  where id=:2  "
	except = "begin name='1'  where id='colin'  "
	actual = orcl.Replace(tpl, input)
	if actual != except {
		t.Error("Replace解析参数有误", actual)
	}

}
