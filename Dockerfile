FROM ubuntu:20.04

RUN DEBIAN_FRONTEND=noninteractive apt-get update && \
       apt-get -y install -y \
       ca-certificates libssl1.1 vim htop iotop sysstat \
       dstat strace lsof curl jq tzdata && \
       rm -rf /var/cache/apt /var/lib/apt/lists/*

RUN rm /etc/localtime && ln -snf /usr/share/zoneinfo/America/Montreal /etc/localtime && dpkg-reconfigure -f noninteractive tzdata

RUN mkdir /tmp/wasmer-install && cd /tmp/wasmer-install && \
       curl -L https://github.com/wasmerio/wasmer/releases/download/2.2.1/wasmer-linux-amd64.tar.gz | tar xzf - && \
       mv lib/libwasmer.a lib/libwasmer.so /usr/lib/ && cd / && rm -rf /tmp/wasmer-install

ADD /sfeth /app/sfeth

COPY tools/sfeth/motd_generic /etc/motd
COPY tools/sfeth/99-sfeth.sh /etc/profile.d/

# On SSH connection, /root/.bashrc is invoked which invokes '/root/.bash_aliases' if existing,
# so we hijack the file to "execute" our specialized bash script
RUN echo ". /etc/profile.d/99-sfeth.sh" > /root/.bash_aliases

ENV PATH "$PATH:/app"

ENTRYPOINT ["/app/sfeth"]
