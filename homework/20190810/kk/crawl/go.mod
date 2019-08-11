module bosscrawl

go 1.12

replace (
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190701094942-4def268fd1a4
	golang.org/x/net => github.com/golang/net v0.0.0-20190724013045-ca1201d0de80
	golang.org/x/sync => github.com/golang/sync v0.0.0-20190423024810-112230192c58
	golang.org/x/sys => github.com/golang/sys v0.0.0-20190804053845-51ab0e2deafa
	golang.org/x/text => github.com/golang/text v0.3.2
	golang.org/x/tools => github.com/golang/tools v0.0.0-20190809145639-6d4652c779c4
	golang.org/x/xerrors => github.com/golang/xerrors v0.0.0-20190717185122-a985d3407aa7
)

require github.com/PuerkitoBio/goquery v1.5.0
