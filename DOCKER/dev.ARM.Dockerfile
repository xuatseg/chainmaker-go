# the second stage
FROM ubuntu:20.04
RUN rm /bin/sh && ln -s /bin/bash /bin/sh
RUN sed -i "s@http://.*archive.ubuntu.com@http://mirrors.tuna.tsinghua.edu.cn@g" /etc/apt/sources.list && \
    sed -i "s@http://.*security.ubuntu.com@http://mirrors.tuna.tsinghua.edu.cn@g" /etc/apt/sources.list && \
    apt-get update && \
    apt-get install -y vim net-tools tree gcc g++ p7zip-full
ENV TZ "Asia/Shanghai"
RUN DEBIAN_FRONTEND=noninteractive apt-get install -y tzdata && \
    echo $TZ > /etc/timezone && \
    ln -fs /usr/share/zoneinfo/$TZ /etc/localtime && \
    dpkg-reconfigure tzdata -f noninteractive

COPY ./main/arm/libwasmer_runtime_c_api.so /usr/lib/libwasmer.so
COPY ./main/prebuilt/linux/wxdec /usr/bin
COPY ./bin/chainmaker /chainmaker-go/bin/
COPY ./bin/cmc /usr/bin/
COPY ./config /chainmaker-go/config/
RUN mkdir -p /chainmaker-go/log/
RUN chmod 755 /usr/bin/wxdec

WORKDIR /chainmaker-go/bin