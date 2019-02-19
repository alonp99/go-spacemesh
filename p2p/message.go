package p2p

import (
	"github.com/alonp99/go-spacemesh/p2p/config"
	"github.com/alonp99/go-spacemesh/p2p/p2pcrypto"
	"github.com/alonp99/go-spacemesh/p2p/pb"
	"time"
)

// NewProtocolMessageMetadata creates meta-data for an outgoing protocol message authored by this node.
func NewProtocolMessageMetadata(author p2pcrypto.PublicKey, protocol string) *pb.Metadata {
	return &pb.Metadata{
		NextProtocol:  protocol,
		ClientVersion: config.ClientVersion,
		Timestamp:     time.Now().Unix(),
		AuthPubkey:    author.Bytes(),
	}
}
