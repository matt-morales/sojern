import { useNavigate } from "react-router-dom";
import { useEffect, useState } from "react";
import { getDogs } from "../services/dogService";
import Dog from "./Dog";
import Header from "./Header";
import SharedGallery from "./SharedGallery";
import { AiFillHeart } from "react-icons/ai";
import { IoMdRefreshCircle } from "react-icons/io";

import "../styles/dogs.css";

const Gallery = () => {
  const [dogs, setDogs] = useState([]);
  const [bigDog, setBigDog] = useState();

  const navigate = useNavigate();

  useEffect(() => {
    const fetchDogs = async () => {
      const dogs = await getDogs();
      setDogs(dogs);
      setBigDog(dogs[0]);
    };

    fetchDogs();
  }, []);

  const handleOnClick = (dog) => {
    setBigDog(dog);
  };

  const likeIcon = (
    <AiFillHeart
      size="3em"
      style={{ position: "absolute", color: "red", left: "5%", top: "5%" }}
    ></AiFillHeart>
  );

  const refreshIcon = (
    <IoMdRefreshCircle
      size="3em"
      style={{ position: "absolute", color: "black", right: "5%", top: "5%" }}
    ></IoMdRefreshCircle>
  );

  const handleRefresh = () => {
    const fetchDogs = async () => {
      const dogs = await getDogs();
      setDogs(dogs);
      setBigDog(dogs[0]);
    };

    fetchDogs();
  };

  return (
    <>
      <button className="nav-button" onClick={() => navigate("favorites")}>
        {likeIcon}
      </button>
      <Header title="Chonk Gallery"></Header>
      <button className="nav-button" onClick={handleRefresh}>
        {refreshIcon}
      </button>
      <SharedGallery dogs={dogs} handleOnClick={handleOnClick} />
      <div className="bigDog">
        <Dog dogSrc={bigDog} />
      </div>
    </>
  );
};

export default Gallery;
