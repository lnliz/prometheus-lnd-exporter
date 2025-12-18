# prometheus lnd exporter

Prometheus exporter for lnd (https://github.com/lightningnetwork/lnd)

[Docker Hub](https://hub.docker.com/r/lnliz/prometheus-lnd-exporter/tags)

### run it

build
```
go build .
```

run
```
./prometheus-lnd-exporter -rpc.addr=lnd-host:20009 -lnd.macaroon-path=/path/to/admin.macaroon -lnd.tls-cert-path=path/to/tls.cert
```



docker 
```
docker pull lnliz/prometheus-lnd-exporter:latest

docker run ...
```
