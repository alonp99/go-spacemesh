package hare

import (
	"bytes"
	"github.com/spacemeshos/go-spacemesh/hare/pb"
)

type ProposalTracker struct {
	proposal      *pb.HareMessage // maps PubKey->Proposal
	isConflicting bool            // maps PubKey->ConflictStatus
}

func NewProposalTracker(expectedSize int) *ProposalTracker {
	pt := &ProposalTracker{}
	pt.proposal = nil
	pt.isConflicting = false

	return pt
}

// Tracks and detects conflicting messages
// Returns true if it is the first time we see m, false otherwise
func (pt *ProposalTracker) TrackConflict(m *pb.HareMessage) bool {


	return true
}

func (pt *ProposalTracker) OnProposal(msg *pb.HareMessage) {
	if pt.proposal == nil { // first leader
		pt.proposal = msg // just update
		return
	}

	// if same sender then we should check for equivocation
	if bytes.Equal(pt.proposal.PubKey, msg.PubKey) {
		s := NewSet(msg.Message.Values)
		g := NewSet(pt.proposal.Message.Values)
		if !s.Equals(g) { // equivocation detected
			pt.isConflicting = true
		}

		return // process done
	}

	// TODO: ignore msgs with higher ranked role proof

	// first time we see msg
	pt.proposal = msg        // update leader msg
	pt.isConflicting = false // assume no conflict
}

func (pt *ProposalTracker) OnLateProposal(msg *pb.HareMessage) {
	// if same sender then we should check for equivocation
	if bytes.Equal(pt.proposal.PubKey, msg.PubKey) {
		s := NewSet(msg.Message.Values)
		g := NewSet(pt.proposal.Message.Values)
		if !s.Equals(g) { // equivocation detected
			pt.isConflicting = true
		}
	}

	// not equal check rank
	// TODO: if msg is ranked lower than proposal set isConflicting to true
}

func (pt *ProposalTracker) IsConflicting() bool {
	return pt.isConflicting
}

func (pt *ProposalTracker) ProposedSet() *Set {
	if pt.proposal == nil {
		return nil
	}

	if pt.isConflicting {
		return nil
	}

	return NewSet(pt.proposal.Message.Values)
}

