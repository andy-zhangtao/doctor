# doctor

Easy login remote machines, like jumpserver. But doctor is lightweight, and simple.

## How to install doctor?

1. Install Postgres.

```
docker run  -d  --name postgres \
				-e POSTGRES_PASSWORD=123456 \
				-p 5432:5432 \
				-v <local dir>:/var/lib/postgresql/data \
				postgres:alpine
```

2. Install doctor

> DOCTOR_GRPC_SERVER Port is 50000

```
docker run  -d  --name doctor\
				-v /root/doctor_data:/data \
				-v /root/doctor_key:/home \
				--net host \
				-e DOCTOR_GRPC_SERVER=<host ip:port> \
				--add-host <postgres ip> \
				vikings/doctor:v0.1 -server
```

## About RemoteNode

1. First Remote Node should allow root login via ssh. (for doctor init)
2. Remote Node should use 22 for ssh port.
