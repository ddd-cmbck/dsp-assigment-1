module github.com/ddd-cmbck/dsp-assigment-1/core

go 1.24.0

toolchain go1.24.9

require (
	github.com/ddd-cmbck/dsp-assigment-1/proto v0.1.0
	google.golang.org/grpc v1.76.0 // direct
)

require (
	golang.org/x/net v0.46.0 // indirect
	golang.org/x/sys v0.37.0 // indirect
	golang.org/x/text v0.30.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20251022142026-3a174f9686a8 // indirect
	google.golang.org/protobuf v1.36.10 // indirect
)

replace github.com/ddd-cmbck/dsp-assigment-1/proto => ../proto
