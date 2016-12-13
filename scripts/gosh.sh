#!/usr/bin/env bash
if [ ! -f install ]; then
echo 'install must be run within its container folder' 1>&2
exit 1
fi
CURDIR=`pwd`
OLDGOPATH="$GOPATH"
export GOPATH="$CURDIR"
gofmt -w src
go install test
export GOPATH="$OLDGOPATH"
echo 'finished'

package main

import (
	"fmt"
	"os/exec"
)

func main() {
	#//执行【ls /】并输出返回文本
	f, err := exec.Command("ls", "/").Output()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(f))
}


#ubuntu :
sudo apt-get install golang
#redhat :
sudo yum install golang

#源码安装
#安装依赖软件包 Mercurial，同样使用网络安装即可，例如
#$ sudo yum install mercurial
#解压缩下载的go源码包
tar xvf go.xxx.tar.gz
#编译
cd /PATH_TO_GO_DIR/src
./all.bash
#出现"ALL TESTS PASSED"，代表编译成功
#配置环境变量，修改profile或.bashrc文件，加入如下代码
export GOROOT=/PATH_TO_GO/go
export GOBIN=$GOROOT/bin
export PATH=$PATH:$GOBIN
