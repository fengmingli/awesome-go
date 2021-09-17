# awesome-go
GO入门到实践

## 初始化Go的开发环境
```shell
> brew update && brew upgrade && brew install go
> mkdir -p awesome-go && cd awesome-go
> go mod init
```
到这里GO语言开发环境就已经完成了!!!

## go的包管理器 module
使用go mod时，下载的依赖文件在 $GOPATH/pkg/mod/ 目录下：
```shell
cd $GOPATH/pkg/mod/
```
### go mod 语法：
> 大多数命令式软件，都存在 xxx help （例如：docker help，kubectl help，git help , go help ,go mod help And so on）
> 对于初学者，--help 刚开始使用可能不是很友好。其实 Google 也是不错的选择
```shell
go mod help

go mod init   //初始化一个新项目
go mod tidy   //自动下载依赖
go mod vendor //把mod中的依赖copy到vendor文件下

```


