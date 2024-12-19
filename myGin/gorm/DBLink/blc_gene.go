package DBLink

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"strconv"
	"strings"
	"sync"
)

type Config struct {
	Database struct {
		Host     string `json:"host"`
		Port     string `json:"port"`
		User     string `json:"user"`
		Password string `json:"password"`
	} `json:"database"`
}

// 初始化配置
func init() {
	fmt.Println("执行初始化")
	//设置默认值
	viper.SetDefault("user", "root")
	viper.SetDefault("dbname", "blc_name")
	//读取配置
	fmt.Printf("user:%v\n", viper.Get("user"))
	fmt.Printf("dbname:%v\n", viper.Get("dbname"))

	//TODO：路径还是在go.mod那里开始算！
	viper.SetConfigFile("./conf.yml") // 指定配置文件路径
	viper.SetConfigName("conf")       // 配置文件名称(无扩展名)
	viper.SetConfigType("yml")        // 如果配置文件的名称中没有扩展名，则需要配置此项
	viper.AddConfigPath(".")          // 指定在那里照配置文件，可以写多个路径。这里“.”代表在工作目录中查找配置
	err := viper.ReadInConfig()       // 查找并读取配置文件
	if err != nil {                   // 处理读取配置文件的错误
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	//TODO:yml文件又缩进的通过下面的方式.读取配置
	fmt.Printf("host:%v\n", viper.Get("db-blc-gene.host"))
}

/*
sync.Once 是 Go 语言中的一个类型，用于确保某个操作只执行一次。它提供了一个 Do 方法，该方法接受一个函数作为参数，
并保证该函数在程序的生命周期中只执行一次，即使在多线程环境下也是如此。  在你的代码中，once sync.Once 定义了
一个 sync.Once 类型的变量 once，然后在 GetDB 函数中使用 once.Do 来确保数据库连接只被初始化一次。
*/
var (
	db   *gorm.DB
	once sync.Once
)

func GetDB() *gorm.DB {
	/*
		TODO：这里传给他的是匿名函数
		once.Do方法确保传入的函数只会被执行一次，即使在多线程环境下也是如此。因此，数据库连接池的初始化只会执行一次。
		如果 once.Do 已经执行过，那么后续调用 GetDB 方法时，once.Do 内的代码不会再执行，直接返回已经初始化的 db 对象。
	*/
	once.Do(func() {
		var err error
		user := viper.GetString("db-blc-gene.user")
		password := viper.GetString("db-blc-gene.password")
		host := viper.GetString("db-blc-gene.host")
		dbname := viper.GetString("db-blc-gene.dbname")
		port := viper.GetString("db-blc-gene.port")
		var sb strings.Builder
		sb.WriteString(user + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8mb4&parseTime=True&loc=Local")
		//dsn := "user:password@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
		dsn := sb.String()
		fmt.Println(dsn)
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalf("Failed to connect to database: %v", err)
		}
		sqlDB, err := db.DB()
		if err != nil {
			log.Fatalf("Failed to get database instance: %v", err)
		}
		// Set connection pool settings
		sqlDB.SetMaxOpenConns(25)
		sqlDB.SetMaxIdleConns(25)
		sqlDB.SetConnMaxLifetime(5 * 60)
	})
	return db
}

//上面的代码就是第一次调用GetDB时会初始化数据库连接池，后面的调用直接返回db对象

/*
定义了一个performDatabaseOperations函数，它启动了10个goroutines，每个goroutine都通过调用GetDB()来获取数据库连接，
并执行数据库操作（例如创建新用户）。由于GetDB()函数使用了sync.Once，所以无论多少个goroutines调用它，数据库连接只会被初始化一次。

每个goroutine执行的操作是独立的，GORM会为每个操作从连接池中分配一个连接。当操作完成时，连接会被自动释放回连接池，供其他操作或goroutines使用。
这样，你就可以在多线程环境中安全地使用GORM进行数据库操作。
*/
func performDatabaseOperations() {
	// 在这里获取数据库连接
	db := GetDB()

	// 假设我们有一个模型User
	type User struct {
		gorm.Model
		Name string
		Age  int
	}

	// 在goroutine中执行数据库操作
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			// 创建一个新用户
			user := User{Name: "User" + strconv.Itoa(id), Age: 20 + id}
			result := db.Create(&user)
			if result.Error != nil {
				log.Printf("Error creating user: %v", result.Error)
				return
			}

			// ... 执行其他数据库操作 ...

		}(i)
	}

	wg.Wait()
}
