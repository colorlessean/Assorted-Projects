// Question from: https://www.reddit.com/r/dailyprogrammer/comments/onfehl/20210719_challenge_399_easy_letter_value_sum/

// just a symbol to debug the code
const VERBOSITY = false;

// You are not allowed to put spaces in anyways...
function getLetterValueSum(word) {
    if (word == 'NaN') return 0;
    // need to map each character with a value of 1-26
    // then sum each value into the value 26
    let sum = 0;
    for (let i = 0; i < word.length; i++) {
        // get the ASCII value of the letters
        let val = word[i].charCodeAt(0);
        if (VERBOSITY) console.log(word[i], val);
        // if its upper case we need to adjust back by 64
        if (val >= 65 && val <= 90) {
            sum += (val - 64);
        }
        // if its lower case we need to adjust back by 96
        else if (val >= 90 && val <= 122) {
            sum += (val - 96);
        }
        else {
            throw 'InvalidCharacter';
        }
    }
    return sum;
}


// shows that the uppercase ASCII start at 65 and go to 90
// lowercase ASCII start at 97 and go to 122
// console.log(getLetterValueSum("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"));
console.log(getLetterValueSum(""));
console.log(getLetterValueSum("a"));
console.log(getLetterValueSum("z"));
console.log(getLetterValueSum("cab"));
console.log(getLetterValueSum("excellent"));
console.log(getLetterValueSum("microspectrophotometries"));

