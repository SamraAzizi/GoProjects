# Weather API in Go

This is a simple weather API server built using Go that fetches weather data for a given city using the [OpenWeatherMap API](https://openweathermap.org/).

## Features

- Fetch current weather data for any city.
- Simple REST API endpoints.
- Lightweight and fast server built using Go's `net/http` package.

---

## Prerequisites

Before running the project, ensure you have the following:

1. **Go Installed:**  
   Download and install Go from [golang.org](https://go.dev/dl/).
   
2. **OpenWeatherMap API Key:**  
   Sign up at [OpenWeatherMap](https://home.openweathermap.org/users/sign_up) to get your free API key.

---

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/yourusername/weather-go-app.git
   cd weather-go-app

2. Create a file named apiConfig.json in the project root directory and add your API key:

```bash
{
    "OpenWeatherMapApiKey": "your_api_key_here"
}
```

3. Run the Go Application

```bash
go run main.go

```
4. You should see the message:
```bash
Starting server on port 8000...
```


## API Usage
The server exposes the following endpoints:

## 1. Hello Route

- Endpoint: /hello
- Method: GET
- Example Request:

```bash
curl http://localhost:8000/hello

```
- Response:
```bash
Hello from Go!
```


## Project Structure
