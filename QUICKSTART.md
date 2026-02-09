## 🚀 QUICK START GUIDE - CTF CYBER MAP

### Cara Paling Mudah untuk Menjalankan Aplikasi

#### Opsi 1: Quick Start (Recommended)

1. Buka folder project di terminal/command prompt
2. Double-click file `quick-start.bat`
3. Browser akan otomatis terbuka ke `http://localhost:8000`

#### Opsi 2: Manual Start

**Terminal 1 - Backend:**

```powershell
cd backend
go mod tidy
go run main.go
```

**Terminal 2 - Frontend:**

```powershell
# Gunakan salah satu:

# Python (paling mudah)
python -m http.server 8000

# Atau Node.js http-server
npx http-server -p 8000

# Atau PHP
php -S localhost:8000
```

**Akses di browser:**

- Frontend: http://localhost:8000
- Backend API: http://localhost:8080

#### Opsi 3: VS Code Live Server

1. Install extension "Live Server" di VS Code
2. Right-click pada `index.html`
3. Pilih "Open with Live Server"
4. Di terminal terpisah, jalankan backend:
   ```powershell
   cd backend
   go run main.go
   ```

### ⚠️ Requirements

- **Go 1.21+** (untuk backend)
- **Python 3.x** atau **Node.js** (untuk frontend server)
- **Browser modern** (Chrome, Firefox, Edge)

### 🔧 Troubleshooting

**Port sudah digunakan:**

```powershell
# Ganti port frontend
python -m http.server 8001

# Atau ganti port backend di main.go
port := ":8081"
```

**Backend gagal start:**

```powershell
# Install dependencies
cd backend
go mod tidy
go get github.com/gorilla/mux
```

**Map tidak muncul:**

- Pastikan koneksi internet aktif
- Check browser console (F12) untuk error
- Refresh page (Ctrl+R)

### 📱 Test di Browser

Buka: http://localhost:8000

Anda akan melihat:

- ✅ Cyber Map dengan 5 team markers
- ✅ Progress bars untuk setiap team
- ✅ Real-time clock
- ✅ Network activity visualization
- ✅ Team cards dengan detail lokasi

### 🎯 Next Steps

1. Test API endpoints: `http://localhost:8080/api/teams`
2. Submit flag via API
3. Customize teams di `js/config.js`
4. Deploy ke production server

### 📞 Support

Untuk bantuan lebih lanjut, lihat file `README.md` atau hubungi tim developer.

---

**Happy Hacking! 🎉**
