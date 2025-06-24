# Load Shedder Demo  

## Overview  

This project demonstrates **load shedding**—a resilience pattern that prioritizes critical system functionality during high traffic by selectively rejecting non-essential requests. The implementation includes:  

- **Essential services** that always remain available  
- **Non-essential services** that get disabled under heavy load  
- **Configurable thresholds** for request handling  
- **Interactive frontend** to simulate traffic spikes  

## Key Features  

✔ **Priority-Based Request Handling**  
- Essential endpoints (e.g., payments, auth) always process requests  
- Non-essential endpoints (e.g., recommendations) shed load when thresholds are exceeded  

✔ **Real-Time Load Monitoring**  
- Atomic counters track concurrent requests  
- Middleware enforces global limits  

✔ **Interactive Simulation**  
- Web interface to generate artificial load  
- Visual feedback on system behavior  

✔ **Graceful Degradation**  
- Clean HTTP 503 responses for rejected requests  
- No silent failures or timeouts  

## Getting Started  

### Prerequisites  
- Go 1.20+  
- Modern web browser  

### Installation  
1. Clone the repository:  
   ```sh
   git clone https://github.com/johneliud/load_shedder.git
   ```  
2. Navigate to the project directory  
3. Run the server from the root directory:  
   ```sh
   go run backend/main.go
   ```  
4. Open `http://localhost:8080` in your browser  

## How It Works  

1. **Request Flow**  
   - All requests pass through the shedding middleware  
   - The system tracks active requests in real-time  

2. **Threshold Enforcement**  
   - Default limit: 50 concurrent requests  
   - Essential services bypass the limit  
   - Non-essential services reject requests when overloaded  

3. **Simulation Controls**  
   - Generate 10-100 concurrent requests  
   - Test both service types under load  
   - Observe live request counter  

## Use Cases  

This pattern is valuable for:  
- **E-commerce platforms** during flash sales  
- **APIs** experiencing traffic spikes  
- **Financial systems** prioritizing transactions  
- **IoT networks** handling emergency signals  

## Configuration  

Adjust in `shed_wrapper.go`:  
- `maxRequests`: Change the concurrent request limit  
- `canShed`: Toggle shedding per endpoint  

## Contributing  

Pull requests welcome! For major changes, please open an issue first.  

## License  

[MIT](https://github.com/johneliud/load_shedder/blob/main/LICENSE)  