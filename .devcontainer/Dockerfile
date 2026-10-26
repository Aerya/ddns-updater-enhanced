FROM ghcr.io/qdm12/godevcontainer:v0.23.0-alpine

RUN addgroup -g 1000 developer && adduser -D -u 1000 -G developer -s /bin/zsh developer \
    && cp -a /root/. /home/developer/ \
    && ln -snf /home/developer/.ssh.sh /home/developer/.windows.sh \
    && chown -R developer:developer /home/developer /go
ENV HOME=/home/developer
USER developer
