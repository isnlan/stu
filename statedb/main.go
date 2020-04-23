package main

import (
	"fmt"

	"github.com/hyperledger/fabric/core/ledger"
	"github.com/hyperledger/fabric/core/ledger/kvledger/txmgmt/rwsetutil"
	"github.com/hyperledger/fabric/core/ledger/kvledger/txmgmt/statedb"
	"github.com/hyperledger/fabric/core/ledger/kvledger/txmgmt/statedb/stateleveldb"
	"github.com/hyperledger/fabric/core/ledger/kvledger/txmgmt/version"
)

// ledger.TxSimulator
//var _ ledger.TxSimulator = &LockBasedTxSimulator{}

type LockBasedTxSimulator struct {
	TxID                      string
	rwsetBuilder              *rwsetutil.RWSetBuilder
	simulationResultsComputed bool
	db                        statedb.VersionedDB
}

func (s *LockBasedTxSimulator) GetState(ns string, key string) ([]byte, error) {
	state, err := s.db.GetState(ns, key)
	if err != nil {
		return nil, err
	}
	s.rwsetBuilder.AddToReadSet("default-ns", key, state.Version)
	return state.Value, err
}

func (s *LockBasedTxSimulator) SetState(ns string, key string, value []byte) error {
	s.rwsetBuilder.AddToWriteSet(ns, key, value)
	return nil
}

func (s *LockBasedTxSimulator) DeleteState(ns string, key string) error {
	return s.SetState(ns, key, nil)
}

func (s *LockBasedTxSimulator) GetTxSimulationResults() (*ledger.TxSimulationResults, error) {
	return s.rwsetBuilder.GetTxSimulationResults()
}

func StateDBUse(vdbProivder statedb.VersionedDBProvider) {
	versionedDB, err := vdbProivder.GetDBHandle("chain_id")
	check(err)
	sim := &LockBasedTxSimulator{
		TxID:                      "tx1",
		rwsetBuilder:              rwsetutil.NewRWSetBuilder(),
		simulationResultsComputed: false,
		db:                        versionedDB,
	}

	{
		// simulator
		err = sim.SetState("contract_name", "k1", []byte("v1"))
		check(err)
	}

	results, err := sim.GetTxSimulationResults()
	check(err)

	// Validate
	txRwSetTemp, err := rwsetutil.TxRwSetFromProtoMsg(results.PubSimulationResults)
	check(err)

	err = validateWriteset(txRwSetTemp, versionedDB.ValidateKeyValue)
	check(err)

	height := &version.Height{
		BlockNum: 0,
		TxNum:    1,
	}
	batch, err := ApplyWriteSet(txRwSetTemp, height)
	check(err)

	err = versionedDB.ApplyUpdates(batch, height)
	check(err)

	{
		// simulator
		state, err := versionedDB.GetState("contract_name", "k1")
		check(err)
		fmt.Println(state)
	}
}

func ApplyWriteSet(txRWSet *rwsetutil.TxRwSet, txHeight *version.Height) (*statedb.UpdateBatch, error) {
	txops := txOps{}
	err := txops.applyTxRwset(txRWSet)
	if err != nil {
		return nil, err
	}

	batch := statedb.NewUpdateBatch()

	for compositeKey, keyops := range txops {
		if compositeKey.coll == "" {
			ns, key := compositeKey.ns, compositeKey.key
			if keyops.isDelete() {
				batch.Delete(ns, key, txHeight)
			} else {
				batch.PutValAndMetadata(ns, key, keyops.value, keyops.metadata, txHeight)
			}
		}
	}
	return batch, nil
}

func validateWriteset(txRWSet *rwsetutil.TxRwSet, validateKVFunc func(key string, value []byte) error) error {
	for _, nsRwSet := range txRWSet.NsRwSets {
		pubWriteset := nsRwSet.KvRwSet
		if pubWriteset == nil {
			continue
		}
		for _, kvwrite := range pubWriteset.Writes {
			if err := validateKVFunc(kvwrite.Key, kvwrite.Value); err != nil {
				return err
			}
		}
	}
	return nil
}

func main() {
	vdbProvider, err := stateleveldb.NewVersionedDBProvider("./data")
	check(err)

	StateDBUse(vdbProvider)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
