package main

import (
	"fmt"
	"testing"

	. "github.com/bjartek/overflow"
	"github.com/zeebo/assert"
)

func TestNFL(t *testing.T) {
	o, err := OverflowTesting()
	assert.NoError(t, err)

	if err != nil {
		fmt.Printf("Error: %s\n", err)
		t.Error(err)
	}

	o.Tx("user/setup_allday_account", WithSignerServiceAccount()).AssertSuccess(t)

	o.Tx("admin/series/create_series", WithSignerServiceAccount(),
		WithArgs("name", "test-series"),
	).AssertSuccess(t)

	series, _ := o.Script("series/read_series_by_name", WithSignerServiceAccount(),
		WithArgs("seriesName", "test-series"),
	).GetAsInterface()

	fmt.Printf("Created Series: %v \n", series)

	o.Tx("admin/sets/create_set", WithSignerServiceAccount(),
		WithArgs("name", "test-set"),
	).AssertSuccess(t)

	metadata := map[string]string{"type": "wow"}

	o.Tx("admin/plays/create_play", WithSignerServiceAccount(),
		WithArgs("name", "test-play"),
		WithArgs("metadata", metadata),
	).AssertSuccess(t)

	o.Tx("admin/editions/create_edition", WithSignerServiceAccount(),
		WithArgs("seriesID", 1),
		WithArgs("setID", 1),
		WithArgs("playID", 1),
		WithArgs("tier", "a-tier"),
		WithArgs("maxMintSize", 1),
	).AssertSuccess(t)

	o.Tx("admin/nfts/mint_moment_nft", WithSignerServiceAccount(),
		WithArgs("recipientAddress", "account"),
		WithArgs("editionID", 1),
	).AssertSuccess(t)

	// 	metaData := map[string]string{
	// 		"type": "SlamDunk",
	// 	}

	// 	o.Tx("admin/create_play", WithSignerServiceAccount(),
	// 		WithArgs("metadata", metaData),
	// 	).AssertSuccess(t)

	// 	o.Tx("admin/create_set", WithSignerServiceAccount(),
	// 		WithArgs("setName", "Overflow"),
	// 	).AssertSuccess(t)

	// 	playId, _ := o.Script("plays/get_nextPlayId").GetAsInterface()
	// 	assert.NotNil(t, playId)

	// 	setId, _ := o.Script("sets/get_nextSetId").GetAsInterface()
	// 	assert.NotNil(t, setId)

	// 	playIdUint := playId.(uint32)
	// 	playIdUint -= 1

	// 	setIdUint := setId.(uint32)
	// 	setIdUint -= 1

	// 	fmt.Printf("PlayId: %v, SetId: %v \n", playIdUint, setIdUint)

	// 	o.Tx("admin/add_play_to_set", WithSignerServiceAccount(),
	// 		WithArgs("setID", setIdUint),
	// 		WithArgs("playID", playIdUint),
	// 	).AssertSuccess(t).Print()

	// 	o.Tx("admin/mint_moment", WithSignerServiceAccount(),
	// 		WithArgs("setID", setIdUint),
	// 		WithArgs("playID", playIdUint),
	// 		WithArgs("recipientAddr", "account"),
	// 	).AssertSuccess(t).Print()

	// 	nftMetaData, err := o.Script("get_nft_metadata", WithSignerServiceAccount(),
	// 		WithArgs("address", "account"),
	// 		WithArgs("id", "1"),
	// 	).GetAsInterface()
	// 	assert.NoError(t, err)

	// 	fmt.Printf("NFT Metadata: %+v \n", nftMetaData)

	// 	setData, err := o.Script("sets/get_set_data", WithSignerServiceAccount(),
	// 		WithArgs("setID", setIdUint),
	// 	).GetAsInterface()
	// 	assert.NoError(t, err)

	// 	topshotData, err := o.Script("get_topshot_metadata", WithSignerServiceAccount(),
	// 		WithArgs("address", "account"),
	// 		WithArgs("id", "1"),
	// 	).GetAsInterface()
	// 	assert.NoError(t, err)

	// 	fmt.Printf("Topshot Metadata: %+v \n", topshotData)
	// 	fmt.Printf("Set Data: %+v \n", setData)
}
