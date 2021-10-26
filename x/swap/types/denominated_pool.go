package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// DenominatedPool implements a denominated constant-product liquidity pool
type DenominatedPool struct {
	// all pool operations are implemented in a unitless base pool
	*BasePool
	// track units of the reserveA and reserveB in base pool
	DenomA string
	DenomB string // TODO fields do not need to be exported
}

// NewDenominatedPool creates a new denominated pool from reserve coins
func NewDenominatedPool(reserves sdk.Coins) (*DenominatedPool, error) {
	if len(reserves) != 2 {
		return nil, sdkerrors.Wrap(ErrInvalidPool, "reserves must have two denominations")
	}

	// Coins should always sorted, so this is deterministic, though it does not need to be.
	// The base pool calculation results do not depend on reserve order.
	reservesA := reserves[0]
	reservesB := reserves[1]

	pool, err := NewBasePool(reservesA.Amount, reservesB.Amount)
	if err != nil {
		return nil, err
	}

	return &DenominatedPool{
		BasePool: pool,
		DenomA:   reservesA.Denom,
		DenomB:   reservesB.Denom,
	}, nil
}

// NewDenominatedPoolWithExistingShares creates a new denominated pool from reserve coins
func NewDenominatedPoolWithExistingShares(reserves sdk.Coins, totalShares sdk.Int) (*DenominatedPool, error) {
	if len(reserves) != 2 {
		return nil, sdkerrors.Wrap(ErrInvalidPool, "reserves must have two denominations")
	}

	// Coins should always sorted, so this is deterministic, though it does not need to be.
	// The base pool calculation results do not depend on reserve order.
	reservesA := reserves[0]
	reservesB := reserves[1]

	pool, err := NewBasePoolWithExistingShares(reservesA.Amount, reservesB.Amount, totalShares)
	if err != nil {
		return nil, err
	}

	return &DenominatedPool{
		BasePool: pool,
		DenomA:   reservesA.Denom,
		DenomB:   reservesB.Denom,
	}, nil
}

// Reserves returns the reserves held in the pool
func (p *DenominatedPool) Reserves() sdk.Coins {
	return p.coins(p.BasePool.ReservesA, p.BasePool.ReservesB)
}

// TotalShares returns the total shares for the pool
func (p *DenominatedPool) TotalShares() sdk.Int {
	return p.BasePool.TotalShares
}

// IsEmpty returns true if the pool is empty
func (p *DenominatedPool) IsEmpty() bool {
	return p.BasePool.IsEmpty()
}

// AddLiquidity adds liquidity to the reserves and returns the added amount and shares created
func (p *DenominatedPool) AddLiquidity(deposit sdk.Coins) (sdk.Coins, sdk.Int) {
	desiredA := deposit.AmountOf(p.DenomA)
	desiredB := deposit.AmountOf(p.DenomB)

	actualA, actualB, shares := p.BasePool.AddLiquidity(desiredA, desiredB)

	return p.coins(actualA, actualB), shares
}

// RemoveLiquidity removes liquidity from the pool
func (p *DenominatedPool) RemoveLiquidity(shares sdk.Int) sdk.Coins {
	withdrawnA, withdrawnB := p.BasePool.RemoveLiquidity(shares)

	return p.coins(withdrawnA, withdrawnB)
}

// ShareValue returns the value of the provided shares
func (p *DenominatedPool) ShareValue(shares sdk.Int) sdk.Coins {
	valueA, valueB := p.BasePool.ShareValue(shares)

	return p.coins(valueA, valueB)
}

// SwapWithExactInput trades an exact input coin for the other.  Returns the positive other coin amount
// that is removed from the pool and the portion of the input coin that is used for the fee.
// It panics if the input denom does not match the pool reserves.
func (p *DenominatedPool) SwapWithExactInput(swapInput sdk.Coin, fee sdk.Dec) (sdk.Coin, sdk.Coin) {
	var (
		swapOutput sdk.Int
		feePaid    sdk.Int
	)

	switch swapInput.Denom {
	case p.DenomA:
		swapOutput, feePaid = p.BasePool.SwapExactAForB(swapInput.Amount, fee)
		return p.coinB(swapOutput), p.coinA(feePaid)
	case p.DenomB:
		swapOutput, feePaid = p.BasePool.SwapExactBForA(swapInput.Amount, fee)
		return p.coinA(swapOutput), p.coinB(feePaid)
	default:
		panic(fmt.Sprintf("invalid denomination: denom '%s' does not match pool reserves", swapInput.Denom))
	}
}

// SwapWithExactOutput trades a coin for an exact output coin b.  Returns the positive input coin
// that is added to the pool, and the portion of that input that is used to pay the fee.
// Panics if the output denom does not match the pool reserves.
func (p *DenominatedPool) SwapWithExactOutput(swapOutput sdk.Coin, fee sdk.Dec) (sdk.Coin, sdk.Coin) {
	var (
		swapInput sdk.Int
		feePaid   sdk.Int
	)

	switch swapOutput.Denom {
	case p.DenomA:
		swapInput, feePaid = p.BasePool.SwapBForExactA(swapOutput.Amount, fee)
		return p.coinB(swapInput), p.coinB(feePaid)
	case p.DenomB:
		swapInput, feePaid = p.BasePool.SwapAForExactB(swapOutput.Amount, fee)
		return p.coinA(swapInput), p.coinA(feePaid)
	default:
		panic(fmt.Sprintf("invalid denomination: denom '%s' does not match pool reserves", swapOutput.Denom))
	}
}

// coins returns a new coins slice with correct reserve denoms from ordered sdk.Ints
func (p *DenominatedPool) coins(amountA, amountB sdk.Int) sdk.Coins {
	return sdk.NewCoins(p.coinA(amountA), p.coinB(amountB))
}

// coinA returns a new coin denominated in denomA
func (p *DenominatedPool) coinA(amount sdk.Int) sdk.Coin {
	return sdk.NewCoin(p.DenomA, amount)
}

// coinA returns a new coin denominated in denomB
func (p *DenominatedPool) coinB(amount sdk.Int) sdk.Coin {
	return sdk.NewCoin(p.DenomB, amount)
}
