> **创建日期:** 2021/11/15
> 
> **更新日期:** 2021/11/21
> 
> **作者:** 道五七

- [元信息](#元信息)
- [类型转换](#类型转换)
- [变量](#变量)
- [函数](#函数)
- [作用域](#作用域)
- [流程控制](#流程控制)
- [模块](#模块)
- [指针](#指针)
- [迭代器](#迭代器)
- [数据结构: 引用传递](#数据结构-引用传递)
- [异步](#异步)
- [其他](#其他)
- [面向对象](#面向对象)
- [库](#库)
- [CMD](#cmd)

# 元信息
- 变量构成: pair
  - type
    - static type
    - concrete type
  - value
- 结构体属性标签
# 类型转换
- 手动转换
  - `string(<var>)`
# 变量
- 类型
  - `<nil>` `func()` `byte` `bool` `int` `int8` `int16` `int32` `int64` `float32` `float64` `string` `[<num>]<type>` `chan <type>`
  - `[]<type>` `map[<type>]<type>`
  - `struct`
  - `interface` `interface{}`
- 声明
  - `var|const <var...> [<type>] [= <val...>]`
  - `var|const ( ... ... )`
  - `const ( iota )`: 枚举
  - `<var> := <val>`: 只能在函数作用域内声明
  - `func <funcName>(arg <type>, ...) <type>|(<type>, ...) { <exe> }`
- 字面量
  - func() {}
  - 3, 3.14, true, false, "dao77777", [5]int{1, 2, 3}
  - []string{"ha"}, map[string]int{...}
  - struct{}
# 函数
- 兜底: defer, 执行语句之前的关键字, 栈形式, 在函数最后执行
- 异常: `<val>, <err> = <func>`
# 作用域
- 全局作用域
- 函数作用域
- 块级作用域
- 作用域链
- 闭包
# 流程控制
- 循环: `for <var init>; <loop condition>; <loop behavior> { <exe> }`
- 分支: `if <condition> { <exe> }`
# 模块
- 包名: `package <packageName>`
- 引入: `import (<alias> <packageName>...)`
  - 初始化引入: alias为_
  - 全量引入: alias为.
  - 正常引入: alias为别名
- 暴露: 变量大写对外暴露, 变量小写私有变量
- 初始化: init()
- 主函数: main()
# 指针
- `&<var>`: 取地址
- `*<var>`: 解引用
# 迭代器
- `len(<iterator>)`: 返回迭代器长度
- `cap(<iterator>)`: 返回迭代器容量
- `for <index>, <val> := range <iterator> { <exe> }`
# 数据结构: 引用传递
- slice: `[]<type>{elem}`
  - 增: `append(<elem>)`: 给slice增加长度, 若超出容量, 以一个cap为单位扩容
  - 删
  - 改: `<var>[<num>] = <elem>`
  - 查: `<var>[<num>]`
  - 截取: `<var>[左闭:右开]`, 返回的为浅拷贝
- map: `map[<type>]<type>{<key>:<value>, ...,}`
  - 增: `<var>[<key>] = <val>`
  - 删: `delete(<var>, <key>)`
  - 改: `<var>[<key>] = <val>`
  - 查: `<var>[<key>]`
# 异步
- 创建: go <func>
- 同步与通信
  - 无缓冲channel: make(chan Type)
  - 有缓冲channel: make(chan Type, capacity)
  - 释放: channel <- value
  - <- channel
  - x := <- channel
  - x, ok := <- channel
- IO: select
# 其他
- 分配
  - slice: `make(<type>, <len>[, <cap>])`
  - map: `make(map[<type>]<type>, <len>)`
  - channel: `make(chan <type>)`
- 深拷贝: `copy(<dest>, <src>)`
# 面向对象
- `<type>{<propert>, ...}`
- 公有私有: 属性或方法大写公有, 属性或方法小写私有
- 封装
- 继承
- 多态: interface
  - 空接口: interface{}
  - 类型断言: `<var>.(<type>)`可返回`val, ok`
# 库
- fmt
  - `Println`
  - `Printf`: %d, %s, %v, %T
  - `Sprintf`
- strings
  - `Split(<string>, <split>)`
- io
  - `EOF`
- os
  - `O_RDWR`
  - `OpenFile(<path>, <right>, <>)`
- runtime
  - `Goexit()`
- time
  - `Second`
  - `Sleep(<second>)`
- reflect
  - `TypeOf(<var>)`
  - `ValueOf(<var>)`
- json
  - `Marshal(<struct>)`
  - `Unmarshal(<json>, <struct>)`
- net
  - `Listen(<network>, <address>)`
- sync
  - RWMutex
# CMD
- 构建: `go build <filename>`, `-o <目标文件> <源文件...>`
- 运行: `go run <filename>`
- 包管理
  - GOPATH: `go get -u <package_name>`
  - GOMODULES
    - go mod init: 生成go.mod文件
    - go mod download: 下载go.mod文件中知名的依赖
    - go mod tidy: 整理现有依赖
    - go mod graph: 查看现有依赖
    - go mod edit: 编辑go.mod文件
    - go mod vendor: 导出项目的所有依赖到vendor目录
    - go mod verify: 校验一个模块是否被篡改
    - go mod why: 查看为什么需要依赖某模块
    - GO111MODULE
    - GOPROXY
    - GOSUMDB
    - GONOPROXY
    - GONOSUMDB
    - GOPRIVATE
- 配置
  - 查看配置(环境变量): `go env`
  - 设置配置(环境变量): `go env -w <key>=<val>`