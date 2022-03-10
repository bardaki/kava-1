package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/kava-labs/kava/x/evmutil/types"
)

// RegisterInvariants registers the swap module invariants
func RegisterInvariants(ir sdk.InvariantRegistry, bankK types.BankKeeper, k Keeper) {
	ir.RegisterRoute(types.ModuleName, "fully-backed", FullyBackedInvariant(bankK, k))
	ir.RegisterRoute(types.ModuleName, "small-balances", SmallBalancesInvariant(bankK, k))
}

// AllInvariants runs all invariants of the swap module
func AllInvariants(bankK types.BankKeeper, k Keeper) sdk.Invariant { // TODO why?
	return func(ctx sdk.Context) (string, bool) {
		if res, stop := FullyBackedInvariant(bankK, k)(ctx); stop {
			return res, stop
		}
		res, stop := SmallBalancesInvariant(bankK, k)(ctx)
		return res, stop
	}
}

// FullyBackedInvariant ensures all minor balances are backed exactly by the coins in the module account.
func FullyBackedInvariant(bankK types.BankKeeper, k Keeper) sdk.Invariant {
	broken := false
	message := sdk.FormatInvariant(types.ModuleName, "fully backed broken", "minor balances do not match module account")

	return func(ctx sdk.Context) (string, bool) {

		totalMinorBalances := sdk.ZeroInt()
		k.IterateAllAccounts(ctx, func(acc types.Account) bool {
			totalMinorBalances = totalMinorBalances.Add(acc.Balance)
			return false
		})

		bankAddr := authtypes.NewModuleAddress(types.ModuleName)
		bankBalance := bankK.GetBalance(ctx, bankAddr, CosmosDenom).Amount.Mul(ConversionMultiplier)

		broken = !totalMinorBalances.Equal(bankBalance)

		return message, broken
	}
}

func SmallBalancesInvariant(_ types.BankKeeper, k Keeper) sdk.Invariant {
	broken := false
	message := sdk.FormatInvariant(types.ModuleName, "small balances broken", "minor balances not all less than overflow")

	return func(ctx sdk.Context) (string, bool) {

		k.IterateAllAccounts(ctx, func(account types.Account) bool {
			if account.Balance.GTE(ConversionMultiplier) {
				broken = true
				return true
			}
			return false
		})
		return message, broken
	}
}
