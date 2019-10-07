# 服务计算 - selpg

本程序使用 go modules 管理源代码和依赖包。

## 实验过程

### 创建项目

通过使用 `go mod init` 即可创建项目如下：

```bash
go mod init github.com/huanghongxun/selpg
```

### 编写代码

#### 参数处理

使用 `pflag` 包，可以大大方便我们处理命令行参数：

```go
pflag.IntVarP(&startPage, "start-page", "s", -1, "start page")
pflag.IntVarP(&endPage, "end-page", "e", -1, "end page")
pflag.IntVarP(&pageLen, "page-len", "l", 72, "page length")
pageTypeFlag := pflag.BoolP("form-feed-delimited", "f", false, "form-feed-delimited")
pflag.StringVarP(&printDest, "print-dest", "d", "", "printer destination")
pflag.Parse()

if *pageTypeFlag {
    pageType = "f"
} else {
    pageType = "l"
}
```

#### 参数校验

为了让程序具有更好的容错性，我们需要检查参数输入是否合法，包括：

1. 参数是否足够
2. 起始页是否小于结束页
3. 页面是否为正数
4. 输入文件是否存在

```go
if startPage == -1 || endPage == -1 {
    usage() // 参数不够，打印程序使用方法
}

if startPage <= 0 || endPage <= 0 {
    return fmt.Errorf("%s", "start page and end page should be positive")
}

if startPage > endPage {
    return fmt.Errorf("%s", "End page should be greater then start page")
}

if pflag.NArg() == 1 {
    inFileName = pflag.Arg(0)
} else if pflag.NArg() > 1 {
    return fmt.Errorf("%s", "Too many arguments")
} else {
    inFileName = ""
}
```

#### 数据输入

我们根据输入数据的不同产生不同的 `reader`，通过 `reader` 抽象文件输入和标准输入可以简化我们之后的处理代码。

```go
if len(inFileName) > 0 {
    fin, err := os.Open(inFileName)
    if err != nil {
        return fmt.Errorf("selpg: could not open input file \"%s\". Reason:\n%s\n", inFileName, err.Error())
    }
    defer fin.Close()
    reader = bufio.NewReader(fin)
} else {
    reader = bufio.NewReader(os.Stdin)
}
```

#### 数据输出

我们根据不同的输出数据方法产生不同的 `writer`，通过 `writer` 抽象打印输出和标准输出可以简化我们之后的处理代码。

```go
if len(printDest) > 0 {
    cmd := exec.Command("lp", fmt.Sprintf("-d%s", printDest))
    stdinPipe, err := cmd.StdinPipe()
    if err != nil {
        return fmt.Errorf("selpg: could not open pipe to \"%s\". Reason:\n%s\n", printDest, err.Error())
    }
    defer stdinPipe.Close()
    cmd.Stdout = os.Stdout
    if err := cmd.Start(); err != nil {
        return err
    }
    writer = bufio.NewWriter(stdinPipe)
} else {
    writer = bufio.NewWriter(os.Stdout)
}
defer writer.Flush()
```

#### 分页法

对于 `-l` 参数，我们每读入一行，就检查是否是需要打印的行，如果是就输出到 `writer` 中：

```go
line := bufio.NewScanner(reader)
for line.Scan() {
    lineCtr++
    if lineCtr > pageLen {
        pageCtr++
        lineCtr = 1
    }
    if pageCtr >= startPage && pageCtr <= endPage {
        if _, err := writer.Write(line.Bytes()); err != nil {
            return err
        }
        if err := writer.WriteByte('\n'); err != nil {
            return err
        }
    }
}
```

#### 分隔符法

对于 `-f` 参数，我们每读到 `\f` 就检查是否是需要打印的行。

```go
for {
    c, err := reader.ReadByte()
    if err == io.EOF {
        break
    } else if err != nil {
        return err
    }
    if c == '\f' {
        pageCtr++
    }

    if pageCtr >= startPage && pageCtr <= endPage {
        if err := writer.WriteByte(c); err != nil {
            return err
        }
    }
}
```

#### 检查是否产生了输出

我们还需要检查 `start-page` 和 `end-page` 范围内是否有输出，否则范围不合法，需要提醒用户

```go
if pageCtr < startPage {
    return fmt.Errorf("start_page (%d) greater than total pages (%d), no output written\n", startPage, pageCtr)
} else if pageCtr < endPage {
    return fmt.Errorf("end_page (%d) greater than total pages (%d), less output than expected\n", endPage, pageCtr)
}
```

## 实验结果

### 使用说明

![usage](assets/usage.PNG)

### 测试结果

![run](assets/run.PNG)