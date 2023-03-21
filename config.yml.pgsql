base:
    is_dev : false
out_dir : ./model  # 输出目录
url_tag : json # web url tag(json,db(https://github.com/google/go-querystring))
language :  # 语言(English,中 文)
db_tag : gorm # 数据库标签名(gorm,db)
simple : true # 简单输出(默认只输出gorm主键和字段标签)
user_gorm_model : false # model是否使用gorm.Model
is_db_tag : true # 是否输出 数据库标签(gorm,db)
is_out_sql : false # 是否输出 sql 原信息
is_out_func : true # 是否输出 快捷函数
is_web_tag : true # 是否打web标记(json标记前提条件)
is_web_tag_pk_hidden: true # web标记是否隐藏主键
is_foreign_key : true # 是否导出外键关联
is_gui : false # 是否ui模式显示
is_table_name : true # 是否直接生成表名
is_column_name : true # 是否直接生成列名
is_null_to_point : false # 数据库默认 'DEFAULT NULL' 时设置结构为指针类型
is_null_to_sql_null: false # 数据库默认 'DEFAULT NULL' 时设置结构为sql.NULL  is_null_to_point如果为true，则is_null_to_sql_null不生效
table_prefix : "" # 表前缀, 如果有则使用, 没有留空(如果表前缀以"-"开头，则表示去掉该前缀，struct、文件名都会去掉该前缀)
table_names: "" # 指定表生成，多个表用,隔开
is_out_file_by_table_name: false # 是否根据表名生成多个model
is_out_page: true # 是否输出分页函数 

db_info_bak:
    host : 106.55.106.191 # type=1的时候，host为yml文件全路径
    port : 26257
    username : iotsaas
    password : 
    database : gfast
    schema	: public
    type: 3 # 数据库类型:0:mysql , 1:sqlite , 2:mssql,3:pgsql
db_info:
    host : 127.0.0.1 # type=1的时候，host为yml文件全路径
    port : 26257
    username : root
    password : 
    database : gfast
    schema	: public
    type: 3 # 数据库类型:0:mysql , 1:sqlite , 2:mssql,3:pgsql
self_type_define: # 自定义数据类型映射
    datetime: time.Time
    time: time.Time
    ^(int)[(]\d+[)]: int
out_file_name: "" # 自定义生成文件名
web_tag_type: 0 # json tag类型 0: 小驼峰 1: 下划线

# sqlite
# db_info:
#     host : /Users/xxj/Downloads/caoguo # type=1的时候，host为yml文件全路径
#     port : 
#     username : 
#     password : 
#     database : 
#     type: 1 # 数据库类型:0:mysql , 1:sqlite , 2:mssql