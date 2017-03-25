var current_time = new Date().getTime();

console.log("********* Bechmark start *********");
var largest_number = 0;
for(var i = 0; i < 10000000; i++) {
  if (largest_number < i) {
    largest_number = i;
  }
}
duration = new Date().getTime() - current_time;
console.log("It takes "+duration+" ms\n**********************************\n");


// ********* Bechmark start *********
// It takes 29 ms
// **********************************