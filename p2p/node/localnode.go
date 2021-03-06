package node

import (
	"github.com/alonp99/go-spacemesh/filesystem"
	"github.com/alonp99/go-spacemesh/log"
	"github.com/alonp99/go-spacemesh/p2p/config"
	"github.com/alonp99/go-spacemesh/p2p/p2pcrypto"
)

// LocalNode implementation.
type LocalNode struct {
	Node
	privKey p2pcrypto.PrivateKey

	networkID int8

	log.Log
}

// NetworkID returns the local node's network id (testnet/mainnet, etc..)
func (n *LocalNode) NetworkID() int8 {
	return n.networkID
}

// PrivateKey returns this node's private key.
func (n *LocalNode) PrivateKey() p2pcrypto.PrivateKey {
	return n.privKey
}

// NewLocalNode creates a local node with a provided ip address.
// Attempts to set node node from persisted data in local store.
// Creates a new node if none was loaded.
func NewLocalNode(config config.Config, address string, persist bool) (*LocalNode, error) {

	if len(config.NodeID) > 0 {
		// user provided node id/pubkey via the cli - attempt to start that node w persisted data
		data, err := readNodeData(config.NodeID)
		if err != nil {
			return nil, err
		}

		return newLocalNodeFromFile(address, data, persist)
	}

	// look for persisted node data in the nodes directory
	// load the node with the data of the first node found
	nodeData, err := readFirstNodeData()
	if err != nil {
		log.Warning("failed to read node data from local store")
	}

	if nodeData != nil {
		// create node using persisted node data
		return newLocalNodeFromFile(address, nodeData, persist)
	}

	// generate new node
	return NewNodeIdentity(config, address, persist)
}

// NewNodeIdentity creates a new local node without attempting to restore node from local store.
func NewNodeIdentity(config config.Config, address string, persist bool) (*LocalNode, error) {
	priv, pub, err := p2pcrypto.GenerateKeyPair()
	if err != nil {
		return nil, err
	}
	return newLocalNodeWithKeys(pub, priv, address, config.NetworkID, persist)
}

func newLocalNodeWithKeys(pubKey p2pcrypto.PublicKey, privKey p2pcrypto.PrivateKey, address string, networkID int8, persist bool) (*LocalNode, error) {

	n := &LocalNode{
		Node: Node{
			pubKey:  pubKey,
			address: address,
		},
		networkID: networkID,
		privKey:   privKey,
	}

	if !persist {
		n.Log = log.New(n.pubKey.String(), "", "")
		return n, nil
	}

	dataDir, err := filesystem.EnsureNodesDataDirectory(config.NodesDirectoryName)
	if err != nil {
		return nil, err
	}

	nodeDir, err := filesystem.EnsureNodeDataDirectory(dataDir, n.pubKey.String())
	if err != nil {
		return nil, err
	}

	// persistent logging
	n.Log = log.New(n.pubKey.String(), nodeDir, "node.log")

	n.Info("Local node identity >> %v", n.String())

	// persist store data so we can start it on future app sessions
	err = n.persistData()
	if err != nil { // no much use of starting if we can't store node private key in store
		n.Error("failed to persist node data to local store", err)
		return nil, err
	}

	return n, nil
}

// Creates a new node from persisted NodeData.
func newLocalNodeFromFile(address string, d *nodeFileData, persist bool) (*LocalNode, error) {

	priv, err := p2pcrypto.NewPrivateKeyFromBase58(d.PrivKey)
	if err != nil {
		return nil, err
	}

	pub, err := p2pcrypto.NewPublicKeyFromBase58(d.PubKey)
	if err != nil {
		return nil, err
	}

	log.Info(">>>> Creating node identity from filesystem existing key %s", pub.String())

	return newLocalNodeWithKeys(pub, priv, address, d.NetworkID, persist)
}
