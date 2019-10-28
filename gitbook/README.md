## 内容简介

- 知识总结

- 学习笔记

- 知识分享


## 服务启动

```bash

cd /home/; git clone  https://github.com/msupers/gitbook.git

docker pull registry.cn-beijing.aliyuncs.com/meowbite/gitbook:3.2.3

docker run -v /home/gitbook/gitbook:/gitbook \
 -p 8000:4000 registry.cn-beijing.aliyuncs.com/meowbite/gitbook:3.2.3
```

## 服务验证

```bash

http://${Docker主机}:8000

```