import numbers from "./numbers.json";

const hashMap = {};

console.clear();

const findMult = () => {
  for (const n of numbers) {
    hashMap[n] = n;
    const other = 2020 - n;
    if (hashMap[other]) {
      console.log(n, other);
      return other * n;
    }
  }
};

document.body.innerHTML = findMult();
