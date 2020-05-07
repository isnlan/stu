package main

import "github.com/hyperledger/fabric/core/ledger"

var _ ledger.TxSimulator = &TxSimulator{}

type TxSimulator struct {
}

func (t TxSimulator) GetState(namespace string, key string) ([]byte, error) {
	panic("implement me")
}

func (t TxSimulator) GetStateRangeScanIterator(namespace string, startKey string, endKey string) (ledger2.ResultsIterator, error) {
	panic("implement me")
}

func (t TxSimulator) GetPrivateDataHash(namespace, collection, key string) ([]byte, error) {
	panic("implement me")
}

func (t TxSimulator) GetStateMetadata(namespace, key string) (map[string][]byte, error) {
	panic("implement me")
}

func (t TxSimulator) GetStateMultipleKeys(namespace string, keys []string) ([][]byte, error) {
	panic("implement me")
}

func (t TxSimulator) GetStateRangeScanIteratorWithMetadata(namespace string, startKey, endKey string, metadata map[string]interface{}) (ledger.QueryResultsIterator, error) {
	panic("implement me")
}

func (t TxSimulator) ExecuteQuery(namespace, query string) (ledger2.ResultsIterator, error) {
	panic("implement me")
}

func (t TxSimulator) ExecuteQueryWithMetadata(namespace, query string, metadata map[string]interface{}) (ledger.QueryResultsIterator, error) {
	panic("implement me")
}

func (t TxSimulator) GetPrivateData(namespace, collection, key string) ([]byte, error) {
	panic("implement me")
}

func (t TxSimulator) GetPrivateDataMetadata(namespace, collection, key string) (map[string][]byte, error) {
	panic("implement me")
}

func (t TxSimulator) GetPrivateDataMetadataByHash(namespace, collection string, keyhash []byte) (map[string][]byte, error) {
	panic("implement me")
}

func (t TxSimulator) GetPrivateDataMultipleKeys(namespace, collection string, keys []string) ([][]byte, error) {
	panic("implement me")
}

func (t TxSimulator) GetPrivateDataRangeScanIterator(namespace, collection, startKey, endKey string) (ledger2.ResultsIterator, error) {
	panic("implement me")
}

func (t TxSimulator) ExecuteQueryOnPrivateData(namespace, collection, query string) (ledger2.ResultsIterator, error) {
	panic("implement me")
}

func (t TxSimulator) Done() {
	panic("implement me")
}

func (t TxSimulator) SetState(namespace string, key string, value []byte) error {
	panic("implement me")
}

func (t TxSimulator) DeleteState(namespace string, key string) error {
	panic("implement me")
}

func (t TxSimulator) SetStateMultipleKeys(namespace string, kvs map[string][]byte) error {
	panic("implement me")
}

func (t TxSimulator) SetStateMetadata(namespace, key string, metadata map[string][]byte) error {
	panic("implement me")
}

func (t TxSimulator) DeleteStateMetadata(namespace, key string) error {
	panic("implement me")
}

func (t TxSimulator) ExecuteUpdate(query string) error {
	panic("implement me")
}

func (t TxSimulator) SetPrivateData(namespace, collection, key string, value []byte) error {
	panic("implement me")
}

func (t TxSimulator) SetPrivateDataMultipleKeys(namespace, collection string, kvs map[string][]byte) error {
	panic("implement me")
}

func (t TxSimulator) DeletePrivateData(namespace, collection, key string) error {
	panic("implement me")
}

func (t TxSimulator) SetPrivateDataMetadata(namespace, collection, key string, metadata map[string][]byte) error {
	panic("implement me")
}

func (t TxSimulator) DeletePrivateDataMetadata(namespace, collection, key string) error {
	panic("implement me")
}

func (t TxSimulator) GetTxSimulationResults() (*ledger.TxSimulationResults, error) {
	panic("implement me")
}
