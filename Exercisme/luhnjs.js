

let a = [3, 5, 10];

let sum = 0;
for (let index = 0; index < a.length; index++) {
  let element = a[index];

  if (a[index] === 5) {
    element *= 2;
  }

  sum += element;
  
}

console.log(sum);