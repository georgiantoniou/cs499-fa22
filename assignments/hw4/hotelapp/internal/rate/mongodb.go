//go:build mongodb

package rate

import (
	log "github.com/sirupsen/logrus"

	pb "github.com/ucy-coast/hotel-app/internal/rate/proto"

	"github.com/ucy-coast/hotel-app/pkg/util"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type DatabaseSession struct {
	MongoSession *mgo.Session
}

func NewDatabaseSession(db_addr string) *DatabaseSession {
	// TODO: Implement me

	session, err := mgo.Dial(db_addr)
	if err != nil {
		log.Fatal(err)
	}
	log.Info("New session successfull...")

	return &DatabaseSession{
		MongoSession: session,
	}
}

func (db *DatabaseSession) LoadDataFromJsonFile(rateJsonPath string) {
	util.LoadDataFromJsonFile(db.MongoSession, "rate-db", "inventory", rateJsonPath)
}

// GetRates gets rates for hotels for specific date range.
func (db *DatabaseSession) GetRates(hotelIds []string) (RatePlans, error) {
	// TODO: Implement me
	session := db.MongoSession.Copy()
	defer session.Close()
	c := session.DB("rate-db").C("inventory")

	RatePlans := make([]*pb.RatePlan, 0)

	for _, id := range hotelIds {
		rate_plan := new(pb.RatePlan)
		err := c.Find(bson.M{"hotelId": id}).One(&rate_plan)
		if err != nil {
                        continue
			log.Fatalf("Failed get hotels data: ", err)
		}
		RatePlans = append(RatePlans, rate_plan)
	}
	return RatePlans, nil

}
