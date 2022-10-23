import Dog from "./Dog";

const SharedGallery = ({ dogs, handleOnClick }) => {
  return (
    <div className="gallery">
      {dogs.map((dog) => (
        <div className="smallDogs" onClick={() => handleOnClick(dog)}>
          <Dog dogSrc={dog} />
        </div>
      ))}
    </div>
  );
};

export default SharedGallery;
