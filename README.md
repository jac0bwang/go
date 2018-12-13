# Go 笔记

一下是用到的环境

```Bash
$ export "GOPATH=$PWD"
$ cat /etc/lsb-release
DISTRIB_ID=Ubuntu
DISTRIB_RELEASE=18.04
DISTRIB_CODENAME=bionic
DISTRIB_DESCRIPTION="Ubuntu 18.04.1 LTS"
$ go version
go version go1.10.4 linux/amd64
```

1. [工作区和GOPATH](01-工作区和GOPATH.md)
2. [命令源码文件](02-commandsrc.md)
3. [库源码文件](03-libsrc.md)
4. [程序实体的那些事儿（上）](04-程序实体的那些事儿（上）.md)
5. [程序实体的那些事儿（中）](05-Program_entity-2.md)
6. [程序实体的那些事儿（下）](06-Program_entity-3.md)
7. [数组和切片](07-array_slice.md)
8. [container包中的那些容器](08-container.md)
9. [字典的操作和约束](09-map.md)
10. [通道的基本操作](10-channel-1.md)
11. [通道的高级玩法](11-channel-2.md)
12. [使用函数的正确姿势](12-func.md)