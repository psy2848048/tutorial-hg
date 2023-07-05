package util

import (
	"fmt"
	"testing"

	hedera "github.com/hashgraph/hedera-sdk-go/v2"
	"github.com/stretchr/testify/assert"
)

const (
	testAccountID      = "0.0.2"
	testAccountPrivKey = "302e020100300506032b65700422042091132178e72057a1d7528025956fe39b0b847f200ab59b2fdd367017f3087137"
)

var client *Client

func TestCreateAccount(t *testing.T) {
	newAccountInfo, err := client.CreateNewAccount(10000000)
	assert.NoError(t, err)

	fmt.Println("New account ID:", newAccountInfo.AccountID.String())
}

func setUp() {
	node := make(map[string]hedera.AccountID, 1)
	node["127.0.0.1:50211"] = hedera.AccountID{Account: 3}

	mirrorNode := []string{"127.0.0.1:5600"}

	var err error
	client, err = NewClient(node, mirrorNode, testAccountID, testAccountPrivKey)
	if err != nil {
		panic(err)
	}
}

func tearDown() {}

func TestMain(m *testing.M) {
	setUp()
	m.Run()
	tearDown()
}
