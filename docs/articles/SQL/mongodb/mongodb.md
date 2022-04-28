## MongoDB



## MongoDB连接

```go
mongodb://[username:password@]host1[:port1][,host2[:port2],...[,hostN[:portN]]][/[database][?options]]
```

Go:

使用mgo.v2, `go get gopkg.in/mgo.v2`

```go
import (
  mgo "gopkg.in/mgo.v2"
)
url := "mongo://:@127.0.0.1:27017/model?connect=direct"
session, err := Dial(url)

dial := session.Clone()
```



## 操作

1. 创建数据库

   `use dbname`创建数据库，若不存在，则创建。

2. 删除数据库

   `db.dropDatabase()` `db`即当前数据库，可以用`db`命令查看当前数据库。

3. 创建集合

   `db.createCollection(name, options)`

   `name`: 要创建的集合的名称

   `options:` 可选参数，指定有关内存大小即索引的选项。

   比如`autoIndexID` 如为true,自动在`_id`字段创建索引，默认为false.

4. 删除集合

   `db.collection.drop()` 若成功，返回true,否则返回false.

5. 插入文档

   文档的数据结构和JSON基本一致。

   所有存储在集合中的数据都是BSON格式。

   BSON是一种类似JSON的二进制形式的存储格式，是Binary JSON的简称。

   使用insert()/save()方法向集合中插入文档：

   `db.COLLECTION_NAME.insert(document)`

   插入多条：

   `db.collection.insertMany([{"b":3}, {'c':4}])`

6. 更新文档

   `update`方法

   ```go
   db.collection.update(
      <query>,
      <update>,
      {
        upsert: <boolean>,
        multi: <boolean>,
        writeConcern: <document>
      }
   )
   ```

   * **`query`**: update 的查询条件
   * **update** : update的对象和一些更新的操作符（如$,$inc...）等，也可以理解为sql update查询内set后面的
   * **upsert** : 可选，这个参数的意思是，如果不存在update的记录，是否插入objNew,true为插入，默认是false，不插入。
   * **multi** : 可选，mongodb 默认是false,只更新找到的第一条记录，如果这个参数为true,就把按条件查出来多条记录全部更新。
   * **writeConcern** :可选，抛出异常的级别。

7. 删除文档

   `db.collection.deleteMany({})`删除此集合下所有文档。

   `db.collection.deleteMany({xxx:xxx})` 删除 xxx等于xxx的全部文档

   `db.collection.deleteOne({xx:xx})`

8. 查询文档

   `db.collection.find()`

   * And条件

   `find({"ada":xxx, "sasda":xxx})`

   * Or条件

   `find($or[{},{}])`

   * 比较大小

   `{<key>:<value>}` =

   `{<key>:{$lt:<value>}}` <

   lte gt gte ne

   *  

   查询 title 包含"教"字的文档：

   ```
   db.col.find({title:/教/})
   ```

   查询 title 字段以"教"字开头的文档：

   ```
   db.col.find({title:/^教/})
   ```

   查询 titl e字段以"教"字结尾的文档：

   ```
   db.col.find({title:/教$/})
   ```

