package modules

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/cosmos/cosmos-sdk/client"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	"github.com/stretchr/testify/assert"

	"github.com/kava-labs/kava/app"
)

func TestCosmosMigrate_Gov(t *testing.T) {
	// The gov json contains a gov app state with 3 different proposals.
	// Two of the proposals are from cosmos while the 3rd one is kavadist/CommunityPoolMultiSpendProposal.
	original := GetTestDataJSON("appstate-gov-v15.json")
	expected := GetTestDataJSON("appstate-gov-v16.json")
	actual := MustMigrateAppStateJSON(original)
	assert.JSONEq(t, expected, actual)
}

// MustMigrateAppStateJSON migrate v15 app state json to v16
func MustMigrateAppStateJSON(appStateJson string) string {
	var appState genutiltypes.AppMap
	if err := json.Unmarshal([]byte(appStateJson), &appState); err != nil {
		panic(err)
	}
	newGenState := MigrateCosmosAppState(appState, NewClientContext())
	actual, err := json.Marshal(newGenState)
	if err != nil {
		panic(err)
	}
	return string(actual)
}

func GetTestDataJSON(filename string) string {
	file := filepath.Join("testdata", filename)
	data, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func NewClientContext() client.Context {
	config := app.MakeEncodingConfig()
	return client.Context{}.
		WithCodec(config.Marshaler).
		WithLegacyAmino(config.Amino)
}