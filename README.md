Berikut adalah dokumentasi lengkap untuk API Anda, yang mencakup semua endpoint dari kontroler `menucontroller` dan `menufavoritecontroller`.

## Eatdah API Documentation

### Base URL
```
http://your-api-base-url.com
```

### Menus Endpoints

#### Get All Menus
```
GET /api/menus
```
- **Description:** Mengambil daftar semua menu yang tersedia.
- **Response:**
  - Status Code: 200
  - Body:
    ```json
    {
        "menu": [
            {
                "id_menu": 1,
                "kategori": "Makanan",
                "nama_menu": "Nasi Goreng",
                "warung": "Warung Makmur",
                "harga": "15000",
                "harga_diskon": "",
                "rating": "4.5",
                "deskripsi": "Nasi goreng dengan bumbu rempah"
            },
            {
                "id_menu": 2,
                "kategori": "Minuman",
                "nama_menu": "Es Teh Manis",
                "warung": "Warung Segar",
                "harga": "5000",
                "harga_diskon": "",
                "rating": "4.2",
                "deskripsi": "Es teh manis segar dengan es batu"
            }
        ]
    }
    ```

#### Get Menu by ID
```
GET /api/menu/:id
```
- **Description:** Mengambil detail menu berdasarkan ID.
- **Parameters:**
  - `id` (path) - ID dari menu yang ingin diambil.
- **Response:**
  - Status Code: 200
  - Body:
    ```json
    {
        "menu": {
            "id_menu": 1,
            "kategori": "Makanan",
            "nama_menu": "Nasi Goreng",
            "warung": "Warung Makmur",
            "harga": "15000",
            "harga_diskon": "",
            "rating": "4.5",
            "deskripsi": "Nasi goreng dengan bumbu rempah"
        }
    }
    ```

#### Create Menu
```
POST /api/menu
```
- **Description:** Membuat menu baru.
- **Request Body:**
  ```json
  {
      "kategori": "Makanan",
      "nama_menu": "Mie Kuah",
      "warung": "Warung Enak",
      "harga": "13000",
      "harga_diskon": "",
      "rating": "4.2",
      "deskripsi": "Mie kuah dengan campuran rempah"
  }
  ```
- **Response:**
  - Status Code: 200
  - Body:
    ```json
    {
        "menu": {
            "id_menu": 3,
            "kategori": "Makanan",
            "nama_menu": "Mie Kuah",
            "warung": "Warung Enak",
            "harga": "13000",
            "harga_diskon": "",
            "rating": "4.2",
            "deskripsi": "Mie kuah dengan campuran rempah"
        }
    }
    ```

#### Update Menu
```
PUT /api/menu/:id
```
- **Description:** Mengubah menu yang sudah ada berdasarkan ID.
- **Parameters:**
  - `id` (path) - ID dari menu yang ingin diubah.
- **Request Body:**
  ```json
  {
      "kategori": "Makanan",
      "nama_menu": "Nasi Uduk",
      "warung": "Warung Makmur",
      "harga": "17000",
      "harga_diskon": "",
      "rating": "4.7",
      "deskripsi": "Nasi uduk dengan santan dan rempah"
  }
  ```
- **Response:**
  - Status Code: 200
  - Body:
    ```json
    {
        "message": "Data Berhasil Diperbarui"
    }
    ```

#### Delete Menu
```
DELETE /api/menu/:id
```
- **Description:** Menghapus menu berdasarkan ID.
- **Parameters:**
  - `id` (path) - ID dari menu yang ingin dihapus.
- **Request Body:**
  ```json
  {
      "id": 1
  }
  ```
- **Response:**
  - Status Code: 200
  - Body:
    ```json
    {
        "message": "Data Berhasil Dihapus"
    }
    ```

### Menus Favorite Endpoints

#### Get All Favorite Menus
```
GET /api/menusfavorite
```
- **Description:** Mengambil daftar semua menu favorit yang tersedia.
- **Response:**
  - Status Code: 200
  - Body:
    ```json
    {
        "menu_favorite": [
            {
                "id_menu_favorite": 1
            },
            {
                "id_menu_favorite": 2
            }
        ]
    }
    ```

#### Get Favorite Menu by ID
```
GET /api/menufavorite/:id
```
- **Description:** Mengambil detail menu favorit berdasarkan ID.
- **Parameters:**
  - `id` (path) - ID dari menu favorit yang ingin diambil.
- **Response:**
  - Status Code: 200
  - Body:
    ```json
    {
        "MenuFavorite": {
            "id_menu_favorite": 1
        }
    }
    ```

#### Create Favorite Menu
```
POST /api/menufavorite
```
- **Description:** Membuat menu favorit baru.
- **Request Body:**
  ```json
  {
      "id_menu_favorite": 3
  }
  ```
- **Response:**
  - Status Code: 200
  - Body:
    ```json
    {
        "MenuFavorite": {
            "id_menu_favorite": 3
        }
    }
    ```

#### Update Favorite Menu
```
PUT /api/menufavorite/:id
```
- **Description:** Mengubah menu favorit yang sudah ada berdasarkan ID.
- **Parameters:**
  - `id` (path) - ID dari menu favorit yang ingin diubah.
- **Request Body:**
  ```json
  {
      "id_menu_favorite": 1
  }
  ```
- **Response:**
  - Status Code: 200
  - Body:
    ```json
    {
        "message": "Data Berhasil Diperbarui"
    }
    ```

#### Delete Favorite Menu
```
DELETE /api/menufavorite/:id
```
- **Description:** Menghapus menu favorit berdasarkan ID.
- **Parameters:**
  - `id` (path) - ID dari menu favorit yang ingin dihapus.
- **Request Body:**
  ```json
  {
      "id": 1
  }
  ```
- **Response:**
  - Status Code: 200
  - Body:
    ```json
    {
        "message": "Data Berhasil Dihapus"
    }
    ```

Pastikan untuk mengintegrasikan informasi ini

 dalam dokumentasi API Anda agar pengguna dan pengembang dapat dengan mudah memahami cara menggunakan API Anda.
