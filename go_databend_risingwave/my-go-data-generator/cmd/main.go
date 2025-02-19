package main

import (
	"flag"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"

	"my-go-data-generator/internal/db"
	"my-go-data-generator/internal/generator"
)

func main() {
	// 加载环境变量（例如 MYSQL_DSN）
	godotenv.Load()
	
	action := flag.String("action", "generate", "操作类型：migrate 或 generate")
	flag.Parse()

	// 从环境变量中获取DSN，如果没有则使用默认配置
	dsn := os.Getenv("MYSQL_DSN")
	if dsn == "" {
		// 格式：user:password@tcp(127.0.0.1:3306)/dbname?charset=utf8&parseTime=True&loc=Local
		dsn = "databend:iloveDatabend#!$@tcp(ilove-databend2025.rwlb.rds.aliyuncs.com:3306)/mydb?charset=utf8&parseTime=True&loc=Local"
	}

	// 连接数据库
	dbConn := db.Connect(dsn)

	// 自动执行数据库迁移逻辑，确保所需表已经存在
	db.Migrate(dbConn)

	if *action == "migrate" {
		return
	} else if *action == "generate" {
		startTime := time.Now()
		log.Println("开始批量生成数据...")
		if err := generator.GenerateData(dbConn); err != nil {
			log.Fatalf("生成数据失败: %v", err)
		}
		elapsedTime := time.Since(startTime)
		log.Printf("批量生成数据完成，总耗时: %s", elapsedTime)

		// 启动定时任务，每30秒插入一条数据
		generator.StartTimer(dbConn, startTime)
	}

	// 阻塞主线程，保持定时任务运行
	select {}
}
