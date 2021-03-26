package db

// _ "github.com/go-sql-driver/mysql"

// func TestQuery(t *testing.T) {

// 	db, err := NewDB("mysql", "hydra:123456@tcp(192.168.0.36:3306)/hydra?charset=utf8", 1, 1, 60)
// 	assert.Equal(t, nil, err)
// 	datas, err := db.Query("select * from dds_area_info limit 2", nil)
// 	assert.Equal(t, nil, err)
// 	assert.Equal(t, 2, datas.Len())

// 	v, err := db.Scalar(`SELECT
// 	a.account_id
//   FROM
// 	beanpay_account_info a
//   WHERE a.account_id = @account_id
//   FOR UPDATE`, map[string]interface{}{"account_id": 86000})
// 	assert.Equal(t, nil, err)
// 	assert.Equal(t, int64(86000), types.GetInt64(v))

// 	v, err = db.Execute(`update beanpay_account_info a set a.create_time = now()  WHERE a.account_id = @account_id `, map[string]interface{}{"account_id": 86000})
// 	assert.Equal(t, nil, err)
// 	assert.Equal(t, int64(1), types.GetInt64(v))

// }
