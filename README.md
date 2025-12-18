# prometheus lnd exporter

Prometheus exporter for lnd (https://github.com/lightningnetwork/lnd)



### run it

build
```
go build .
```


run
```
./prometheus-lnd-exporter -rpc.addr=lnd-host:20009 -lnd.macaroon-path=/path/to/admin.macaroon -lnd.tls-cert-path=path/to/tls.cert
```


