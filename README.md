# 🌩️ Go-meteor: Scalable Weather Forecast CLI in Go
![Go](https://img.shields.io/badge/Go-1.20%2B-00ADD8?logo=go&logoColor=white) 
![License](https://img.shields.io/badge/License-MIT-blue) 
![Terminal](https://img.shields.io/badge/CLI-Tool-%2344cc11)  


🚀 **go-meteor** is a high-performance, modular, and extensible **command-line application** that fetches real-time and forecasted weather data using the **WeatherAPI**. Designed with **clean architecture principles**, it provides robust error handling, caching, and a plugin-based architecture for future enhancements.  



## **System Overview & Architecture**  

###  🔹 High-Level System Design
go-meteor is designed with a **layered architecture** to ensure scalability, maintainability, and flexibility.  



### **🔹Key Design Principles**  

 **Clean Architecture** – Separation of concerns across CLI, business logic, and API layers.  
 **Caching (LRU-based)** – Improves performance by avoiding redundant API calls.  
 **Rate Limiting** – Prevents excessive API requests to stay within free-tier limits.  
 **Configuration-Driven** – Uses `.env` and flags to control API key and location.  
 **Error Resilience** – Graceful fallbacks for API failures or network issues.  


## **🌟 Features**  

### **🔹 Core Functionalities**  
 **Current Weather** – Displays real-time temperature, humidity, and conditions.  
 **Hourly Forecast** – Predicts the next 12 hours' weather.  
 **Location Search** – Fetches data for any global location.  

### **🔹 Advanced Features**  
 **LRU Caching** – Reduces API calls and speeds up response times.  
 **Rate Limiting** – Ensures API usage stays within free-tier limits.  
 **Multiple API Support** – Can be extended to fetch data from different providers.  
 **Modular Design** – Easily extend functionality with new modules.  

---
![image](https://github.com/user-attachments/assets/910c664b-75e9-451c-9ca4-8de109fba1bc)


## **🚀 Installation**  

### **1️ Clone the Repository**  
```sh
git clone https://github.com/your-username/go-meteor.git
cd go-meteor
```

### **2. Build and Run the application**
```sh
go build -o go-meteor

./go-meteor "San Francisco"
          (OR)

go run ./ Mumbai
```

## 🔑 API Key Setup
This app uses WeatherAPI. Get a free API key and create a .env file:

```
WEATHER_API_KEY=your_api_key_here
```
### Then Load it into your environment
```
source .env
```

## 🤝 Contribute
- Fork the repo

- Create your feature branch

- Commit changes

- Push and open a PR

---
