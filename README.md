# MyGram API
---

## Table of Contents

- [Deskripsi](#deskripsi)
- [Struktur Folder](#struktur-folder)
- [Dokumentasi API](#dokumentasi-api)
---

## Deskripsi
MyGram API merupakan sebuah service yang dapat melakukan pengoperasian CRUD (Create Read Update Delete) pada beberapa entitas yang dibangun. MyGram API dibangun dengan framework gin dan gorm sebagai ORM guna mapping data dari database yang pada project ini menggunakan PostgreSQL. 

---
## Struktur Folder
Adapun struktur folder project ini mengadopsi struktur dari Clean Architecture yang secara sederhana merupakan konsep pemisahan kode kedalam layer yang saling terkait dan bergantung. Dengan begitu, maka pemeliharaan kode kedepannya akan mudah diubah, diuji, dan dipelihara. Berikut adalah deskripsi singkat dari setiap layer,

| Layer                | Directory      |
|----------------------|----------------|
| Frameworks & Drivers | app            |
| Entities             | models         |
| Interface            | handlers       |
| Usecases             | services       |
| Repository           | repositories   |

---

## Dokumentasi API

