module todolist

go 1.12

require (
	github.com/astaxie/beego v1.12.0
	github.com/go-sql-driver/mysql v1.4.1
	github.com/shiena/ansicolor v0.0.0-20151119151921-a422bbe96644 // indirect
)

replace (
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190701094942-4def268fd1a4
	golang.org/x/net => github.com/golang/net v0.0.0-20190628185345-da137c7871d7
	golang.org/x/sync => github.com/golang/sync v0.0.0-20190423024810-112230192c58
	golang.org/x/sys => github.com/golang/sys v0.0.0-20190626221950-04f50cda93cb
	golang.org/x/text => github.com/golang/text v0.3.2
	golang.org/x/tools => github.com/golang/tools v0.0.0-20190706070813-72ffa07ba3db

)
