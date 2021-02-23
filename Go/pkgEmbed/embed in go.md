# pkg embed in go

embed包 提供了访问正在运行的go程序的功能。

* 嵌入一个文件到string

```go
import _ "embed"

//go:embed hello.txt
var s string
print(s)
```

* 嵌入一个文件到一个字节切片

```go
import _ "embed"

//go:embed hello.txt
var b []byte
print(string(b))
```

* 嵌入一个或多个文件作为一个文件系统

```go
import "embed"

//go:embed hello.txt
var f embed.FS
data, _ := f.ReadFile("hello.txt")
print(string(data))
```



## usage

变量声明上面的`go:embed`指令使用一个或多个路径指定要嵌入的文件。匹配模式。

该指令必须紧接在包含单个变量声明的行之前。 指令和声明之间仅允许使用**空行**和**`//`**行注释。

变量的类型必须是这三种的其中一个`string`, `[]byte`, `FS`(或者FS的别名)。

```go
package server

import "embed"

// content holds our static web server content.
//go:embed image/* template/*
//go:embed html/index.html
var content embed.FS
```

`//go:embed`指令为简洁起见接受多个空格分隔的模式，但它也可以重复，以避免有许多模式时出现很长的行。模式是相对于包含源文件的包目录进行解释的。即使在windows上，路径分隔符也是`/`(forward slash)。模式不能包含`.`或`..`或空路径元素，也不能以`/`开头或者结尾。使用`*`来匹配当前路径下的所有内容。为了允许命名名称中有空格的文件，可以将模式写为Go的双引号`""`或者反引号```` ``字符串文本

如果一个模式命名为一个目录，则该目录下的子树中所有文件都将被递归嵌入，但是名称以`.`（隐藏文件）或`\`开头的文件除外。所以上面例子中的变量几乎等价于：

```go
// content is our static web server content.
//go:embed image template html/index.html
var content embed.FS
```

区别就是`image/*`嵌入了`image/.tempfile`,但是`image`没有。

`//go:embed`指令可以与导出和未导出的变量一起使用，具体取决于软件包是否希望使数据可用于其他软件包。 它只能与程序包范围的全局变量一起使用，而**不能与局部变量一起使用**。

```go
func Partial() {
    //go:embed test.txt
    vat a string
}
```

编译将会报错：

```go
//go:embed only allowed in Go files that import "embed"
```

模式不得与包模块外部的文件匹配，例如` .git/*`或符号链接。 空目录的匹配将被忽略。 之后，`//go:embed`行中的每个模式必须至少匹配一个文件或非空目录。

如果**任何模式无效或匹配无效**，则构建将失败。

### File System

`FS`实现了`io/fs`包的`FS`接口，所以它能在任何理解文件系统的包内使用,包括`net/http`, `text/template`, 和`html/template`。

```go
http.Handle("/static/",http.StripPrefix("/static/",http.FileServer(http.FS(content))))

template.ParseFS(content, "*.tmpl")
```



## 无效模式

https://github.com/golang/go/issues/44486

embed的有效文件名不能为`.git`, `.hg`, `.bzr`, `.svn`，并且不能包含``* < > ? ` ' | / \ : ``等在某些Shell或操作系统中具有特殊含义的符号。

```go
// isBadEmbedName reports whether name is the base name of a file that
// can't or won't be included in modules and therefore shouldn't be treated
// as existing for embedding.
func isBadEmbedName(name string) bool {
	if err := module.CheckFilePath(name); err != nil {
		return true
	}
	switch name {
	// Empty string should be impossible but make it bad.
	case "":
		return true
	// Version control directories won't be present in module.
	case ".bzr", ".hg", ".git", ".svn":
		return true
	}
	return false
}

// CheckFilePath checks that a slash-separated file path is valid.
// The definition of a valid file path is the same as the definition
// of a valid import path except that the set of allowed characters is larger:
// all Unicode letters, ASCII digits, the ASCII space character (U+0020),
// and the ASCII punctuation characters
// “!#$%&()+,-.=@[]^_{}~”.
// (The excluded punctuation characters, " * < > ? ` ' | / \ and :,
// have special meanings in certain shells or operating systems.)
//
// CheckFilePath may be less restrictive in the future, but see the
// top-level package documentation for additional information about
// subtleties of Unicode.
func CheckFilePath(path string) error {
	if err := checkPath(path, true); err != nil {
		return fmt.Errorf("malformed file path %q: %v", path, err)
	}
	return nil
}
```

如果非要支持带特殊字符的话，可以通过重新编译go二进制来实现。

