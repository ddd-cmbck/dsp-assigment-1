module github.com/ddd-cmbck/dsp-assigment-1/database

go 1.24.0

toolchain go1.24.9

require (
	github.com/ddd-cmbck/dsp-assigment-1/proto v0.1.0
	google.golang.org/grpc v1.76.0
)

require (
	golang.org/x/net v0.46.0 // indirect
	golang.org/x/sys v0.37.0 // indirect
	golang.org/x/text v0.30.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250804133106-a7a43d27e69b // indirect
	google.golang.org/protobuf v1.36.10 // indirect
)

replace github.com/ddd-cmbck/dsp-assigment-1/proto => ../proto
