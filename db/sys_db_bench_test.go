/*
测试数据为创建一个连接，执行次数，通过100次测试的平均值，结果受网络影响比较大
=========================================================================================
数据库最小连接数	数据库最大连接数	创建一个对象总共执行的次数	执行100次的平均时间
1						1					100				192.760584ms
											200				320.599664ms
											300				460.860102ms【极限值】

1						2					100				120.911301ms
											200				181.185771ms
											300				265.58756ms
											400				312.995867ms
											500				397.115957ms
											600				452.255876ms【极限值】

1						3					100				98.462513ms
											200				142.574024ms
											300				188.185131ms
											400				240.945498ms
											500				303.599712ms
											600				353.686304ms
											700				383.946214ms
											800				437.705034ms【极限值】

1						5					200				134.261366ms
											400				206.37301ms
											600				291.973872ms
											800				367.36614ms
											1000			430.189444ms
											1200			504.232834ms【极限值】

1						10					500				249.471914ms
											1000			360.359579ms
											1500			478.114352ms
											2000			628.062569ms【极限值】

1						20					500				232.310445ms
											1000			383.906315ms
											2000			564.33213ms【极限值】
=========================================================================================
以上测试结果不仅供参考，不推荐使用极限值，测试时候本机ping数据库的结果：
2139 packets transmitted, 2139 received, 0% packet loss, time 2173437ms
rtt min/avg/max/mdev = 0.173/0.827/21.232/1.220 ms



测试数据为每个并发创建一个连接，通过100次测试的平均值，结果受网络影响比较大
=========================================================================================
数据库最小连接数	数据库最大连接数	并发数		执行100次的平均时间
1						1			10			94.572378ms
									20			176.884323ms
									30			238.729181ms
									50			354.432595ms
									80			568.760759ms
									100			665.402301ms
									150			1207.104297ms

1						2			10			80.286042ms
									20			174.795353ms
									30			236.468163ms
									50			342.223012ms
									80			570.88296ms
									100			686.044418ms
									150			1189.003588ms

1						5			10			109.533951ms
									20			188.302289ms
									30			234.065368ms
									50			328.681431ms
									80			573.614511ms
									100			679.108308ms
									150			1241.176618ms

1						10			10			123.949325ms
									20			180.021309ms
									30			260.598765ms
									50			398.30377ms
									80			519.415628ms
									100			653.308291ms
									150			1215.371327ms

1						20			10			106.705388ms
									20			158.357563ms
									30			230.254467ms
									50			357.754897ms
									80			519.842163ms
									100			690.460256ms
									150			1245.593939ms
=========================================================================================
以上测试结果不仅供参考，测试时候本机ping数据库的结果：
2884 packets transmitted, 2884 received, 0% packet loss, time 2930088ms
rtt min/avg/max/mdev = 0.154/0.743/6.852/0.590 ms
*/

package db

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// BenchmarkSysDbTest 测试数据库连接的并发性
func TestSysDbTest(b *testing.T) {
	sql := "select * from test where id = :1"
	args := []interface{}{"1"}

	times := time.Second * 0

	for j := 0; j < 100; j++ {
		wg := &sync.WaitGroup{}
		start := time.Now()
		n := 10
		wg.Add(n)
		for i := 0; i < n; i++ {
			go func() {
				obj, err := NewSysDB("oracle", dbConnectStr, 10)
				if obj == nil || err != nil {
					b.Error("创建数据库连接失败:", err)
				}
				_, _, err = obj.Query(sql, args...)
				if err != nil {
					b.Errorf("test fail %v", err)
				}
				//obj.Close()
				// fmt.Printf("%+v\t%+v\n", dataRows, colus)
				wg.Done()
			}()
		}
		wg.Wait()
		fmt.Println("总共耗时：", time.Now().Sub(start))
		times += time.Now().Sub(start)
	}

	fmt.Println("10并发：", times/100)
}

func TestSysDbTest2(b *testing.T) {
	sql := "select * from test where id = :1"

	args := []interface{}{"1"}
	wg := &sync.WaitGroup{}
	times := time.Second * 0

	for j := 0; j < 100; j++ {
		obj, err := NewSysDB("oracle", dbConnectStr, 10)
		if obj == nil || err != nil {
			b.Error("创建数据库连接失败:", err)
		}
		n := 10
		wg.Add(n)
		start := time.Now()
		for i := 0; i < n; i++ {
			go func() {
				_, _, err := obj.Query(sql, args...)
				if err != nil {
					b.Errorf("test fail %v", err)
				}
				// obj.Print()
				// obj.Close()
				// fmt.Printf("%+v\t%+v\n", dataRows, colus)
				wg.Done()
			}()
		}
		wg.Wait()
		// fmt.Println("总共耗时：", time.Now().Sub(start))
		times += time.Now().Sub(start)
	}

	fmt.Println("10:", times/100)
}
