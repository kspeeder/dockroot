# dockroot

https://www.koolcenter.com/t/topic/7778


## DockRoot: 一个轻量版Docker

用法跟 docker 相似，但是减少了功能以及简化了用法
不需要内核支持，支持更多 Linux 系统
没后台服务进程，更省内存
核心代码来源于 skopoe
借助工具：ruri 并能支持 Android

## 缺点
必须以超级权限运行，相当于 Docker 的 privilege
没任何网络隔离，相当于 Docker 的 host 模式

## 原理
下载 Docker 镜像
解压为 rootfs
生成 ruri 需要运行的配置文件
ruri -c xxx/ruri.conf 运行镜像
