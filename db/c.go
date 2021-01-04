package db

import (
	"fmt"
	"path/filepath"
	"strings"
)

var errCodes = []string{
	"1051", "1050", "1061", "1091", "ORA-02289", "ORA-00942",
}

//ParseConnectString 解析数据库连接串
//输入串可以是:"[用户名]:[密码]@[tns名称]" hydra:123456@hydra
//也可以是:"[用户名]:[密码]@[数据库名]/数据库ip"  hydra:123456@hydra/123456
func ParseConnectString(tp string, conn string) (string, error) {
	var uName, pwd, db, ip string
	ips := strings.SplitN(conn, "/", 2)
	if len(ips) > 1 {
		ip = ips[1]
	}
	dbs := strings.SplitN(ips[0], "@", 2)
	if len(dbs) > 1 {
		db = dbs[1]
	}
	up := strings.SplitN(dbs[0], ":", 2)
	if len(up) > 1 {
		pwd = up[1]
	}
	uName = up[0]
	switch tp {
	case "oracle", "ora":
		if uName == "" || pwd == "" || db == "" {
			return "", fmt.Errorf("数据为连接串错误:%s(格式:%s)", conn, `"[用户名]:[密码]@[tns名称]" hydra:123456@hydra`)
		}
		return fmt.Sprintf("%s/%s@%s", uName, pwd, db), nil
	case "mysql":
		if uName == "" || pwd == "" || db == "" || ip == "" {
			return "", fmt.Errorf("数据为连接串错误:%s(格式:%s)", conn, `"[用户名]:[密码]@[数据库名]/数据库ip"  hydra:123456@hydra/123456`)
		}
		return fmt.Sprintf("%s:%s@tcp(%s)/%s", uName, pwd, ip, db), nil
	default:
		return "", fmt.Errorf("不支持的数据库类型:%s", tp)
	}

}

func join(elem ...string) string {
	path := filepath.Join(elem...)
	return strings.Replace(path, "\\", "/", -1)

}
