package main

import (
	"context"
	"flag"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"os"
	"strings"
	"time"
)

type DepositCertificate struct {
	//Id primitive.ObjectID `json:"uid" bson:"_id"` // id
	//Type         string             `json:"type" bson:"type"`                 // 数据格式
	//BduId        string             `json:"bduId" bson:"bduId"`               // 数字身份唯一标识
	//IpfsAddr     string             `json:"ipfsAddr" bson:"ipfsAddr"`         // ipfs地址
	//RecordId     string             `json:"recordId" bson:"recordId"`         // 特征值
	//Size         int                `json:"size" bson:"size"`                 // 数据大小
	//Mtime        time.Time          `json:"mtime" bson:"mtime"`               // 存证时间
	//TxId         string             `json:"txId" bson:"txId"`                 // 区块链交易号
	//BlockNum     string             `json:"blockNum" bson:"blockNum"`         // 区块号
	//BlockHash    string             `json:"blockHash" bson:"blockHash"`       // 区块哈希
	//PreBlockHash string             `json:"preBlockHash" bson:"preBlockHash"` // 前区块哈希
	//DcType       string             `json:"dcType" bson:"dcType"`             // 存证类型
	//Name         string             `json:"name" bson:"name"`                 // 文件名
	//FileAlias    string             `json:"fileAlias" bson:"fileAlias"`       // 文件别名
	//FileId       string             `json:"fileId" bson:"fileId"`             // 企业云盘文件id
	CreateTime time.Time `json:"createTime" bson:"createTime"` // 创建时间
	//UpdateTime   time.Time          `json:"updateTime" bson:"updateTime"`     // 更新时间
	//Ext          string             `json:"ext" bson:"ext"`                   // 扩展名
	Key string `json:"key" bson:"key"` // 用户自定义的数据键值
	//Deleted      bool               `json:"deleted" bson:"deleted"`           // 是否删除
}

func initMongo(uri string) (*mongo.Client, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}
	return client, nil
}

func GetKedasDB(cli *mongo.Client) *mongo.Database {
	return cli.Database("kedas")
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func find(collection *mongo.Collection) (map[string]*DBID, error) {
	cur, err := collection.Find(context.Background(), bson.M{"key": bson.M{"$exists": true}})
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())

	dbids := map[string]*DBID{}

	for cur.Next(context.Background()) {
		var dc DepositCertificate
		err := cur.Decode(&dc)
		if err != nil {
			return nil, err
		}

		keys := strings.Split(dc.Key, "_")
		if len(keys) == 4 {
			id := keys[0]
			if _, found := dbids[id]; !found {
				dbids[id] = &DBID{
					Count: 0,
					Start: dc.CreateTime,
					End:   dc.CreateTime,
				}
			}

			dbid := dbids[id]
			dbid.Count += 1
			if dc.CreateTime.Before(dbid.Start) {
				dbid.Start = dc.CreateTime
			}
			if dc.CreateTime.After(dbid.End) {
				dbid.End = dc.CreateTime
			}
		}
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return dbids, nil
}

type DBID struct {
	Count int
	Start time.Time
	End   time.Time
}

func main() {
	f, err := os.Create("dbids.csv")
	check(err)
	defer f.Close()

	f.WriteString("dbid, count, start, end,\n")

	fmt.Println("use:", "./mongo-tool -uri='mongodb://root:4LY0Ul9yNw@127.0.0.1:27017/admin?authSource=admin'")
	url := flag.String("uri", "mongodb://kbcs-formal:NBm2Hg2kLIQUi9lb@10.244.4.188:27017/kedas?authSource=admin", "mongo uri")
	flag.Parse()
	fmt.Println("URI:", *url)
	client, err := initMongo(*url)
	check(err)
	db := GetKedasDB(client).Collection("deposit_certificate")
	dbids, err := find(db)
	check(err)

	total := 0
	max := 0
	for k, v := range dbids {
		max += v.Count
		_, err := f.WriteString(fmt.Sprintf("%s, %d, %v, %v,\n", k, v.Count, v.Start, v.End))
		if err != nil {
			fmt.Printf("ERROR: %v\n", err)
		}
		f.Sync()
	}

	f.WriteString("00000, ")
	f.Sync()
	fmt.Printf("Total: %d, DBID count: %d\n", total, len(dbids))
}
