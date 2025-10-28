# [模块](https://golang.halfiisland.com/essential/senior/115.module.html)

### 一、[开发文档](https://golang.halfiisland.com/essential/senior/115.module.html#%E5%88%9B%E5%BB%BA)

#### 1. 创建项目和初始化模块

```shell
mkdir hello
cd hello

# 初始化
go mod init github.com/OneHundred86/hello

```

#### 2. 编写代码和测试


#### 3. 编写命令行程序（非必要）

对于命令行程序而言，按照规范是在项目的 `cmd/<cmd-name>/` 目录中进行创建；

开发完成代码后，测试：

```shell

go run cmd/hello/*.go --help

go run cmd/hello/*.go -name jack -cnt 2

```



### 二、使用文档

#### 1. Install code

```bash

go get github.com/OneHundred86/hello

```

#### 2. Install cmd

```bash
# 安装到 $GOPATH/bin 目录
go install github.com/OneHundred86/hello/cmd/hello


```

#### 3. 使用示例

```go
package main

import (
  "fmt"
  "github.com/OneHundred86/hello"
)

func main() {
  result := hello.SayHelloTo("jack")
  fmt.Println(result)
}

```
