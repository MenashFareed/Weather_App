# Weather App

A simple weather application built with Go and JavaScript.

## Setup

1. Clone the repository
2. Copy `.env.example` to `.env`:
   ```bash
   cp .env.example .env
   ```
3. Get an API key from [OpenWeatherMap](https://openweathermap.org/api)
4. Add your API key to the `.env` file
5. Install dependencies:
   ```bash
   go mod download
   ```
6. Run the application:
   ```bash
   go run cmd/main.go
   ```

## Development

For hot reloading during development:
bash
go install github.com/air-verse/air@latest
air
