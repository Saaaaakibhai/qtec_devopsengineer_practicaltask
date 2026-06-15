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

## Project Structure

```text
qtec_devopsengineer_practicaltask/
│
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
└── README.md
```

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
screen shot
1. ec2 instance
<img width="1473" height="839" alt="Web_Photo_Editor(1)" src="https://github.com/user-attachments/assets/ab6f0d65-e60b-43d0-81aa-5b97bf916ea8" />


2. grafana

   <img width="1911" height="959" alt="Screenshot from 2026-06-15 11-15-53" src="https://github.com/user-attachments/assets/926b8564-152a-4340-a5b6-e8c157303dbb" />

3. promethues <img width="1911" height="959" alt="Screenshot from 2026-06-15 11-16-15" src="https://github.com/user-attachments/assets/584383a6-9206-4f6b-909f-47c8aa6d8a11" />

4. github action ci cd pipeline <img width="1911" height="959" alt="Screenshot from 2026-06-15 11-16-30" src="https://github.com/user-attachments/assets/a3c37256-4cb4-4889-bf3a-c97f8672db9b" />


5. Handle 100 requeest/sec
   <img width="1912" height="960" alt="Screenshot from 2026-06-15 11-36-36" src="https://github.com/user-attachments/assets/3dccfec0-df61-4355-ab52-b6e6d590eb41" />
<img width="1912" height="960" alt="Screenshot from 2026-06-15 11-36-27" src="https://github.com/user-attachments/assets/5aef9a4d-5a00-4111-8c9f-35220db87450" />


6. docker
   <img width="1912" height="960" alt="Screenshot from 2026-06-15 11-40-38" src="https://github.com/user-attachments/assets/49e01455-4825-4ba5-9444-a84db1a46c81" />

7. zero-downtime deployment works
<img width="1912" height="960" alt="Screenshot from 2026-06-15 11-52-58" src="https://github.com/user-attachments/assets/89f5a473-1569-40e2-9543-40959524dd27" />

8. Nginix Arcitecture and network

<img width="1912" height="960" alt="Screenshot from 2026-06-15 12-15-27" src="https://github.com/user-attachments/assets/c6d6ee1c-8bed-4dd4-b2ca-9bd128552daf" />
<img width="1912" height="960" alt="Screenshot from 2026-06-15 12-15-19" src="https://github.com/user-attachments/assets/9eb266bd-c40c-4f56-a5a7-6b2d1b37b8a2" />



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
