## 并行循环

这里我们探讨一些通用的并行模式，来并行执行所有的循环迭代。需要解决的问题是，生成一批全尺寸图像的缩略图。

[thumbnail.go](example/thumbnail/thumbnail.go)

```go
package thumbnail
// ImageFile从infile中读取一个图像并将它的缩略图写入同一个目录中
// 它返回生成的文件名，例如"foo.thumb.jpg"
func ImageFile(infile string) (string, error) { /* ... */ }
```

  * 第一版程序

  ```go
  func makeThumbnails(filenames []string) {
    for _, f := range filenames {
      if _, err := thumbnail.ImageFile(f); err != nil {
        log.Println(err)
      }
    }
  }
  ```

  这是简单的情况，直接在文件名字slice上循环，然后给每一个图像生成一个缩略图。

  * 第二版程序

  ```go
  func makeThumbnails2(filenames []string) {
    for _, f := range filenames {
      go thumbnail.ImageFile(f)
    }
  }
  ```

  这次，我们为每个图像的处理，创建一个goroutine，可以提高效率。注意，**它启动了所有的goroutine，每一个图像对应一个，但是没有等它们执行完毕**。

  * 第三版程序

  ```go
  func makeThumbnails3(filenames []string) {
    ch := make(chan struct{})
    for _, f := range filenames {
      // 同样，为每个图像处理启动一个goroutine
      go func(f string) {
        thumbnail.ImageFile(f)
        ch <- struct{}{}
      }(f)
    }
    // 等待goroutine完成
    for range filenames {
      <-ch
    }
  }
  ```

  * 第四版程序

  ```go
  // makeThumbnails4为指定文件列表并行地生成缩略图
  // 如果任何步骤出现，则返回一个错误
  func makeThumbnails4(filenames []string) {
    errors := make(chan error)
    for _, f := range filenames {
      go func(f string) {
        _, err := thumbnail.ImageFile(f)
        // 将错误信息写入errors通道
        errors <- err
      }(f)
    }
    for range filenames {
      // 处理errors通道接收到的err
      if err := <- errors; err != nil {
        return err  // 因为只有主goroutine处理err信息，造成goroutine泄漏
      }
    }
    return nil
  }
  ```

  这个版本虽然进行错误处理，但是仍然有问题。当遇到第一个非nil错误时，它将错误返回给调用者，这里并没有goroutine继续从errors返回通道上的进行接收，直到读完。那么，**每一个现存的工作goroutine在试图发送值到此通道的时候将永远阻塞，永不终止**。这个情况下goroutine泄漏可能导致整个程序卡住或系统内存耗尽。

  * 第五版程序

  使用缓冲通道处理第四版程序的的问题。

  ```go
  // makeThumbnails5为指定文件并行地生成缩略图
  // 它以任意顺序返回生成的文件名
  // 如果任何步骤出错就返回一个错误
  func makeThumbnails5(filenames []string) (thumbfiles []string, err error) {
    type item struct {
      thumbfile string
      err error
    }

    // 创建一个通道ch，用于接收生成的文件名和错误信息
    ch := make(chan item, len(filenames))

    for _, f := range filenames {
      go func(f string) {
        var it item
        it.thumbfile, it.err = thumbfile.ImageFile(f)
        ch <- it
      }(f)
    }

    // 通道的接收操作
    for range filenames {
      it := <-ch
      if it.err != nil {
        return nil, it.err
      }
      thumbfiles = append(thumbfiles, it.thumbfile)
    }
    return thumbfiles, nil
  }
  ```

  * 最终版程序

  这个版本中返回新文件所占用的总字节数。而且，它不是使用slice接收文件名，而是借助 *一个字符串通道*，因此，我们也就无法知道迭代次数，需要单独处理。

  为了知道什么时候最后一个goroutine结束，需要 **在每一个goroutine启动前递增计数，在每一个goroutine结束时递减计数**。这就需要一个特殊类型的计数器，它可以被多个goroutine安全地操作，然后有一个方法一直等到它变为0.

  这个计数器类型就是`sync.WaitGroup`

  ```go
  func makeThumbnails6(filenames <-chan string) int64 {
    // 创建一个chan int64类型的通道，用来接收每个文件处理后返回的大小
    sizes := make(chan int64)

    var wg sync.WaitGroup  // 负责为每个goroutine计数
    for f := range filenames {
      wg.Add(1)
      // worker
      go func(f string) {
        defer wg.Done()   // 递减计数，等价于Add(-1)
        thumb, err := thumbnail.ImageFile(f)
        if err != nil {
          log.Println(err)
          return
        }
        info, _ := os.Stat(thumb)
        // 将生成新文件的大小写入size通道
        sizes <- info.Size()
      }(f)
    }

    // closer
    go func() {
      wg.Wait()
      // 关闭sizes通道，这样主goroutine就可以使用range sizes
      close(size)
    }()

    // 主goroutine
    // 计算总文件大小
    var total int64
    for size := range sizes {
      total += size
    }
    return total
  }
  ```

  这个版本的程序中，sizes通道将每一个新生成的文件大小带回主goroutine，然后使用range循环进行累计求和。

  注意，在closer的这个goroutine中，在关闭sizes通道前，等待所有工作结束。等待和关闭两个操作必须和在size通道上面的迭代(range sizes)并行执行。

  如果我们将等待操作放到循环之前的主goroutine中，因为通道会满，它将永远不能结束；如果放在循环之后，它将不可达，因为没有任何东西可用来关闭通道，循环可能永不结束。

  ![事件序列](http://on64c9tla.bkt.clouddn.com/Go/thumbnail.jpg)

  记住：**主程序也是一个goroutine**。
