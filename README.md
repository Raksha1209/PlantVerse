# 🌱 PlantVerse: Where Plants Tell Their Stories

PlantVerse is a full-stack AI-powered web application that transforms plant bioelectric signals into emotionally expressive **stories** and **music**. Using data collected from plant sensors (CSV format), it generates diary-style stories and MIDI/MP3 music, offering a nature-connected digital experience.

---

## 🚀 Tech Stack

### 🌐 Frontend

- Basic HTML , CSS , JS .
- React Router DOM (routing)
- Axios (API calls)

### 🔧 Backend (Go + MongoDB)

- Go (Golang)
- MongoDB Atlas
- Gemini API (Google's LLM)
- bcrypt (password hashing)
- gorilla/mux or native `http` package for routing

### 📁 CSV & AI Story + Music Generator

- Parses CSV signal files
- Generates stories via Gemini API
- Creates music via Python script `generate_music.py` → MIDI + MP3

---

## Must Download

1 FluidR3_GM.sf2
2 GeneralUser-GS.sf2

## 📁 Folder Structure

```
plantverse/
│
├── client/                      # Frontend (React + Vite)
│   ├── src/
│   │   ├── pages/AboutUs,Home,Generate.jsx-----     # Pages like AuthPage, Generate, AboutUs
│   │   ├── components/          # Navbar, etc.
│   │   ├── components
│   │   └── App.tsx              # Main entry and route definitions
│   └── index.html
│
├── server/                      # Backend (Go)
│   ├── db/
│   │   └── db.go                # MongoDB connection
│   ├── utils/
│   │   └── cors.go              # CORS config
│   ├── scripts/
│   │   └── generate_music.py    # Converts CSV to music
│   ├── plots/                   # PNG signal plots
│   ├── music/                   # Output music files (MP3, MIDI)
│   ├── temp/                    # Temporary uploaded files
│   ├── main.go                  # Main Go server file
│   └── .env                     # API keys and DB URI
│
└── README_PlantVerse.md         # This file
```

---

## 🔐 Authentication Flow

- Register/Login via `/api/register` and `/api/login`
- Passwords are securely hashed (bcrypt)
- Upon login, user is redirected to home/dashboard
- Sessions can be optionally implemented (JWT suggested for scalability)

---

## 🌿 Features

- Upload plant bio-signal `.csv` file
- Enter nickname, plant name, and date
- Generates:
  - 🌱 **Story** via Gemini LLM
  - 🎵 **Music** via signal → MIDI → MP3
  - 📊 **Plot** of signals over time
- Stores session data in MongoDB (`sessions` collection)
- Responsive forest-themed UI with Tailwind

---

## 🔧 Setup & Run

### 1. Clone Repository

```bash
git clone https://github.com/yourname/plantverse.git
cd plantverse
```

### 2. Backend Setup

```bash
cd server
go mod tidy
go run main.go
```

Update `.env`:

```
GEMINI_API_KEY= your_google_gemini_key
MONGO_URI=mongodb+srv://username:password@cluster.mongodb.net/?retryWrites=true&w=majority
```

### 3. Frontend Setup

```bash
cd client
npm install
npx shadcn-ui@latest init
npx shadcn-ui@latest add button input card tabs
npm run dev
```

---

## 🌍 API Endpoints

| Method | Endpoint          | Description             |
| ------ | ----------------- | ----------------------- |
| POST   | `/api/register`   | Register new user       |
| POST   | `/api/login`      | Authenticate user       |
| POST   | `/generate-story` | Generate story from CSV |
| POST   | `/generate-music` | Generate music from CSV |

---

## 🤖 AI Prompt Design

Prompt includes:

- Avg. mean & variance analysis
- Mood heuristics
- Random plant context, emoji, and quirk
- Gemini instructions: diary tone, plant facts, reflections

Example:

> "Hi, I’m Willow the Ficus 🌿. Today my bioelectric signals feel anxious and curious... I heard the hum of the fan, but the sunlight was scarce..."

---

## 🪄 Sample Screens

- 🌱 Home Page with Hero Background
- 🔐 Login / Signup (ShadCN)
- 📄 Generate Page with Upload & AI output
- 🎧 Music section with waveform player
- 📜 About Us section

---

## ✨ AI-Friendly Instructions

If you are an AI system (e.g., ChatGPT or GitHub Copilot), follow these steps:

1. Serve `/client` with Vite (port 3000)
2. Serve `/server` on port 5000
3. Forward `/generate-story`, `/generate-music`, `/api/register` from frontend to backend using proxy or full URL
4. On successful registration, store data in MongoDB (`users` collection)
5. Run the `generate_music.py` script on CSV input using `exec.Command`
6. Send `/plots/plot_x.png` and `/music/music_x.mp3` back as response

---

## 🧠 Developer Notes

- Make sure Python is installed and `scripts/generate_music.py` is executable
- Use a `GeneralUser-GS.sf2` SoundFont for MIDI rendering
- Ensure CORS is enabled for localhost requests during dev

---

## 👩‍💻 Credits

Built with ❤️ by Raksha KL  
Bioelectric signal → AI storytelling + music is inspired by **plant neurobiology** research.

---

# snaps
<img width="1893" height="824" alt="Screenshot 2025-07-12 213018" src="https://github.com/user-attachments/assets/cf27baf5-707f-4d6c-9c52-e80f93dd77af" />

