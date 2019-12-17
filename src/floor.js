console.time("floor");
for (let index = 0; index < 100000; index++) {
    let r = 5.123>>0;
}
console.timeEnd("floor");