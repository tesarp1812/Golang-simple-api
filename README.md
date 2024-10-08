# API Go Project

Ini adalah proyek API yang ditulis dalam bahasa Go. Proyek ini menggunakan GORM untuk interaksi dengan database dan dilengkapi dengan Docker dan Docker Compose untuk kemudahan pengembangan dan penerapan.

## Daftar Isi

- [Fitur](#fitur)
- [Rute API](#rute-api)
- [Pengaturan Lingkungan](#pengaturan-lingkungan)
- [Menjalankan Aplikasi](#menjalankan-aplikasi)

## Fitur

- Menggunakan GORM untuk ORM dengan MySQL.
- API sederhana dengan rute CRUD.
- Containerisasi dengan Docker.

## Rute API

| Metode | Rute           | Deskripsi                  |
|--------|----------------|----------------------------|
| GET    | `/animals`     | Mendapatkan semua hewan    |
| GET    | `/animals/:id` | Mendapatkan hewan berdasarkan ID |
| POST   | `/animals`     | Menambahkan hewan baru      |
| PUT    | `/animals/:id` | Memperbarui hewan berdasarkan ID |
| DELETE | `/animals/:id` | Menghapus hewan berdasarkan ID |

## Pengaturan Lingkungan

Sebelum menjalankan aplikasi, pastikan untuk mengatur variabel lingkungan berikut dalam file `.env` atau langsung di `docker-compose.yaml`:

```plaintext
DB_USER=root
DB_PASSWORD=rootpassword
DB_NAME=db_animal
DB_HOST=db
```

## Menjalankan Aplikasi
1. Pastikan Anda memiliki Docker dan Docker Compose terinstal.
2. Jalankan perintah berikut untuk membangun dan menjalankan aplikasi:
   ```
   docker-compose up --build -d
   
Perintah ini akan:
 - Membangun image Docker dari Dockerfile.
 - Menjalankan aplikasi dan database di latar belakang.

3. Akses API Anda melalui:
   ```
   http://localhost:8080

## Licensi
Proyek ini dilisensikan di bawah MIT License.

### Penjelasan Struktur `README.md`
- **Judul dan Deskripsi**: Menyediakan informasi dasar tentang proyek.
- **Daftar Isi**: Memudahkan navigasi dalam dokumentasi.
- **Fitur**: Menjelaskan fitur utama aplikasi.
- **Struktur Proyek**: Menyediakan informasi tentang struktur direktori dan fungsinya.
- **Rute API**: Menyediakan tabel rute API untuk referensi cepat.
- **Pengaturan Lingkungan**: Menginstruksikan pengguna untuk mengatur variabel lingkungan.
- **Menjalankan Aplikasi**: Langkah-langkah untuk membangun dan menjalankan aplikasi.
- **Kontribusi**: Menyediakan petunjuk untuk kontribusi.
- **Lisensi**: Menyebutkan lisensi proyek.

Anda dapat menyesuaikan bagian mana pun sesuai kebutuhan proyek Anda.
