# PDDA Backend REST API

**PDDA Backend** adalah backend REST API untuk proyek *Plant Diseases Dataset Archive* yang ditulis menggunakan bahasa pemrograman **Go**. API ini bertujuan untuk menyediakan antarmuka yang cepat, aman, dan handal dalam pengelolaan dataset penyakit tanaman.

## Fitur Utama

- **Performa Tinggi**: Dibangun menggunakan Go untuk memastikan API berjalan dengan cepat dan efisien.
- **RESTful API**: Endpoint lengkap untuk kebutuhan CRUD dataset.
- **Keamanan Modern**: Dilengkapi dengan mekanisme autentikasi berbasis token.
- **Arsitektur Sederhana dan Scalable**: Mudah diadaptasi untuk pengembangan lebih lanjut.

## Teknologi yang Digunakan

- **Go**: Bahasa pemrograman utama untuk membangun backend.
- **GORM**: ORM untuk integrasi database.
- **Swagger**: Dokumentasi API otomatis.

## Instalasi dan Menjalankan Proyek

### Prasyarat

- Pastikan Anda telah menginstal **Go** (versi terbaru).
- Database seperti **PostgreSQL** atau **MySQL** untuk penyimpanan data.

### Langkah Instalasi

1. Clone repositori ini:
   ```bash
   git clone https://github.com/username/pdda-backend.git
   cd pdda-backend
   ```
2. Install dependensi:
   ```bash
   go mod tidy
   ```
3. Jalankan migrasi database (jika menggunakan PostgreSQL):
   ```bash
   go run cmd/migrate.go
   ```
4. Jalankan server:
   ```bash
   go run main.go
   ```
5. API akan berjalan di `http://localhost:8080` secara default.

## Endpoint API

Berikut adalah daftar endpoint utama yang disediakan oleh PDDA Backend:

| Method | Endpoint          | Deskripsi                     |
|--------|-------------------|-------------------------------|
| GET    | /datasets         | Mendapatkan daftar dataset    |
| POST   | /datasets         | Mengunggah dataset baru       |
| GET    | /datasets/{id}    | Mendapatkan dataset spesifik  |
| PUT    | /datasets/{id}    | Memperbarui dataset           |
| DELETE | /datasets/{id}    | Menghapus dataset             |

Untuk dokumentasi lengkap, lihat [API Documentation](#).

## Struktur Proyek

```plaintext
pdda-backend/
├── cmd/                # Command utilities (e.g., migrasi)
├── controllers/        # Logika bisnis API
├── models/             # Definisi model database
├── routes/             # Routing endpoint API
├── utils/              # Fungsi utilitas
├── main.go             # Entry point aplikasi
├── go.mod              # Dependency manager
└── README.md           # Dokumentasi
```

## Kontribusi

Kontribusi Anda sangat kami hargai! Silakan ikuti langkah-langkah berikut:

1. Fork repositori ini.
2. Buat branch baru untuk fitur atau perbaikan Anda.
   ```bash
   git checkout -b fitur-anda
   ```
3. Lakukan perubahan dan commit:
   ```bash
   git commit -m "Menambahkan fitur X"
   ```
4. Kirim pull request ke branch `main`.

## Lisensi

Proyek ini dilisensikan di bawah lisensi [MIT](LICENSE).

## Kontak

Untuk pertanyaan atau kolaborasi, hubungi:
- **Email**: [email@example.com](mailto:email@example.com)
- **Issues**: Gunakan tab *Issues* untuk melaporkan bug atau meminta fitur baru.

---

Mari bersama-sama membangun backend REST API yang efisien dan scalable untuk mendukung inovasi dalam pengelolaan dataset penyakit tanaman!

