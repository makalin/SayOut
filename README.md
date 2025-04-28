# 📣 Sayo — Minimal Micro Twitter Clone

![Sayo Logo](./assets/sayo-logo.png)

[![License](https://img.shields.io/github/license/makalin/Sayo)](https://github.com/makalin/Sayo/blob/main/LICENSE)
[![Build](https://img.shields.io/badge/build-passing-brightgreen)]()
[![Go Backend](https://img.shields.io/badge/backend-Go%20%2B%20Fiber-blue)]()
[![Frontend](https://img.shields.io/badge/frontend-HTMX%20%2B%20TailwindCSS-blueviolet)]()
[![Docker](https://img.shields.io/badge/deploy-Docker%20%2B%20Caddy-2496ED)]()
[![Made with ❤️](https://img.shields.io/badge/made%20by-makalin-red)](https://github.com/makalin)

---

Sayo is a lightweight **micro social network** where you can **Say Out** your thoughts — either to the **public** or just to your **friends**. Designed for minimal resources with blazing fast performance.

---

## 🚀 Features
- ✨ Dual-Mode "Says"  
  • **Public Says** — Share with everyone  
  • **Private Says** — Only friends or extended circles
- ⚡ Minimal, fast UI (no SPA bloat)
- 🗄️ Embedded **SQLite** database
- 🔒 JWT Authentication
- 🌐 Easy deployment via Docker + Caddy

---

## 🛠️ Tech Stack
- **Backend:** Go + Fiber  
- **Frontend:** HTMX + TailwindCSS  
- **Database:** SQLite  
- **Auth:** JWT  
- **Deploy:** Docker + Caddy Server  

---

## 📂 Project Structure
```
sayo/
├── backend/          # Go Fiber API
├── frontend/         # HTMX + Tailwind templates
├── database/         # SQLite schema & migrations
├── assets/           # Logos, images
├── docker-compose.yml
└── README.md
```

---

## 🚧 Installation & Run

```bash
git clone https://github.com/makalin/Sayo.git
cd Sayo
docker-compose up --build
```

Access at: [http://localhost:8080](http://localhost:8080)

---

## ⚡ Roadmap
- [x] Core "Say" system
- [x] Public / Private modes
- [ ] Friend & Circle management
- [ ] Notifications & Mentions
- [ ] UI/UX enhancements
- [ ] Live Demo Deployment

---

## 🤝 Contributing
Pull requests are welcome! Open an issue for feature suggestions or bug reports.

---

## 📄 License
[MIT License](https://github.com/makalin/Sayo/blob/main/LICENSE)

---

## 🌐 Live Demo
_Coming soon..._
