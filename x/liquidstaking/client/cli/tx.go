package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/version"

	"github.com/kava-labs/kava/x/liquidstaking/types"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	liquidstakingTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "liquidstaking transactions subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmds := []*cobra.Command{
		getCmdMintDerivative(),
		getCmdBurnDerivative(),
	}

	for _, cmd := range cmds {
		flags.AddTxFlagsToCmd(cmd)
	}

	liquidstakingTxCmd.AddCommand(cmds...)

	return liquidstakingTxCmd
}

func getCmdMintDerivative() *cobra.Command {
	return &cobra.Command{
		Use:   "mint [validator-addr] [amount]",
		Short: "mint stKava derivative from a delegation",
		Example: fmt.Sprintf(
			`%s tx %s mint kavavaloper16lnfpgn6llvn4fstg5nfrljj6aaxyee9z59jqd 10000000ukava --from <key>`, version.AppName, types.ModuleName,
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			valAddr, err := sdk.ValAddressFromBech32(args[0])
			if err != nil {
				return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, err.Error())
			}

			amount, err := sdk.ParseCoinNormalized(args[1])
			if err != nil {
				return err
			}

			msg := types.NewMsgMintDerivative(clientCtx.GetFromAddress(), valAddr, amount)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}
}

func getCmdBurnDerivative() *cobra.Command {
	return &cobra.Command{
		Use:   "burn [validator-addr] [amount]",
		Short: "burn stKava derivative to redeem a delegation",
		Example: fmt.Sprintf(
			`%s tx %s burn kavavaloper16lnfpgn6llvn4fstg5nfrljj6aaxyee9z59jqd 10000000ukava --from <key>`, version.AppName, types.ModuleName,
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			valAddr, err := sdk.ValAddressFromBech32(args[0])
			if err != nil {
				return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, err.Error())
			}

			amount, err := sdk.ParseCoinNormalized(args[1])
			if err != nil {
				return err
			}

			msg := types.NewMsgBurnDerivative(clientCtx.GetFromAddress(), valAddr, amount)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}
}