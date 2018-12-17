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
13. [结构体及其方法的使用法门](13-struct.md)
14. [接口类型的合理运用](14-interface.md)
15. [关于指针的有限操作](15-point.md)
16. [go语句及其执行规则（上）](16-goroutine-1.md)
17. [go语句及其执行规则（下）](17-goroutine-2.md)
18. [if语句、for语句和switch语句](18-if_for_switch.md)
19. [错误处理（上）](19-error-1.md)
20. [错误处理（下）](20-error-2.md)
21. [panic函数、recover函数以及defer语句 （上）](21-panic_recover_defer.md)
22. [panic函数、recover函数以及defer语句 （下）](22-panic_recover_defer-2.md)
23. [测试的基本规则和流程 （上）](23-test-1.md)
24. [测试的基本规则和流程 （下）](24-test-2.md)
25. [更多的测试手法](25-more-test.md)
26. [sync.Mutex与sync.RWMutex](26-sync.Mutex.md)
27. [条件变量sync.Cond](27-sync.Cond.md)
28. [原子操作](28-atomic.md)
29. [sync.WaitGroup和sync.Once](29-sync.WaitGroup-Once.md)
30. [context.Context类型](30-context.md)
31. [临时对象池sync.Pool](31-sync.Pool.md)
32. [并发安全字典sync.Map](32-sync.Map.md)