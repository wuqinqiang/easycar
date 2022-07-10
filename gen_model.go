package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

// generate code
func main() {
	// specify the output directory (default: "./query")
	// ### if you want to query without context constrain, set mode gen.WithoutContext ###

	outPath := "./core/dao/gorm/query"

	g := gen.NewGenerator(gen.Config{
		OutPath: outPath,
		OutFile: outPath + "/query.go",
		/* Mode: gen.WithoutContext|gen.WithDefaultQuery*/
		//if you want the nullable field generation property to be pointer type, set FieldNullable true
		/* FieldNullable: true,*/
		//If you need to generate index tags from the database, set FieldWithIndexTag true
		/* FieldWithIndexTag: true,*/
	})

	// reuse the database connection in Project or create a connection here
	// if you want to use GenerateModel/GenerateModelAs, UseDB is necessary or it will panic
	db, err := gorm.Open(mysql.Open("root:Passw0rd@(127.0.0.1:3306)/easycar?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		panic("gorm err")
	}
	g.UseDB(db)

	g.ApplyBasic(
		g.GenerateModelAs("branch", "Branch"),
		g.GenerateModelAs("global", "Global"),
	)
	g.Execute()
}
