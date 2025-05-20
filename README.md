# Golang Gin Project

Project ini adalah contoh aplikasi backend REST API menggunakan [Gin](https://github.com/gin-gonic/gin), sebuah web framework ringan dan cepat untuk Golang.

---

## Fitur

- Routing HTTP dengan Gin
- Uplad data from excel
- CRUD sederhana dengan GORM 
- Validasi input request menggunakan [go-playground/validator](https://github.com/go-playground/validator/v10)
- Struktur project repositories pattern yang scalable
- Logging request dan error handling

---

## Prasyarat

- Go versi 1.18 atau lebih tinggi
- Database MySQL / PostgreSQL / SQLite (sesuaikan konfigurasi)
- Git (untuk cloning repo)

---

## Instalasi

1. Clone repository ini

   ```bash
   git clone  git@github.com:afryn123/golang-service-upload.git
   cd golang-service-upload
   ```
2. Intall package/library yang ada di dalam project
    ```bash
    go mod tidy
    ```
3. Run aplikasi di root project(Devlopment)
    ```bash
    go run main.go
    ```
4. (opsional/production) Buld
    ```bash
    go build
    ```