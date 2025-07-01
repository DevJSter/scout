package config

func PrepareDefaultAnvilConfig() *NetworkConfig {
	return &NetworkConfig{
		Chains: []*ChainConfig{
			{
				Name:       "Thane Testnet Chain",
				RPCUrl:     "http://13.233.251.224:9650/ext/bc/22RWkERgqVKS42gJMR7MUAEqi1vQkaYC2yiqkeggRMaeuju7gN/rpc",
				FirstBlock: 0,
				ChainID:    1143689,
			},
		},
	}
}
