module cosmossdk.io/server/v2/appmanager

go 1.21

replace (
	cosmossdk.io/core => ../../../core
	cosmossdk.io/server/v2/stf => ../stf
)

require (
	cosmossdk.io/core v0.12.0
	cosmossdk.io/server/v2/stf v0.0.0-00010101000000-000000000000
)

require (
	github.com/cosmos/gogoproto v1.5.0 // indirect
	github.com/google/go-cmp v0.6.0 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	golang.org/x/exp v0.0.0-20240613232115-7f521ea00fb8 // indirect
	google.golang.org/protobuf v1.34.2 // indirect
)
