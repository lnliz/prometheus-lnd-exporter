# LND RPC stubs

This package contains generated LND Lightning gRPC client/message stubs copied
from `github.com/lightningnetwork/lnd/lnrpc` at tag `v0.20.1-beta.rc1`.

Only the generated `lightning.pb.go` and `lightning_grpc.pb.go` files are
vendored here. The broader upstream `lnrpc` package also includes helper files
that pull in LND database packages not needed by this exporter.
