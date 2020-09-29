var numbers = [0, 1, 3, 4, 5, 7, 8]; // Missing 2,6
var missing = [];

// Find the missing array items
for (var i = 0; i < numbers.length; i++) {

    if ((numbers[i + 1] - numbers[i]) > 1) {
        missing.push(numbers[i + 1] - numbers[1]);
    }
}

console.log(missing);