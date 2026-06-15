# QTEC DevOps Engineer Practical Task - Final Submission

## Production-Style Containerized API with CI/CD, Reverse Proxy & Observability

---

### Objective
This project fully meets **all core requirements** of the DevOps Engineer practical task. I built a simple API, containerized it, added Nginx as reverse proxy, implemented CI/CD, enabled zero-downtime deployments, added monitoring, and showed how it can handle ~100 requests/sec. Everything is deployed on AWS EC2.

---

### System Architecture

### Client → Nginx Reverse Proxy (Port 80) → Go API Container (Port 8080)

- **Frontend Entry Point**: Nginx handles incoming traffic and forwards it to the backend.
- **Backend**: Lightweight Go application running inside Docker.
- **Orchestration**: Docker Compose (local & on EC2).
- **Monitoring**: Prometheus + Grafana.
- **Cloud**: AWS EC2.

**Screenshots**:
- AWS EC2 Instance: ![EC2](https://github.com/user-attachments/assets/ab6f0d65-e60b-43d0-81aa-5b97bf916ea8)
---

### API Endpoints
- **GET /status** – Returns service health status.
- **POST /data** – Accepts JSON data and returns confirmation.

Example:
```json
// GET /status
{"status": "ok", "timestamp": "2026-06-15T12:00:00Z"}

// POST /data
{"name": "qtec"}  →  {"message": "received", "data": {"name": "qtec"}}
```

---

### Containerization Approach
I used **Docker** with a **multi-stage Dockerfile**:
- First stage builds the Go application.
- Second stage uses lightweight Alpine image for runtime.
- Runs as non-root user for security.
- Supports environment variables.
- Small and efficient final image.

**Dockerfile** is in the `app/` folder.

**How to run locally**:
```bash
git clone https://github.com/Saaaaakibhai/qtec_devopsengineer_practicaltask.git
cd qtec_devopsengineer_practicaltask
docker compose up --build -d
```

**Screenshot - Running Docker Containers**:
![Docker](https://github.com/user-attachments/assets/49e01455-4825-4ba5-9444-a84db1a46c81)

---

### Reverse Proxy Setup
**Nginx** is used as reverse proxy:
- Listens on port 80.
- Forwards all requests to the Go API container.
- Configured for basic load balancing (ready for multiple API instances).

Configuration file: `nginx/default.conf`

**Screenshot - Nginx Network**:
![Nginx Network](https://github.com/user-attachments/assets/9eb266bd-c40c-4f56-a5a7-6b2d1b37b8a2)
<img width="1912" height="960" alt="Screenshot from 2026-06-15 12-15-27" src="https://github.com/user-attachments/assets/6da66c9a-d1fc-4b3b-a5ea-dc96abf1932f" />


---

### CI/CD Pipeline
Used **GitHub Actions** for automation.
- Triggered on every push to `main` branch.
- Builds the Go app.
- Runs basic validation.
- Builds the Docker image.

Pipeline file: `.github/workflows/ci-cd.yml`

**Screenshot - Successful CI/CD Pipeline**:
![CI/CD](https://github.com/user-attachments/assets/a3c37256-4cb4-4889-bf3a-c97f8672db9b)

---

### Deployment Process
1. Code pushed to GitHub → CI/CD pipeline runs.
2. Docker image is built.
3. On AWS EC2: `docker compose up --build -d` deploys the new version.
4. Nginx continues serving traffic during update.

**Deployment setup** is in `docker-compose.yml` (includes API, Nginx, Prometheus, Grafana).

---

### How Zero-Downtime Deployment Works
- Docker Compose starts the new container version first.
- Nginx keeps routing requests to the old container until the new one is healthy.
- Once new container is ready, old one is stopped.
- Users experience **no interruption**.

**Screenshot - Zero-Downtime in Action**:
![Zero-Downtime](https://github.com/user-attachments/assets/89f5a473-1569-40e2-9543-40959524dd27)

---

### Logging & Monitoring Setup
- **Application logs**: Printed to stdout (visible via `docker logs`).
- **Nginx logs**: Access and error logs.
- **Prometheus**: Collects metrics from the system and API.
- **Grafana**: Provides nice dashboards for monitoring.

**Screenshots**:
- Grafana Dashboard: ![Grafana](https://github.com/user-attachments/assets/926b8564-152a-4340-a5b6-e8c157303dbb)
- Prometheus: ![Prometheus](https://github.com/user-attachments/assets/584383a6-9206-4f6b-909f-47c8aa6d8a11)

---

### How the System Can Handle ~100 Requests/Sec
- **Stateless design**: Go API uses goroutines for high concurrency.
- **Horizontal scaling**: Easy to run multiple API containers.
- **Nginx load balancing**: Distributes traffic across containers.
- **Lightweight & efficient**: Go + Alpine Docker image.

I performed load testing and it successfully handled over 100 requests per second.

**Screenshots - Load Testing**:
![Load Test Result](https://github.com/user-attachments/assets/5aef9a4d-5a00-4111-8c9f-35220db87450)
![Load Test](https://github.com/user-attachments/assets/3dccfec0-df61-4355-ab52-b6e6d590eb41)


---

### Project Structure
```text
├── app/                    # Go source code + Dockerfile
├── nginx/                  # Nginx configuration
├── .github/workflows/      # CI/CD pipeline
├── docker-compose.yml      # Deployment setup
├── prometheus/ & grafana/  # Monitoring
└── README.md
```

---

### Security Practices
- No hardcoded credentials
- Non-root user in container
- Multi-stage builds (smaller attack surface)

---

### How to Run Locally
```bash
git clone https://github.com/Saaaaakibhai/qtec_devopsengineer_practicaltask.git
cd qtec_devopsengineer_practicaltask
docker compose up --build -d
```
Then open `http://localhost` in your browser.

---

**All submission requirements are completed**: Source code, Dockerfile, Proxy config, CI/CD, Deployment setup, and detailed README.

**Repository**: https://github.com/Saaaaakibhai/qtec_devopsengineer_practicaltask

---
*Submitted by: S. M. Mahedi Hasan*  
*Date: June 15, 2026*
```
