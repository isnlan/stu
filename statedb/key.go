package main

import (
	"github.com/hyperledger/fabric-protos-go/ledger/rwset/kvrwset"
	"github.com/hyperledger/fabric/core/ledger/kvledger/txmgmt/rwsetutil"
	"github.com/hyperledger/fabric/core/ledger/kvledger/txmgmt/storageutil"
)

// applyTxRwset records the upsertion/deletion of a kv and updatation/deletion
// of associated metadata present in a txrwset
func (txops txOps) applyTxRwset(rwset *rwsetutil.TxRwSet) error {
	for _, nsRWSet := range rwset.NsRwSets {
		ns := nsRWSet.NameSpace
		for _, kvWrite := range nsRWSet.KvRwSet.Writes {
			txops.applyKVWrite(ns, "", kvWrite)
		}
		for _, kvMetadataWrite := range nsRWSet.KvRwSet.MetadataWrites {
			txops.applyMetadata(ns, "", kvMetadataWrite)
		}

		// apply collection level kvwrite and kvMetadataWrite
		for _, collHashRWset := range nsRWSet.CollHashedRwSets {
			coll := collHashRWset.CollectionName
			for _, hashedWrite := range collHashRWset.HashedRwSet.HashedWrites {
				txops.applyKVWrite(ns, coll,
					&kvrwset.KVWrite{
						Key:      string(hashedWrite.KeyHash),
						Value:    hashedWrite.ValueHash,
						IsDelete: hashedWrite.IsDelete,
					},
				)
			}

			for _, metadataWrite := range collHashRWset.HashedRwSet.MetadataWrites {
				txops.applyMetadata(ns, coll,
					&kvrwset.KVMetadataWrite{
						Key:     string(metadataWrite.KeyHash),
						Entries: metadataWrite.Entries,
					},
				)
			}
		}
	}
	return nil
}

// applyKVWrite records upsertion/deletion of a kvwrite
func (txops txOps) applyKVWrite(ns, coll string, kvWrite *kvrwset.KVWrite) {
	if kvWrite.IsDelete {
		txops.delete(compositeKey{ns, coll, kvWrite.Key})
	} else {
		txops.upsert(compositeKey{ns, coll, kvWrite.Key}, kvWrite.Value)
	}
}

// applyMetadata records updatation/deletion of a metadataWrite
func (txops txOps) applyMetadata(ns, coll string, metadataWrite *kvrwset.KVMetadataWrite) error {
	if metadataWrite.Entries == nil {
		txops.metadataDelete(compositeKey{ns, coll, metadataWrite.Key})
	} else {
		metadataBytes, err := storageutil.SerializeMetadata(metadataWrite.Entries)
		if err != nil {
			return err
		}
		txops.metadataUpdate(compositeKey{ns, coll, metadataWrite.Key}, metadataBytes)
	}
	return nil
}

type keyOpsFlag uint8

const (
	upsertVal      keyOpsFlag = 1 // 1 << 0
	metadataUpdate            = 2 // 1 << 1
	metadataDelete            = 4 // 1 << 2
	keyDelete                 = 8 // 1 << 3
)

type compositeKey struct {
	ns, coll, key string
}

type txOps map[compositeKey]*keyOps

type keyOps struct {
	flag     keyOpsFlag
	value    []byte
	metadata []byte
}

////////////////// txOps functions

func (txops txOps) upsert(k compositeKey, val []byte) {
	keyops := txops.getOrCreateKeyEntry(k)
	keyops.flag += upsertVal
	keyops.value = val
}

func (txops txOps) delete(k compositeKey) {
	keyops := txops.getOrCreateKeyEntry(k)
	keyops.flag += keyDelete
}

func (txops txOps) metadataUpdate(k compositeKey, metadata []byte) {
	keyops := txops.getOrCreateKeyEntry(k)
	keyops.flag += metadataUpdate
	keyops.metadata = metadata
}

func (txops txOps) metadataDelete(k compositeKey) {
	keyops := txops.getOrCreateKeyEntry(k)
	keyops.flag += metadataDelete
}

func (txops txOps) getOrCreateKeyEntry(k compositeKey) *keyOps {
	keyops, ok := txops[k]
	if !ok {
		keyops = &keyOps{}
		txops[k] = keyops
	}
	return keyops
}

////////////////// keyOps functions

func (keyops keyOps) isDelete() bool {
	return keyops.flag&(keyDelete) == keyDelete
}

func (keyops keyOps) isUpsertAndMetadataUpdate() bool {
	if keyops.flag&upsertVal == upsertVal {
		return keyops.flag&metadataUpdate == metadataUpdate ||
			keyops.flag&metadataDelete == metadataDelete
	}
	return false
}

func (keyops keyOps) isOnlyUpsert() bool {
	return keyops.flag|upsertVal == upsertVal
}
