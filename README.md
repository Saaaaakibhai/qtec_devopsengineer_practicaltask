# QTEC DevOps Engineer Practical Task

## Production-Style Containerized API with CI/CD, Reverse Proxy & Observability

---

## Objective

This project demonstrates a production-style system implementing:

- Containerization using Docker
- CI/CD automation using GitHub Actions
- Traffic management using Nginx reverse proxy
- Basic observability using logs
- Scalable API design (~100 requests/sec)
- Zero-downtime deployment strategy

---

## System Architecture

Client → Nginx (Reverse Proxy) → Go API (Containerized) → Docker Runtime

---

## API Endpoints

### GET /status

Returns service health:

{"status":"ok"}

---

### POST /data

Request:
{"name":"qtec"}

Response:
{
  "message": "received",
  "data": {
    "name": "qtec"
  }
}

---

## Containerization

This service is containerized using Docker.

Features:
- Multi-stage build
- Lightweight Alpine Linux image
- Production-ready runtime
- Environment variable support

Run locally:

docker compose up --build

---

## Reverse Proxy (Nginx)

Nginx is used as a reverse proxy to:

- Route incoming traffic
- Forward requests to the Go API container
- Act as a single entry point (port 80)
- Support future load balancing

Flow:
Client → Nginx → API Container

---

## CI/CD Pipeline (GitHub Actions)

Automated pipeline performs:

On every push to main:

- Checkout repository
- Setup Go environment
- Build application
- Validate build
- Build Docker image

Pipeline file:
.github/workflows/ci-cd.yml

---

## Zero-Downtime Deployment

Zero-downtime deployment is achieved through:

- Starting new container before stopping old one
- Nginx continuously routing traffic
- Docker Compose controlled service updates

This ensures users never experience downtime during deployment.

---

## Scalability (~100 requests/sec)

This system can handle ~100 requests/sec using:

- Stateless Go API design
- Horizontal scaling using multiple containers
- Nginx load balancing
- Lightweight Go runtime with high concurrency support

---

## Logging & Monitoring

Currently implemented:

- Application logs via stdout
- Docker logs
- Nginx access logs

Future improvements:

- Prometheus for metrics
- Grafana for dashboards
- ELK stack for centralized logging

---

## How to Run

git clone https://github.com/Saaaaakibhai/qtec_devopsengineer_practicaltask.git

cd qtec_devopsengineer_practicaltask

docker compose up --build

---

## Project Structure

qtec_devopsengineer_practicaltask/
├── app/
│   ├── main.go
│   ├── go.mod
│   ├── Dockerfile
│
├── nginx/
│   └── default.conf
│
├── .github/
│   └── workflows/
│       └── ci-cd.yml
│
├── docker-compose.yml
├── .gitignore
├── README.md

---

## Security Practices

- No hardcoded credentials
- Uses environment variables
- Lightweight minimal container image
- Non-root container execution

---

## Cloud Deployment Options

This system can be deployed on:

- AWS (EC2 / ECS / EKS)
- Google Cloud Platform
- DigitalOcean

---

## Skills Demonstrated

- Docker containerization
- CI/CD automation (GitHub Actions)
- Reverse proxy configuration (Nginx)
- Scalable API design
- Microservice architecture
- Zero-downtime deployment concept
- Basic system observability

---

## Conclusion

This project demonstrates a production-ready DevOps workflow including automation, scalability, and deployment strategies used in real-world systems.
