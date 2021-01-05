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
		`1where dual`:               {`1where dual`, 0},
		`2where name=@=`:            {`2where name=@=`, 0},
		`3where name=@(2`:           {`3where name=@(2`, 0},
		`4where name=@!`:            {`4where name=@!`, 0},
		`5where name=@@`:            {`5where name=@@`, 0},
		`6where name=@w`:            {`6where name=:`, 1},
		`7where name=@w id=@id`:     {`7where name=: id=:`, 2},
		`8where name=@w\r\n id=@id`: {`8where name=:\r\n id=:`, 2},
		`9where id in(#ids)`:        {`9where id in(NULL)`, 0},
		`10where name='#name'`:      {`10where name='colin'`, 0},
		`11where id=0 &name`:        {`11where id=0 and name=:`, 1},
		`12where id=0 &id`:          {`12where id=0 `, 0},
		`13where id=0 |name`:        {`13where id=0 or name=:`, 1},
		`14where id=0 |id`:          {`14where id=0 `, 0},
		`15where id=0 !id`:          {`15where id=0 !id`, 0},
		`16set name=colin~id`:       {`16set name=colin`, 0},
		`17set id=0~name`:           {`17set id=0,name=:`, 1},
		`18where name=@name3_`:      {`18where name=:`, 1},
		`19where name=@t.name3_`:    {`19where name=:`, 1},
		`20where name=$name3_`:      {`20where name=name3_`, 0},
		`21where ?name`:             {`21where name like '%'||:||'%'`, 1},
		`24where \<name`:            {`24where <name`, 0},
		`25where \@name`:            {`25where @name`, 0},
		`26where &t.name`:           {`26where t.name=:`, 1},
		`27where |t.name`:           {`27where t.name=:`, 1},
		`28where &t.sex`:            {`28`, 0},
		`29where name=@t.name`:      {`29where name=:`, 1},
		`30where name=#t.name`:      {`30where name=colin`, 0},
		`32where ~t.name`:           {`32where ,t.name=:`, 1},
		`33where &p.name`:           {`33where p.name=:`, 1},
		`34where &t.sex order by`:   {`34order by`, 0},
		`35where &t.sex group by`:   {`35group by`, 0},
		`36where &t.sex limit`:      {`36limit`, 0},
		`37where |t.sex order by`:   {`37order by`, 0},
		`38where |t.sex group by`:   {`38group by`, 0},
		`39where |t.sex limit`:      {`39limit`, 0},
		`40where order by`:          {`40order by`, 0},
		`41where
and if(isnull(?),1=1,t.kw=?)`: {`41where if(isnull(?),1=1,t.kw=?)`, 0},
		`42where
				&t.storage_mode
			)`: {`42)`, 0},
		`42email='yanglei\@100bm.cn'`: {`42email='yanglei@100bm.cn'`, 0},
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
	if actual != except || len(params) != 1 || params[0].(string) != "" {
		t.Errorf("AnalyzeTPL解析参数有误,except:%s,actual:%s,param:%+v", except, actual, params[0].(string))
	}

	//多个相同属性
	tpl = "select seq_wxaccountmenu_auto_id.nextval from where name=@id and id=@id"
	except = "select seq_wxaccountmenu_auto_id.nextval from where name=: and id=:"
	actual, params, _ = AnalyzeTPL(tpl, input, f)
	if actual != except || len(params) != 2 || params[0].(string) != "" || params[1].(string) != "" {
		t.Errorf("AnalyzeTPL解析参数有误,except:%s,actual:%s,params:%+v", except, actual, params)
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
