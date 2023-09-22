# Kelas ORM 1

## Library yang dipakai

```go
    go get -u gorm.io/gorm
    go get -u gorm.io/driver/mysql // sesuaikan db masing-masing
```

## Format ENV
```env
    SERVER=
    DBPORT=
    DBHOST=
    DBUSER=
    DBPASS=
    DBNAME=
```

## JWT Notes

JWT Token akan dihasilkan ketika melakukan login.
Proses untuk membuat JWT Token tertera pada `helper > jwt.go`. Untuk data yang disimpan dalam token, bisa disesuaikan kembali pada bagian claims. Fungsi dasar untuk melakukan extract token juga sudah disediakan.

Pada contoh penggunaan JWT pada endpoint ada pada bagian routing. Selamat mendalami middleware!!!