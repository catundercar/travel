# SQLite

## 常用命令

* .help: 帮助
* .databases: 列出数据库
* .tables: 列出表名
* .open dbname 打开数据库
* .save dbname 保存为数据库
* .exit 退出
* .schema[tbname] 列出表，索引，触发器的创建语句。
* .output frame.txt 写结果到文件
* .show 显示各种设置的默认值
* .indices tbnames  列出某表的索引



## 数据类型

| 存储类  | 描述 |
| ------- | ---- |
| NULL    |      |
| INTEGER |      |
| REAL    |      |
| TEXT    |      |
| BLOB    |      |

INT CHAR TEXT FLOAT DOUBLE BOOLEAN DATE DATETIME TIMESTAMP

## 常用操作

* sqlite注册模糊功能

```go
regex := func(re string, s interface{}) (bool, error){
  switch x := s.(type) {
    case string:
    	return regexp.MatchString(re, x)
    default:
    	return false, nil
  }
}

sql.Register("sqlite3_with_go_func",
             &sqlite3.SQLiteDriver{
               ConnectHook: func(conn *sqlite3.SQLiteConn) error {
                 return conn.RegisterFunc("REGEXP", regex, false)
               },
             })
```



* 初始化索引

create index idx_xxx on tablename(fieldname)

### SQL Json

* Function Details
* * json_extract()

```sql
json_extract('{"a":2,"c":[4,5,{"f":7}]}', '$') → '{"a":2,"c":[4,5,{"f":7}]}'
```

