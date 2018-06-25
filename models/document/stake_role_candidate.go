package document

import (
	"time"

	"github.com/irisnet/irishub-server/models"
	"github.com/irisnet/irishub-server/modules/logger"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	CollectionNmStakeRoleCandidate = "stake_role_candidate"
)

type Candidate struct {
	Address     string      `json:"address" bson:"address"` // owner
	PubKey      string      `json:"pub_key" bson:"pub_key"`
	Shares      int64       `json:"shares" bson:"shares"`
	Description Description `json:"description" bson:"description"`  // Description terms for the candidate
    UpdateTime  time.Time   `json:"update_time" bson:"update_time"`

	VotingPower float64     `json:"voting_power"` // Voting power if pubKey is a considered a validator
	Delegators []Delegator  `json:"delegators"`

}

func (d Candidate) Name() string {
	return CollectionNmStakeRoleCandidate
}

func (d Candidate) PkKvPair() map[string]interface{} {
	return bson.M{"pub_key": d.PubKey}
}

func (d Candidate) Query(
	query bson.M, skip int, limit int, sorts ...string,
	) (results []Candidate, err error) {
	exop := func(c *mgo.Collection) error {
		return c.Find(query).Sort(sorts...).Skip(skip).Limit(limit).All(&results)
	}
	return results, models.ExecCollection(d.Name(), exop)
}


func (d Candidate) GetCandidatesList(q string, sorts []string, skip int, limit int) ([]Candidate, error)  {
	query := bson.M{
		"shares": &bson.M{
			"$ne": 0,
		},
	}
	if q != "" {
		query["description.moniker"] = &bson.M{
			"$regex": q,
			"$options": "$i",
		}
	}
	candidates, err := d.Query(query, skip, limit, sorts...)

	if err != nil {
		logger.Error.Println(err)
	}

	return candidates, err
}

func (d Candidate) GetCandidatesListByPubKeys(pubKeys []string) ([]Candidate, error)  {
	query := bson.M{
		"pub_key": &bson.M{
			"$in": pubKeys,
		},
	}
	sorts := make([]string, 0)

	candidates, err := d.Query(query, 0, len(pubKeys), sorts...)

	if err != nil {
		logger.Error.Println(err)
	}

	return candidates, err
}

func (d Candidate) GetTotalShares() (uint64, error)  {
	type result struct {
		Id string `bson:"_id"`
		TotalShares uint64 `bson:"total_shares"`
	}
	var value result

	q := func(c *mgo.Collection) error {
		m := []bson.M{
			{"$group": bson.M{"_id": "test", "total_shares": bson.M{"$sum": "$shares"}}},
		}
		return c.Pipe(m).One(&value)
	}

	err := models.ExecCollection(d.Name(), q)

	if err !=  nil {
		return 0, err
	}
	return value.TotalShares, nil
}

func (d Candidate) GetCandidateDetail(pubKey string) (Candidate, error) {
	query := bson.M{
		"pub_key": pubKey,
	}
	sorts := make([]string, 0)

	candidates, err := d.Query(query, 0, 1, sorts...)

	if err != nil {
		logger.Error.Println(err)
	}
	if len(candidates) > 0 {
		return candidates[0], err
	}
	return Candidate{}, err
}

