module github.com/wideway/gormt

go 1.20

require (
	github.com/jroimartin/gocui v0.5.0
	github.com/nicksnyder/go-i18n/v2 v2.2.1
	github.com/spf13/cobra v1.6.1
	github.com/wideway/public v0.0.0-20230311050521-6d756397f72b
	golang.org/x/text v0.8.0
	gopkg.in/go-playground/validator.v9 v9.31.0
	gopkg.in/yaml.v3 v3.0.1
	gorm.io/driver/mysql v1.4.7
	gorm.io/driver/sqlite v1.4.4
	gorm.io/driver/sqlserver v1.4.2
	gorm.io/gorm v1.24.7-0.20230306060331-85eaf9eeda11
)

//replace gitlab.com/tools86/lib => ../../tools86/lib
replace github.com/wideway/public => ../public

//replace github.com/wideway/public => /media/zk008/data/gowork/src/github.com/wideway/public

require (
	github.com/axgle/mahonia v0.0.0-20180208002826-3358181d7394 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-sql-driver/mysql v1.7.0 // indirect
	github.com/golang-sql/civil v0.0.0-20220223132316-b832511892a9 // indirect
	github.com/golang-sql/sqlexp v0.1.0 // indirect
	github.com/gookit/color v1.5.2 // indirect
	github.com/inconshreveable/mousetrap v1.0.1 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/pgx/v5 v5.3.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/leodido/go-urn v1.2.2 // indirect
	github.com/mattn/go-runewidth v0.0.9 // indirect
	github.com/mattn/go-sqlite3 v1.14.15 // indirect
	github.com/microsoft/go-mssqldb v0.19.0 // indirect
	github.com/nsf/termbox-go v1.1.1 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/xo/terminfo v0.0.0-20210125001918-ca9a967f8778 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/zap v1.24.0 // indirect
	golang.org/x/crypto v0.7.0 // indirect
	golang.org/x/sys v0.6.0 // indirect
	gopkg.in/go-playground/assert.v1 v1.2.1 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.2.1 // indirect
	gorm.io/driver/postgres v1.5.0 // indirect
)
