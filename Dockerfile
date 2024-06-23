FROM alpine:3.18.4

RUN mkdir /go && cd /go \
    && wget --no-check-certificate https://golang.google.cn/dl/go1.21.0.linux-amd64.tar.gz \
    && tar -C /usr/local -zxf go1.21.0.linux-amd64.tar.gz \
    && rm -rf /go/go1.21.0.linux-amd64.tar.gz \
    # 注意！这一步是因为alpine系统对go程序运行的时候默认查找的类库与提供的类库不一致
    && mkdir /lib64 \
    && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2 
   
# 配置系统环境变量
ENV GOPATH /go
ENV PATH /usr/local/go/bin:$GOPATH/bin:$PATH

CMD ["ping", "www.baidu.com"]
