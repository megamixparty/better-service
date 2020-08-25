# Rekomendasi pengembangan servis / aplikasi

1. Logging ketika menyalakan dan mematikan service serta untuk setiap HTTP call yang berhasil diterima oleh service
Logging ini diperlukan agar developer dapat mengetahui jika service yang dijalankannya telah berfungsi dan menerima request

2. Ada sebuah bug ketika melakukan login, yaitu service tidak mendapatkan nilai username dan password sehingga tidak mendapatkan hasil ketika melakukan query ke database. Perbaikannya adalah dengan membaca nilai dari body request. Selain itu, terdapat kesalahan dalam penulisan query untuk mencari administrator (kurang petik untuk mengencapsulate string password)

3. Pengecekan token login hanya perlu dilakukan ketika user mengakses resource selain resource login. Selain itu, responsenya dibuat juta dalam bentuk JSON. Untuk response gagal login dan token tidak valid, balikan status code 401

4. Tabel `customers` dan `customer_addresses` belum menggunakan tipe PRIMARY KEY dan AUTO INCREMENT pada kolom ID nya. Penggunaan ini diperlukan agar setiap record yang masuk memiliki ID yang unique

5. Dalam proses penulisan ke tabel `customer_addresses`, sistem belum meletakkan value customer_id sehingga hubungan antara customers dan customer_addresses menjadi hilang. Selain itu, ada variable `cid` yang tidak diperlukan karena value id ketika melakukan insert ke tabel `customers` dapat langsung di-assign ke variable `id`
