# Game Management Microservice API

A high-performance RESTful microservice built with Go that manages video game data through a PostgreSQL database. This service provides comprehensive CRUD operations for video game management with a focus on reliability and efficiency.

## 🎯 Features

- **Full CRUD Operations**: Create, Read, Update, and Delete video game entries
- **RESTful Architecture**: Standard HTTP methods and status codes
- **PostgreSQL Integration**: Robust data persistence
- **JSON Communication**: Standardized data exchange format
- **Gorilla Mux Router**: Efficient request routing
- **Containerized**: Docker-ready PostgreSQL configuration
- **Error Handling**: Comprehensive error responses

## 🛠️ Technical Stack

- **Language**: Go
- **Router**: Gorilla Mux
- **Database**: PostgreSQL
- **Data Format**: JSON
- **Container**: Docker support

## 📊 API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/games` | Create a new game entry |
| GET | `/games` | Retrieve all games |
| GET | `/games/{id}` | Retrieve a specific game |
| PUT/PATCH | `/games/{id}` | Update a game entry |
| DELETE | `/games/{id}` | Delete a game entry |

## 🔧 Installation

1. Clone the repository
```bash
git clone https://github.com/yourusername/game-management-api.git
cd game-management-api
```

2. Set up PostgreSQL database
```bash
docker run --name pgdb -e POSTGRES_PASSWORD=mypassword -p 5432:5432 -d postgres
```

3. Install dependencies
```bash
go mod download
```

4. Build and run
```bash
go build
./game-management-api
```

## 💻 Usage

### Create a Game
```bash
curl -X POST http://localhost:8000/games \
  -H "Content-Type: application/json" \
  -d '{
    "title": "The Legend of Zelda",
    "console": "Nintendo Switch",
    "rating": 9.8,
    "complete": false
  }'
```

### Retrieve All Games
```bash
curl http://localhost:8000/games
```

### Retrieve Specific Game
```bash
curl http://localhost:8000/games/1
```

### Update a Game
```bash
curl -X PUT http://localhost:8000/games/1 \
  -H "Content-Type: application/json" \
  -d '{
    "title": "The Legend of Zelda",
    "console": "Nintendo Switch",
    "rating": 10.0,
    "complete": true
  }'
```

### Delete a Game
```bash
curl -X DELETE http://localhost:8000/games/1
```

## 📝 Data Model

```go
type Game struct {
    ID       int       `json:"id"`
    Title    string    `json:"title"`
    Console  string    `json:"console"`
    Rating   float64   `json:"rating"`
    Complete bool      `json:"complete"`
    Created  time.Time `json:"created"`
    Updated  time.Time `json:"updated"`
}
```

## 🔒 Environment Configuration

```go
const (
    DatabaseUser     = "postgres"
    DatabasePassword = "mypassword"
    DatabaseHost     = "pgdb"
    DatabaseName     = "postgres"
)
```

## 📁 Project Structure
```
game-management-api/
├── main.go          # Application entry point and database setup
├── handler.go       # HTTP request handlers
├── go.mod           # Go module definition
├── go.sum           # Go module checksums
└── README.md        # Documentation
```

## 🔍 Error Handling

The API returns appropriate HTTP status codes:

- `200 OK`: Successful GET, PUT, PATCH requests
- `201 Created`: Successful POST requests
- `204 No Content`: Successful DELETE requests
- `400 Bad Request`: Invalid request data
- `404 Not Found`: Resource not found
- `500 Internal Server Error`: Server-side errors

## 🚀 Performance Considerations

- Connection pooling with `database/sql`
- Prepared statements for SQL queries
- JSON encoding/decoding optimization
- Gorilla Mux router for efficient routing

## 🔐 Security Features

- SQL injection prevention through parameterized queries
- Input validation for all endpoints
- Proper error message handling without exposing system details

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the [LICENSE.txt](LICENSE.txt) file for details.

## 📚 API Documentation

For detailed API documentation, see the following examples:

<details>
<summary>Sample Response Formats</summary>

```json
{
  "id": 1,
  "title": "The Legend of Zelda",
  "console": "Nintendo Switch",
  "rating": 9.8,
  "complete": false,
  "created": "2024-12-04T10:00:00Z",
  "updated": "2024-12-04T10:00:00Z"
}
```
</details>

## 📞 Contact

For questions and feedback, please open an issue in the repository.
