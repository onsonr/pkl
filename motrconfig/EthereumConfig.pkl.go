// Code generated from Pkl module `MotrConfig`. DO NOT EDIT.
package motrconfig

type EthereumConfig struct {
	ChainId string `pkl:"chainId"`

	ValApiUri string `pkl:"valApiUri"`

	ValRpcUri *string `pkl:"valRpcUri"`

	ValGrpcUri *string `pkl:"valGrpcUri"`
}
