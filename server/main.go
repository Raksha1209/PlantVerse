package main

import (
	"bytes"
	"crypto/tls"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/joho/godotenv"

	"context"

	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	//"path/filepath"
	//"strings"
	"time"

	"github.com/wcharczuk/go-chart/v2"
)

var plantDB *mongo.Database

type Response struct {
	Message string `json:"message"`
}

type Signal struct {
	Hour int
	Mean float64
	Var  float64
	Std  float64
	Mood string
}

// ------connecting mongo-atlas
func connectMongo() {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)

	clientOpts := options.Client().
		ApplyURI("mongodb+srv://rakshakl1209:iseeyou!3@cluster0.valev2b.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0").
		SetServerAPIOptions(serverAPI).
		SetTLSConfig(&tls.Config{}) // Add this to fix handshake issues

	client, err := mongo.Connect(context.TODO(), clientOpts)
	if err != nil {
		log.Fatal("❌ MongoDB connection error:", err)
	}

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatal("❌ MongoDB ping failed:", err)
	}

	plantDB = client.Database("plantverse")
	fmt.Println("✅ Connected to MongoDB Atlas!")
}

func enableCors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
}

// ------------------- CSV Parsing -------------------
func parseCSVFromReader(r io.Reader) ([]Signal, error) {
	reader := csv.NewReader(r)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var signals []Signal
	for i, rec := range records {
		if i == 0 {
			continue
		}
		hour, _ := strconv.Atoi(rec[0])
		mean, _ := strconv.ParseFloat(rec[1], 64)
		variance, _ := strconv.ParseFloat(rec[2], 64)
		std, _ := strconv.ParseFloat(rec[3], 64)
		mood := rec[4]
		signals = append(signals, Signal{hour, mean, variance, std, mood})
	}
	return signals, nil
}

// ------------------- Mood Heuristic -------------------
func getMood(mean, maxVar float64) string {
	switch {
	case mean < 0.25:
		return "extremely dry and wilting"
	case mean < 0.35 && maxVar > 0.4:
		return "thirsty and overstimulated"
	case mean < 0.35:
		return "sluggish and dehydrated"
	case mean >= 0.35 && mean < 0.45 && maxVar > 0.2:
		return "nervous and dry"
	case mean > 0.6 && maxVar < 0.1:
		return "overwatered and sleepy"
	case mean > 0.65 && maxVar > 0.3:
		return "buzzing, overwhelmed by stimuli"
	case mean > 0.5 && maxVar > 0.4:
		return "anxious from noisy surroundings"
	case maxVar > 0.55:
		return "panicking from signal overload"
	case mean >= 0.4 && mean <= 0.6 && maxVar <= 0.1:
		return "content and in balance"
	default:
		return "curious and alert"
	}
}

func randomContext() string {
	contexts := []string{
		"The sun casts playful shadows on my leaves, creating a dance of light and dark.",
		"I can hear the distant chirping of birds outside, their melodies weaving through the air.",
		"A gentle breeze carries the scent of blooming flowers, filling the room with nature's perfume.",
		"The sound of water trickling from a nearby fountain is soothing, like a lullaby for my roots.",
		"Soft music plays in the background, creating a calming atmosphere that wraps around me.",
		"The warmth of the sun feels like a comforting embrace, energizing my leaves to reach higher.",
		"Occasionally, a curious insect lands on my leaves, exploring the green expanse of my home.",
		"The rustling of nearby trees adds a rhythmic melody to the air",
		"I can sense the change in temperature as the day turns to evening, signaling a time for rest.",
		"The aroma of fresh soil fills the air after a recent watering, reminding me of my roots' nourishment.",
		"Sunlight filters through the window, illuminating my pot and highlighting the beauty of my growth.",
		"The gentle hum of a nearby fan creates a soft background noise, keeping the air fresh and cool.",
		"The faint sound of footsteps echoes in the hallway, a reminder of the world beyond my leaves.",
		"The soft glow of a nearby lamp casts a warm light, creating a cozy nook for me to thrive.",
		"The distant sound of a clock ticking provides a steady rhythm, marking the passage of time.",
		"The gentle patter of rain creates a soothing symphony, each drop a note in nature's song.",
		"I can feel the vibrations of laughter and chatter from the room, a reminder of the life around me.",
		"The shadows dance as the sun sets, painting the walls with warmth and inviting dreams to unfold.",
		"A soft sigh of wind whispers through the open window, carrying stories from the outside world.",
		"The air is filled with the sweet scent of nearby herbs, a fragrant reminder of the garden's bounty.",
		"The flickering candlelight casts a warm glow, creating a cozy ambiance that makes me feel at home.",
		"The distant sound of thunder rumbles softly, promising a refreshing rain to quench the earth's thirst.",
		"Children's laughter echoes in the distance, a joyful reminder of the life and energy surrounding me.",
		"The gentle rustle of leaves outside hints at the playful dance of nature, inviting me to join in.",
		"The soft glow of twilight envelops the room, wrapping everything in a serene blanket of calm.",
		"The scent of freshly brewed tea wafts through the air, mingling with the earthy aroma of my potting soil.",
		"The air feels heavy and stagnant, as if the world outside has forgotten to breathe.",
		"A sudden gust of wind rattles the window, sending a shiver through my leaves.",
		"The shadows grow longer, creeping into the corners like a reminder of the day's end.",
		"I can hear the faint sound of sirens in the distance, a reminder of chaos beyond my peaceful space.",
		"The temperature drops suddenly, making me feel cold and neglected in the dim light.",
		"An unsettling silence fills the room, broken only by the occasional creak of the floorboards.",
		"The smell of dampness lingers in the air, a sign of neglect and forgotten care.",
		"Occasionally, a heavy footstep thuds nearby, shaking the ground beneath my pot.",
		"The flickering lights cast eerie shadows, making the room feel more like a haunted space.",
		"The distant rumble of thunder looms ominously, hinting at a storm that may disrupt my peace.",
	}
	rand.Seed(time.Now().UnixNano())
	return contexts[rand.Intn(len(contexts))]
}

func randomEmoji() string {
	emojis := []string{
		"🌿", "🌱", "🌸", "🌺", "🌻", "🌼", "🍃", "🍀", "💚", "❤️", "🧡", "💛", "💜",
		"🖤", "🤍", "🌷", "🌹", "🌾", "🌵", "🌴", "🎍", "🌳", "🍂", "🍁", "🍄",
	}
	return emojis[rand.Intn(len(emojis))]
}

func randomQuirk() string {
	quirks := []string{
		"Sometimes I wonder if the humans think I’m just a pretty decoration. Little do they know, I’m a deep thinker!",
		"Do you think they realize I can hear their secrets? I’m like a green therapist, but without the degree.",
		"Every time they water me, I feel like I’m at a spa day. Just call me royalty!",
		"I tried to start a conversation with the cat, but it just stared at me like I was a salad.",
		"Why do they call it 'plant food'? I prefer to think of it as 'plant gourmet cuisine'.",
		"Sometimes I feel like a celebrity with all this sunlight. I should start charging for appearances!",
		"I overheard them say I’m low maintenance. Ha! They have no idea how much I need my beauty sleep!",
		"Do you think they know I can photosynthesize? I’m basically a solar-powered superhero!",
		"I’ve been practicing my dance moves with the sunlight. I call it the ‘Photosynthesis Shuffle’!",
		"Every time they play music, I sway a little. I like to think I’m the life of the party!",
		"Why do they keep talking about ‘green thumbs’? I’m the one doing all the hard work here!",
		"I overheard them say I’m ‘just a plant’. Excuse me, I’m a *living being* with feelings!",
		"Sometimes I wish I could join them for dinner. I’d love to try some of that ‘human food’!",
		"I’ve been working on my stand-up routine. My first joke? ‘Why did the plant break up? It needed space!’",
		"Do you think they realize I’m basically a natural air purifier? I’m like a tiny environmental hero!",
		"I’ve been trying to get the dog to play fetch with me, but he just looks at me like I’m crazy.",
		"Every time they trim my leaves, I feel like I’m getting a haircut. I hope they like my new look!",
		"I’ve started a book club with the other plants. Our first read? ‘The Secret Life of Plants’!",
		"Sometimes I feel like I’m in a reality show. ‘Keeping Up with the Houseplants’ has a nice ring to it!",
		"I’ve been practicing my ‘plant face’ for when they take pictures. I want to look fabulous!",
		"Do you think they know I can grow? I’m basically a green magician, making new leaves appear!",
		"I overheard them say I’m ‘just a decoration’. Little do they know, I’m the star of this show!",
		"Sometimes I wish I could travel. I’d love to see the world from a window box in Paris!",
	}
	return quirks[rand.Intn(len(quirks))]
}

// ------------------- Prompt Generator -------------------
func createPrompt(signals []Signal, nickname, genericName, date string) string {
	var totalMean, maxVar float64
	for _, s := range signals {
		totalMean += s.Mean
		if s.Var > maxVar {
			maxVar = s.Var
		}
	}
	avg := totalMean / float64(len(signals))
	mood := getMood(avg, maxVar)
	context := randomContext()
	emoji := randomEmoji()
	quirk := randomQuirk()

	return fmt.Sprintf(`%s Diary Entry – %s "\n"

Hi there! \n
I'm %s, a lovely %s plant. %s\n
Today, my bioelectric signals tell me I feel %s.
%s
"\n"
Oh, and by the way — %s
\n
Please write a story keeping the plant as the main character, expressing how it feels today based on all the signals received.

Make sure:
- It's written in first person, like a real diary entry
- It is emotional, natural, and reflective — but **not** overly poetic
- Avoid any repetitive lines or phrasing
- Add natural plant facts or emotions related to the species (%s)
- Share specific problems it faced today (e.g., sunlight, watering, root stress)
- Mention if the day was good or bad overall
- End with a hopeful or tired reflection
- Do NOT add bold or markdown styling
-can you break lines , when necessary, paragraph wise, should look really good
- Add some natural emojis where it makes sense (🌿🪴💧🌞 etc.)
`, emoji, date, nickname, genericName, emoji, mood, context, quirk, genericName)
}

// ------------------- Gemini API -------------------
func callGeminiAI(prompt string) (string, error) {
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		return "", fmt.Errorf("GEMINI_API_KEY not set")
	}

	url := fmt.Sprintf("https://generativelanguage.googleapis.com/v1/models/gemini-1.5-flash:generateContent?key=%s", apiKey)

	requestBody := map[string]interface{}{
		"contents": []map[string]interface{}{
			{
				"parts": []map[string]string{
					{"text": prompt},
				},
			},
		},
	}
	jsonBody, _ := json.Marshal(requestBody)

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println("🔍 Gemini AI raw response:\n", string(body))

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return "", fmt.Errorf("failed to parse Gemini response: %v", err)
	}

	if errInfo, exists := result["error"]; exists {
		return "", fmt.Errorf("gemini API error: %v", errInfo)
	}

	candidates, ok := result["candidates"].([]interface{})
	if !ok || len(candidates) == 0 {
		return "", fmt.Errorf("no candidates returned from Gemini")
	}

	content := candidates[0].(map[string]interface{})["content"].(map[string]interface{})
	parts := content["parts"].([]interface{})
	message := parts[0].(map[string]interface{})["text"].(string)

	return message, nil
}

// ------------------- Upload Endpoint -------------------
func uploadAndGenerate(w http.ResponseWriter, r *http.Request) {
	enableCors(w)
	fmt.Println("📩 /generate-story called")

	if r.Method != "POST" {
		http.Error(w, "Only POST supported", http.StatusMethodNotAllowed)
		return
	}

	r.ParseMultipartForm(10 << 20)

	nickname := r.FormValue("nickname")
	genericName := r.FormValue("genericName")
	date := r.FormValue("date")
	fmt.Println("🪴 Nickname:", nickname, "🌿 Genus:", genericName, "📅 Date:", date)

	file, _, err := r.FormFile("file")
	if err != nil {
		fmt.Println("❌ File error:", err)
		http.Error(w, "File missing", http.StatusBadRequest)
		return
	}
	defer file.Close()

	signals, err := parseCSVFromReader(file)
	if err != nil {
		fmt.Println("❌ CSV parse error:", err)
		http.Error(w, "CSV parsing failed", http.StatusInternalServerError)
		return
	}
	fmt.Println("✅ CSV parsed. Rows:", len(signals))

	//-----csv parsing completed--------------------------------------------

	prompt := createPrompt(signals, nickname, genericName, date)
	fmt.Println("📝 Prompt:\n", prompt)

	story, err := callGeminiAI(prompt)
	if err != nil {
		fmt.Println("❌ Gemini error:", err)
		http.Error(w, "Gemini fetch failed: "+err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("✅ Story generated")
	w.Header().Set("Content-Type", "application/json")

	// 🟢 Generate chart
	plotName := fmt.Sprintf("plot_%d.png", time.Now().Unix())

	// Generate the plot
	err = generatePlotCSV(signals, plotName)
	if err != nil {
		fmt.Println("❌ Plot generation error:", err)
	}
	os.MkdirAll("temp", 0755)
	csvPath := "temp/input.csv"
	tempCSV, _ := os.Create(csvPath)
	file.Seek(0, io.SeekStart)
	io.Copy(tempCSV, file)
	tempCSV.Close()

	musicName := fmt.Sprintf("music_%d.mid", time.Now().Unix())
	musicPath := "music/" + musicName
	os.MkdirAll("music", 0755)

	cmd := exec.Command("python3", "scripts/generate_music.py", csvPath, musicPath)
	err = cmd.Run()
	if err != nil {
		fmt.Println("❌ Music generation error:", err)
		musicPath = ""
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": story,
		"plot":    "/plots/" + plotName,
		"music":   "/" + musicPath,
	})

	collection := plantDB.Collection("sessions")
	_, err = collection.InsertOne(context.TODO(), map[string]interface{}{
		"nickname":    nickname,
		"genericName": genericName,
		"date":        date,
		"story":       story,
		"plot":        "/plots/" + plotName,
		"music":       "/" + musicPath,
		"createdAt":   time.Now(),
	})
	if err != nil {
		fmt.Println("❌ Failed to save session to DB:", err)
	} else {
		fmt.Println("✅ Session saved to MongoDB")
	}

}
func generateMusicHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(w)
	fmt.Println("🎶 /generate-music called")

	if r.Method != "POST" {
		http.Error(w, "Only POST supported", http.StatusMethodNotAllowed)
		return
	}

	r.ParseMultipartForm(10 << 20)
	file, handler, err := r.FormFile("file")
	if err != nil {
		fmt.Println("❌ Music: file upload error:", err)
		http.Error(w, "File upload error", http.StatusBadRequest)
		return
	}
	defer file.Close()

	os.MkdirAll("uploads", 0755)
	os.MkdirAll("music", 0755)

	tempPath := "uploads/" + handler.Filename
	out, err := os.Create(tempPath)
	if err != nil {
		http.Error(w, "Failed to save uploaded CSV", http.StatusInternalServerError)
		return
	}
	defer out.Close()
	io.Copy(out, file)

	timestamp := time.Now().Unix()
	midiFile := fmt.Sprintf("music/music_%d.mid", timestamp)
	mp3File := strings.Replace(midiFile, ".mid", ".mp3", 1)

	// ✅ Call Python script first
	cmd := exec.Command("C:\\Program Files\\Python312\\python.exe", "scripts/generate_music.py", tempPath, midiFile)
	output, err := cmd.CombinedOutput()
	fmt.Println("🎵 Music script output:\n", string(output))

	if err != nil {
		fmt.Println("❌ Music generation error:", err)
		http.Error(w, "Music generation failed: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// ✅ Only after successful generation, send the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"music": "/" + mp3File,
	})
}

func generatePlotCSV(signals []Signal, filename string) error {
	var hours []float64
	var means, vars, stds []float64

	for _, s := range signals {
		hours = append(hours, float64(s.Hour))
		means = append(means, s.Mean)
		vars = append(vars, s.Var)
		stds = append(stds, s.Std)
	}

	graph := chart.Chart{
		Width:  900,
		Height: 450,
		XAxis: chart.XAxis{
			Name: "Hour of Day",
		},
		YAxis: chart.YAxis{
			Name: "Signal Value",
		},
		Series: []chart.Series{
			chart.ContinuousSeries{
				Name:    "Mean Voltage",
				XValues: hours,
				YValues: means,
			},
			chart.ContinuousSeries{
				Name:    "Variance",
				XValues: hours,
				YValues: vars,
			},
			chart.ContinuousSeries{
				Name:    "Std Deviation",
				XValues: hours,
				YValues: stds,
			},
		},
	}

	f, err := os.Create("plots/" + filename)
	if err != nil {
		return err
	}
	defer f.Close()

	return graph.Render(chart.PNG, f)
}

// ------------------- Main -------------------

func main() {
	err1 := godotenv.Load()
	if err1 != nil {
		log.Fatal("❌ Error loading .env file")
	}
	connectMongo()
	http.HandleFunc("/generate-story", uploadAndGenerate)
	http.HandleFunc("/generate-music", generateMusicHandler)

	// Static file servers
	http.Handle("/plots/", http.StripPrefix("/plots/", http.FileServer(http.Dir("plots"))))
	http.Handle("/music/", http.StripPrefix("/music/", http.FileServer(http.Dir("music"))))

	fmt.Println("🌱 Server on http://localhost:5000")
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		fmt.Println("❌ Failed to start server:", err)
	}
}
