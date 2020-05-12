# k8s-grpc

基于grpc搭建k8s集群

## 克隆代码

```shell
git clone https://github.com/ganodermaking/k8s-grpc
```

## 编译可执行文件

scratch 是一个特殊的镜像，它是一个虚拟镜像，也就是一个空白镜像；利用Golang的静态化编译无依赖性，可以大幅度减少编译时间和镜像大小。

```shell
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o grpc .
```

> GOOS=linux 是将交叉编译的目标设置为Linux，这样你在Mac或者Win下也不会出现问题。 -installsuffix cgo 是为了在静态编译中导入net

## 打包镜像

```shell
docker build -t scratch/grpc:v1 .
```

## 启动集群

```shell
kubectl create -f deployment.yaml
```

## 测试服务

go test -v -count=1 k8s-grpc/tests -run TestSayHello

> mac docker desktop 2.3.0.2自带kubernetes v1.16.5演示
