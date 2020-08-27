# Rekomendasi pengembangan servis / aplikasi

1. Logging ketika menyalakan dan mematikan service serta untuk setiap HTTP call yang berhasil diterima oleh service
Logging ini diperlukan agar developer dapat mengetahui jika service yang dijalankannya telah berfungsi dan menerima request

2. Ada sebuah bug ketika melakukan login, yaitu service tidak mendapatkan nilai username dan password sehingga tidak mendapatkan hasil ketika melakukan query ke database. Perbaikannya adalah dengan membaca nilai dari body request. Selain itu, terdapat kesalahan dalam penulisan query untuk mencari administrator (kurang petik untuk mengencapsulate string password)

3. Pengecekan token login hanya perlu dilakukan ketika user mengakses resource selain resource login. Selain itu, responsenya dibuat juta dalam bentuk JSON. Untuk response gagal login dan token tidak valid, balikan status code 401

4. Tabel `customers` dan `customer_addresses` belum menggunakan tipe PRIMARY KEY dan AUTO INCREMENT pada kolom ID nya. Penggunaan ini diperlukan agar setiap record yang masuk memiliki ID yang unique

5. Dalam proses penulisan ke tabel `customer_addresses`, sistem belum meletakkan value customer_id sehingga hubungan antara customers dan customer_addresses menjadi hilang. Selain itu, ada variable `cid` yang tidak diperlukan karena value id ketika melakukan insert ke tabel `customers` dapat langsung di-assign ke variable `id`

6. Dalam proses penghapusan data di tabel `customers`, sistem perlu menghapus juga data di tabel `customer_addresses` dikarenakan customer_addresses memiliki hubungan dengan data di tabel `customers`

7. ketika mengakses endpoint `/get/{id}` dan `/delete/{id}`, sistem perlu memberi balikan 404 jika record dengan id tersebut tidak ditemukan

8. Menambahkan header `Content-Type` `application/json` untuk setiap response yang dikembalikan ke user

9. Mengikuti salah satu kaidah RESTful API, penamaan url untuk resource customer tidak memerlukan nama CRUD function (get, create, delete), cukup gunakan HTTP Method sebagai pemisah fungsi. Selain itu, penamaan resource customer perlu menggunakan plural noun (customers)

10. Mencegah terjadinya SQL injection dengan menggunakan argument placeholder yang disediakan oleh package `database/sql`

11. Mengimplementasikan enkripsi satu arah dalam penyimpanan password administrator. Dalam poin ini, akan diimplementasikan a) sistem untuk register administrator, b) password hashing dengan menggunakan bcrypt, c) penambahan index unique pada tabel `administrator` kolom `username`, d) membuat struct RegisterRequest dan RegisterResponse, e) Mengubah cara komparasi password dengan menggunakan bcrypt, f) meloloskan semua path yang mengandung `auth` dari validasi token

12. Menerapkan pagination pada endpoint index customers. Pagination diletakkan di Header response. Batasi jumlah data yang bisa diambil sebanyak maksimal 50
