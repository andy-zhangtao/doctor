FROM vikings/alpine
RUN  apk update && \
    apk add openssh && \
    ssh-keygen -f /etc/ssh/ssh_host_rsa_key -N '' -t rsa && \
    rm /var/cache/apk/*
COPY motd /etc/motd
COPY bwidow_pq.toml /bwidow_pq.toml
COPY entrypoint.sh /entrypoint.sh
COPY cli.sh /cli.sh
COPY doctor /doctor
COPY nurse /tmp/nurse
ENTRYPOINT [ "/entrypoint.sh" ]