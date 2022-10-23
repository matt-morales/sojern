export async function getDogs() {
  const urls = Array(6).fill("https://random.dog/woof.json");
  const dogs = await Promise.all(
    urls.map((url) =>
      fetch(url)
        .then((res) => res.json())
        .then((obj) => obj.url)
    )
  );
  return dogs;
}

export async function getDog() {
  const url = "https://random.dog/woof.json";
  const dog = await fetch(url).then((res) => res.json());
  return dog.url;
}
