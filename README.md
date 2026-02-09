# CTF GIS - Cyber Map

## 🌐 Live Hacking CTF Geographic Information System

Aplikasi CTF (Capture The Flag) dengan visualisasi peta geografis real-time yang menampilkan aktivitas tim, skor, dan network activity map.

## ✨ Fitur

- 🗺️ **Interactive Global Map** - Visualisasi lokasi tim di seluruh dunia
- 📊 **Real-time Progress Tracking** - Monitor progress setiap tim secara real-time
- 🏆 **Live Scoreboard** - Papan skor yang update otomatis
- 🌐 **Network Activity Visualization** - Visualisasi aktivitas jaringan antar tim
- 🎯 **Team Management** - Managemen tim dan submission flag
- ⚡ **Live Updates** - Update data secara real-time

## 🛠️ Teknologi

### Frontend

- **Vanilla JS ES6+ Module** dengan framework [jscroot](https://croot.js.org)
- **Leaflet.js** untuk interactive maps
- **CSS3** dengan theme cyber/dark mode
- **Server**: [PingBox](https://github.com/noz0/pingbox)

### Backend

- **Golang** dengan [Lemes Framework](https://github.com/noz0/lemes)
- **RESTful API**
- **In-memory data store**

### DNS

- **DNS Server**: [Alodek](https://github.com/n0z0/alodek)

## 📁 Struktur Proyek

```
UAS/
├── index.html              # Main HTML file
├── css/
│   └── style.css          # Styling dengan tema cyber
├── js/
│   ├── app.js             # Main application logic
│   ├── config.js          # Configuration & teams data
│   ├── map.js             # Map visualization module
│   ├── teams.js           # Team management module
│   └── api.js             # API integration module
├── backend/
│   ├── main.go            # Golang backend API
│   └── go.mod             # Go dependencies
└── README.md              # Documentation
```

## 🚀 Instalasi & Setup

### 1. Setup Backend (Lemes)

```powershell
# Masuk ke direktori backend
cd backend

# Install dependencies
go mod tidy

# Jalankan backend server
go run main.go
```

Backend akan berjalan di: `http://localhost:8080`

### 2. Setup Frontend (PingBox)

#### Install PingBox

```powershell
# Clone PingBox repository
git clone https://github.com/noz0/pingbox.git
cd pingbox

# Build PingBox (sesuaikan dengan instruksi repo)
go build -o pingbox.exe
```

#### Jalankan PingBox Server

```powershell
# Kembali ke direktori utama project
cd "d:\SEMESTER 5\SISTEM INFORMASI GEOGRAFIS\UAS"

# Jalankan PingBox untuk serve frontend
pingbox serve -p 8000 -d .
```

Frontend akan berjalan di: `http://localhost:8000`

### 3. Setup DNS Server (Alodek)

#### Install Alodek

```powershell
# Clone Alodek repository
git clone https://github.com/n0z0/alodek.git
cd alodek

# Build Alodek
go build -o alodek.exe
```

#### Konfigurasi DNS

Buat file `alodek.conf`:

```
# Alodek DNS Configuration
domain cybermap.local
server 127.0.0.1
port 53

# DNS Records
A cybermap.local 127.0.0.1
A api.cybermap.local 127.0.0.1
```

#### Jalankan Alodek

```powershell
# Jalankan DNS server (butuh admin privileges)
.\alodek.exe -config alodek.conf
```

### 4. Akses Aplikasi

Setelah semua service berjalan:

1. **Via localhost**: `http://localhost:8000`
2. **Via DNS**: `http://cybermap.local:8000` (jika DNS sudah dikonfigurasi)

## 📡 API Endpoints

### Teams

- `GET /api/teams` - Get all teams
- `GET /api/teams/:id` - Get team by ID
- `PUT /api/teams/:id` - Update team data

### Challenges

- `GET /api/challenges` - Get all challenges

### Submissions

- `POST /api/submit` - Submit flag
  ```json
  {
    "teamId": 1,
    "challengeId": 1,
    "flag": "FLAG{example}"
  }
  ```

### Network Activity

- `GET /api/activity` - Get network activity data

## 🎮 Cara Penggunaan

### 1. Monitoring Tim

- Lihat progress setiap tim di progress bars bagian atas
- Klik marker di peta untuk melihat detail lokasi tim
- Monitor skor real-time di team cards

### 2. Submit Flag

Gunakan API endpoint untuk submit flag:

```javascript
const response = await fetch("http://localhost:8080/api/submit", {
  method: "POST",
  headers: { "Content-Type": "application/json" },
  body: JSON.stringify({
    teamId: 1,
    challengeId: 1,
    flag: "FLAG{example}",
  }),
});
```

### 3. Update Team Data

```javascript
const response = await fetch("http://localhost:8080/api/teams/1", {
  method: "PUT",
  headers: { "Content-Type": "application/json" },
  body: JSON.stringify({
    score: 300,
    solved: 3,
    progress: 95,
  }),
});
```

## 🎨 Kustomisasi

### Menambah Tim Baru

Edit file `js/config.js`:

```javascript
export const teams = [
  // ... existing teams
  {
    id: 6,
    name: "Team Baru",
    location: "Jakarta, Indonesia",
    ip: "192.168.1.106",
    lat: -6.2088,
    lng: 106.8456,
    color: "blue", // red, orange, green, yellow, purple, blue
    members: 5,
    score: 0,
    solved: 0,
    progress: 0,
  },
];
```

### Menambah Challenge Baru

Edit file `backend/main.go`:

```go
var challenges = []Challenge{
    // ... existing challenges
    {
        ID: 4,
        Name: "New Challenge",
        Category: "Web",
        Points: 150,
        Description: "Challenge description",
        Flag: "FLAG{new_challenge}",
    },
}
```

### Mengubah Tema Warna

Edit file `css/style.css` untuk mengubah color scheme.

## 🔧 Troubleshooting

### Backend tidak bisa diakses

- Pastikan port 8080 tidak digunakan aplikasi lain
- Cek firewall settings
- Jalankan dengan: `go run main.go`

### Frontend tidak tampil

- Pastikan PingBox sudah terinstall
- Cek apakah port 8000 sudah terpakai
- Gunakan port lain: `pingbox serve -p 8001 -d .`

### DNS tidak resolve

- Jalankan Alodek dengan administrator privileges
- Edit file `C:\Windows\System32\drivers\etc\hosts` (Windows) dan tambahkan:
  ```
  127.0.0.1  cybermap.local
  127.0.0.1  api.cybermap.local
  ```

### Map tidak muncul

- Pastikan koneksi internet aktif (untuk load Leaflet tiles)
- Check browser console untuk error
- Coba refresh halaman

## 📝 Notes

- Aplikasi ini menggunakan in-memory storage, data akan hilang ketika server restart
- Untuk production, gunakan database seperti PostgreSQL atau MongoDB
- Sesuaikan CORS settings di backend jika deploy ke domain berbeda

## 🎯 Development

### Menjalankan dalam Development Mode

```powershell
# Terminal 1 - Backend
cd backend
go run main.go

# Terminal 2 - Frontend
cd ..
pingbox serve -p 8000 -d .
```

### Build untuk Production

```powershell
# Build backend
cd backend
go build -o ctf-backend.exe main.go

# Backend siap di deploy
.\ctf-backend.exe
```

## 📄 License

MIT License - Feel free to use for educational purposes

## 👥 Teams Data

Default teams yang tersedia:

1. **Team Glend** - Tokyo, Japan (192.168.1.101)
2. **Team Bona** - San Francisco, USA (192.168.1.102)
3. **Team Arief** - Berlin, Germany (192.168.1.103)
4. **Team Irfan** - Singapore (192.168.1.104)
5. **Team Deni** - Sydney, Australia (192.168.1.105)

## 🌟 Features Highlight

- ✅ Real-time clock
- ✅ Live packet counter
- ✅ Interactive world map dengan Leaflet
- ✅ Animated network connections
- ✅ Team progress bars
- ✅ Responsive team cards
- ✅ CTF scoring system
- ✅ RESTful API backend

---

**Dibuat untuk UAS Sistem Informasi Geografis**
Semester 5 - 2025
