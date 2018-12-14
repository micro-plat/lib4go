package db

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/micro-plat/lib4go/db"
)

func TestT1(t *testing.T) {

	db := db.NewDB("oci", "rcoupon/123456@ORC136", 1, 1, 300)
	var v_name string
	var v_result string

	_, _, err := db.ExecuteSP("create_test_proc(@id, :v_name, :v_result)", map[string]interface{}{
		"id": 200,
	}, sql.Named("v_name", sql.Out{Dest: &v_name}),
		sql.Named("v_result", sql.Out{Dest: &v_result}))
	fmt.Println("err:", err)
	fmt.Println("v_name:", v_name)
	fmt.Println("v_result:", v_result)
}
