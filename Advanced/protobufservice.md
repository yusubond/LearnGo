## ProtoBuf语法

### Message Type

```ProtoBuf
syntax = "proto3"

message SearchRequest{
  string query = 1;
  int32 page = 2;
  int32 result = 3;
}
```

`1`,`2`成为分配标签，从`1`开始。

### 
