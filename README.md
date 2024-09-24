# E-commerce High Concurrency System

This project is an e-commerce flash sale system developed in Golang, designed to handle high concurrency and ensure high availability using a distributed microservices architecture. The system is optimized for high performance and scalability, making it suitable for handling large volumes of traffic during flash sale events.

## Features

- High concurrency handling
- Distributed microservices architecture
- Product management
- Order processing
- User authentication and authorization
- Inventory management
- Real-time monitoring and observability

## Tech Stack

- Backend: Golang
- Database: MySQL/PostgreSQL
- Message Queue: Kafka/RabbitMQ
- Cache: Redis
- API Gateway: Kong/NGINX
- Containerization: Docker
- Orchestration: Kubernetes
- Observability: Prometheus, Grafana, ELK Stack (Elasticsearch, Logstash, Kibana)

## Project Structure
![14](https://github.com/user-attachments/assets/67aa94cc-2596-45ff-bc6e-a0b3049755da)

- `back-end`: Backend services
- `front-end`: Frontend static files
- `iris`: Web framework related code
- `rabbitmq`: RabbitMQ integration
- `distributed verification`: Distributed system verification
- `cookies verify`: User authentication
- `order background`: Order processing backend
- `oversold message queue`: Queue for handling oversold situations
- `product exhibition`: Product display functionality
- `receptionist user login`: User login and reception

## Requirements

- Golang 1.16+
- Docker
- Kubernetes
- Helm (optional)
- Prometheus
- Grafana
- ELK Stack
- Kafka/RabbitMQ
- Redis

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## Acknowledgements

- All contributors to this project
- Open-source libraries and tools used in this project
