package controllers

import (
	"context"
	"encoding/json"
	"fiber-mongo-api/configs"
	"fiber-mongo-api/models"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Get_message(c *fiber.Ctx) error {

	cacheKey := "messages"

	cacheData, err := configs.GetFromCache(cacheKey)
	if err == nil {
		c.Send(cacheData)
		log.Println("get from cache !!")
		return nil
	}

	collection, err := configs.GetMongoDbCollection(configs.DbName, configs.CollectionName)
	if err != nil {
		c.Status(500).Send(nil)
		return err
	}

	var filter bson.M = bson.M{}

	if c.Params("id") != "" {
		id := c.Params("id")
		log.Print(id)
		objID, _ := primitive.ObjectIDFromHex(id)
		filter = bson.M{"_id": objID}
		log.Print(objID)
	}

	var results []bson.M
	cur, err := collection.Find(context.Background(), filter)
	defer cur.Close(context.Background())

	if err != nil {
		c.Status(500).Send(nil)
		return err
	}

	cur.All(context.Background(), &results)

	if results == nil {
		c.SendStatus(404)
		return err
	}

	jsonData, _ := json.Marshal(results)

	err = configs.SetToCache(cacheKey, jsonData, 5*time.Minute)
	log.Println("set to cache !!")
	if err != nil {
		log.Println("Failed to set data in cache:", err)
	}

	c.Send(jsonData)

	return nil
}

func Create_message(c *fiber.Ctx) error {
	collection, err := configs.GetMongoDbCollection(configs.DbName, configs.CollectionName)
	if err != nil {
		c.Status(500).Send(nil)
	}

	var message models.Message
	json.Unmarshal([]byte(c.Body()), &message)

	res, err := collection.InsertOne(context.Background(), message)
	if err != nil {
		c.Status(500).Send(nil)
		return err
	}

	response, _ := json.Marshal(res)
	c.Send(response)

	return nil
}

func Update_message(c *fiber.Ctx) error {
	collection, err := configs.GetMongoDbCollection(configs.DbName, configs.CollectionName)
	if err != nil {
		c.Status(500).Send(nil)
		return err
	}
	var message models.Message
	json.Unmarshal([]byte(c.Body()), &message)

	update := bson.M{
		"$set": message,
	}

	objID, _ := primitive.ObjectIDFromHex(c.Params("id"))
	res, err := collection.UpdateOne(context.Background(), bson.M{"_id": objID}, update)

	if err != nil {
		c.Status(500).Send(nil)
		return err
	}

	response, _ := json.Marshal(res)
	c.Send(response)

	return nil
}

func Delete_message(c *fiber.Ctx) error {
	collection, err := configs.GetMongoDbCollection(configs.DbName, configs.CollectionName)

	if err != nil {
		c.Status(500).Send(nil)
		return err
	}

	objID, _ := primitive.ObjectIDFromHex(c.Params("id"))
	res, err := collection.DeleteOne(context.Background(), bson.M{"_id": objID})

	if err != nil {
		c.Status(500).Send(nil)
		return err
	}

	jsonResponse, _ := json.Marshal(res)
	c.Send(jsonResponse)

	return nil
}
