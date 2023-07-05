package util

import (
	"fmt"

	hedera "github.com/hashgraph/hedera-sdk-go/v2"
)

type Client struct {
	hederaClient    *hedera.Client
	nodeAccMap      map[string]hedera.AccountID
	mirrorNodesList []string
	accountId       hedera.AccountID
	privateKey      hedera.PrivateKey
}

func NewClient(nodeAccMap map[string]hedera.AccountID, mirrorNodesList []string, strAccountId string, strPrivKey string) (*Client, error) {
	hdrClient := hedera.ClientForNetwork(nodeAccMap)
	hdrClient.SetMirrorNetwork(mirrorNodesList)

	accountId, err := hedera.AccountIDFromString(strAccountId)
	if err != nil {
		return nil, err
	}

	privateKey, err := hedera.PrivateKeyFromString(strPrivKey)
	if err != nil {
		return nil, err
	}

	client := &Client{
		hederaClient:    hdrClient,
		nodeAccMap:      nodeAccMap,
		mirrorNodesList: mirrorNodesList,
		accountId:       accountId,
		privateKey:      privateKey,
	}

	client.hederaClient.SetOperator(client.accountId, client.privateKey)

	return client, nil
}

func (client *Client) CreateNewAccount(asset int64) (*hedera.AccountInfo, error) {
	newAccount, err := hedera.NewAccountCreateTransaction().
		SetKey(client.privateKey).
		SetInitialBalance(hedera.HbarFromTinybar(asset)).
		Execute(client.hederaClient)

	if err != nil {
		return nil, err
	}

	receipt, err := newAccount.GetReceipt(client.hederaClient)
	if err != nil {
		return nil, err
	}

	// FIXME: into logger
	fmt.Println("Tx ID:", receipt.TransactionID)

	accInfo, err := hedera.NewAccountInfoQuery().
		SetAccountID(*receipt.AccountID).
		Execute(client.hederaClient)

	if err != nil {
		return nil, err
	}

	return &accInfo, nil
}
