package genpgsql

/*
GenModel() DBInfo
GetDbName() string
GetPkgName() string    // Getting package names through config outdir configuration.通过config outdir 配置获取包名
GetTableNames() string // Getting tableNames by config. 获取设置的表名
*/

import (
	"database/sql"
	"fmt"
	"sort"

	"github.com/wideway/public/pgsqldb"

	//	"database/sql"
	//	"fmt"
	//	"sort"
	"strings"

	"github.com/wideway/gormt/data/config"
	"github.com/wideway/gormt/data/view/model"

	//	"github.com/wideway/public/pgsqldb"
	//"github.com/wideway/public/github.com/wideway/public/mysqldb"
	"github.com/wideway/public/mylog"
	"github.com/wideway/public/tools"
	//	"gorm.io/driver/postgres"
	//	"gorm.io/gorm"
)

var PgSQLModel pgsqlModel

type pgsqlModel struct {
}

func (m *pgsqlModel) GenModel() model.DBInfo {

	/*	DbName      string    // database name
		PackageName string    // package name
		TabList     []TabInfo // table list .表列表
	*/
	/*
		dbURL := "postgres://iotsaas:''@106.55.106.191:26257/gfast"

		db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

		dsn := fmt.Sprintf("server=%v;database=%v;user id=%v;password=%v;port=%v;encrypt=disable",
			config.GetDbInfo().Host, config.GetDbInfo().Database, config.GetDbInfo().Username, config.GetDbInfo().Password, config.GetDbInfo().Port)

		db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	*/
	//pwd =
	/*
		dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
			config.GetDbInfo().Username,
			config.GetDbInfo().Password,
			config.GetDbInfo().Host,
			config.GetDbInfo().Port,
			config.GetDbInfo().Database)

		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			mylog.Error(err)
			return model.DBInfo{}
		}
		defer func() {
			sqldb, _ := db.DB()
			sqldb.Close()
		}()
	*/
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		config.GetDbInfo().Username,
		config.GetDbInfo().Password,
		config.GetDbInfo().Host,
		config.GetDbInfo().Port,
		config.GetDbInfo().Database)
	db := pgsqldb.OnInitDBOrm(dsn, true)

	var dbInfo model.DBInfo
	m.getPackageInfo(db, &dbInfo)
	dbInfo.PackageName = m.GetPkgName()
	dbInfo.DbName = m.GetDbName()
	return dbInfo
	return model.DBInfo{}
}

// GetDbName get database name.获取数据库名字
func (m *pgsqlModel) GetDbName() string {
	return config.GetDbInfo().Database
}

// GetDbName get database name.获取数据库名字
func (m *pgsqlModel) GetDbSchema() string {
	return config.GetDbInfo().Schema
}

// GetTableNames get table name.获取格式化后指定的表名
func (m *pgsqlModel) GetTableNames() string {
	return config.GetTableNames()
}

// GetOriginTableNames get table name.获取原始指定的表名
func (m *pgsqlModel) GetOriginTableNames() string {
	return config.GetOriginTableNames()
}

// GetPkgName package names through config outdir configuration.通过config outdir 配置获取包名
func (m *pgsqlModel) GetPkgName() string {
	dir := config.GetOutDir()
	dir = strings.Replace(dir, "\\", "/", -1)
	if len(dir) > 0 {
		if dir[len(dir)-1] == '/' {
			dir = dir[:(len(dir) - 1)]
		}
	}
	var pkgName string
	list := strings.Split(dir, "/")
	if len(list) > 0 {
		pkgName = list[len(list)-1]
	}

	if len(pkgName) == 0 || pkgName == "." {
		list = strings.Split(tools.GetModelPath(), "/")
		if len(list) > 0 {
			pkgName = list[len(list)-1]
		}
	}

	return pkgName
}

//mysqldb.MySqlDB,
func (m *pgsqlModel) getPackageInfo(orm *pgsqldb.PgSqlDB, info *model.DBInfo) {
	tabls := m.getTables(orm) // get table and notes
	// if m := config.GetTableList(); len(m) > 0 {
	// 	// 制定了表之后
	// 	newTabls := make(map[string]string)
	// 	for t := range m {
	// 		if notes, ok := tabls[t]; ok {
	// 			newTabls[t] = notes
	// 		} else {
	// 			fmt.Printf("table: %s not found in db\n", t)
	// 		}
	// 	}
	// 	tabls = newTabls
	// }
	for tabName, notes := range tabls {
		var tab model.TabInfo
		tab.Name = tabName
		tab.Notes = notes

		if config.GetIsOutSQL() {
			// Get create SQL statements.获取创建sql语句
			rows, err := orm.Raw("show create table " + assemblyTable(tabName)).Rows()
			//defer rows.Close()
			if err == nil {
				if rows.Next() {
					var table, CreateTable string
					rows.Scan(&table, &CreateTable)
					tab.SQLBuildStr = CreateTable
				}
			}
			rows.Close()
			// ----------end
		}

		// build element.构造元素
		tab.Em = m.getTableElement(orm, tabName)
		// --------end

		info.TabList = append(info.TabList, tab)
	}
	// sort tables
	sort.Slice(info.TabList, func(i, j int) bool {
		return info.TabList[i].Name < info.TabList[j].Name
	})
}

// getTables Get columns and comments.获取表列及注释
//func (m *pgsqlModel) getTables(orm *mysqldb.MySqlDB) map[string]string {
func (m *pgsqlModel) getTables(orm *pgsqldb.PgSqlDB) map[string]string {
	tbDesc := make(map[string]string)
	// Get column names.获取列名
	var tables []string
	if m.GetOriginTableNames() != "" {
		sarr := strings.Split(m.GetOriginTableNames(), ",")
		if len(sarr) != 0 {
			for _, val := range sarr {
				tbDesc[val] = ""
			}
		}
	} else {

		var listTable []genTableName
		orm.Raw("show tables").Scan(&listTable)
		for _, v := range listTable {
			if v.SchemaName != config.GetDbInfo().Schema {
				continue
			}
			tables = append(tables, v.TableName)
			tbDesc[v.TableName] = ""
		}
	}

	// Get table annotations.获取表注释
	var err error
	var rows1 *sql.Rows
	if m.GetTableNames() != "" {
		//rows1, err = orm.Raw("SELECT TABLE_NAME,TABLE_COMMENT FROM information_schema.TABLES WHERE table_schema= '" + m.GetDbName() + "'and TABLE_NAME IN(" + m.GetTableNames() + ")").Rows()
		rows1, err = orm.Raw("SELECT TABLE_NAME,'' as TABLE_COMMENT FROM information_schema.TABLES WHERE table_schema= '" + m.GetDbSchema() + "' and TABLE_NAME IN('" + m.GetTableNames() + "')").Rows()
		//fmt.Println("getTables:" + m.GetTableNames())
		//fmt.Println("SELECT TABLE_NAME,TABLE_COMMENT FROM information_schema.TABLES WHERE table_schema= '" + m.GetDbName() + "'and TABLE_NAME IN(" + m.GetTableNames() + ")")
	} else {
		//rows1, err = orm.Raw("SELECT TABLE_NAME,TABLE_COMMENT FROM information_schema.TABLES WHERE table_schema= '" + m.GetDbName() + "'").Rows()

		//SELECT table_name FROM information_schema.TABLES WHERE table_schema= 'public'
		rows1, err = orm.Raw("SELECT TABLE_NAME , '' as TABLE_COMMENT FROM information_schema.TABLES WHERE table_schema= '" + m.GetDbSchema() + "'").Rows()

		//SELECT table_name FROM information_schema.TABLES WHERE table_schema= 'public'
	}

	if err != nil {
		if !config.GetIsGUI() {
			fmt.Println(err)
		}
		return tbDesc
	}

	for rows1.Next() {
		var table, desc string
		rows1.Scan(&table, &desc)
		tbDesc[table] = desc
	}
	rows1.Close()

	return tbDesc
}
func (m *pgsqlModel) getTableElement(orm *pgsqldb.PgSqlDB, tab string) (el []model.ColumnsInfo) {
	keyNameCount := make(map[string]int)
	KeyColumnMp := make(map[string][]keys)
	// get keys
	var Keys []keys
	orm.Raw("show keys from " + assemblyTable(tab)).Scan(&Keys)
	for _, v := range Keys {
		keyNameCount[v.KeyName]++
		KeyColumnMp[v.ColumnName] = append(KeyColumnMp[v.ColumnName], v)
	}
	// ----------end

	var list []genColumns
	// Get table annotations.获取表注释
	//  mysql:orm.Raw("show FULL COLUMNS from " + assemblyTable(tab)).Scan(&list)
	//         Field,Type,Collation,NULL,Key,Default,Extra,Privileges,Comment

	//pgsql:SELECT * FROM information_schema.columns WHERE table_name ='sys_dept' and table_schema ='public'
	//
	//select   column_name as Field,udt_name as Type,'' as collation, is_nullable  as Null, '' as key, '' as Default, '' as Extra,'' as Privileges,'' as comment,*
	//FROM information_schema.columns WHERE table_name ='books' and table_schema ='public'
	sql := fmt.Sprintf(`select   column_name as Field,crdb_sql_type as Type,'' as collation, is_nullable  as Null, '' as key, '' as Default, '' as Extra,'' as Privileges,'' as comment
	 	FROM information_schema.columns WHERE table_name ='%s' and table_schema ='%s'`,
		tab, config.GetDbInfo().Schema) //"public")

	//var listPgsql []genColumnsPgsql
	orm.Raw(sql).Scan(&list)
	//fmt.Println("sql::::", sql)
	//fmt.Println("list::::", len(list))
	for index, v := range list {
		switch v.Type {
		case "VARCHAR":
			list[index].Type = "varchar(100)"
		case "INT8":
			list[index].Type = "bigint unsigned"
		case "INT4":
			list[index].Type = "int"
		case "TIMESTAMP":
			list[index].Type = "datetime"
		case "DATE":
			list[index].Type = "datetime"
		default:
			mylog.Errorf("pgsql type (%s) Undefined ", v.Type)
		}
	}

	// filter gorm.Model.过滤 gorm.Model
	if filterModel(&list) {
		el = append(el, model.ColumnsInfo{
			Type: "gorm.Model",
		})
	}
	// -----------------end

	// ForeignKey
	var foreignKeyList []genForeignKey
	if config.GetIsForeignKey() {
		sql := fmt.Sprintf(`select table_schema as table_schema,table_name as table_name,column_name as column_name,referenced_table_schema as referenced_table_schema,referenced_table_name as referenced_table_name,referenced_column_name as referenced_column_name
		from INFORMATION_SCHEMA.KEY_COLUMN_USAGE where table_schema = '%v' AND REFERENCED_TABLE_NAME IS NOT NULL AND TABLE_NAME = '%v'`, m.GetDbName(), tab)
		//fmt.Println("sql:", sql)
		orm.Raw(sql).Scan(&foreignKeyList)
	}
	// ------------------end

	for _, v := range list {
		var tmp model.ColumnsInfo
		tmp.Name = v.Field
		tmp.Type = v.Type
		tmp.Extra = v.Extra
		FixNotes(&tmp, v.Desc) // 分析表注释

		if v.Default != nil {
			if *v.Default == "" {
				tmp.Gormt = "default:''"
			} else {
				tmp.Gormt = fmt.Sprintf("default:%s", *v.Default)
			}
		}

		// keys
		if keylist, ok := KeyColumnMp[v.Field]; ok { // maybe have index or key
			for _, v := range keylist {
				if v.NonUnique == 0 { // primary or unique
					if strings.EqualFold(v.KeyName, "PRIMARY") { // PRI Set primary key.设置主键
						tmp.Index = append(tmp.Index, model.KList{
							Key:     model.ColumnsKeyPrimary,
							Multi:   (keyNameCount[v.KeyName] > 1),
							KeyType: v.IndexType,
						})
					} else { // unique
						if keyNameCount[v.KeyName] > 1 {
							tmp.Index = append(tmp.Index, model.KList{
								Key:     model.ColumnsKeyUniqueIndex,
								Multi:   (keyNameCount[v.KeyName] > 1),
								KeyName: v.KeyName,
								KeyType: v.IndexType,
							})
						} else { // unique index key.唯一复合索引
							tmp.Index = append(tmp.Index, model.KList{
								Key:     model.ColumnsKeyUnique,
								Multi:   (keyNameCount[v.KeyName] > 1),
								KeyName: v.KeyName,
								KeyType: v.IndexType,
							})
						}
					}
				} else { // mut
					tmp.Index = append(tmp.Index, model.KList{
						Key:     model.ColumnsKeyIndex,
						Multi:   true,
						KeyName: v.KeyName,
						KeyType: v.IndexType,
					})
				}
			}
		}

		tmp.IsNull = strings.EqualFold(v.Null, "YES")

		// ForeignKey
		fixForeignKey(foreignKeyList, tmp.Name, &tmp.ForeignKeyList)
		// -----------------end
		el = append(el, tmp)
	}

	return
}
func assemblyTable(name string) string {
	return "`" + name + "`"
}
