package document

import (
	"time"

	"github.com/irisnet/irishub-server/models"
	"github.com/irisnet/irishub-server/utils/constants"
	"github.com/irisnet/irishub-server/utils/helper"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	CollectionNmCommonTx = "tx_common"
)

type CommonTx struct {
	TxHash    string            `json:"tx_hash" bson:"tx_hash"`
	Time      time.Time         `json:"time" bson:"time"`
	Height    int64             `json:"height" bson:"height"`
	From      string            `json:"from" bson:"from"`
	To        string            `json:"to" bson:"to"`
	Amount    Coins             `json:"amount" bson:"amount"`
	Type      string            `json:"type" bson:"type"`
	Fee       Fee               `bson:"fee"`
	Memo      string            `bson:"memo"`
	Status    string            `bson:"status"`
	Log       string            `bson:"log"`
	GasUsed   int64             `bson:"gas_used"`
	GasPrice  float64           `bson:"gas_price"`
	ActualFee ActualFee         `bson:"actual_fee"`
	Tags      map[string]string `bson:"tags"`

	Candidate Candidate `json:"candidate"`
}

func (d CommonTx) Name() string {
	return CollectionNmCommonTx
}

func (d CommonTx) PkKvPair() map[string]interface{} {
	return bson.M{"tx_hash": d.TxHash}
}

func (d CommonTx) Query(
	query bson.M, fields bson.M,
	skip int, limit int, sorts ...string) (results []CommonTx, err error) {

	exop := func(c *mgo.Collection) error {
		return c.Find(query).Select(fields).Sort(sorts...).Skip(skip).Limit(limit).All(&results)
	}
	return results, models.ExecCollection(d.Name(), exop)
}

func (d CommonTx) GetList(address string, txType string,
	startTime time.Time, endTime time.Time,
	skip int, limit int, sorts []string, ext string, height int64) (
	[]CommonTx, error) {

	query := bson.M{}

	if txType == "" {
		query = bson.M{
			"$or": []bson.M{
				bson.M{"from": address},
				bson.M{"to": address},
			},
		}
		query["type"] = bson.M{
			"$in": []string{
				constants.TxTypeFrontMapDb[constants.TxTypeCoinSend],
				constants.TxTypeFrontMapDb[constants.TxTypeCoinReceive],
				constants.TxTypeFrontMapDb[constants.TxTypeStakeDelegate],
				constants.TxTypeFrontMapDb[constants.TxTypeStakeBeginUnBonding],
			},
		}
	} else {
		switch txType {
		case constants.TxTypeCoinReceive:
			query["to"] = address
			query["type"] = constants.TxTypeFrontMapDb[txType]
			break
		case constants.TxTypeCoinSend, constants.TxTypeStakeDelegate,
			constants.TxTypeStakeBeginUnBonding:
			query["from"] = address
			query["type"] = constants.TxTypeFrontMapDb[txType]
			if ext != "" && txType != constants.TxTypeCoinSend {
				query["to"] = ext
			}
			break
		case constants.TxTypeStakeUnbond:
			query["from"] = address
			query["type"] = bson.M{
				"$in": []string{
					constants.TxTypeFrontMapDb[constants.TxTypeStakeBeginUnBonding],
				},
			}
			if ext != "" {
				query["to"] = ext
			}
			break
		case constants.TxTypeStake:
			query["from"] = address
			query["type"] = bson.M{
				"$in": []string{
					constants.TxTypeFrontMapDb[constants.TxTypeStakeDelegate],
					constants.TxTypeFrontMapDb[constants.TxTypeStakeBeginUnBonding],
				},
			}
			if ext != "" {
				query["to"] = ext
			}
		default:
			return nil, nil
		}
	}

	if startTime.IsZero() {
		startTime, _ = helper.ParseFullTime(constants.TIME_START)
	}
	if endTime.IsZero() {
		endTime = time.Now()
	}
	query["time"] = bson.M{
		"$gte": startTime,
		"$lte": endTime,
	}
	if height > 0 {
		query["height"] = bson.M{
			"$gt": height,
		}
	}

	fields := bson.M{}

	var txs []CommonTx
	commTxs, err := d.Query(query, fields, skip, limit, sorts...)
	if err == nil {
		for _, tx := range commTxs {
			if txType == "" {
				if tx.Type != constants.TxTypeFrontMapDb[constants.TxTypeStakeDelegate] {
					txs = append(txs, tx)
				} else {
					if tx.From == address {
						txs = append(txs, tx)
					}
				}
			} else {
				txs = append(txs, tx)
			}
		}
	}

	return txs, err
}

func (d CommonTx) GetRewardList(delAddr string) (results []CommonTx) {
	query := bson.M{}
	query["from"] = delAddr
	query["type"] = bson.M{
		"$in": []string{
			constants.DbTxTypeSetWithdrawAddress,
			constants.DbTxTypeWithdrawDelegatorReward,
			constants.DbTxTypeWithdrawDelegatorRewardsAll,
		},
	}
	fields := bson.M{}
	results, _ = d.Query(query, fields, 0, 1000, "-time")
	return results
}

func (d CommonTx) GetDetail(txHash string) (CommonTx, error) {
	query := bson.M{
		"tx_hash": txHash,
	}
	fields := bson.M{}
	var (
		sorts []string
	)

	txs, err := d.Query(query, fields, 0, 1, sorts...)
	if err != nil || len(txs) == 0 {
		return CommonTx{}, err
	}
	return txs[0], nil
}
