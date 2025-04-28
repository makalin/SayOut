# 📣 SayOut — Minimal Micro Twitter Clone

[![License](https://img.shields.io/github/license/makalin/Sayo)](https://github.com/makalin/Sayo/blob/main/LICENSE)
[![Build](https://img.shields.io/badge/build-passing-brightgreen)]()
[![Go Backend](https://img.shields.io/badge/backend-Go%20%2B%20Fiber-blue)]()
[![Frontend](https://img.shields.io/badge/frontend-HTMX%20%2B%20TailwindCSS-blueviolet)]()
[![Docker](https://img.shields.io/badge/deploy-Docker%20%2B%20Caddy-2496ED)]()
[![Made with ❤️](https://img.shields.io/badge/made%20by-makalin-red)](https://github.com/makalin)

---

SayOut is a lightweight **micro social network** where you can **Say Out** your thoughts — either to the **public** or just to your **friends**. Designed for minimal resources with blazing fast performance.

---

## 🚀 Features
- ✨ Dual-Mode "Says"  
  • **Public Says** — Share with everyone  
  • **Private Says** — Only friends or extended circles
- ⚡ Minimal, fast UI (no SPA bloat)
- 🗄️ Embedded **SQLite** database
- 🔒 JWT Authentication
- 🌐 Easy deployment via Docker + Caddy
- 🌓 Dark/Light theme support
- 🔄 Real-time updates with HTMX
- 👥 User profiles and friend system

---

## 🛠️ Tech Stack
- **Backend:** Go + Fiber  
- **Frontend:** HTMX + TailwindCSS  
- **Database:** SQLite  
- **Auth:** JWT + bcrypt  
- **Deploy:** Docker + Caddy Server  

---

## 📂 Project Structure
```
SayOut/
├── backend/
│   ├── main.go           # Main application entry
│   ├── models/           # Database models
│   │   └── models.go     # User, Say, and Friend models
│   ├── auth/             # Authentication
│   │   └── auth.go       # JWT and password handling
│   └── handlers/         # API handlers
│       └── handlers.go   # Route handlers
├── frontend/
│   ├── index.html        # Main page
│   ├── login.html        # Login page
│   ├── register.html     # Registration page
│   ├── templates/        # HTML templates
│   │   └── say.html      # Say template
│   └── css/
│       └── styles.css    # Custom styles
├── database/
│   └── schema.sql        # Database schema
├── assets/               # Static assets
├── Dockerfile           # Backend container
├── docker-compose.yml   # Service definitions
├── Caddyfile           # Reverse proxy config
└── README.md
```

---

## 🚧 Installation & Run

1. Clone the repository:
```bash
git clone https://github.com/makalin/SayOut.git
cd SayOut
```

2. Set up environment variables:
```bash
cp .env.example .env
# Edit .env with your configuration
```

3. Start the application:
```bash
docker-compose up --build
```

Access at: [http://localhost:8080](http://localhost:8080)

---

## 🔑 Authentication
- JWT-based authentication
- Secure password hashing with bcrypt
- Protected API endpoints
- Session management

## 📝 Features in Detail
- **User Management**
  - Registration and login
  - Profile customization
  - Friend system
- **Content**
  - Public and private posts
  - Real-time updates
  - Responsive design
- **UI/UX**
  - Dark/Light theme
  - Mobile-friendly
  - Minimal and clean interface

---

## ⚡ Roadmap
- [x] Core "Say" system
- [x] Public / Private modes
- [x] User authentication
- [x] Dark theme support
- [ ] Friend & Circle management
- [ ] Notifications & Mentions
- [ ] UI/UX enhancements
- [ ] Live Demo Deployment

---

## 🤝 Contributing
Pull requests are welcome! Open an issue for feature suggestions or bug reports.

---

## 📄 License
[MIT License](https://github.com/makalin/SayOut/blob/main/LICENSE)

---

## 🌐 Live Demo
_Coming soon..._
