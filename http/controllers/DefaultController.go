package controllers

import (
	"App/database"
	"fmt"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

type Controller struct {
}

func (Controller) Collection(collectionName string) *mongo.Collection {
	return database.Collection(collectionName)
}

func (Controller) Log(c *gin.Context) *log.Entry {
	contextualizedLogger := log.WithFields(
		log.Fields{
			"route": c.Request.RequestURI,
		})
	return contextualizedLogger
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func PatchValid(patch map[string]interface{}, schema map[string][]string) bool {
	for patchKey, patchValues := range patch {
		schemaAllowedValues, ok := schema[patchKey]
		if !ok { return false }

		fmt.Println(patchValues, schemaAllowedValues)
		for patchSectionKey, _ := range patchValues.(bson.M) {
			if !stringInSlice(patchSectionKey, schemaAllowedValues) { return false }
		}
	}

	return true
}
