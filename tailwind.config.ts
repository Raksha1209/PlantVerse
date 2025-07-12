import type { Config } from "tailwindcss";

export default {
  darkMode: "class",
  content: ["./client/index.html", "./client/src/**/*.{js,jsx,ts,tsx}"],
  theme: {
    extend: {
      colors: {
        'plant-green': '#3db714',
        'dark-green': '#111b0e',
        'light-green': '#60974e',
        'pale-green': '#eaf3e7',
        'background': '#f9fcf8',
      },
    },
  },
  plugins: [
    require("tailwindcss-animate"),
    require("@tailwindcss/typography")
  ],
} satisfies Config;
