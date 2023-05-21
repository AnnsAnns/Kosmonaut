import { useState } from "react";
import "./App.css";

function App() {
  const [count, setCount] = useState(0);

  return (
    <div
      className="absolute grid min-h-full min-w-full grid-cols-3 content-center justify-items-center bg-no-repeat"
      style={{
        backgroundImage: `url(https://cdn.discordapp.com/attachments/993141623443685376/1109930153469034736/image.png)`,
      }}
    >
      <div></div>
      <div
        className="
          flex
          flex-col
        "
      >
        <div
          className="
          grid
          grow
          grid-rows-2
          content-center
          justify-items-center
        "
        >
          <img src="https://cdn.discordapp.com/attachments/993141623443685376/1109932676938465340/wordart2.png" />
          <button
            className="
          group 
          flex
          h-16
          w-2/4
          place-content-center
          items-center
          gap-2
          whitespace-nowrap
          rounded
          bg-red-500
          fill-slate-300
          px-4
          text-center
          text-lg
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
        </div>
        <div
          className="
      self-center
      text-center
      text-gray-500
      "
        >
          <p>Credits:</p>
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
            <a
              href="#"
              className="text-gray-400"
            >
              Bernv3
            </a>{" "}
            for the assets
          </p>
        </div>
      </div>
    </div>
  );
}

export default App;
