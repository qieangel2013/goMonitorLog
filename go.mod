module goMonitorLog

go 1.13

require (
	github.com/BurntSushi/toml v0.3.1 // indirect
	github.com/fsnotify/fsnotify v1.4.7 // indirect
	github.com/hpcloud/tail v1.0.0 // indirect
	github.com/juju/errors v0.0.0-20190930114154-d42613fe1ab9 // indirect
	golang.org/x/sys v0.0.0-00010101000000-000000000000 // indirect
	gopkg.in/fsnotify.v1 v1.4.7 // indirect
	gopkg.in/tomb.v1 v1.0.0-20141024135613-dd632973f1e7 // indirect
	julive.com/handle v0.0.0-00010101000000-000000000000 // indirect
)

replace (
	golang.org/x/sys => github.com/golang/sys v0.0.0-20191105231009-c1f44814a5cd
	julive.com/handle => ./handle
)
