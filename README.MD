# GoLang语言学习
### imooc文件夹主要学习目录

### 一个Go工程中主要包含以下三个目录：
```
src：源代码文件
pkg：包文件
bin：相关bin文件
```
### 测试文件
```
cd /src/example/
go test  
go test -run=TestAdd
```
### 

### 运行单个go文件
```
go run 文件名.go 

```
### 

### git 
```
 git status
 git add .
 git commit -m "更新文件"
 git status
 git push -u origin master

 拉
  cd src 
  git init
  git remote add origin https://github.com/chaodada/Go.git
  git pull origin master
```

### Go中打印一个变量的数据类型
```
fmt.Println(reflect.TypeOf(a))
```