# E-Commerce Microservices Implementation Roadmap

## Tech Stack Overview

| Component | Choice | Notes |
|-----------|--------|-------|
| Language | Go 1.23+ | |
| Web Framework | Gin | Fast, minimal, well-documented |
| Message Broker | Apache Kafka | Event-driven communication between services |
| Serialization | Apache Avro | Schema evolution, compact binary format for Kafka |
| API Gateway | Nginx | Reverse proxy, load balancing, rate limiting |
| User Service DB | PostgreSQL | Relational, ACID for auth & profiles |
| Product Service DB | PostgreSQL | Relational, complex queries, inventory |
| Cart Service DB | MongoDB | Document model, flexible schema, high read/write |
| Order Service DB | PostgreSQL | ACID for transactions |
| Payment Service DB | PostgreSQL | Financial data integrity |
| Notification Service | PostgreSQL or MongoDB | Event logs, delivery status |

---

## Architecture Diagram (Target State)

```
                    ┌─────────────────┐
                    │   Nginx (API    │
                    │    Gateway)     │
                    └────────┬────────┘
                             │
        ┌────────────────────┼────────────────────┐
        │                    │                    │
        ▼                    ▼                    ▼
┌───────────────┐   ┌───────────────┐   ┌───────────────┐
│ User Service  │   │ Product Svc   │   │ Cart Service  │
│   (Postgres)  │   │  (Postgres)   │   │  (MongoDB)    │
└───────┬───────┘   └───────┬───────┘   └───────┬───────┘
        │                   │                   │
        └───────────────────┼───────────────────┘
                            │
                    ┌───────▼───────┐
                    │    Kafka      │
                    └───────┬───────┘
                            │
        ┌───────────────────┼───────────────────┐
        │                   │                   │
        ▼                   ▼                   ▼
┌───────────────┐   ┌───────────────┐   ┌───────────────┐
│ Order Service │   │ Payment Svc    │   │ Notification  │
│  (Postgres)   │   │  (Postgres)   │   │    Service     │
└───────────────┘   └───────────────┘   └───────────────┘
```

---

## Day-by-Day Implementation Plan

### Day 1: User Service
**Goal:** Production-ready user registration, authentication, profile management.

**What you'll learn:**
- Gin routing, middleware, validation
- Password hashing (bcrypt)
- JWT for authentication
- PostgreSQL with GORM
- Environment-based config
- Docker for the service

**Deliverables:**
- [ ] REST API: `POST /register`, `POST /login`, `GET /users/:id`, `PUT /users/:id`
- [ ] Password hashing with bcrypt
- [ ] JWT token generation & validation middleware
- [ ] Dockerfile + docker-compose for User Service + Postgres
- [ ] Config via env vars (DB URL, JWT secret)

---

### Day 2: Product Catalog Service
**Goal:** Admin CRUD for products, categories, inventory.

**What you'll learn:**
- RESTful resource design
- GORM relations (products ↔ categories)
- Admin auth (JWT validation from User Service)
- Inventory constraints

**Deliverables:**
- [ ] REST API: Products CRUD, Categories CRUD
- [ ] Inventory management (update quantity)
- [ ] Validate JWT from User Service (shared secret or public key)
- [ ] Docker + Postgres

---

### Day 3: Shopping Cart Service
**Goal:** Create cart, add/remove items, update quantities.

**What you'll learn:**
- MongoDB with Go (mongo-driver)
- Document model for cart
- Service-to-service calls (validate user, check product existence)

**Deliverables:**
- [ ] REST API: Create cart, add item, remove item, update quantity, get cart
- [ ] MongoDB integration
- [ ] User validation via User Service or JWT
- [ ] Docker + MongoDB

---

### Day 4: Kafka + Avro Setup & Order Service
**Goal:** Event-driven plumbing + Order placement flow.

**What you'll learn:**
- Kafka producers/consumers in Go
- Avro schema design & serialization
- Order state machine (pending → paid → shipped)

**Deliverables:**
- [ ] Kafka + Zookeeper in docker-compose
- [ ] Avro schemas for events (e.g., `OrderPlaced`, `OrderPaid`)
- [ ] Order Service: place order, check inventory (via Product Service), publish to Kafka
- [ ] Docker for Order Service + Postgres

---

### Day 5: Payment Service
**Goal:** Process payments, integrate Stripe (or mock), publish `OrderPaid`.

**What you'll learn:**
- Stripe SDK (or mock server)
- Kafka consumer for `OrderPlaced`, producer for `OrderPaid`
- Idempotency for payment handling

**Deliverables:**
- [ ] Consume `OrderPlaced` from Kafka
- [ ] Create Stripe payment intent
- [ ] Publish `OrderPaid` on success
- [ ] Store payment records in Postgres

---

### Day 6: Notification Service
**Goal:** Consume events, send email/SMS via SendGrid/Twilio.

**What you'll learn:**
- Kafka consumer for multiple event types
- SendGrid/Twilio APIs
- Template-based notifications

**Deliverables:**
- [ ] Consume `OrderPlaced`, `OrderPaid`, `OrderShipped`
- [ ] Send email (order confirmation, shipping)
- [ ] Optional: SMS via Twilio
- [ ] Docker + config for API keys

---

### Day 7: Nginx API Gateway
**Goal:** Single entry point, route to services, rate limiting.

**What you'll learn:**
- Nginx reverse proxy configuration
- Path-based routing
- CORS, rate limiting

**Deliverables:**
- [ ] `/users/*` → User Service
- [ ] `/products/*` → Product Service
- [ ] `/carts/*` → Cart Service
- [ ] `/orders/*` → Order Service
- [ ] `/payments/*` → Payment Service (if needed)
- [ ] Rate limiting, CORS

---

## Shared Components (Build Incrementally)

- **`pkg/config`** – Load env vars, DB URLs, Kafka brokers
- **`pkg/auth`** – JWT validation middleware (shared across services)
- **`pkg/kafka`** – Producer/consumer helpers with Avro
- **`pkg/avro`** – Schema definitions and codegen

---

## Current State vs Target

| Component | Current | Target |
|-----------|---------|--------|
| User Service | Basic CRUD, plaintext password | Bcrypt + JWT, proper validation |
| Product Service | Skeleton | Full CRUD, categories, inventory |
| Cart Service | Partial (Postgres) | MongoDB, full cart operations |
| Order Service | Skeleton | Kafka events, inventory check |
| Payment Service | Skeleton | Stripe + Kafka |
| Notification Service | Skeleton | Kafka consumer + SendGrid |
| Kafka | Not used | Central event bus |
| Avro | Not used | Event serialization |
| Nginx | Not used | API Gateway |

---

## Where to Start

**Day 1 (User Service)** is the right starting point because:
1. Other services depend on it (auth, user IDs)
2. You already have a base implementation to extend
3. You’ll set up patterns (config, Docker, JWT) that every service will use

Next step: **Implement Day 1 – User Service** with password hashing, JWT, Docker, and config.
