import { useState } from 'react';

function GeneratePage() {
  const [story, setStory] = useState('');
  const [musicLink, setMusicLink] = useState('');
  const [file, setFile] = useState(null);
  const [plantName, setPlantName] = useState('');
  const [nickname, setNickname] = useState('');
  const [date, setDate] = useState('');
  const [showImage, setShowImage] = useState(false);
  const [plotPath, setPlotPath] = useState('');
  const [musicLoading, setMusicLoading] = useState(false);
  const [activeTab, setActiveTab] = useState('story');

  // 🌱 Handle Story Generation
  const handleUpload = async (e) => {
    e.preventDefault();

    if (!file || !plantName || !nickname || !date) {
      alert("Please fill all fields and select a CSV file.");
      return;
    }

    const formData = new FormData();
    formData.append("file", file);
    formData.append("genericName", plantName);
    formData.append("nickname", nickname);
    formData.append("date", date);

    try {
      const res = await fetch("http://localhost:5000/generate-story", {
        method: "POST",
        body: formData,
      });

      if (!res.ok) throw new Error("Failed to fetch story from backend.");

      const data = await res.json();
      setStory(data.message);
      setPlotPath("http://localhost:5000" + data.plot);
      setShowImage(true);
    } catch (error) {
      console.error("❌ Error:", error);
      alert("Something went wrong while generating the story.");
    }
  };

  // 🎵 Handle Music Generation
  const handleMusicGenerate = async () => {
    if (!file) {
      alert("Please upload a CSV file first.");
      return;
    }

    setMusicLoading(true);

    const formData = new FormData();
    formData.append("file", file);

    try {
      const res = await fetch("http://localhost:5000/generate-music", {
        method: "POST",
        body: formData,
      });

      if (!res.ok) throw new Error("Failed to fetch music from backend.");

      const data = await res.json();
      const fullMusicPath = "http://localhost:5000" + data.music;
      setMusicLink(fullMusicPath);
    } catch (error) {
      console.error("🎵 Music Error:", error);
      alert("Failed to generate music.");
    } finally {
      setMusicLoading(false);
    }
  };

  return (
    <div className="space-y-6">
      <h1 className="text-3xl font-bold">
        Nature has a voice. You just need to listen.
      </h1>
      <br />
      <br />

      {/* Tabs */}
      <div className="flex gap-4 border-b border-[#d6e7d0] mb-6">
        
         <button 
                type="submit" 
                className="btn px-8 py-3 bg-green-600 text-white rounded-lg hover:bg-green-700 transition-colors"
                onClick={() => setActiveTab('story')}
              >
                Generate Story
        </button>
        {/* Spacer between buttons */}
        
        

       
        <button 
                type="submit" 
                className="btn px-8 py-3 bg-green-600 text-white rounded-lg hover:bg-green-700 transition-colors"
                onClick={() => setActiveTab('music')}
              >
                Generate Music
        </button>
      </div>
<br />
      {activeTab === 'story' && (
        <div className="card">
          <h2 className="text-2xl font-bold mb-6">Generate Story</h2>
          
          <form onSubmit={handleUpload} className="space-y-6">
            <div className="grid grid-cols-2 gap-4">
              <div className="input-group">
                <label>Nickname</label>
                <input
                  type="text"
                  placeholder="Enter your nickname"
                  value={nickname}
                  onChange={(e) => setNickname(e.target.value)}
                  className="w-full p-2 border border-gray-300 rounded"
                />
              </div>
              
              <div className="input-group">
                <label>Plant Name</label>
                <input
                  type="text"
                  placeholder="Enter the plant's generic name"
                  value={plantName}
                  onChange={(e) => setPlantName(e.target.value)}
                  className="w-full p-2 border border-gray-300 rounded"
                />
              </div>
            </div>
            
            <div className="grid grid-cols-2 gap-4">
              <div className="input-group">
                <label>Date</label>
                <input
                  type="date"
                  value={date}
                  onChange={(e) => setDate(e.target.value)}
                  className="w-full p-2 border border-gray-300 rounded"
                />
              </div>
              
              <div className="input-group">
                <label>Upload File</label>
                <input
                  type="file"
                  accept=".csv"
                  onChange={(e) => setFile(e.target.files[0])}
                  className="w-full p-2"
                />
              </div>
            </div>
            
            <div className="text-right">
              <button 
                type="submit" 
                className="btn px-8 py-3 bg-green-600 text-white rounded-lg hover:bg-green-700 transition-colors"
              >
                Generate Story
              </button>
            </div>
          </form>

          {story && (
            <div className="mt-8 bg-[#f0fff4] rounded-xl p-6 border border-[#d6e7d0]">
              <br />
              <h3 className="text-xl font-bold mb-4">Generated Story</h3>
              <br />
              <p className="text-base">{story}</p>
            </div>
          )}

          {showImage && (
            <div className="mt-8">
              <h3 className="text-xl font-bold mb-4">Signal Visualization</h3>
              <img
                src={plotPath}
                alt="Plant Signal Plot"
                className="w-full max-w-2xl rounded-xl border border-[#d6e7d0]"
              />
            </div>
          )}
        </div>
      )}

      {activeTab === 'music' && (
        <div className="card">
          <h2 className="text-2xl font-bold mb-6">Generate Music</h2>
          
          <div className="input-group">
            <label>Upload File</label>
            <input
              type="file"
              accept=".csv"
              onChange={(e) => setFile(e.target.files[0])}
              className="w-full p-2"
            />
          </div>
          
          {file && (
            <div className="text-right mt-4">
              <button 
                onClick={handleMusicGenerate} 
                disabled={musicLoading}
                className="btn px-8 py-3 bg-green-600 text-white rounded-lg hover:bg-green-700 transition-colors disabled:opacity-50"
              >
                {musicLoading ? "⏳ Generating..." : "🎶 Generate Music"}
              </button>
              <br />
            </div>
          )}
          
          {musicLink && (
            <div className="mt-8">
              <br />
              <h3 className="text-xl font-bold mb-4">🎧 Generated Music</h3>
              
              <div className="bg-[#eaf3e7] rounded-xl p-6 border border-[#d6e7d0]">
                <audio 
                  key={musicLink} 
                  controls 
                  src={musicLink} 
                  className="w-full max-w-md mb-4"
                />
                
                <div className="text-center">
                  <a 
                    href={musicLink} 
                    download 
                    className="inline-flex items-center justify-center rounded-full h-10 px-6 text-sm font-bold bg-white text-[#111b0e] border border-[#d6e7d0] hover:bg-[#f0f8ec] transition-colors"
                  >
                    <br />
                    ⬇️ Download Music
                  </a>
                </div>
              </div>
            </div>
          )}
        </div>
      )}
    </div>
  );
}

export default GeneratePage;