# go-albums-api

## Introduction
go-albums-api is a simple API written in Golang using the Gin web framework. This API comes from the official Golang tutorial; it was my first contact with the language.
This API allows users to manage a collection of albums, where each album has an ID, title, artist, list of musics, and price.

### GitHub Repository
[https://github.com/yan-pi/go-albums-api/](https://github.com/yan-pi/go-albums-api/)

## Endpoints

### 1. Get All Albums

#### Endpoint
```
GET /albums
```

#### Description
Returns a list of all albums stored in the API.

#### Response
- **Status Code**: 200 OK
- **Body**:
  ```json
  [
    {
      "id": "1",
      "title": "Dialogos",
      "artist": "Bardek",
      "musics": ["Dialogos", "Agora", "Amanhã", "Sempre"],
      "price": 19.99
    }
  ]
  ```

### 2. Get Album by ID

#### Endpoint
```
GET /albums/:id
```

#### Parameters
- **id**: The ID of the album to retrieve.

#### Description
Returns the album with the specified ID.

#### Response
- **Status Code**: 200 OK
- **Body**:
  ```json
  {
    "id": "1",
    "title": "Dialogos",
    "artist": "Bardek",
    "musics": ["Dialogos", "Agora", "Amanhã", "Sempre"],
    "price": 19.99
  }
  ```

- **Status Code**: 404 Not Found
- **Body**:
  ```json
  {
    "message": "album not found"
  }
  ```

### 3. Add a New Album

#### Endpoint
```
POST /albums
```

#### Request Body
```json
{
  "id": "2",
  "title": "New Album",
  "artist": "New Artist",
  "musics": ["Song1", "Song2"],
  "price": 9.99
}
```

#### Description
Adds a new album to the collection.

#### Response
- **Status Code**: 201 Created
- **Body**:
  ```json
  {
    "id": "2",
    "title": "New Album",
    "artist": "New Artist",
    "musics": ["Song1", "Song2"],
    "price": 9.99
  }
  ```

## Running the API

To run the API, execute the following command:

```bash
go run main.go
```

After running the command, the API will start and listen on `localhost:8080`.

## Conclusion

This is a simple API for managing a collection of albums using Golang and Gin. You can use the provided endpoints to get all albums, get an album by ID, and add a new album to the collection.
