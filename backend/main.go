package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDB client
var client *mongo.Client
var teamsCollection *mongo.Collection
var challengesCollection *mongo.Collection
var submissionsCollection *mongo.Collection

// Team represents a CTF team
type Team struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	TeamID   int                `json:"teamId" bson:"teamId"`
	Name     string             `json:"name" bson:"name"`
	Location string             `json:"location" bson:"location"`
	IP       string             `json:"ip" bson:"ip"`
	Lat      float64            `json:"lat" bson:"lat"`
	Lng      float64            `json:"lng" bson:"lng"`
	Color    string             `json:"color" bson:"color"`
	Members  int                `json:"members" bson:"members"`
	Score    int                `json:"score" bson:"score"`
	Solved   int                `json:"solved" bson:"solved"`
	Progress int                `json:"progress" bson:"progress"`
}

// Challenge represents a CTF challenge
type Challenge struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	ChallengeID int                `json:"challengeId" bson:"challengeId"`
	Name        string             `json:"name" bson:"name"`
	Category    string             `json:"category" bson:"category"`
	Points      int                `json:"points" bson:"points"`
	Description string             `json:"description" bson:"description"`
	Flag        string             `json:"flag" bson:"flag"`
}

// Submission represents a flag submission
type Submission struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	TeamID      int                `json:"teamId" bson:"teamId"`
	ChallengeID int                `json:"challengeId" bson:"challengeId"`
	Flag        string             `json:"flag" bson:"flag"`
	IsCorrect   bool               `json:"isCorrect" bson:"isCorrect"`
	Timestamp   time.Time          `json:"timestamp" bson:"timestamp"`
}

// NetworkActivity represents network activity data
type NetworkActivity struct {
	PacketsPerSec int       `json:"packetsPerSec"`
	ActiveNodes   int       `json:"activeNodes"`
	Timestamp     time.Time `json:"timestamp"`
}

// Connect to MongoDB
func connectMongoDB() (*mongo.Client, error) {
	mongoURI := os.Getenv("MONGODB_URI")
	if mongoURI == "" {
		mongoURI = "mongodb://localhost:27017"
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	// Ping to verify connection
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	log.Println("✅ Connected to MongoDB!")
	return client, nil
}

// Seed initial data if collections are empty
func seedData() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Check if teams collection is empty
	count, err := teamsCollection.CountDocuments(ctx, bson.M{})
	if err != nil {
		log.Printf("Error counting teams: %v", err)
		return
	}

	if count == 0 {
		log.Println("📦 Seeding initial teams data...")
		teams := []interface{}{
			Team{TeamID: 1, Name: "Phoenix Hackers", Location: "Tokyo, Japan", IP: "192.168.1.101", Lat: 35.6762, Lng: 139.6503, Color: "red", Members: 5, Score: 0, Solved: 0, Progress: 88},
			Team{TeamID: 2, Name: "Cyber Ninjas", Location: "San Francisco, USA", IP: "192.168.1.102", Lat: 37.7749, Lng: -122.4194, Color: "orange", Members: 5, Score: 0, Solved: 0, Progress: 59},
			Team{TeamID: 3, Name: "Binary Wolves", Location: "Berlin, Germany", IP: "192.168.1.103", Lat: 52.5200, Lng: 13.4050, Color: "green", Members: 5, Score: 0, Solved: 0, Progress: 88},
			Team{TeamID: 4, Name: "Shadow Raiders", Location: "Singapore", IP: "192.168.1.104", Lat: 1.3521, Lng: 103.8198, Color: "yellow", Members: 5, Score: 0, Solved: 0, Progress: 59},
			Team{TeamID: 5, Name: "Quantum Knights", Location: "Sydney, Australia", IP: "192.168.1.105", Lat: -33.8688, Lng: 151.2093, Color: "purple", Members: 5, Score: 0, Solved: 0, Progress: 69},
		}
		_, err := teamsCollection.InsertMany(ctx, teams)
		if err != nil {
			log.Printf("Error seeding teams: %v", err)
		} else {
			log.Println("✅ Teams seeded successfully!")
		}
	}

	// Check if challenges collection is empty
	count, err = challengesCollection.CountDocuments(ctx, bson.M{})
	if err != nil {
		log.Printf("Error counting challenges: %v", err)
		return
	}

	if count == 0 {
		log.Println("📦 Seeding initial challenges data...")
		challenges := []interface{}{
			Challenge{ChallengeID: 1, Name: "Web Challenge 1", Category: "Web", Points: 100, Description: "Find the hidden flag", Flag: "FLAG{w3b_ch4ll3ng3_1}"},
			Challenge{ChallengeID: 2, Name: "Crypto Challenge 1", Category: "Crypto", Points: 200, Description: "Decrypt the message", Flag: "FLAG{crypt0_m4st3r}"},
			Challenge{ChallengeID: 3, Name: "Reverse Challenge 1", Category: "Reverse", Points: 300, Description: "Reverse engineer the binary", Flag: "FLAG{r3v_3ng1n33r}"},
			Challenge{ChallengeID: 4, Name: "Forensics Challenge 1", Category: "Forensics", Points: 150, Description: "Analyze the network packet", Flag: "FLAG{p4ck3t_hunt3r}"},
			Challenge{ChallengeID: 5, Name: "PWN Challenge 1", Category: "PWN", Points: 250, Description: "Exploit the buffer overflow", Flag: "FLAG{buff3r_0v3rfl0w}"},
		}
		_, err := challengesCollection.InsertMany(ctx, challenges)
		if err != nil {
			log.Printf("Error seeding challenges: %v", err)
		} else {
			log.Println("✅ Challenges seeded successfully!")
		}
	}
}

// CORS Middleware
func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// Get all teams
func getTeams(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := teamsCollection.Find(ctx, bson.M{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	defer cursor.Close(ctx)

	var teams []Team
	if err := cursor.All(ctx, &teams); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	json.NewEncoder(w).Encode(teams)
}

// Get team by ID
func getTeam(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	teamID, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid team ID"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var team Team
	err = teamsCollection.FindOne(ctx, bson.M{"teamId": teamID}).Decode(&team)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{"error": "Team not found"})
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	json.NewEncoder(w).Encode(team)
}

// Create new team
func createTeam(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var team Team
	if err := json.NewDecoder(r.Body).Decode(&team); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := teamsCollection.InsertOne(ctx, team)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	team.ID = result.InsertedID.(primitive.ObjectID)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(team)
}

// Update team
func updateTeam(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	teamID, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid team ID"})
		return
	}

	var updatedTeam Team
	if err := json.NewDecoder(r.Body).Decode(&updatedTeam); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	update := bson.M{
		"$set": bson.M{
			"score":    updatedTeam.Score,
			"solved":   updatedTeam.Solved,
			"progress": updatedTeam.Progress,
		},
	}

	result, err := teamsCollection.UpdateOne(ctx, bson.M{"teamId": teamID}, update)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	if result.MatchedCount == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Team not found"})
		return
	}

	// Fetch updated team
	var team Team
	teamsCollection.FindOne(ctx, bson.M{"teamId": teamID}).Decode(&team)
	json.NewEncoder(w).Encode(team)
}

// Delete team
func deleteTeam(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	teamID, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid team ID"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := teamsCollection.DeleteOne(ctx, bson.M{"teamId": teamID})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	if result.DeletedCount == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Team not found"})
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Team deleted successfully"})
}

// Get all challenges
func getChallenges(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := challengesCollection.Find(ctx, bson.M{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	defer cursor.Close(ctx)

	var challenges []Challenge
	if err := cursor.All(ctx, &challenges); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	json.NewEncoder(w).Encode(challenges)
}

// Submit flag
func submitFlag(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var submission Submission
	if err := json.NewDecoder(r.Body).Decode(&submission); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Find challenge
	var challenge Challenge
	err := challengesCollection.FindOne(ctx, bson.M{"challengeId": submission.ChallengeID}).Decode(&challenge)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{"error": "Challenge not found"})
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	// Validate flag
	submission.Timestamp = time.Now()
	submission.IsCorrect = submission.Flag == challenge.Flag

	// Save submission
	_, err = submissionsCollection.InsertOne(ctx, submission)
	if err != nil {
		log.Printf("Error saving submission: %v", err)
	}

	if submission.IsCorrect {
		// Update team score
		update := bson.M{
			"$inc": bson.M{
				"score":  challenge.Points,
				"solved": 1,
			},
		}

		result, err := teamsCollection.UpdateOne(ctx, bson.M{"teamId": submission.TeamID}, update)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
			return
		}

		if result.MatchedCount == 0 {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{"error": "Team not found"})
			return
		}

		// Fetch updated team
		var team Team
		teamsCollection.FindOne(ctx, bson.M{"teamId": submission.TeamID}).Decode(&team)

		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": true,
			"message": "Flag accepted! 🎉",
			"points":  challenge.Points,
			"team":    team,
		})
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": false,
		"message": "Wrong flag! Try again.",
	})
}

// Get submissions history
func getSubmissions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Get query parameters
	teamIDStr := r.URL.Query().Get("teamId")
	filter := bson.M{}

	if teamIDStr != "" {
		teamID, err := strconv.Atoi(teamIDStr)
		if err == nil {
			filter["teamId"] = teamID
		}
	}

	// Sort by timestamp descending
	opts := options.Find().SetSort(bson.D{{Key: "timestamp", Value: -1}})

	cursor, err := submissionsCollection.Find(ctx, filter, opts)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	defer cursor.Close(ctx)

	var submissions []Submission
	if err := cursor.All(ctx, &submissions); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	json.NewEncoder(w).Encode(submissions)
}

// Get network activity
func getNetworkActivity(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Count active teams
	activeTeams, _ := teamsCollection.CountDocuments(ctx, bson.M{})

	activity := NetworkActivity{
		PacketsPerSec: 100 + (int(time.Now().Unix()) % 50), // Simulate variation
		ActiveNodes:   int(activeTeams),
		Timestamp:     time.Now(),
	}

	json.NewEncoder(w).Encode(activity)
}

func main() {
	log.Println("🚀 Starting CTF Backend API with MongoDB...")

	// Connect to MongoDB
	var err error
	client, err = connectMongoDB()
	if err != nil {
		log.Fatalf("❌ Failed to connect to MongoDB: %v", err)
	}
	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			log.Printf("Error disconnecting from MongoDB: %v", err)
		}
	}()

	// Get database name from env or use default
	dbName := os.Getenv("MONGODB_DATABASE")
	if dbName == "" {
		dbName = "ctf_gis"
	}

	// Initialize collections
	db := client.Database(dbName)
	teamsCollection = db.Collection("teams")
	challengesCollection = db.Collection("challenges")
	submissionsCollection = db.Collection("submissions")

	// Seed initial data
	seedData()

	// Setup router
	router := mux.NewRouter()

	// API routes
	router.HandleFunc("/api/teams", getTeams).Methods("GET")
	router.HandleFunc("/api/teams", createTeam).Methods("POST")
	router.HandleFunc("/api/teams/{id}", getTeam).Methods("GET")
	router.HandleFunc("/api/teams/{id}", updateTeam).Methods("PUT")
	router.HandleFunc("/api/teams/{id}", deleteTeam).Methods("DELETE")

	router.HandleFunc("/api/challenges", getChallenges).Methods("GET")

	router.HandleFunc("/api/submit", submitFlag).Methods("POST")
	router.HandleFunc("/api/submissions", getSubmissions).Methods("GET")

	router.HandleFunc("/api/activity", getNetworkActivity).Methods("GET")

	// Health check
	router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  "healthy",
			"mongodb": "connected",
			"time":    time.Now(),
		})
	}).Methods("GET")

	// Enable CORS
	handler := enableCORS(router)

	// Get port from env or use default
	port := os.Getenv("API_PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("🚀 CTF Backend API running on http://localhost:%s", port)
	log.Printf("📊 Database: %s", dbName)
	log.Println("📡 API Endpoints:")
	log.Println("   GET    /api/teams")
	log.Println("   POST   /api/teams")
	log.Println("   GET    /api/teams/{id}")
	log.Println("   PUT    /api/teams/{id}")
	log.Println("   DELETE /api/teams/{id}")
	log.Println("   GET    /api/challenges")
	log.Println("   POST   /api/submit")
	log.Println("   GET    /api/submissions")
	log.Println("   GET    /api/activity")
	log.Println("   GET    /api/health")

	if err := http.ListenAndServe(":"+port, handler); err != nil {
		log.Fatalf("❌ Server failed: %v", err)
	}
}
