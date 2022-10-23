import { useNavigate } from "react-router-dom";
import { useState, useEffect } from "react";
import { AiFillHome } from "react-icons/ai";

import Header from "./Header";
import SharedGallery from "./SharedGallery";

const Favorites = () => {
  const [favorites, setFavorites] = useState([]);

  const navigate = useNavigate();

  const homeIcon = (
    <AiFillHome
      size="3em"
      style={{ position: "absolute", left: "5%", top: "5%" }}
    ></AiFillHome>
  );

  useEffect(() => {
    let item = window.localStorage.getItem("favorites");
    setFavorites(item ? JSON.parse(item) : []);
  }, []);

  return (
    <>
      <button className="nav-button" onClick={() => navigate("/")}>
        {homeIcon}
      </button>
      <Header title="Favorite Chonks" />
      <SharedGallery dogs={favorites} handleSetFavorites={setFavorites} />
    </>
  );
};

export default Favorites;
