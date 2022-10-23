import { BrowserRouter, Routes, Route } from "react-router-dom";
import Favorites from "./components/Favorites";
import Gallery from "./components/Gallery";

import "./App.css";

function App() {
  return (
    <div className="App">
      <BrowserRouter>
        <Routes>
          <Route path="/" element={<Gallery />} />
          <Route path="/favorites" element={<Favorites />} />
        </Routes>
      </BrowserRouter>
    </div>
  );
}

export default App;
