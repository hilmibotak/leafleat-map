# MongoDB Setup Guide

## Prasyarat

1. **Install MongoDB** di komputer Anda:
   - **Windows**: Download MongoDB Community Server dari [mongodb.com](https://www.mongodb.com/try/download/community)
   - **Mac**: `brew install mongodb-community`
   - **Linux**: `sudo apt-get install mongodb`

2. **Install MongoDB Compass** (Optional - GUI Tool):
   - Download dari [mongodb.com/products/compass](https://www.mongodb.com/products/compass)

## Menjalankan MongoDB

### Windows

```powershell
# Start MongoDB service
net start MongoDB

# Atau jika install manual, jalankan:
mongod --dbpath "C:\data\db"
```

### Mac/Linux

```bash
# Start MongoDB service
brew services start mongodb-community

# Atau
sudo systemctl start mongod
```

## Konfigurasi Environment

File `.env` sudah tersedia dengan konfigurasi default:

```env
# MongoDB Database
MONGODB_URI=mongodb://localhost:27017
MONGODB_DATABASE=ctf_gis
MONGODB_TIMEOUT=10
```

## Struktur Database

Database `ctf_gis` memiliki 3 collections:

### 1. Teams Collection

```json
{
  "_id": ObjectId,
  "teamId": 1,
  "name": "Phoenix Hackers",
  "location": "Tokyo, Japan",
  "ip": "192.168.1.101",
  "lat": 35.6762,
  "lng": 139.6503,
  "color": "red",
  "members": 5,
  "score": 0,
  "solved": 0,
  "progress": 88
}
```

### 2. Challenges Collection

```json
{
  "_id": ObjectId,
  "challengeId": 1,
  "name": "Web Challenge 1",
  "category": "Web",
  "points": 100,
  "description": "Find the hidden flag",
  "flag": "FLAG{w3b_ch4ll3ng3_1}"
}
```

### 3. Submissions Collection

```json
{
  "_id": ObjectId,
  "teamId": 1,
  "challengeId": 1,
  "flag": "FLAG{w3b_ch4ll3ng3_1}",
  "isCorrect": true,
  "timestamp": ISODate("2026-02-07T...")
}
```

## Menjalankan Backend

1. Pastikan MongoDB sudah berjalan
2. Dari root directory, jalankan:

```powershell
cd backend
go run main.go
```

Backend akan:

- Connect ke MongoDB di `localhost:27017`
- Membuat database `ctf_gis` jika belum ada
- Seed data awal (teams & challenges) jika collection kosong
- Start server di `http://localhost:8080`

## API Endpoints

### Teams

- `GET /api/teams` - Get all teams
- `POST /api/teams` - Create new team
- `GET /api/teams/{id}` - Get team by ID
- `PUT /api/teams/{id}` - Update team
- `DELETE /api/teams/{id}` - Delete team

### Challenges

- `GET /api/challenges` - Get all challenges

### Submissions

- `POST /api/submit` - Submit a flag
- `GET /api/submissions?teamId=1` - Get submissions (optional filter by teamId)

### Activity

- `GET /api/activity` - Get network activity stats

### Health Check

- `GET /api/health` - Check server and MongoDB status

## Menggunakan MongoDB Compass

1. Buka MongoDB Compass
2. Connect ke: `mongodb://localhost:27017`
3. Pilih database `ctf_gis`
4. Explore collections: `teams`, `challenges`, `submissions`

## Contoh CRUD Operations

### Create Team

```bash
curl -X POST http://localhost:8080/api/teams \
  -H "Content-Type: application/json" \
  -d '{
    "teamId": 6,
    "name": "New Team",
    "location": "Jakarta, Indonesia",
    "ip": "192.168.1.106",
    "lat": -6.2088,
    "lng": 106.8456,
    "color": "blue",
    "members": 5,
    "score": 0,
    "solved": 0,
    "progress": 0
  }'
```

### Update Team Score

```bash
curl -X PUT http://localhost:8080/api/teams/1 \
  -H "Content-Type: application/json" \
  -d '{
    "score": 350,
    "solved": 3,
    "progress": 75
  }'
```

### Submit Flag

```bash
curl -X POST http://localhost:8080/api/submit \
  -H "Content-Type: application/json" \
  -d '{
    "teamId": 1,
    "challengeId": 1,
    "flag": "FLAG{w3b_ch4ll3ng3_1}"
  }'
```

## Troubleshooting

### MongoDB Connection Error

```
❌ Failed to connect to MongoDB: connection refused
```

**Solusi**: Pastikan MongoDB service sudah berjalan

### Port Already in Use

```
❌ Server failed: listen tcp :8080: bind: address already in use
```

**Solusi**:

1. Kill process yang menggunakan port 8080, atau
2. Ubah `API_PORT` di file `.env`

### Database Permission Error

**Solusi**: Pastikan user memiliki read/write permission ke database

## Reset Database

Jika ingin reset database ke kondisi awal:

```javascript
// Melalui MongoDB Compass atau mongo shell:
use ctf_gis
db.teams.drop()
db.challenges.drop()
db.submissions.drop()

// Restart backend untuk auto-seed data baru
```

## Production Deployment

Untuk production, update `.env`:

```env
MONGODB_URI=mongodb://username:password@your-mongodb-server:27017/ctf_gis?authSource=admin
MONGODB_DATABASE=ctf_gis
API_PORT=8080
```

Gunakan MongoDB Atlas untuk managed database di cloud:

- [mongodb.com/cloud/atlas](https://www.mongodb.com/cloud/atlas)
