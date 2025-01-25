async function getWeather() {
    const cityInput = document.getElementById('cityInput');
    const weatherResult = document.getElementById('weatherResult');
    const city = cityInput.value.trim();

    if (!city) {
        alert('Please enter a city name');
        return;
    }

    try {
        const response = await fetch(`/weather?city=${encodeURIComponent(city)}`);
        const data = await response.json();

        if (!response.ok || data.error) {
            throw new Error(data.error || 'Failed to fetch weather data');
        }

        weatherResult.innerHTML = `
            <div class="weather-card">
                <h2>${data.name}</h2>
                <img src="https://openweathermap.org/img/wn/${data.weather[0].icon}@2x.png" alt="Weather icon">
                <p>${data.weather[0].description}</p>
                <p>Temperature: ${Math.round(data.main.temp)}°C</p>
                <p>Feels like: ${Math.round(data.main.feels_like)}°C</p>
                <p>Humidity: ${data.main.humidity}%</p>
            </div>
        `;
    } catch (error) {
        weatherResult.innerHTML = `<p style="color: red;">Error: ${error.message}</p>`;
    }
} 