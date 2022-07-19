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

# ERD
![ERD-MNRoom](/erd/ERD-MNroom.jpg)