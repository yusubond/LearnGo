### 并发目录遍历

我们构建这样一个程序，根据命令行指定的输入，报告一个或多个目录的磁盘使用情况，类似于UNIX du命令。

```go
// wakjDir递归遍历以dir为根目录的整个文件树
// 并在fileSizes上发送每个已经找到的文件的大小
func wakjDir(dir string, fileSizes chan<- int64) {
  for _, entry := range dirents(dir) {
    if entry.IsDir() {
      subdir := filepath.Join(dir, entry.Name())
      wakjDir(subdir, fileSizes)
    } else {
      fileSizes <- entry.Size()
    }
  }
}
```

对于每一个子目录，wakjDir函数递归调用自己，对于每一个文件，wakjDir发送一条消息(即，文件所占用的字节数)到fileSizes通道。

```go
// dirents返回dir目录中的条目
func dirents(dir string) []os.FileInfo {
  entries, err := ioutil.ReadDir(dir)
  if err != nil {
    fmt.Fprintf(os.Stderr, "du1: %v\n", err)
    return nil
  }
  return entries
}
```

ioutil.ReadDir函数返回一个os.FileInfo类型的slice。对于单个文件，可以通过调用os.Stat函数获得。
