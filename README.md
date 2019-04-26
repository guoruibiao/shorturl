## shorturl [![shorturl-build](https://travis-ci.org/guoruibiao/shorturl.svg?branch=master)](https://travis-ci.org/guoruibiao/shorturl/)
`shorturl`是一个`demo`，其实是用来测试`travis-ci`的集成效果。上面的**build history** 记录了每次的`build`历程。下面一点点来看看一个`golang`项目是如何添加持续集成的吧。

## 1 `Travis-ci`支持
Travis-ci的工作流程大致有这么几步：
- 从关联了权限的GitHub仓库中拉取代码
- 根据.travis.yml中设置好的环境建造容器
- 把repository扔到容器中，按照定制的stages一点点去执行
- 反馈集成结果

所以在决定给某一个项目添加`travis`集成之前，我们是需要先有一个`repository`的，然后在 https://travis-ci.org 上关联`GitHub`对应的仓库，打开持续集成开关即可。

## 2 `.travis.yml`
第一部分说到 `Travis-CI`会根据`.travis.yml` 设置好的环境参数去创建对应的容器，可见这个文件是很重要的。一个待持续集成的项目根目录下必然需要提供一个这样的文件，官网有一个教程：https://docs.travis-ci.com/user/languages/go/ 不过我看的云里雾里的，还是从[阮一峰-travis-ci持续集成指南](http://www.ruanyifeng.com/blog/2017/12/travis_ci_tutorial.html) 中大致摸索了一个结构。


```ymal
# .travis.yml
language:    # 指明使用的语言
  - go 

go:          # 语言版本号
  - "1.x"    # x表示对应前缀的最新版本
  - "1.10"   # 注意，需要 "1.10" 版本的时候必须表示为字符串形式，如果写成 1.10 则会使用 1.1 版本
  - "1.10.x"
  - master   # 默认使用最新版本

script:      # 执行的脚步，但是go默认会执行下面的这些命令的，所以可以不用写
  - go get -v
  - go test -v ./...
```

## 3 Hello World
架子搭好了，下面先来一个简单的测试。还好有`Git`，可以方便的乘坐时光机到某一个版本的快照，具体的代码可以参考：https://github.com/guoruibiao/shorturl/tree/298b7bfc978cda8dc4ae6a53c0186e9a6060c53a

对应的`.travis.yml` 内容也在里面了，这里先单独摘出来，如下：
```ymal
# .travis.yml
language:    # 指明使用的语言
  - go 

go:          # 语言版本号
  - "1.11.5"   # 默认使用最新版本,注意，需要 "1.10" 版本的时候必须表示为字符串形式，如果写成 1.10 则会使用 1.1 版本;x表示对应前缀的最新版本

script:      # 执行的脚步，但是go默认会执行下面的这些命令的，所以可以不用写
  - go get -v
  - go test ./...
```

通过push提交代码后，Travis-ci会受到hook通知，开启build， 对应的构建号为：[524402743](https://travis-ci.org/guoruibiao/shorturl/builds/524402743) 

详细的构建日志可以参考链接为： [build-log](https://api.travis-ci.org/v3/job/524402744/log.txt)

至此，一个简单的集成测试流程就算走通了。但是如果只是这样，持续集成也就有点太鸡肋了不是，下面一点点来丰富它。

## 4 短链服务
短链接服务在某些比如短信，邮件，微博等限制字数的业务场景下还是使用的比较频繁的，下面借助新浪短链的服务来代理下，做一个简单的短链生成服务。

代码：https://github.com/guoruibiao/shorturl/tree/8c5b567a2a72f6b79572445cf7026c0d0f46c92f

本地测试
```golang
func Test_SinaURLShort(t *testing.T) {
	originurl := "https://github.com"
	dao, _ := New()
	response, _ := dao.SinaURLShort(originurl)
	if response.URLS[0].ShortURL == "http://t.cn/RxnlTYR" {
		t.Log("SinaURLShort passed")
	} else {
		t.Error("SinaURLShort failed", originurl, " to ", response.URLS[0].ShortURL)
	}
}
```
输出内容：
```
Running tool: /usr/local/Cellar/go/1.11.5/libexec/bin/go test -timeout 30s shorturl/dao -run ^(Test_SinaURLShort)$

ok  	shorturl/dao	0.631s
Success: Tests passed.
```

## 5 添加持续集成
本地代码编写，测试通过了，但不是说就一定可以通过集成测试套件，因为我就发现了一个问题。

> 现象是：本地好好的，push到GitHub后，触发Travis-CI的构建 => 失败

查看下构建记录：[524408121](https://travis-ci.org/guoruibiao/shorturl/builds/524408121)
```
$ go version
go version go1.11.5 linux/amd64
go.env
$ go env
GOARCH="amd64"
GOBIN=""
GOCACHE="/home/travis/.cache/go-build"
GOEXE=""
GOFLAGS=""
GOHOSTARCH="amd64"
GOHOSTOS="linux"
GOOS="linux"
GOPATH="/home/travis/gopath"
GOPROXY=""
GORACE=""
GOROOT="/home/travis/.gimme/versions/go1.11.5.linux.amd64"
GOTMPDIR=""
GOTOOLDIR="/home/travis/.gimme/versions/go1.11.5.linux.amd64/pkg/tool/linux_amd64"
GCCGO="gccgo"
CC="gcc"
CXX="g++"
CGO_ENABLED="1"
GOMOD=""
CGO_CFLAGS="-g -O2"
CGO_CPPFLAGS=""
CGO_CXXFLAGS="-g -O2"
CGO_FFLAGS="-g -O2"
CGO_LDFLAGS="-g -O2"
PKG_CONFIG="pkg-config"
GOGCCFLAGS="-fPIC -m64 -pthread -fmessage-length=0 -fdebug-prefix-map=/tmp/go-build040675356=/tmp/go-build -gno-record-gcc-switches"
install
0.00s$ travis_install_go_dependencies 1.11.5 -v
Makefile detected
0.34s$ go get -v
github.com/guoruibiao/shorturl
The command "go get -v" exited with 0.
0.08s$ go test ./...
dao/shortdao.go:7:2: cannot find package "shorturl/model" in any of:
	/home/travis/.gimme/versions/go1.11.5.linux.amd64/src/shorturl/model (from $GOROOT)
	/home/travis/gopath/src/shorturl/model (from $GOPATH)
The command "go test ./..." exited with 1.
Done. Your build exited with 1.
```

终于发现问题了，是`package`未在`GOPATH`中找到引起的。难怪会构建失败了。

## 6 问题fix实践
既然构建日志都说是`package`没找到了，那么让它找到应该就好了吧。看`GOPATH`的内容，发现`test`集合中的包是从**github.com/xxx/yyy**中去查找的。那这么说，`import`的时候直接写死应该就可以了吧。

代码：[596c9ebd84ffbbd68a8febfd52551d40dc665eab](https://github.com/guoruibiao/shorturl/tree/596c9ebd84ffbbd68a8febfd52551d40dc665eab)
构建结果： [524410672](https://travis-ci.org/guoruibiao/shorturl/builds/524410672)

发现构建成功了，也在一定程度上验证了我们的猜想。不过这种解法总觉得怪怪的，因为本地代码中**github.com/xxx/yyy**中是没有这个`package`的，直接写到代码中，有点不伦不类。

## 7 还得是GOPATH
弄完第6部分已经凌晨一点了，还是头发比较重要，睡觉去。不过脑袋里却依然萦绕着这个问题，想了想，还得是`GOPATH`的锅。在`Travis-ci`的运行容器中，有这么一条内容： 

```
git clone github.com/xxx/yyy xxx/yyy
```

也就是说它最终进到了`yyy`目录，即`$GOPATH/src/github.com/xxx/yyy`和我本地开发环境中 `yyy` 在 `$GOPATH/src/yyy` 的路径是不一致的。这也是为啥一直在提示找不到`package`的问题所在。

那么如何确定就是这个问题呢？还得是从容器本身的运行流程下手了,切入点还是**.travis.yml**
```YAML
# .travis.yml
language:    # 指明使用的语言
  - go 

go:          # 语言版本号
  - "1.11.5"   # 默认使用最新版本,注意，需要 "1.10" 版本的时候必须表示为字符串形式，如果写成 1.10 则会使用 1.1 版本;x表示对应前缀的最新版本


before_script:
  - go get github.com/gin-gonic/gin

script:      # 执行的脚步，但是go默认会执行下面的这些命令的，所以可以不用写
  #- go get -v
  - ls
  - echo $GOPATH
  - echo $PWD
  - cd $GOPATH/src
  - mv $GOPATH/src/github.com/guoruibiao/shorturl $GOPATH/src/
  - echo $PWD
  - cd ./shorturl/
  - go test -v ./...

```
构建号：[524756302](https://travis-ci.org/guoruibiao/shorturl/builds/524756302)

从日志的输出结果看，的确是这个问题。
```
$ echo $GOPATH
/home/travis/gopath
The command "echo $GOPATH" exited with 0.
0.00s$ echo $PWD
/home/travis/gopath/src/github.com/guoruibiao/shorturl
The command "echo $PWD" exited with 0.
0.00s$ cd $GOPATH/src
The command "cd $GOPATH/src" exited with 0.
0.00s$ mv $GOPATH/src/github.com/guoruibiao/shorturl $GOPATH/src/
The command "mv $GOPATH/src/github.com/guoruibiao/shorturl $GOPATH/src/" exited with 0.
0.00s$ echo $PWD
/home/travis/gopath/src
The command "echo $PWD" exited with 0.
0.00s$ cd ./shorturl/
The command "cd ./shorturl/" exited with 0.
7.67s$ go test -v ./...
?   	shorturl	[no test files]
=== RUN   Test_SinaURLShort
--- PASS: Test_SinaURLShort (1.36s)
    shortdao_test.go:12: SinaURLShort passed
PASS
ok  	shorturl/dao	1.369s
?   	shorturl/model	[no test files]
=== RUN   Test_ShortURL
--- FAIL: Test_ShortURL (1.10s)
    shorturl_test.go:18: service shorturl result failed
FAIL
FAIL	shorturl/service	1.108s
?   	shorturl/utils	[no test files]
The command "go test -v ./..." exited with 1.
Done. Your build exited with 1.
```

`finally`，那接下来修改下容器运行就好了。

## 8 模拟重试
有问题不可怕，一点点去试错，去修复就好了。

代码：[4a5fd91316c5428fec5b793e6ef341083aa6743d](https://github.com/guoruibiao/shorturl/tree/4a5fd91316c5428fec5b793e6ef341083aa6743d)

构建号：[524757287](https://travis-ci.org/guoruibiao/shorturl/builds/524757287)

```
7.83s$ go test -v ./...
?   	shorturl	[no test files]
=== RUN   Test_SinaURLShort
--- PASS: Test_SinaURLShort (1.44s)
    shortdao_test.go:12: SinaURLShort passed
PASS
ok  	shorturl/dao	1.448s
?   	shorturl/model	[no test files]
=== RUN   Test_ShortURL
--- PASS: Test_ShortURL (1.10s)
PASS
ok  	shorturl/service	1.108s
?   	shorturl/utils	[no test files]
The command "go test -v ./..." exited with 0.
```
这样就搞定了，本地代码就和Travis-CI容器中的一致了。

## 9 反思
再来瞅瞅这个目录结构。
```
➜  shorturl git:(master) ✗ tree -L 2
.
├── Dockerfile
├── Makefile
├── README.md
├── app.go
├── dao
│   ├── shortdao.go
│   └── shortdao_test.go
├── model
│   └── model.go
├── service
│   ├── shorturl.go
│   └── shorturl_test.go
└── utils
    └── utils.go

```
最后回顾这个问题的时候，我也在想,这中`import`到底合不合适。
```golang
package dao

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"shorturl/model"
	// model "github.com/guoruibiao/shorturl/model"
)
```
或许本地开发，使用没啥问题，但是做出来的项目有时候不仅仅是用一次就完事了的，基于重用的原则，下面的那个引入应该才是最为规范的。然后我去`GitHub`上搜了搜，看看人家的项目中都是怎么`import`的，不出意外，人家都是用的全路径引入的模式。我目前对`dep`还不太了解，或许会有一个更好的解决方案，这里暂且先不写了。

**// TODO Learn go dep ...**

后来前大佬同事[todd](https://github.com/wncbb) 聊这个话题来着，发现是开发的时候选址就错了，一般来讲，规范的开发目录本身就得是：

```
$GOPATH/src/github.com/username/repository-name
```
这样引入package依赖也好，持续构建也好，都挺合适。

还是印证了那句话：
> 一个人可以走得很快，但两个人可以走的很远。

分享、交流，才能碰撞出思维的闪光点。

参考链接：  
- 阮一峰 Travis-CI 指南 http://www.ruanyifeng.com/blog/2017/12/travis_ci_tutorial.html
- travis-ci-yml https://docs.travis-ci.com/user/languages/go/
