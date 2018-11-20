package models

import (
	"github.com/astaxie/beego"
	"fmt"
	"github.com/go-xorm/xorm"
	"os"
	_ "github.com/go-sql-driver/mysql"
)

//运行标识
const (
	RUN_MODE_DEV  = "dev"
	RUN_MODE_TEST = "test"
	RUN_MODE_PROD = "prod"
)

//数据库参数、redis参数、服务地址
var dbType, dbHost, dbPort, dbUser, dbPassword, dbName, dbCharset, redisConn, serverURL string

//系统初始化
func init(){

	if beego.BConfig.RunMode == RUN_MODE_DEV {
		serverURL = "http://106.14.202.179:8888"
		dbType = "mysql"
		dbHost = "106.14.202.179"
		dbPort = ":3306"
		dbUser = "startapi"
		dbPassword = "startapi"
		dbName = "startapi"
		dbCharset = "utf8mb4"
		redisConn = "106.14.202.179:6379"
	} else if beego.BConfig.RunMode == RUN_MODE_TEST {
		serverURL = "http://106.14.202.179:8888"
		dbType = "mysql"
		dbHost = "106.14.202.179"
		dbPort = ":3306"
		dbUser = "startapi"
		dbPassword = "startapi"
		dbCharset = "utf8mb4"
		redisConn = "106.14.202.179:6379"
	} else if beego.BConfig.RunMode == RUN_MODE_PROD {
		serverURL = "http://106.14.202.179:8888"
		dbType = "mysql"
		dbHost = "106.14.202.179"
		dbPort = ":3306"
		dbUser = "startapi"
		dbPassword = "startapi"
		dbCharset = "utf8mb4"
		redisConn = "106.14.202.179:6379"
	} else {
		panic("运行标识错误")
	}

	initDB(dbConfig{dbType, dbUser, dbPassword, dbName, dbCharset, dbHost, dbPort})

}

type dbConfig struct {
	DbType   		string
	DbUser     		string
	DbPassword 		string
	DbName 			string
	DbCharset  		string
	DbHost     		string
	DbPort     		string
}

//数据库初始化
func initDB(dbConfig dbConfig){
	var err error
	//"root:123@/test?charset=utf8"
	dbUrl := fmt.Sprintf("%s:%s@tcp(%s%s)/%s?charset=%s",
		dbConfig.DbUser, dbConfig.DbPassword, dbConfig.DbHost, dbConfig.DbPort, dbConfig.DbName, dbConfig.DbCharset)
	fmt.Println(dbUrl)
	engine, err := xorm.NewEngine(dbConfig.DbType, dbUrl)
	if err != nil {
		panic("创建数据库连接Engine失败! err:"+err.Error())
	}
	engine.ShowSQL(false)			//在控制台打印出生成的SQL
	engine.SetMaxIdleConns(20)	//设置闲置的连接数
	engine.SetMaxOpenConns(100)	//设置最大打开的连接数，默认值为0表示不限制
	cacher := xorm.NewLRUCacher(xorm.NewMemoryStore(), 1000)	//启用一个全局的内存缓存，存放到内存中，缓存struct的记录数为1000条
	engine.SetDefaultCacher(cacher)

	//SnakeMapper为默认值，结构体驼峰结构，表名转为下划线，可以不写。SameMapper为结构体与表名一致
	//表名前后缀 core.NewPrefixMapper(core.SnakeMapper{}, "prefix")  core.NewSufffixMapper(core.SnakeMapper{}, "suffix")
	//engine.SetMapper(core.SnakeMapper{})

	//engine.DBMetas()	//获取到数据库中所有的表，字段，索引的信息

	//engine.CreateTables()
	//engine.IsTableEmpty()
	//engine.IsTableExist()
	//engine.DropTables()
	//engine.CreateIndexes()
	//engine.CreateUniques()

	//engine.DumpAll()		//导出
	//engine.DumpAllToFile()
	//engine.Import()		//导入
	//engine.ImportFile()

	err = engine.Ping()
	if err != nil {
		panic("数据库连接ping失败! err:"+err.Error())
	}

	//将sql写入到文件中
	f, err := os.Create("sql.log")
	if err != nil {
		panic("创建sql.log文件失败! err:"+err.Error())
	}
	 defer f.Close()
	engine.SetLogger(xorm.NewSimpleLogger(f))

	//同步表结构
	err = engine.Sync2(new(User))
	if err != nil {
		panic("同步表结构失败! err:"+err.Error())
	}
}