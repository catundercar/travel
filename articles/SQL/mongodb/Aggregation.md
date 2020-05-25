# Aggregation

## Map-Reduce

### 介绍

Map-reduce是一种数据处理范例，用于将大量数据压缩为有用的聚合结果。

* 编程模型

  该计算任务将一个键值对集合作为输入，并生成一个键值对集合作为输出。MapReduce这个库的用户将这种计算任务以两个函数进行表达，即**Map**和**Reduce**。

  由用户所编写的**Map**函数接收输入，并生成一个中间键值对集合。MapReduce这个库会将所有共用一个键的值组合在一起，并将它们传递给**Reduce**函数。

  **Reduce**函数也是由用户所编写。它接受一个中间键以及该键的值的集合作为输入。它会将这些值合并在一起，以此来生成一组更小的值的集合。通常每次调用**Reduce**函数所产生的值的结果只有0个或者1个。中间值通过一个迭代器来传递给用户所编写的**Reduce**函数。这使得我们可以处理这些因为数据量太大而无法存放在内存中的存储值的list列表了。

例如：

从大量的文档中计算出每个单词的出现次数，伪代码如下:

```c++
map(String key, String value):
// key: document name
// value: document contents
for each word w in value:
EmitIntermediate(w,"1");
reduce(String key, Iterator values):
// key: a word
// values: a list of counts
int result = 0;
for each v in values:
result += ParseInt(v);
Emit(AsString(result));
```

### 食用方式

在MongoDB中，map-reduce由三部分组成，map函数、reduce函数、结果。

命令格式如下:

```
db.runCommand(
               {
                 mapReduce: <collection>,
                 map: <function>,
                 reduce: <function>,
                 finalize: <function>,
                 out: <output>,
                 query: <document>,
                 sort: <document>,
                 limit: <number>,
                 scope: <document>,
                 jsMode: <boolean>,
                 verbose: <boolean>,
                 bypassDocumentValidation: <boolean>,
                 collation: <document>,
                 writeConcern: <document>
               }
             )
```

其中，map函数和reduce函数使用js语法。另外，`mapReduce:<collection>`可直接用`db.Collection.mapReduce()`.

例如：

分析以下map-reduce操作：

```
db.orders.mapReduce(
	function() { emit( this.cust_id, this.amount ); },
	function(key, values) { return Array.sum( values ) },
	{
		query: { status: "A" },
		out: "order_totals"
	}
)
```

![map-reduce](/Users/xbrother/Desktop/map-reduce.bakedsvg.svg)



## Aggregation Pileline

### 介绍

聚合管道是一个基于数据处理管道概念的数据聚合框架。文档进入多阶段管道，该管道将文档转换成汇总结果。

### 食用方式

命令格式:

```
{
  aggregate: "<collection>" || 1,
  pipeline: [ <stage>, <...> ],
  explain: <boolean>,
  allowDiskUse: <boolean>,
  cursor: <document>,
  maxTimeMS: <int>,
  bypassDocumentValidation: <boolean>,
  readConcern: <document>,
  collation: <document>,
  hint: <string or document>,
  comment: <string>,
  writeConcern: <document>
}
```

大多数用户选择`db.collection.aggregate()`。