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

	o.Script("series/read_series_by_name", WithSignerServiceAccount(),
		WithArgs("seriesName", "test-series"),
	).GetAsInterface()

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
		WithArgs("tier", "Legendary"),
		WithArgs("maxMintSize", 1),
	).AssertSuccess(t)

	o.Tx("admin/nfts/mint_moment_nft", WithSignerServiceAccount(),
		WithArgs("recipientAddress", "account"),
		WithArgs("editionID", 1),
	).AssertSuccess(t)

	editions, err := o.Script(
		"editions/read_all_editions", WithSignerServiceAccount()).GetAsInterface()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		t.Error(err)
	}

	fmt.Printf("Editions: %v\n", editions)

	matchingEditions, err := o.Script("editions/get_editions_by_field", WithSignerServiceAccount(),
		WithArgs("query", "Legendary"),
	).GetAsInterface()

	fmt.Printf("Matching Editions: %v\n", matchingEditions)
}
