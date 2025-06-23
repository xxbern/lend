package main

import (
	"database/sql"
	"fmt"
	"github.com/fergusstrange/embedded-postgres"
	"gorm.io/driver/postgres"
	"gorm.io/gen"
	"gorm.io/gorm"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

func GenGorm() {
	os.RemoveAll("./gen/db/")
	killProcessOnPort(5432)
	err2 := os.RemoveAll("../gen/model/")
	println(err2)
	config := gen.Config{
		OutPath:       "../gen/db",
		Mode:          gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
		FieldNullable: true,
	}
	config.WithDbNameOpts(func(*gorm.DB) string {
		return "public"
	})
	g := gen.NewGenerator(config)

	const (
		dbUser     = "postgres"
		dbPassword = "password123"
		dbName     = "postgres"
		dbPort     = 5432
	)

	// 启动嵌入式 Postgres
	psql := embeddedpostgres.NewDatabase(embeddedpostgres.DefaultConfig().
		Username(dbUser).
		Password(dbPassword).
		Port(uint32(dbPort)))

	if err := psql.Start(); err != nil {
		panic(fmt.Sprintf("failed to start embedded postgres: %v", err))
	}
	defer psql.Stop()

	// 连接参数
	dsn := fmt.Sprintf("host=localhost port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbPort, dbUser, dbPassword, dbName)
	sqlDB, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	// 执行初始化 SQL
	file, err := os.ReadFile("sql/00001_init.sql")
	if err != nil {
		panic(err)
	}
	_, err = sqlDB.Exec(string(file))
	if err != nil {
		panic(err)
	}

	// 使用 GORM 连接
	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	g.UseDB(gormDB) // reuse your gorm db

	tables, err := gormDB.Migrator().GetTables()
	if err != nil {
		panic(err)
	}
	//g.ApplyBasic(g.GenerateAllTable())
	for _, table := range tables {
		fmt.Println("🔍 Generating model for table:", table)
		func(tbl string) {
			defer func() {
				if r := recover(); r != nil {
					fmt.Printf("⛔️ Failed to generate model for table %s: %v\n", tbl, r)
				}
			}()
			g.ApplyBasic(g.GenerateModel(tbl, gen.FieldType("STATE", "dm.State")))
		}(table)
	}
	//
	g.Execute()
}

func killProcessOnPort(port int) {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		// Windows系统
		cmd = exec.Command("netstat", "-ano")
		output, _ := cmd.Output()
		for _, line := range strings.Split(string(output), "\n") {
			if strings.Contains(line, "LISTENING") && strings.Contains(line, ":"+strconv.Itoa(port)) {
				fields := strings.Fields(line)
				if len(fields) > 4 {
					pid, _ := strconv.Atoi(fields[len(fields)-1])
					exec.Command("taskkill", "/F", "/PID", strconv.Itoa(pid)).Run()
				}
			}
		}
	case "linux", "darwin":
		// Linux/Mac系统
		cmd = exec.Command("lsof", "-i", ":"+strconv.Itoa(port), "-t")
		if pidStr, err := cmd.Output(); err == nil {
			pid, _ := strconv.Atoi(strings.TrimSpace(string(pidStr)))
			exec.Command("kill", "-9", strconv.Itoa(pid)).Run()
		}
	}
}

//version: "0.1"
//database:
//# consult[https://gorm.io/docs/connecting_to_the_database.html]"
//#  dsn : "username:password@tcp(address:port)/db?charset=utf8mb4&parseTime=true&loc=Local"
//dsn : "/mnt/c/Users/xx/DataGripProjects/Rental.db"
//# input mysql or postgres or sqlite or sqlserver. consult[https://gorm.io/docs/connecting_to_the_database.html]
//db  : "sqlite"
//# enter the required data table or leave it blank.You can input : orders,users,goods
//#  tables  : "*"
//# specify a directory for output
//outPath :  "./db"
//# query code file name, default: gen.go
//outFile :  ""
//# generate unit test for query code
//withUnitTest  : false
//# generated model code's package name
//modelPkgName  : ""
//# generate with pointer when field is nullable
//fieldNullable : false
//# generate field with gorm index tag
//fieldWithIndexTag : false
//# generate field with gorm column type tag
//fieldWithTypeTag  : false
