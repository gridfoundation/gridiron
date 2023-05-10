package irctesting

import (
	"time"

	"github.com/cosmos/ibc-go/v3/modules/core/exported"
	ibctesting "github.com/cosmos/ibc-go/v3/testing"

	sdk "github.com/cosmos/cosmos-sdk/types"

	rollapptypes "github.com/gridfoundation/gridiron/x/rollapp/types"
)

var (
	hash32 = []byte("12345678901234567890123456789012")
)

// FuryintTestChainClient is a wrapper for baseTestChainClient
// it is used to intersept the NextBlock() function
// in order to track after the app.Commit() and get an AppHash.
// The AppHash is used for generating BD
type FuryintTestChainClient struct {
	baseTestChainClient ibctesting.TestChainClientI
	baseTestChain       *ibctesting.TestChain
	// track rollapp BDs for update
	bds rollapptypes.BlockDescriptors
}

func (furyintC *FuryintTestChainClient) GetContext() sdk.Context {
	return furyintC.baseTestChainClient.GetContext()
}
func (furyintC *FuryintTestChainClient) NextBlock() {
	// NextBlock sets the last header to the current header and increments the current header to be
	// at the next block height.
	// This function must only be called after app.Commit() occurs
	// The app.Commit() calculate the new AppHash which gets into the BlockHeader
	furyintC.baseTestChainClient.NextBlock()
	// app.Commit() resets the deliver state (which holds the changes of the multistore as created in the block)
	// NextBlock opens a new block and calls to BeginBlock, which initializes the deliver state
	// so we can access the last committed block only after NextBlock() and not between app.Commit() & NextBlock()
	// otherwise, GetContext() returns nil
	bd := rollapptypes.BlockDescriptor{
		Height:                 uint64(furyintC.baseTestChain.GetContext().BlockHeader().Height),
		StateRoot:              furyintC.baseTestChain.GetContext().BlockHeader().AppHash,
		IntermediateStatesRoot: hash32,
	}
	// add new BD to list
	furyintC.bds.BD = append(furyintC.bds.BD, bd)
}
func (furyintC *FuryintTestChainClient) BeginBlock() {
	furyintC.baseTestChainClient.BeginBlock()
}
func (furyintC *FuryintTestChainClient) UpdateCurrentHeaderTime(t time.Time) {
	furyintC.baseTestChainClient.UpdateCurrentHeaderTime(t)
}
func (furyintC *FuryintTestChainClient) ClientConfigToState(ClientConfig ibctesting.ClientConfig) exported.ClientState {
	return furyintC.baseTestChainClient.ClientConfigToState(ClientConfig)
}
func (furyintC *FuryintTestChainClient) GetConsensusState() exported.ConsensusState {
	return furyintC.baseTestChainClient.GetConsensusState()
}
func (furyintC *FuryintTestChainClient) NewConfig() ibctesting.ClientConfig {
	return furyintC.baseTestChainClient.NewConfig()
}
func (furyintC *FuryintTestChainClient) GetSelfClientType() string {
	return furyintC.baseTestChainClient.GetSelfClientType()
}
func (furyintC *FuryintTestChainClient) GetLastHeader() interface{} {
	return furyintC.baseTestChainClient.GetLastHeader()
}
