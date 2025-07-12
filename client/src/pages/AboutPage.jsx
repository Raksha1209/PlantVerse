function AboutPage() {
  return (
    <div className="space-y-8">
      <div className="card">
        <h1 className="text-3xl font-bold mb-6">About PlantVerse</h1>
        <p>
          <br />
          Welcome to PlantVerse, a heartfelt fusion of nature and technology. This platform was born from a simple yet profound idea:<br />
          <br />
          <strong>“What if plants could tell their stories?”</strong><br />
          <br />
          Imagine a world where plants could share their experiences, emotions, and the beauty of their existence through stories and music.
          <br />
          My PlantVerse makes that dream real — by capturing the bioelectric signals of plants and translating them into beautiful stories and music using advanced AI.
          PlantVerse is where that silence is transformed into art.
        </p>
      </div>

      <div className="card">
        <h2>My Inspiration - Mother's Garden</h2>
        <p>
          PlantVerse was inspired by someone very special — my mom, whose love for plants shaped my childhood. 
          I watched her care for every leaf like it had a soul. She would talk to them, gently touch their leaves, move them into sunlight, and even sing to them.
          To her, they were family.
          That single question became the root of this project. This is my way of giving back to her — and to every plant lover who’s ever whispered to a leaf.
        </p>
      </div>

      <div className="card">
        <h2>Meet Me</h2>
        <div className="grid grid-cols-1 md:grid-cols-3 gap-6 mt-4">
          <div className="text-center">
            <div 
              className="w-40 h-40 mx-auto bg-cover bg-center rounded-full mb-4"
              style={{ backgroundImage: 'url("https://lh3.googleusercontent.com/aida-public/AB6AXuBwk9OykL_DmUuGUwsUZV32l2YIEETwXuWCmEtGIYSseNaNApkOu1cywnOlBRKkQ7YNMtPkR0IZOsGkRyiJv_8VuTD79wY7dtJnBjWT5oFZsvNb4frW8z1EWnANCRI7iSm8iD-sYG4rOkpPE4TxS-fLLXmDXe-Sy6iAFJnypqTcPGjeWbclD2A275BE4c-VrWGyRx-XXYN7g054duy5dZlSCxH0mNBrawTsIJttYVr5kJWTloYbpmPc-85GQEXuUunzuG9V_IihLwI")' }}
            ></div>
            <p className="font-bold">Raksha K L (Creator)</p>
            <p className="text-[#60974e]"></p>
            <p>A daughter, inspired by her mother's love of plants.</p>
          </div>
        </div>
      </div>
      <div className="card">
        <h2>Contact Me</h2>
        <p>
          I'd love to hear from you! If you have any questions, feedback, or suggestions, 
          
        </p>
      </div>
    </div>
  );
}

export default AboutPage;