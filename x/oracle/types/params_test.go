package types

import (
	"testing"

	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestParamsEqual(t *testing.T) {
	p1 := DefaultParams()
	err := p1.Validate()
	require.NoError(t, err)

	// minus vote period
	p1.VotePeriod = 0
	err = p1.Validate()
	require.Error(t, err)

	// small vote threshold
	p2 := DefaultParams()
	p2.VoteThreshold = sdk.ZeroDec()
	err = p2.Validate()
	require.Error(t, err)

	// negative reward band
	p3 := DefaultParams()
	p3.RewardBand = sdk.NewDecWithPrec(-1, 2)
	err = p3.Validate()
	require.Error(t, err)

	// negative slash fraction
	p4 := DefaultParams()
	p4.SlashFraction = sdk.NewDec(-1)
	err = p4.Validate()
	require.Error(t, err)

	// negative min valid per window
	p5 := DefaultParams()
	p5.MinValidPerWindow = sdk.NewDec(-1)
	err = p5.Validate()
	require.Error(t, err)

	// small slash window
	p6 := DefaultParams()
	p6.SlashWindow = 0
	err = p6.Validate()
	require.Error(t, err)

	// small distribution window
	p7 := DefaultParams()
	p7.RewardDistributionWindow = 0
	err = p7.Validate()
	require.Error(t, err)

	// empty name
	p9 := DefaultParams()
	p9.AcceptList[0].BaseDenom = ""
	p9.AcceptList[0].SymbolDenom = "ATOM"
	err = p9.Validate()
	require.Error(t, err)

	// empty
	p10 := DefaultParams()
	p10.AcceptList[0].BaseDenom = "uatom"
	p10.AcceptList[0].SymbolDenom = ""
	err = p10.Validate()
	require.Error(t, err)

	p11 := DefaultParams()
	require.NotNil(t, p11.ParamSetPairs())
	require.NotNil(t, p11.String())
}
