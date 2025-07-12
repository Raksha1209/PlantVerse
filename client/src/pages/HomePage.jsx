import { Link } from 'react-router-dom';

function HomePage() {
  return (
    <div className="space-y-8">
      
      <div 
        className="relative rounded-2xl overflow-hidden min-h-[500px] flex items-center justify-center"
        style={{ 
          background: 'linear-gradient(rgba(0, 0, 0, 0.3), url("https://lh3.googleusercontent.com/aida-public/AB6AXuAkcdrJC5BitW-cQz08zpUwavRJmXChOI0ap7ypYAlK1ZjZGH3629fQ5endNTMg7rDx6bW-y1JVKYDac2BpoCI03p1-EaVExpELY2wb6QsYR2WXeq9B-AJJvFnHZA6IQ_p6q0nhscZWRE5alucPAzUKYblfCkudPorMBupziOvmM6DN_AZT_6cjcrXySEheXYHpAfI_G89gep-Nzj4MjQLgyIqMCxbsay8nl6qFWwkVg-cc1ZjmStrSND8pQEc-5X9Kker4Eqq2BLc")',
          backgroundSize: 'cover',
          backgroundPosition: 'center'
        }}
      >
        <br />
        <div className="text-center max-w-2xl p-6">
          <h1 className="text-white text-4xl md:text-5xl font-bold mb-6 pl-6">
            Can I have a moment of talk with you, or do you want me to sing?
          </h1>
          <br />
          <div style={{ display: 'flex', justifyContent: 'center' }}>
          <Link
            to="/generate"
            className="btn text-lg px-8 py-3"
            style={{ backgroundColor: '#60974e', color: '#fff' }}
          >
            Get Started
          </Link>
        </div>
          <br />
        </div>
      </div>
        <br />
      {/* About Section */}
      <div className="card">
        <h2 className="text-2xl font-bold mb-4">About PlantVerse</h2>
        <p className="mb-6">
          A platform where your plants will talk and sing to you! How cool is that?
        </p>
        <div className="text-center">
          <Link to="/about" className="btn" style={{ backgroundColor: '#60974e', color: '#fff'}}>
            Learn More
          </Link>
        </div>
      </div>

      {/* Generate Options */}
      <div className="card">
        <h2 className="text-2xl font-bold mb-6">Create with Nature</h2>
        <div className="grid-cols-2">
          {/* Story Card */}
          <div className="bg-[#f9fcf8] rounded-xl p-6 border border-[#d6e7d0]">
            <div 
              className="w-full h-48 bg-cover bg-center rounded-xl mb-4"
              style={{ backgroundImage: 'url("https://lh3.googleusercontent.com/aida-public/AB6AXuDkyNy97haLrz7gBLIjiz87Z7Wa_OMogusDapewmdHQyWjIPKbGnIBif8LRI-dKIV60X0LYkCFxoh5bbB3RjckO6p7Fcapc_K1m05e6z1r1a3TQUaTc_wEdfMdH-sR7cm7MjRJprFENkaouB2WitiRUEIl4JZt-sAYObPi1tYP-S1dgfQwq5q1NrEq4wrHgaaCkoufomHGdSIS0QAkSSRNLh6CW2p48wndADNZmjJzPczlGcKpZBEblY1b7UkqGORT7F4i4hZEXYc4")' }}
            ></div>
            <h3 className="text-xl font-bold mb-2">Generate Story</h3>
            <p className="text-[#60974e] mb-4">
              Create a unique story inspired by your favorite plant.
            </p>
            <Link to="/generate" className="btn w-full" style={{ backgroundColor: '#60974e', color: '#fff'}}>
              Generate Story
            </Link>
          </div>
          
          {/* Music Card */}
          <div className="bg-[#f9fcf8] rounded-xl p-6 border border-[#d6e7d0]">
            <div 
              className="w-full h-48 bg-cover bg-center rounded-xl mb-4"
              style={{ backgroundImage: 'url("https://lh3.googleusercontent.com/aida-public/AB6AXuBehZ7jIksTTecWUGKIn0QVcXTaQ6Uek0JjDXpRBrLOnIFfbaK55UPP9ejdKjoxSFPR5rP_YjBmYePUgXE6AB5LQVxo_P3aR8ngYSoiW1UYbZbh_sfhSNOV7srqVftHzRW49xbnehNFwao8Lq71CQ7hjJUwN8cMLAxZ6LK2bbFUtIldLn7F6sustYeOe5ZKiMYT06kYGzlxQbMWqnXbZR60PxgQPUXef7Va-o9uEgOUp0xq2nyvrJe5ZcNJkS9J0vzDHJD5IoPV5LA")' }}
            ></div>
            <h3 className="text-xl font-bold mb-2">Generate Music</h3>
            <p className="text-[#60974e] mb-4">
              Create soothing music inspired by your favorite plant.
            </p>
            <Link to="/generate" className="btn w-full" style={{ backgroundColor: '#60974e', color: '#fff'}}>
              Generate Music
            </Link>
          </div>
        </div>
      </div>
    </div>
  );
}

export default HomePage;