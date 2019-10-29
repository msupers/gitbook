# 文档说明

总结一些shell简单且常用的文本处理命令。

## find

- 查找.groovy文件

```bash
~# find . -name *.groovy -print
./ep/jenkins-slaves/dind/nodejs/master.groovy
./zhcspi/leshan-smart-park-fe/master/master.groovy
./zhcspi/smart-island-fe/master/master.groovy
```

- 反向查找 非*.groovy的文件

```bash
~# find . -name *.groovy -print
./zhcsdsj/frontend/master
./zhcsdsj/frontend/master/Dockerfile
./zhcsdsj/frontend/master/ansible-tag
./zhcsdsj/frontend/master/ansible-tag/hosts
...
```



