package main

import (
	"os"

	"github.com/cosmos/cosmos-sdk/server"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/simapp/simd/cmd"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func main() {

	// setBech32Prefixes()

	rootCmd, _ := cmd.NewRootCmd()

	if err := svrcmd.Execute(rootCmd, simapp.DefaultNodeHome); err != nil {
		switch e := err.(type) {
		case server.ErrorCode:
			os.Exit(e.Code)

		default:
			os.Exit(1)
		}
	}
}

func setBech32Prefixes() {
	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount("cosmic", "cosmicpub")
	config.SetBech32PrefixForValidator("cosmicvaloper", "cosmicvaloperpub")
	config.SetBech32PrefixForConsensusNode("cosmicvalcons", "cosmicvalconspub")
	config.Seal()
}
