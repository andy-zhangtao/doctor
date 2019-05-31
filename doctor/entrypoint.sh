#!/bin/sh
[ -f "/data/passwd" ] && cp /data/passwd /etc/passwd
/usr/sbin/sshd && /doctor "$@"