import "./App.css";

function App() {
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
          <img src="https://github.com/tumGER/Kosmonaut/blob/main/src/assets/logo.png?raw=true" />
          <div
            className="
            flex
            flex-col
            justify-center
            justify-items-center
            gap-6
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
                place-content-center
                items-center
                gap-2
                self-center
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
                // eslint-disable-next-line @typescript-eslint/no-unused-vars
                onClick={_ => window.open("https://raw.githubusercontent.com/tumGER/Kosmonaut/gh-pages/assets/bundle.zip")}
              >
                Download
              </button>
            <div
              className="
          text-center
          text-gray-500
           "
            >
              <p className="text-lg font-semibold text-gray-400">Credits:</p>
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
