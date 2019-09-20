module log_analysis

go 1.12

replace (
	cloud.google.com/go => github.com/googleapis/google-cloud-go v0.45.1
	cloud.google.com/go/bigquery => github.com/googleapis/google-cloud-go/bigquery v1.0.1
	cloud.google.com/go/datastore => github.com/googleapis/google-cloud-go/datastore v1.0.0
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190701094942-4def268fd1a4
	golang.org/x/exp => github.com/golang/exp v0.0.0-20190829153037-c13cbed26979
	golang.org/x/image => github.com/golang/image v0.0.0-20190902063713-cb417be4ba39
	golang.org/x/lint => github.com/golang/lint v0.0.0-20190409202823-959b441ac422
	golang.org/x/mobile => github.com/golang/mobile v0.0.0-20190830201351-c6da95954960
	golang.org/x/mod => github.com/golang/mod v0.1.0
	golang.org/x/net => github.com/golang/net v0.0.0-20190628185345-da137c7871d7
	golang.org/x/oauth2 => github.com/golang/oauth2 v0.0.0-20190604053449-0f29369cfe45
	golang.org/x/sync => github.com/golang/sync v0.0.0-20190423024810-112230192c58
	golang.org/x/sys => github.com/golang/sys v0.0.0-20190712062909-fae7ac547cb7
	golang.org/x/text => github.com/golang/text v0.3.2
	golang.org/x/time => github.com/golang/time v0.0.0-20190308202827-9d24e82272b4
	golang.org/x/tools => github.com/golang/tools v0.0.0-20190712213246-8b927904ee0d
	google.golang.org/api => github.com/googleapis/google-api-go-client v0.10.0
	google.golang.org/appengine => github.com/golang/appengine v1.6.2
	google.golang.org/genproto => github.com/googleapis/go-genproto v0.0.0-20190905072037-92dd089d5514
	google.golang.org/grpc => github.com/grpc/grpc-go v1.23.0
)

require (
	github.com/go-sql-driver/mysql v1.4.1
	github.com/jinzhu/gorm v1.9.10
)
