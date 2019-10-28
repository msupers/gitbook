## quick start

使用docker快速启动一个Prometheus

***prom/proethues是官方镜像，建议使用的时候指定版本名称***

```bash
docker run -p 9090:9090 -v /prometheus-data \
       prom/prometheus:v2.13.1 --config.file=/prometheus-data/prometheus.yml
```

## 访问地址

```bash
http://localhost:9090
```

##参考信息

-官方镜像地址

```bash
https://hub.docker.com/r/prom/prometheus/tags

```

-自定义镜像地址

```bash
registry.cn-beijing.aliyuncs.com/meowbite/prometheus:v2.3.1
```