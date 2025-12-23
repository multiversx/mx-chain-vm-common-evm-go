module github.com/multiversx/mx-chain-vm-common-evm-go

go 1.23.0

replace (
	github.com/multiversx/mx-chain-core-go => github.com/multiversx/mx-chain-core-sovereign-go v1.2.25-0.20251223075042-2f2f8092dc66
	github.com/multiversx/mx-chain-vm-common-go => github.com/multiversx/mx-chain-vm-common-sovereign-go v1.5.17-0.20251223080659-b248455c9554
)

require (
	github.com/ethereum/go-ethereum v1.13.15
	github.com/multiversx/mx-chain-core-evm-go v0.0.0-20251223082923-3e86e35ee16d
	github.com/multiversx/mx-chain-core-go v1.4.0
	github.com/multiversx/mx-chain-vm-common-go v0.0.0-00010101000000-000000000000
	github.com/multiversx/mx-sdk-abi-go v0.3.1-0.20250423092559-f01fdd10b35d
	github.com/stretchr/testify v1.8.4
)

require (
	github.com/btcsuite/btcd/btcec/v2 v2.2.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.0.1 // indirect
	github.com/denisbrodbeck/machineid v1.0.1 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/holiman/uint256 v1.2.4 // indirect
	github.com/mitchellh/mapstructure v1.4.1 // indirect
	github.com/mr-tron/base58 v1.2.0 // indirect
	github.com/multiversx/mx-chain-logger-go v1.1.0 // indirect
	github.com/multiversx/mx-components-big-int v1.0.0 // indirect
	github.com/pelletier/go-toml v1.9.3 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/crypto v0.33.0 // indirect
	golang.org/x/sys v0.30.0 // indirect
	google.golang.org/protobuf v1.36.5 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
