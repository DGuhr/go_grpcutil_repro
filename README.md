# Reproducer
To reproduce the issue, check out this repository and run `go mod tidy`

It will fail with the following message:

```
go: finding module for package google.golang.org/grpc/test/grpc_testing
go_grpcutil_repro imports
        github.com/authzed/grpcutil tested by
        github.com/authzed/grpcutil.test imports
        google.golang.org/grpc/test/grpc_testing: module google.golang.org/grpc@latest found (v1.55.0), but does not contain package google.golang.org/grpc/test/grpc_testing
```

... because the `grpc_testing` package was moved in v1.55.0 to interop/grpc_testing.



## possible temporary solution to unblock services using the grpcutil
A PR using the changed package was opened here: https://github.com/authzed/grpcutil/pull/47/files - it removes one case from the utils tests though, as the compressiontype is not available anymore in interop/grpc_testing. May be closed in favour of the proper solution (see below)

## proper (maybe mid-term) solution:
the grpc team suggests instead to not rely on the package at all, see discussion here: https://github.com/grpc/grpc-go/issues/6264#issuecomment-1540638126

## Workaround:
switch the `	google.golang.org/grpc` to v1.54.0 instead of v1.55.0 or import / c&p the actual code of the grpcutil inside your service.