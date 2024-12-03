# Artikel API Documentation

## Endpoint

Base URL: `http://localhost:8081`

### Routes

- **`/getAllArticles`**  
  Menampilkan semua artikel.

- **`/getLatestArticles`**  
  Menampilkan 5 artikel terbaru.

- **`/getArticleWithKey/:title`**  
  Menampilkan artikel berdasarkan judul. Digunakan untuk fitur pencarian artikel dan menampilkan artikel yang relevan setelah klasifikasi sampah.

- **`/showArticle/:slug`**  
  Digunakan untuk membuka detail artikel.

---

## Setup Golang

1. Install Golang dari [situs resmi Golang](https://go.dev/doc/install).
2. Jalankan perintah `go version` di terminal untuk memastikan bahwa Golang telah terinstall dengan benar.

3. Buka project ini di code editor (misalnya Visual Studio Code).

4. Buka terminal di dalam project, lalu jalankan perintah berikut untuk menjalankan aplikasi:

   ```bash
   go run main.go
   ```
