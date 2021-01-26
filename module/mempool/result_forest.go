// (c) 2019 Dapper Labs - ALL RIGHTS RESERVED

package mempool

import (
	"github.com/onflow/flow-go/model/flow"
)

// ResultForest represents a concurrency-safe memory pool for execution receipts.
// Its is aware of the tree structure formed by execution results. To enable
// this, the mempool utilizes knowledge about the height of the block the result
// is for. Hence, the Mempool can only store and process Receipts whose block
// is known.
type ResultForest interface {

	// Add the given execution receipt to the memory pool. Requires height
	// of the block the receipt is for. We enforce data consistency on an API
	// level by using the block header as input.
	Add(receipt *flow.ExecutionReceipt, block *flow.Header) (bool, error)

	// ReachableReceipts returns a slice of ExecutionReceipt, whose result
	// is computationally reachable from resultID. Context:
	// * Conceptually, the Execution results form a tree, which we refer to as
	//   Execution Tree. A branch in the execution can be due to a fork in the main
	//   chain. Furthermore, the execution branches if ENs disagree about the result
	//   for the same block.
	// * As the ID of an execution result contains the BlockID, which the result
	//   for, all Execution Results with the same ID necessarily are for the same
	//   block. All Execution Receipts committing to the same result from an
	//   equivalence class and can be represented as one vertex in the Execution
	//   Tree.
	// * An execution result r1 points (field ExecutionResult.ParentResultID) to
	//   its parent result r0 , whose end state was used as the starting state
	//   to compute r1. Formally, we have an edge r0 -> r1 in the Execution Tree,
	//   if a result r1 is stored in the mempool, whose ParentResultID points to
	//   r0.
	// ReachableReceipts implements a tree search on the Execution Tree starting
	// from the provided resultID. Execution Receipts are traversed in a
	// parent-first manner, meaning that a the parent result is traversed
	// _before_ any of its derived results. The algorithm only traverses to
	// results, for which there exists a sequence of interim result in the
	// mempool without any gaps.
	//
	// Two filters are supplied:
	// * blockFilter: the tree search will only travers to results for
	//   blocks which pass the filter. Often higher-level logic is only
	//   interested in results for blocks in a specific fork. Such can be
	//   implemented by a suitable blockFilter.
	// * receiptFilter: for a reachable result (subject to the restrictions
	//   imposed by blockFilter, all known receipts are returned.
	//   While _all_ Receipts for the parent result are guaranteed to be
	//   listed before the receipts for the derived results, there is no
	//   specific ordering for the receipts committing to the same result
	//   (random order). If only a subset of receipts for a result is desired
	//   (e.g. for de-duplication with parent blocks), receiptFilter should
	//   be used.
	// Note the important difference between the two filters:
	// * The blockFilter suppresses traversal to derived results.
	// * The receiptFilter does _not_ suppresses traversal to derived results.
	//   Only individual receipts are dropped.
	ReachableReceipts(resultID flow.Identifier, blockFilter BlockFilter, receiptFilter ReceiptFilter) ([]*flow.ExecutionReceipt, error)

	//// Has checks if the given receipt is part of the memory pool.
	//Has(receiptID flow.Identifier) bool
	//
	//// Rem will remove a receipt by ID.
	//Rem(receiptID flow.Identifier) bool
	//
	//// ByID retrieve the execution receipt with the given ID from the memory
	//// pool. It will return false if it was not found in the mempool.
	//ByID(receiptID flow.Identifier) (*flow.ExecutionReceipt, bool)
	//
	//// Size will return the current size of the memory pool.
	//Size() uint
	//
	//// All will return a list of all receipts in the memory pool.
	//All() []*flow.ExecutionReceipt
}

// BlockFilter is used for controlling the ResultForest's Execution Tree search.
// The search only traverses to results for blocks which pass the filter.
// If an the block for an execution result does not pass the filter, the entire
// sub-tree of derived results is not traversed.
type BlockFilter func(header *flow.Header) bool

// ReceiptFilter is used to drop specific receipts from. It does NOT
// affect the ResultForest's Execution Tree search.
type ReceiptFilter func(receipt *flow.ExecutionReceipt) bool