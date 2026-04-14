# TaskFlow Backend Assignment

## 1. Overview
TaskFlow is a backend service designed to manage tasks, users, and workflows efficiently. It exposes RESTful APIs to create, update, assign, and track tasks.

This project demonstrates backend engineering fundamentals including API design, database modeling, authentication, and containerized deployment.

---

## 2. Tech Stack
- **Language:** Node.js (JavaScript/TypeScript)
- **Framework:** Express.js
- **Database:** MongoDB
- **Authentication:** JWT
- **Containerization:** Docker

---

## 3. Features
- User authentication (JWT-based login/signup)
- CRUD operations for tasks
- Task assignment to users
- Status tracking (pending, in-progress, completed)
- Role-based access control (optional)

---

## 4. Project Structure
```
taskflow/
│── src/
│   ├── controllers/
│   ├── models/
│   ├── routes/
│   ├── middleware/
│   ├── services/
│   └── app.js
│
│── config/
│── Dockerfile
│── docker-compose.yml
│── package.json
│── README.md
```

---

## 5. Setup Instructions

### Prerequisites
- Node.js installed
- MongoDB running locally or cloud instance
- Docker (optional)

### Installation
```bash
git clone <repo-url>
cd taskflow
npm install
```

### Run Locally
```bash
npm start
```

---

## 6. Environment Variables
Create a `.env` file:

```
PORT=5000
MONGO_URI=your_mongo_connection_string
JWT_SECRET=your_secret_key
```

---

## 7. API Endpoints

### Auth
- POST `/api/auth/register`
- POST `/api/auth/login`

### Tasks
- GET `/api/tasks`
- POST `/api/tasks`
- PUT `/api/tasks/:id`
- DELETE `/api/tasks/:id`

---

## 8. Docker Setup
```bash
docker build -t taskflow .
docker run -p 5000:5000 taskflow
```

---

## 9. Future Improvements
- Add pagination & filtering
- Implement notifications
- Add unit/integration tests
- Improve logging & monitoring

---

## 10. Author
Saket Jasuja
