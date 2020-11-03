package sqlx

import (
	"fmt"
	// mysql驱动
	_ "github.com/go-sql-driver/mysql"
	// sqlx
	"github.com/jmoiron/sqlx"
	"testing"
)

// 定义数据库对象
var pool *sqlx.DB

// MYSQL dsn格式：
// {username}:{password}@tcp({host}:{port})/{Dbname}?charset=utf8&parseTime=True&loc=Local
func init() {
	// 定义mysql数据源,配置数据库地址，账户密码
	dsn := "root:qwerty..@tcp(localhost:3306)/test2?charset=utf8&parseTime=True&loc=Local"
	var err error
	// 根据数据源dsn和mysql驱动，创建数据库对象
	pool, err = sqlx.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	// 4.数据库连接池设置
	// sqlx.DB内置了数据库连接池，在你调用sql查询函数的时候，自动从连接池申请连接，可以通过下面方式设置连接池参数:

	// 设置连接池最大连接数
	pool.SetMaxOpenConns(100)

	// 设置连接池最大空闲连接数
	pool.SetMaxIdleConns(20)
}

func TestSql1(t *testing.T) {
	sql := "select * from user where name=?"

	rows, err := pool.Queryx(sql, 'a')
	if err != nil {
		panic(err)
	}
	for rows.Next() {

	}

}

func TestSqlx2(t *testing.T) {
	// schema := `CREATE TABLE place(
	// 	id int primary key auto_increment,
	// 	country varchar(50),
	// 	city varchar(50) NULL default '',
	// 	telcode int);`
	// // 调用Exec函数执行sql语句、创建表
	// _, err := pool.Exec(schema)
	// // 错误处理
	// if err != nil {
	// 	panic(err)
	// }
	// 定义sql语句，通过赞为辐 问号（？）定义三个参数
	countryCitySQL := `INSERT INTO place (country, city, telcode) VALUES (?,?,?)`

	// 通过Exec插入数据，这里传入了三个参数，对应sql语句定义的三个问号所在的位置
	result1, err := pool.Exec(countryCitySQL, "china", "hong kong", 852)
	// 错误处理
	if err != nil {
		fmt.Println(err)
		fmt.Println("插入失败")
		return
	}

	// 插入成功后获取，insert id
	id, _ := result1.LastInsertId()

	fmt.Println("insert id1:", id)

	// 通过MustExec插入数据，如果sql语句出错，则直接抛出panic错误
	result2 := pool.MustExec(countryCitySQL, "South Africa", "Johannesburg", 27)

	// 插入成功后，获取从插入id
	id2, _ := result2.LastInsertId()
	fmt.Println("insert id2:", id2)
}

func TestUpdate(t *testing.T) {
	sql := "update place set telcode=?,city=? where id=?"

	result1, err := pool.Exec(sql, 100, "hong kong", 1)
	// 错误处理
	if err != nil {
		fmt.Println("更新失败:", err)
		return
	}

	// 查询更新影响行数
	rowsAffected, _ := result1.RowsAffected()

	fmt.Println("影响行数:", rowsAffected)
}

type Place struct {
	id      int
	country string
	city    string
	telcode int
}

func TestGet(t *testing.T) {
	sql := `SELECT * FROM place LIMIT ?`
	p := Place{}
	// 查询一条数据，并且往sql语句传入参数1，替换sql语句中的问号，最后将查询结果保存到struct对象中
	err := pool.Get(&p, sql, 1)
	if err != nil {
		panic(err)
	}
}
