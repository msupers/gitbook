FROM registry.cn-zhangjiakou.aliyuncs.com/meow/gitbook:v1  
# Set npm registry to China Taobao and install Gitbook
#RUN npm config set registry https://registry.npm.taobao.org && \
#    npm install gitbook-cli -g && \
#    gitbook -V && \
#    mkdir /gitbook
#
# Make port 4000 available to the world outside this container
COPY . gitbook
EXPOSE 4000

CMD ["sh", "-c", "gitbook install /gitbook; gitbook serve /gitbook"]
