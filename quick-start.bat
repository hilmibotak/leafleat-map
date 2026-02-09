@echo off
echo ========================================
echo   CTF CYBER MAP - QUICK START
echo   (Menggunakan Python HTTP Server)
echo ========================================
echo.

REM Check if Python is installed
where python >nul 2>nul
if %errorlevel% neq 0 (
    echo [ERROR] Python tidak terinstall!
    echo Silakan install Python dari https://www.python.org/downloads/
    pause
    exit /b 1
)

echo [1/2] Starting Backend Server...
cd backend
start "CTF Backend API" cmd /k "go run main.go"
timeout /t 3 >nul
cd ..

echo.
echo [2/2] Starting Frontend Server...
start "CTF Frontend" cmd /k "python -m http.server 8000"
timeout /t 3 >nul

echo.
echo ========================================
echo   Setup Complete!
echo ========================================
echo.
echo   Backend API: http://localhost:8080
echo   Frontend:    http://localhost:8000
echo.
echo ========================================
echo.

REM Wait before opening browser
timeout /t 2 >nul

REM Open browser
start "" "http://localhost:8000"

echo.
echo Aplikasi sudah berjalan!
echo Tekan Ctrl+C di window server untuk stop.
echo Atau jalankan stop.bat untuk stop semua server.
echo.
pause
