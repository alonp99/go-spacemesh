package consensus

import (
	"github.com/alonp99/go-spacemesh/config"
	"github.com/alonp99/go-spacemesh/log"
	"github.com/alonp99/go-spacemesh/mesh"
)

type Algorithm struct {
	Tortoise
}

type Tortoise interface {
	handleIncomingLayer(ll *mesh.Layer)
	latestComplete() mesh.LayerID
	getVote(id mesh.BlockID) vec
	getVotes() map[mesh.BlockID]vec
}

func NewAlgorithm(trtl Tortoise) *Algorithm {
	return &Algorithm{Tortoise: trtl}
}

func (alg *Algorithm) HandleLateBlock(b *mesh.Block) {
	//todo feed all layers from b's layer to tortoise
	log.Info("received block with layer Id %v block id: %v ", b.Layer(), b.ID())
}

func (alg *Algorithm) HandleIncomingLayer(ll *mesh.Layer) (mesh.LayerID, mesh.LayerID) {
	oldPbase := alg.latestComplete()
	alg.Tortoise.handleIncomingLayer(ll)
	newPbase := alg.latestComplete()
	updateMetrics(alg, ll)
	return oldPbase, newPbase
}

func updateMetrics(alg *Algorithm, ll *mesh.Layer) {
	pbaseCount.Set(float64(alg.latestComplete()))
	processedCount.Set(float64(ll.Index()))
	var valid float64
	var invalid float64
	for _, k := range alg.getVotes() {
		if k == Support {
			valid++
		} else {
			invalid++
		}
	}
	validBlocks.Set(valid)
	invalidBlocks.Set(invalid)
}

func (alg *Algorithm) ContextualValidity(id mesh.BlockID) bool {
	return alg.getVote(id) == Support
}

func CreateGenesisBlock() *mesh.Block {
	log.Info("Creating genesis")
	bl := &mesh.Block{
		Id:         mesh.BlockID(config.GenesisId),
		LayerIndex: 0,
		Data:       []byte("genesis"),
	}
	return bl
}

func createGenesisLayer() *mesh.Layer {
	log.Info("Creating genesis")
	l := mesh.NewLayer(Genesis)
	l.AddBlock(CreateGenesisBlock())
	return l
}
