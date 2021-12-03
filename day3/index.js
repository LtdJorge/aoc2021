const fs = require("fs");

// let inputFile = fs.readFileSync("test-input.txt", "utf-8");
let inputFile = fs.readFileSync("input.txt", "utf-8");

let oxygenGenRating;
let CO2ScrubberRating;

const lines = inputFile.split(/\r?\n/);

let splitLines = []

lines.forEach((line) => {
    splitLines.push(line.split(""));
});


function findMostCommon(position, lines){
    let oneCounter = 0;
    let zeroCounter = 0;
    lines.forEach((line) => {
        line[position] === '0' ? zeroCounter++ : oneCounter++
    })
    return zeroCounter > oneCounter ? 0 : 1;
}

function findLeastCommon(position, lines){
    let oneCounter = 0;
    let zeroCounter = 0;
    lines.forEach((line) => {
        line[position] === '0' ? zeroCounter++ : oneCounter++
    })
    return zeroCounter <= oneCounter ? 0 : 1;
}

function getBitCriteria(position, getOxy, lines){
    return getOxy
        ? findMostCommon(position, lines).toString()
        : findLeastCommon(position, lines).toString()
}

let mostCommonBuffer = []
let leastCommonBuffer = []

for (let i=0; i<splitLines[0].length; i++){
    const common = findMostCommon(i, splitLines);
    mostCommonBuffer.push(common);
    common===1 ? leastCommonBuffer.push('0') : leastCommonBuffer.push('1');
}

let gammaRate = mostCommonBuffer.join("");
let epsilonRate = leastCommonBuffer.join("");

console.log(`Gamma: ${gammaRate} \nEpsilon: ${epsilonRate}`);
console.log(`Power consumption: ${parseInt(gammaRate, 2)*parseInt(epsilonRate, 2)}`)

function findOxyRating(position, lines){
    if (lines.length === 1) {
        return lines[0].join("");
    }
    let filteredLines = lines.filter((line) => line[position] === getBitCriteria(position, true, lines))
    return findOxyRating(++position, filteredLines);
}

function findCO2Rating(position, lines){
    if (lines.length <= 1) {
        return lines[0].join("");
    }
    let filteredLines = lines.filter((line) => line[position] === getBitCriteria(position, false, lines))
    return findCO2Rating(++position, filteredLines);
}
oxygenGenRating = findOxyRating(0, splitLines);
CO2ScrubberRating = findCO2Rating(0, splitLines);
console.log(`O2: ${oxygenGenRating}`);
console.log(`CO2: ${CO2ScrubberRating}`);

let lifeSupportRating = parseInt(oxygenGenRating, 2) * parseInt(CO2ScrubberRating, 2);
console.log(`Life support: ${lifeSupportRating}`)