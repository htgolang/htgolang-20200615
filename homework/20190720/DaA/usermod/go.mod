module github.com/xxdu521/usermod

go 1.12

require (
	github.com/howeyc/gopass v0.0.0-20170109162249-bf9dde6d0d2c
	golang.org/x/crypto v0.0.0-20190308221718-c2843e01d9a2 // indirect
)

replace (
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190701094942-4def268fd1a4
	golang.org/x/net => github.com/golang/net v0.0.0-20190628185345-da137c7871d7
	golang.org/x/sync => github.com/golang/sync v0.0.0-20190423024810-112230192c58
	golang.org/x/sys => github.com/golang/sys v0.0.0-20190712062909-fae7ac547cb7
	golang.org/x/text => github.com/golang/text v0.3.2
	golang.org/x/tools => github.com/golang/tools v0.0.0-20190712213246-8b927904ee0d
)
