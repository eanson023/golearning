package sqlx

import (
	"database/sql"
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
	Id      int    `db:"id"` //记得Id要大写
	Country string `db:"country"`
	City    string `db:"city"`
	Telcode int    `db:"telcode"`
}

func TestGet(t *testing.T) {
	sql := `SELECT * FROM place LIMIT ?`
	p := Place{}
	// 查询一条数据，并且往sql语句传入参数1，替换sql语句中的问号，最后将查询结果保存到struct对象中
	err := pool.Get(&p, sql, 1)
	if err != nil {
		panic(err)
	}
	t.Log("place:", p)

	var total int
	// 统计表的总记录表，并将查询结果保存到一个变量中
	err = pool.Get(&total, "SELECT count(*) FROM place")
	if err != nil {
		panic(err)
	}
	t.Log("total:", total)

	// 通过Select查询多条记录，并且将结果保存至names变量中
	// 这里仅查询一个字段
	var cities []string
	err = pool.Select(&cities, "SELECT city FROM place LIMIT ?", 10)
	if err != nil {
		panic(err)
	}
	t.Log(cities)
}

// 通过Queryx和QueryRowx查询数据
// 相对于Get和Select函数，Queryx和QueryRowx函数要繁琐一些。
// Queryx可以用于查询多条记录，QueryRowx函数用于查询一条记录。
func TestQueryx1(t *testing.T) {
	// 查询所有的数据，这里返回的是sqlx.Rows对象
	rows, err := pool.Queryx("SELECT country,city,telcode FROM place")
	// 错误检测
	if err != nil {
		panic(err)
	}
	var places []Place = []Place{}
	// 循环遍历每一行记录，rows.Next()函数用于判断是否还有下一行数据
	for rows.Next() {
		// 这里定义三个变量用于接收每一行数据
		var (
			country string
			city    string
			telcode int
		)

		// 调用scan函数，将记录的数据保存到变量中，这里参数的顺序跟上面的sql语句中select后面的字段顺序一致
		err = rows.Scan(&country, &city, &telcode)
		if err != nil {
			panic(err)
		}
		p := Place{}
		p.City = city
		p.Country = country
		p.Telcode = telcode
		places = append(places, p)
	}

	for _, v := range places {
		t.Log(v)
	}

}

type Place2 struct {
	Country       string
	City          sql.NullString //因为city字段允许null,所以这里可以使用sql.NullString类型
	TelephoneCode int            `db:"telcode"` //如果struct字段名的字段名不一样，可以通过db标签设置数据库字段名
}

func TestQueryx2(t *testing.T) {
	// 查询数据
	rows, _ := pool.Queryx("SELECT * FROM place")
	//遍历数据
	for rows.Next() {
		//下面演示如何将数据保存到struct、map和数组中
		//定义struct对象
		var p Place

		//定义map类型
		m := make(map[string]interface{})

		//使用StructScan函数将当前记录的数据保存到struct对象中
		_ = rows.StructScan(&p)
		//保存到map
		_ = rows.MapScan(m)
		//保存到数组
		s, _ := rows.SliceScan()

		t.Log(m)
		t.Log("---------------------------")
		t.Log(s)
		t.Log("---------------------------")
		t.Log(p)
	}
}

// QueryRowx操作跟Queryx类似，区别就是返回一行数据
func TestQueryRowx(t *testing.T) {
	row := pool.QueryRowx("SELECT country,city,telcode FROM place where id = ?", 1)
	var p Place
	//使用StructScan函数将当前记录的数据保存到struct对象中
	row.StructScan(&p)
	t.Log(p)
}

func TestDelete(t *testing.T) {
	// 定义sql语句，通过问号定义了一个参数
	sql := "delete from place where id=?"

	// 通过Exec删除数据，这里传入一个参数，对应sql语句定义的问号所在的位置
	result, err := pool.Exec(sql, 3)

	// 错误处理
	if err != nil {
		t.Fatal(err)
	}
	// 获取删除数据影响行数
	rowsAffected, _ := result.RowsAffected()
	t.Log(rowsAffected)
}

func TestTx(t *testing.T) {
	// 开始一个事务，返回一个事务对象tx
	tx, err := pool.Beginx()
	if err != nil {
		t.Fatal(err)
	}
	// 开始事务对象tx,执行事务
	tx.Commit()
}
