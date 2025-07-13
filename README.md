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
<img width="1380" height="762" alt="Screenshot 2025-07-12 213903" src="https://github.com/user-attachments/assets/7538e84c-d7c2-414b-945a-84af0d7470d2" />
<img width="1783" height="571" alt="Screenshot 2025-07-12 213851" src="https://github.com/user-attachments/assets/8281c98f-21ba-434d-bf30-382b87843cd0" />
<img width="1312" height="562" alt="Screenshot 2025-07-12 213837" src="https://github.com/user-attachments/assets/df8c4d4c-1e7c-45f2-810b-9070e381bab0" />
<img width="940" height="649" alt="Screenshot 2025-07-12 213213" src="https://github.com/user-attachments/assets/046908d7-ccba-456d-abd9-c62a7f46e833" />
<img width="627" height="357" alt="Screenshot 2025-07-12 213200" src="https://github.com/user-attachments/assets/5d1fb99b-ae2c-4a48-9403-d84b8dab0b16" />
<img width="1880" height="824" alt="Screenshot 2025-07-12 213120" src="https://github.com/user-attachments/assets/b2081d69-01fb-4e65-87b9-d67931e74f13" />
<img width="1893" height="764" alt="Screenshot 2025-07-12 213100" src="https://github.com/user-attachments/assets/58765af7-bf2b-4645-aa69-23e434e0919b" />
<img width="1888" height="820" alt="Screenshot 2025-07-12 213039" src="https://github.com/user-attachments/assets/ee88c796-a583-487d-a6ba-f994b2fa6f79" />
<img width="1893" height="824" alt="Screenshot 2025-07-12 213018" src="https://github.com/user-attachments/assets/479a415f-1123-4ca8-a1b3-50539ee8cf83" />

