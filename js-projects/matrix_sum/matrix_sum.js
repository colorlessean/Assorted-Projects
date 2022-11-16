// Take an nxn matrix and select 20 elements from the matrix such that no two elements come from the same row or column
// Find the minimum sum that is possible and return the elements that make up the sum
// Our matrix must be at minimum 20 elements so smallest would be a 5x5 matrix in this case
// https://www.reddit.com/r/dailyprogrammer/comments/oirb5v/20210712_challenge_398_difficult_matrix_sum/
const VERBOSE = true;

function maxSum(matrix) {
    let values = new Array(20).fill(0);
    let sum = Number.MAX_VALUE;


    if (VERBOSE) {
        console.log(sum, values);
    }
    
}

// Gets our matrix value from the file and then returns it as a 2-D array
function getMatrixFromFile() {
    
}

maxSum([]);

