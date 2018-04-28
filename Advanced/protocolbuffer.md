## Protocol Buffer

  * [安装](#安装)
  * [ProtoBuf语法](ProtoBuf语法)
  * [示例文件](#示例文件)

Protocol Buffer，简称ProtoBuf，一套完整的IDL(接口描述语言)。ProtoBuf由google设计，基于C++进行实现，开发人员可以根据ProtoBuf的语言规范生成多种编程语言(例如Python, Java, Go)的接口代码。

这篇文章主要讲述如何在Golang中使用ProtoBuf。

### 安装

1. 根据系统的选择合适zip文件，ProtoBuf下载地址：[https://github.com/google/protobuf/releases](https://github.com/google/protobuf/releases)

以OSX系统为例，将zip文件减压后，分别将`/bin/protoc`拷贝至`/usr/local/bin`目录下，`include`拷贝至`/usr/local/bin/`。

```shell
$ cp bin/protoc /usr/local/bin
$ cp include /usr/local/bin/
# 拷贝完成后，source一下环境变量
# 检查是否安装成功
$ protoc --version
```

2. 安装ProtoBuf相关的golang依赖库

```go
$ go get -u github.com/golang/protobuf/protoc-gen-go
```

### ProtoBuf语法

ProtoBuf文件以`.proto`命名，例如[addressbook.proto](https://github.com/yusubond/LearnGo/blob/master/ProtoBuf/src/addressbook/addressbook.proto)。

通过命令`protoc --go_out=/output/path filename.proto`可以将proto文件生成golang源码，生成的文件名字以`pb.go`结尾。

```protobuf
syntax = "proto3";

package addressbook;

message Person {
    string name = 1;
    int32 id = 2;

    // 枚举类型的第一个字段必须为0
    enum PhoneType {
      MOBILE = 0;
      HOME = 1;
      WORK = 2;
    }

    message Phone {
      string number = 1;
      PhoneType type = 2;
    }

    repeated Phone phones = 3;
}

message AddressBook {
    repeated Person persons = 1;
}
```

`syntax`字段：

    指定protobuf的版本

`package`字段：

    指定包名，即通过protoc生成go源码后的包名

`message`字段：

    定义一个消息类型，允许嵌套；     
    后面的`1`,`2`等数字表示标识号，须从`1`开始    

`enum`字段：

    定义枚举类型，第一个字段必须为`0`  

`repeated`字段：

    表示当前变量可重复，允许多个

通过对比`*.proto`与`*.pb.go`可以发现`proto`文件中定义的消息类型与`golang`语言对应关系。

例如：

proto中`message Person{}`对应golang中的`type Person struct{}`，另外三个变量(`name`, `id`, `phones`)分别对应struct中的三个成员(`Name`, `Id`, `Phones`)

proto中`message Phone{}`作为`message Person{}`的子成员,在golang中则为`type Person_Phone struct{}`，即 *通过下划线的方式来别是成员间的包含关系*。

枚举类型`PhoneType`在golang中成为`type Person_PhoneType int32`类型，其成员分别成为常量(`Person_HOME`, `Person_WORK`, `Person_MOBILE`)。

其次，对应每一种`message`转化为golang源码后，分别生成一些常用的方法，例如`Reset()`, `String()`,`GetName()`等方法。

## 示例文件

  * [addressbook.proto](ProtoBuf/src/addressbook/addressbook.proto)    
  * [addressbook.pb.go](ProtoBuf/src/addressbook/addressbook.pb.go)    
