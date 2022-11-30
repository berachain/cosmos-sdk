package mempool

import (
	crand "crypto/rand" // #nosec // crypto/rand is used for seed generation
	"encoding/binary"
	"fmt"
	"math/rand" // #nosec // math/rand is used for random selection and seeded from crypto/rand

	"github.com/huandu/skiplist"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/signing"
)

var (
	_ Mempool  = (*senderNonceMempool)(nil)
	_ Iterator = (*senderNonceMepoolIterator)(nil)
)

// senderNonceMempool is a mempool that prioritizes transactions within a sender by nonce, the lowest first,
// but selects a random sender on each iteration.  The mempool is iterated by:
//
// 1) Maintaining a separate list of nonce ordered txs per sender
// 2) For each select iteration, randomly choose a sender and pick the next nonce ordered tx from their list
// 3) Repeat 1,2 until the mempool is exhausted
//
// Note that PrepareProposal could choose to stop iteration before reaching the end if maxBytes is reached.
type senderNonceMempool struct {
	senders map[string]*skiplist.SkipList
	rnd     *rand.Rand
}

// NewSenderNonceMempool creates a new mempool that prioritizes transactions by nonce, the lowest first.
func NewSenderNonceMempool() Mempool {
	senderMap := make(map[string]*skiplist.SkipList)
	snp := &senderNonceMempool{
		senders: senderMap,
	}

	var seed int64
	err := binary.Read(crand.Reader, binary.BigEndian, &seed)
	if err != nil {
		panic(err)
	}
	snp.setSeed(seed)

	return snp
}

// NewSenderNonceMempoolWithSeed creates a new mempool that prioritizes transactions by nonce, the lowest first and sets the random seed.
func NewSenderNonceMempoolWithSeed(seed int64) Mempool {
	senderMap := make(map[string]*skiplist.SkipList)
	snp := &senderNonceMempool{
		senders: senderMap,
	}
	snp.setSeed(seed)
	return snp
}

func (snm *senderNonceMempool) setSeed(seed int64) {
	s1 := rand.NewSource(seed)
	snm.rnd = rand.New(s1) //#nosec // math/rand is seeded from crypto/rand by default
}

// Insert adds a tx to the mempool. It returns an error if the tx does not have at least one signer.
// priority is ignored.
func (snm *senderNonceMempool) Insert(_ sdk.Context, tx sdk.Tx) error {
	sigs, err := tx.(signing.SigVerifiableTx).GetSignaturesV2()
	if err != nil {
		return err
	}
	if len(sigs) == 0 {
		return fmt.Errorf("tx must have at least one signer")
	}

	sig := sigs[0]
	sender := sig.PubKey.Address().String()
	nonce := sig.Sequence
	senderTxs, found := snm.senders[sender]
	if !found {
		senderTxs = skiplist.New(skiplist.Uint64)
		snm.senders[sender] = senderTxs
	}
	senderTxs.Set(nonce, tx)

	return nil
}

// Select returns an iterator ordering transactions the mempool with the lowest nonce of a random selected sender first.
func (snm *senderNonceMempool) Select(_ sdk.Context, _ [][]byte) Iterator {
	var senders []string
	senderCursors := make(map[string]*skiplist.Element)

	orderedSenders := skiplist.New(skiplist.String)

	// #nosec
	for s := range snm.senders {
		orderedSenders.Set(s, s)
	}

	s := orderedSenders.Front()
	for s != nil {
		sender := s.Value.(string)
		senders = append(senders, sender)
		senderCursors[sender] = snm.senders[sender].Front()
		s = s.Next()
	}

	iter := &senderNonceMepoolIterator{
		senders:       senders,
		rnd:           snm.rnd,
		senderCursors: senderCursors,
	}

	return iter.Next()
}

// CountTx returns the total count of txs in the mempool.
func (snm *senderNonceMempool) CountTx() int {
	count := 0

	// Disable gosec here since we need neither strong randomness nor deterministic iteration.
	// #nosec
	for _, value := range snm.senders {
		count += value.Len()
	}
	return count
}

// Remove removes a tx from the mempool. It returns an error if the tx does not have at least one signer or the tx
// was not found in the pool.
func (snm *senderNonceMempool) Remove(tx sdk.Tx) error {
	sigs, err := tx.(signing.SigVerifiableTx).GetSignaturesV2()
	if err != nil {
		return err
	}
	if len(sigs) == 0 {
		return fmt.Errorf("tx must have at least one signer")
	}

	sig := sigs[0]
	sender := sig.PubKey.Address().String()
	nonce := sig.Sequence
	senderTxs, found := snm.senders[sender]
	if !found {
		return ErrTxNotFound
	}

	res := senderTxs.Remove(nonce)
	if res == nil {
		return ErrTxNotFound
	}

	if senderTxs.Len() == 0 {
		delete(snm.senders, sender)
	}
	return nil
}

type senderNonceMepoolIterator struct {
	rnd           *rand.Rand
	currentTx     *skiplist.Element
	senders       []string
	senderCursors map[string]*skiplist.Element
}

// Next returns the next iterator state which will contain a tx with the next smallest nonce of a randomly
// selected sender.
func (i *senderNonceMepoolIterator) Next() Iterator {
	for len(i.senders) > 0 {
		senderIndex := i.rnd.Intn(len(i.senders))
		sender := i.senders[senderIndex]
		senderCursor, found := i.senderCursors[sender]
		if !found {
			i.senders = removeAtIndex(i.senders, senderIndex)
			continue
		}

		if nextCursor := senderCursor.Next(); nextCursor != nil {
			i.senderCursors[sender] = nextCursor
		} else {
			i.senders = removeAtIndex(i.senders, senderIndex)
		}

		return &senderNonceMepoolIterator{
			senders:       i.senders,
			currentTx:     senderCursor,
			rnd:           i.rnd,
			senderCursors: i.senderCursors,
		}
	}

	return nil
}

func (i *senderNonceMepoolIterator) Tx() sdk.Tx {
	return i.currentTx.Value.(sdk.Tx)
}

func removeAtIndex[T any](slice []T, index int) []T {
	return append(slice[:index], slice[index+1:]...)
}
