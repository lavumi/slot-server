package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
	"net/http"
	"slot-server/internal/db"
	"slot-server/internal/server/forms"
	"slot-server/internal/server/models"
	"slot-server/internal/slot"
)

type Game struct {
	Slot *slot.Client
	Db   *db.MongoDb
}

func (g *Game) Spin(c *gin.Context) {

	var req forms.SpinRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		//c.JSON(400, gin.H{"msg": err})
		fmt.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		//SendError(c, 400, err.Error())
		return
	}

	//1. Get User Data From Session
	slotId := req.Id
	user := c.MustGet("user").(models.User)
	cash := user.Cash

	//log.Printf("UserID : %s | Cash : %f", user.UUID, user.Cash)

	//2. Get User Data From Db
	var save = models.SavedFeature{}
	filter := bson.D{{"uuid", user.UUID}}
	if err := g.Db.GetCollection(fmt.Sprintf("save_%d", slotId)).FindOne(context.TODO(), filter).Decode(&save); err != nil {
		if err != mongo.ErrNoDocuments {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		} else {
			save.UUID = user.UUID
		}
	}
	//log.Printf("Saved Feature : %v ", save.SaveData == nil)

	//3. Spin
	spin, additionalInfo, err := g.Slot.RequestSpin(slotId, req.BetCash, save.SaveData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	spinObject := make(map[string]interface{})
	err = json.Unmarshal(spin, &spinObject)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	//4. handle response
	user.Cash = cash + additionalInfo.CashDiff
	c.Set("user", user)

	update := bson.D{{"$set", bson.D{{"save", additionalInfo.SavedFeature}, {"c", additionalInfo.Collectable}}}}
	if _, err := g.Db.GetCollection(fmt.Sprintf("save_%d", slotId)).UpdateOne(context.TODO(), filter, update); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		//todo rollback session
		return
	}

	c.JSON(http.StatusOK, forms.SpinResponse{
		SpinResult: spinObject,
		After:      user.Cash,
		Before:     cash,
	})

}
