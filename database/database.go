package database

import (
	"context"
	"log"
	"time"

	"github.com/prasaddls/go-graphql-mongodb-curd/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var connectionString = "mongodb+srv://prasad2023aws:Fu3nqfigSXwJYyCa@cluster0.yscl1ix.mongodb.net/"

type DB struct {
	client *mongo.Client
}

func Connect() *DB {
	client, err := mongo.NewClient(options.Client().ApplyURI(connectionString))

	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, readpref.Primary())

	if err != nil {
		log.Fatal(err)
	}

	return &DB{
		client: client,
	}

}

func (db *DB) GetJob(id string) *model.JobListing {

	jobCollect := db.client.Database("graphql-job-board").Collection("job")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_id, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": _id}
	var jobListing model.JobListing
	err := jobCollect.FindOne(ctx, filter).Decode(&jobListing)
	if err != nil {
		log.Fatal(err)
	}
	return &jobListing
}

func (db *DB) GetJobs() []*model.JobListing {
	jobCollect := db.client.Database("graphql-job-board").Collection("job")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	var jobListing []*model.JobListing
	cursor, err := jobCollect.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	if err = cursor.All(context.TODO(), &jobListing); err != nil {
		panic(err)
	}
	return jobListing
}

func (db *DB) CreateJobListing(jobInfo model.CreateJobListingInput) *model.JobListing {
	jobCollec := db.client.Database("graphql-job-board").Collection("jobs")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	inserg, err := jobCollec.InsertOne(ctx, bson.M{"title": jobInfo.Title, "description": jobInfo.Description, "url": jobInfo.URL, "company": jobInfo.Company})

	if err != nil {
		log.Fatal(err)
	}

	insertedID := inserg.InsertedID.(primitive.ObjectID).Hex()
	returnJobListing := model.JobListing{ID: insertedID, Title: jobInfo.Title, Company: jobInfo.Company, Description: jobInfo.Description, URL: jobInfo.URL}
	return &returnJobListing
}

func (db *DB) UpdateJobListing(jobId string, jobInfo model.UpdateJobListingInput) *model.JobListing {
	jobCollec := db.client.Database("graphql-job-board").Collection("jobs")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	updateJobInfo := bson.M{}

	if jobInfo.Title != nil {
		updateJobInfo["title"] = jobInfo.Title
	}
	if jobInfo.Description != nil {
		updateJobInfo["description"] = jobInfo.Description
	}
	if jobInfo.URL != nil {
		updateJobInfo["url"] = jobInfo.URL
	}

	_id, _ := primitive.ObjectIDFromHex(jobId)
	filter := bson.M{"_id": _id}
	update := bson.M{"$set": updateJobInfo}

	results := jobCollec.FindOneAndUpdate(ctx, filter, update, options.FindOneAndUpdate().SetReturnDocument(1))

	var jobListing model.JobListing

	if err := results.Decode(&jobListing); err != nil {
		log.Fatal(err)
	}

	return &jobListing
}

func (db *DB) DeleteJobListing(jobId string) *model.DeleteJobResponse {
	jobCollec := db.client.Database("graphql-job-board").Collection("jobs")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_id, _ := primitive.ObjectIDFromHex(jobId)
	filter := bson.M{"_id": _id}
	_, err := jobCollec.DeleteOne(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	return &model.DeleteJobResponse{DeletedJobID: jobId}
}
