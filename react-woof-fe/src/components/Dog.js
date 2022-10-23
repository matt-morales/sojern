import { useState, useEffect } from "react";
import { AiFillHeart, AiOutlineHeart } from "react-icons/ai";
import "../styles/dogs.css";

const Dog = ({ dogSrc, handleSetFavorites }) => {
  const [isFavorite, setIsFavorite] = useState(false);

  var isImage = true;

  if (dogSrc) {
    isImage = !["mp4", "webm"].some((s) => dogSrc.endsWith(s));
  }

  useEffect(() => {
    let item = window.localStorage.getItem("favorites");
    let favorites = item ? JSON.parse(item) : [];
    setIsFavorite(favorites.includes(dogSrc));
  }, [dogSrc]);

  const handleOnClick = (e) => {
    let item = window.localStorage.getItem("favorites");
    let favorites = item ? JSON.parse(item) : [];
    const newIsFavorite = !isFavorite;
    setIsFavorite(newIsFavorite);

    if (newIsFavorite && !favorites.includes(dogSrc)) {
      let newFavorites = [...favorites, dogSrc];
      handleSetFavorites && handleSetFavorites(newFavorites);
      window.localStorage.setItem("favorites", JSON.stringify(newFavorites));
    } else if (favorites && !newIsFavorite && favorites.includes(dogSrc)) {
      let newFavorites = favorites.filter((src) => {
        return !(src === dogSrc);
      });
      handleSetFavorites && handleSetFavorites(newFavorites);
      window.localStorage.setItem("favorites", JSON.stringify(newFavorites));
    }

    e.stopPropagation();
  };

  return (
    <>
      {!isImage ? <video src={dogSrc} alt="" /> : <img src={dogSrc} alt="" />}
      <button className="button" onClick={handleOnClick}>
        {isFavorite ? (
          <AiFillHeart size="2em" style={{ color: "red" }}></AiFillHeart>
        ) : (
          <AiOutlineHeart
            size="2em"
            style={{
              top: 0,
              right: 0,
              color: "red",
              stroke: "white",
              strokeWidth: "5",
            }}
          />
        )}
      </button>
    </>
  );
};

export default Dog;
