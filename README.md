# **BOOK STORE API**

by Riyan Naffa Nusafara

<details>
  <summary>Contents</summary>
  <ol>
    <li>
      <a href="#project-overview">Project Overview</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
        <a href="#book-collection">Book Collection</a>
        <ul>
            <li><a href="#api-handler">API Handler</a></li>
            <li>
                <a href="#endpoint-handler">Endpoint Handler</a>
                <ul>
                    <li><a href="#display-all-books">Display All Books</a></li>
                    <li><a href="#book-details">Book Details</a></li>
                    <li><a href="#update-a-book">Update a Book</a></li>
                    <li><a href="#add-a-book">Add a Book</a></li>
                    <li><a href="#delete-a-book">Delete a Book</a></li>
                </ul>
            </li>
        </ul>
    </li>
    <li>
        <a href="#employee-collection">Employee Collection</a>
        <ul>
            <li><a href="#api-handler-1">API Handler</a></li>
            <li>
                <a href="#endpoint-handler-1">Endpoint Handler</a>
                <ul>
                    <li><a href="#display-all-employees">Display All Employees</a></li>
                    <li><a href="#add-an-employee">Add an Employee</a></li>
                    <li><a href="#delete-an-employee">Delete an Employee</a></li>
                </ul>
            </li>
        </ul>
    </li>
  </ol>
</details>

## Project Overview

Repository ini merupakan tempat untuk mengumpulkan final project dari Pelatihan Web Development KMTETI yang diselenggarakan oleh Divisi Workshop KMTETI. Project ini berfokus pada pengembangan website dari sisi backend. Project ini berfokus pada request handling pada 7 API endpoint dan koneksi ke database.

### Built With

Project ini dibangun dengan bahasa **Go** sebagai bahasa pemrograman dan **MongoDB** sebagai database, lebih rincinya **MongoDB Atlas** sebagai remote database.


## [Book Collection][BookCollection]

Koleksi ini merupakan bagian dari database `book-store-dev` yang khusus menyimpan data buku. Buku disimpan dalam bentuk document dengan field dan key:

```json
{
    "_id": ObjectID,
    "title": String,
    "author": String,
    "year": Int32,
    "stock": Int32,
    "price": Decimal128,
}
```

### [API Handler][BookHandler]

```go
func BookHandler(w http.ResponseWriter, r *http.Request)
```

Handler mengambil HTTP response writer sebagai output dari data yang diterima dan HTTP request sebagai data yang ingin disampaikan ke database. Fungsi ini dipanggil ketika terdapat request dari endpoint `/api/book`.

### Endpoint Handler

#### [Display All Books][BookDisplayAll]

```go
func BookDisplayAll() ([]*model.BookDisplay, error)
```

Fungsi ini dipanggil ketika terdapat request dari endpoint `/api/book?action=display`. Keluaran dari fungsi ini adalah slice dari dokumen buku-buku dengan field yang untuk ditampilkan dan error jika terjadi suatu kesalahan dalam koneksi ke database, encoding, maupun request.

#### [Book Details][BookDetails]

```go
func BookDetails(idReq *string) (*model.BookDetailed, error)
```

Fungsi ini dipanggil ketika terdapat request dari endpoint `/api/book?action=details&id=<idReq>`. Fungsi ini mengambil variabel pointer string sebagai kata kunci untuk menemukan dokumen yang dimaksud. Keluaran dari fungsi ini adalah dokumen mendetail dari buku yang diminta dan sebuah error jika terjadi error.

#### [Update a Book][BookUpdate]

```go
func BookUpdate(req io.Reader) (*mongo.UpdateResult, error)
```

Fungsi ini dipanggil ketika terdapat request dari endpoint `/api/book?action=update`. Fungsi ini menerima parameter berupa HTTP Body dengan isi sebuah dokumen JSON yang berisi ID, harga, dan stok yang ingin diubah. Kembalian dari fungsi ini adalah hasil update dokumen dan error jika terjadi error.

#### [Add a Book][BookAdd]

```go
func BookAdd(req io.Reader) (*mongo.InsertOneResult, error)
```

Fungsi ini dipanggil ketika terdapat request dari endpoint `/api/book?action=add`. Fungsi ini menerima parameter berupa HTTP Body dengan isi sebuah dokumen JSON yang berisi judul, pengarang, tahun terbit, stok, dan harga dari buku yang ingin ditambahkan. ID akan ditambahkan secara otomatis oleh program karena ID bersifat unik setiap dokumen. Keluaran fungsi ini adalah hasil memasukkan data dan sebuah error jika terjadi error.

#### [Delete a Book][BookDelete]

```go
func BookDelete(idReq *string) (*model.BookDeleteResponse, error)
```

Fungsi ini dipanggil ketika terdapat request dari endpoint `/api/book?action=delete`. Fungsi ini menerima parameter sebuah variabel pointer string yang menunjukkan ID dari buku yang ingin dihapus. Keluaran dari fungsi ini adalah respons penghapusan dokumen yang berisi ID dan judul buku yang telah berhasil dihapus dan sebuah error jika terjadi error.


## [Employee Collection][EmployeeCollection]

Koleksi ini merupakan bagian dari database `book-store-dev` yang khusus menyimpan data pegawai. Data pegawai disimpan dalam bentuk document dengan field dan key:

```json
{
    "_id": ObjectID,
    "name": String,
    "nik": String,
    "edu": String,
    "date": {
        "day": Int32,
        "month": Int32,
        "year": Int 32,
    },
    "type": String,
}
```

### [API Handler][EmployeeHandler]

```go
func EmployeeHandler(w http.ResponseWriter, r *http.Request)
```

Handler mengambil HTTP response writer sebagai output dari data yang diterima dan HTTP request sebagai data yang ingin disampaikan ke database. Fungsi ini dipanggil ketika terdapat request dari endpoint `/api/employee`.

### Endpoint Handler

#### [Display All Employees][EmployeeDisplayAll]

```go
func EmployeeDisplayAll() ([]*model.EmployeeDisplay, error)
```

Fungsi ini dipanggil ketika terdapat request dari endpoint `/api/employee?action=display`. Keluaran dari fungsi ini adalah slice dokumen pegawai-pegawai yang ada di database dan sebuah error jika terjadi error.

#### [Add an Employee][EmployeeAdd]

```go
func EmployeeAdd(req io.Reader) (*mongo.InsertOneResult, error)
```

Fungsi ini dipanggil ketika terdapat request dari endpoint `/api/employee?action=add`. Fungsi ini mengambil parameter sebuah HTTP Body yang berisi data pegawai yang ingin dimasukkan ke dalam database dalam format JSON. Keluaran dari fungsi ini adalah response keberhasilan memasukkan dokumen dan sebuah error jika terjadi error.

#### [Delete an Employee][EmployeeDelete]

```go
func EmployeeDelete(idReq *string) (*model.EmployeeDeleteResponse, error)
```

Fungsi ini dipanggil ketika terdapat request dari endpoint `/api/employee?action=delete&id=<idReq>`. Fungsi ini mengambil parameter sebuah pointer string ID dari seorang pegawai yang ingin dihapus dari database. Keluaran dari fungsi ini adalah response keberhasilan penghapusan dokumen dan sebuah error jika terjadi error.

[BookHandler]: https://github.com/RiyanNaffa/KMTETI-GO_Final-Project/blob/eb1837dce0252cd525acd177129fe8c97b25584b/src/api/book.go#L15
[EmployeeHandler]: https://github.com/RiyanNaffa/KMTETI-GO_Final-Project/blob/eb1837dce0252cd525acd177129fe8c97b25584b/src/api/employee.go#L14

[BookCollection]: https://github.com/RiyanNaffa/KMTETI-GO_Final-Project/blob/eb1837dce0252cd525acd177129fe8c97b25584b/src/model/book.model.go#L8
[BookDisplayAll]: https://github.com/RiyanNaffa/KMTETI-GO_Final-Project/blob/eb1837dce0252cd525acd177129fe8c97b25584b/src/service/book.service.go#L20
[BookDetails]: https://github.com/RiyanNaffa/KMTETI-GO_Final-Project/blob/eb1837dce0252cd525acd177129fe8c97b25584b/src/service/book.service.go#L57
[BookUpdate]: https://github.com/RiyanNaffa/KMTETI-GO_Final-Project/blob/eb1837dce0252cd525acd177129fe8c97b25584b/src/service/book.service.go#L97
[BookAdd]: https://github.com/RiyanNaffa/KMTETI-GO_Final-Project/blob/eb1837dce0252cd525acd177129fe8c97b25584b/src/service/book.service.go#L144
[BookDelete]: https://github.com/RiyanNaffa/KMTETI-GO_Final-Project/blob/eb1837dce0252cd525acd177129fe8c97b25584b/src/service/book.service.go#L188

[EmployeeCollection]: https://github.com/RiyanNaffa/KMTETI-GO_Final-Project/blob/eb1837dce0252cd525acd177129fe8c97b25584b/src/model/employee.model.go#L12
[EmployeeDisplayAll]: https://github.com/RiyanNaffa/KMTETI-GO_Final-Project/blob/eb1837dce0252cd525acd177129fe8c97b25584b/src/service/employee.service.go#L18
[EmployeeAdd]: https://github.com/RiyanNaffa/KMTETI-GO_Final-Project/blob/eb1837dce0252cd525acd177129fe8c97b25584b/src/service/employee.service.go#L54
[EmployeeDelete]: https://github.com/RiyanNaffa/KMTETI-GO_Final-Project/blob/eb1837dce0252cd525acd177129fe8c97b25584b/src/service/employee.service.go#L91