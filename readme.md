# MN Room - Project Capstone Group 4

<!-- PROJECT LOGO -->
<br/>
<div align="center">
<!--  mengarah ke repo  -->
  <a href="https://github.com/Capstone-MNRoom">
    <img src="https://gitlab.com/davidwahyup/asset-blog/-/raw/main/images/2022/image/MN_Room.png" width="150" height="150">
  </a>

<h3 align="center">MN Room</h3>

  <p align="center">
    MN Room Application Capstone Project 4
    <br />
    <a href="https://app.swaggerhub.com/apis-docs/davidwah/MNROOM/1.0#/"><strong>» Open API »</strong></a>
    <br />
  </p>
</div>


### 🛠 &nbsp;Build App & Tools

![Golang](https://img.shields.io/badge/-Golang-05122A?style=flat&logo=go&logoColor=4479A1)&nbsp;
![Docker](https://img.shields.io/badge/-Docker-05122A?style=flat&logo=docker)&nbsp;
![GitHub](https://img.shields.io/badge/-GitHub-05122A?style=flat&logo=github)&nbsp;
![Visual Studio Code](https://img.shields.io/badge/-Visual%20Studio%20Code-05122A?style=flat&logo=visual-studio-code&logoColor=007ACC)&nbsp;
![Postman](https://img.shields.io/badge/-Postman-05122A?style=flat&logo=postman)&nbsp;
![MySQL](https://img.shields.io/badge/-MySQL-05122A?style=flat&logo=mysql&logoColor=4479A1)&nbsp;
![GDC](https://img.shields.io/badge/-GoogleCloud-05122A?style=flat&logo=google)&nbsp;
![JSON](https://img.shields.io/badge/-JSON-05122A?style=flat&logo=json&logoColor=000000)&nbsp;
![Ubuntu](https://img.shields.io/badge/-Ubuntu-05122A?style=flat&logo=ubuntu)&nbsp;

## About Project
**MN-Room** merupakan projek capstone yang muncul berawal dari keresahan seseorang yang sedang menyewa sebuah ruangan untuk keperluan acara. Namun ketika menjelang hari acara ternyata pihak rental gedung/ruangan tersebut menyampaikan kabar bahwa ruangan tersebut sudah disewa oleh orang lain. Dari masalah tersebut terpikirkan sebuah solusi yang harusnya sebuah penyewaan gedung/ruangan mempunyai catatan secara online yang dapat dilihat oleh berbagai user dan dapat juga di pesan secara online, serta menerapkan pembayaran secara online juga melalui payment gateway. Fitur-fitur yang ada pada **MN-Room** sebagai berikut:

<div>
      <details>
<summary>🙎 Users</summary>
  
  <!---
  | Command | Description |
| --- | --- |
  --->
  
 Endpoint User terdapat fitur untuk membuat Akun dan Login agar mendapat mengakses berbagai layanan di aplikasi MN-Room, 
 selain itu terdapat fitur Update untuk mengedit data profile user, serta fitur delete untuk menghapus akun.
 
<div>
  
| Feature User | Endpoint | Param | JWT Token | Fungsi |
| --- | --- | --- | --- | --- |
| POST | /signup  | - | NO | Melakukan proses registrasi user |
| POST | /login | - | NO | Melakukan proses login user |
| GET | /users | - | YES | Mendapatkan informasi daftar user terdaftar |
| PUT | /users | - | YES | Melakukan update profile user yang sedang login | 
| DEL | /users | - | YES | Menghapus user yang sedang login |

</details>  

<details>
<summary>🛍 &nbsp;Room</summary>
  
  <!---
  | Command | Description |
| --- | --- |
  --->
  
Pada Room ini user dapat melihat beberapa daftar ruangan dan detail ruangan yang disewakan. Selain itu User juga dapat membuat profile ruangan sendiri yang nantinya akan disewakan. Terdapat beberapa fitur tambahan pada **Room** ini, yaitu *facility* dan *categories* yang bertujuan untuk mengelompokkan ruangan dengan berdasarkan kategori dan menampikan fasilitas yang ada pada ruangan tersebut. Sehingga memudahkan User untuk menentukan pilihan ketika hendak menyewa ruangan.
  
| Feature Products | Endpoint | Param | JWT Token | Fungsi |
| --- | --- | --- | --- | --- |
| POST | /rooms  | - | YES | Membuat rooms profile baru |
| GET | /rooms | - | YES | Mendapatkan informasi seluruh product user yang sedang login |
| GET | /rooms/:id | id | NO | Mendapatkan informasi rooms berdasarkan id-rooms |
| PUT | /rooms | - | YES | Melakukan update profile informasi room |
| GET | /user/rooms | id | NO | Mendapatkan informasi rooms yang terlah dibuat oleh user. |
| DEL | /rooms/:id | id | YES | Melakukan delete rooms tertentu berdasarkan id rooms |

</details>


<details>
<summary>🛒 &nbsp;Rent</summary>
  
  <!---
  | Command | Description |
| --- | --- |
  --->
Cart merupakan fitur untuk menampung berbagai product yang akan dibeli oleh user, adapun fiturnya ada GET dimana user bisa melihat barang apa aja yang ada di dalam keranjang, ada fitur history dimana user bisa melihat jumlah product yang sudah dibayar.
  
| Feature cart | Endpoint | Param | JWT Token | Fungsi |
| --- | --- | --- | --- | --- |
| POST | /rents | - | YES | Melakukan sewa pada sebuah rooms |
| GET | /rents | - | YES | Mendapatkan informasi sewa yang telah dilakukan |

</details>

<details>
<summary>🗓&nbsp;Payments</summary>
  
  <!---
  | Command | Description |
| --- | --- |
  --->
Merupakan fitur untuk dimana user melakukan pembayaran sesuai sewa room yang dipilih. Mode pembayaran ini menggunakan payment gateway yang disediakan oleh xendit, payment_method yang digunakan pada MN-Room yaitu pembayaran melalui BANK BCA dan BNI. 
  
| Feature booking | Endpoint | Param | JWT Token | Fungsi |
| --- | --- | --- | --- | --- |
| GET | /payments | id | YES | Mendapatkan informasi pembayaran berdasarkan rent id |
| POST | /payments/status | - | NO | Membuat pembayaran sewa ruang |

</details>


<details>
<summary>📈&nbsp;ERD</summary>
<img src="./erd/ERD-MNroom.jpg">
</details>  

### Contact

[![GitHub Azis](https://img.shields.io/badge/-Azis-white?style=flat&logo=github&logoColor=black)](https://github.com/mohamadazisadnan)
[![LinkedIn Azis](https://img.shields.io/badge/-Azis-blue?style=flat&logo=linkedin)](https://www.linkedin.com/in/azisadnn/)

[![GitHub David](https://img.shields.io/badge/-David-white?style=flat&logo=github&logoColor=black)](https://github.com/davidwah)
[![LinkedIn David](https://img.shields.io/badge/-David-blue?style=flat&logo=linkedin)](https://www.linkedin.com/in/david-wahyu-pratomo/)

[![GitHub Mulya](https://img.shields.io/badge/-Mulya-white?style=flat&logo=github&logoColor=black)](https://github.com/mulyanurdin10)
[![LinkedIn Mulya](https://img.shields.io/badge/-Mulya-blue?style=flat&logo=linkedin)](https://www.linkedin.com/in/mulya-nurdin-473807246/)
