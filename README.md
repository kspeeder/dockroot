# dockroot

https://www.koolcenter.com/t/topic/7778


## DockRoot: 一个轻量版Docker

* 用法跟 docker 相似
* 没原生 Docker 强大
* 但是没任何依赖，是 Linux 就能运行
* 没后台服务进程，更省内存
* 核心代码来源于 skopoe
* 辅助工具：ruri

## 缺点
* 必须以 root 权限运行（有办法不用 root，但是麻烦懒得弄）
* 没任何网络隔离，相当于 Docker 的 host 模式

## 原理
* 下载 Docker 镜像
* 解压为 rootfs
* 生成 ruri 需要运行的配置文件
* ruri -c xxx/ruri.conf 运行镜像
