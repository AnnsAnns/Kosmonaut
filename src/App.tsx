import { useState } from "react";
import "./App.css";

function App() {
  const [count, setCount] = useState(0);

  return (
    <div
      className="absolute grid min-h-full min-w-full grid-cols-3 content-center bg-no-repeat"
      style={{
        backgroundImage: `url(https://github.com/tumGER/Kosmonaut/blob/main/src/assets/Galaxy.png?raw=true)`,
        backgroundColor: "#17172c",
      }}
    >
      <div></div>
      <div
        className="
          flex
          flex-col
        "
      >
        <div>
          <img src="https://cdn.discordapp.com/attachments/993141623443685376/1109932676938465340/wordart2.png" />
          <div
            className="
            flex
            flex-col
            gap-6
            justify-center
            justify-items-center
            rounded-3xl
            bg-slate-800
            bg-opacity-50
            py-8
            text-center
            text-slate-400
            ring
            ring-slate-900
          "
          >
            <div>
              <p className="text-2xl text-slate-300">
                Atmosphere ✕ Hekate Bundler
              </p>
            </div>
            <div>
            <p className="text-lg font-semibold">Setup:</p>
              <p>1. Download ZIP</p>
              <p>2. Extract ZIP onto your SD</p>
              <p>3. Launch Hekate</p>
            </div>
            <button
              className="
          h-28
          w-72
          self-center
          place-content-center
          items-center
          gap-2
          rounded-lg
          bg-red-500
          fill-slate-300
          text-center
          text-2xl
          font-semibold
          text-white
          hover:bg-red-600
          focus:bg-red-600 
          focus:outline-none
          focus:ring-2
          focus:ring-red-300
          active:bg-red-800
        "
            >
              Download
            </button>
            <div
              className="
          text-center
          text-gray-500
           "
            >
              <p className="text-lg text-gray-400 font-semibold">Credits:</p>
              <p>
                <a
                  href="https://git.nicholemattera.com/NicholeMattera"
                  className="text-gray-400"
                >
                  NicholeMattera
                </a>{" "}
                for the Builder
              </p>
              <p>
                <a
                  href="https://github.com/Atmosphere-NX/Atmosphere"
                  className="text-gray-400"
                >
                  Team Atmosphère
                </a>{" "}
                for Atmosphere
              </p>
              <p>
                <a
                  href="https://github.com/CTCaer/hekate"
                  className="text-gray-400"
                >
                  CTCaer
                </a>{" "}
                for Hekate
              </p>
              <p>
                <a href="#" className="text-gray-400">
                  Bernv3
                </a>{" "}
                for the assets
              </p>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}

export default App;
