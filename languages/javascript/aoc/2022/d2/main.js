console.log("AOC2022 Day 2");

// character values
const p1_rock_character = 'A';
const p1_paper_character = 'B';
const p1_scissors_character = 'C';

const p2_rock_character = 'X';
const p2_paper_character = 'Y';
const p2_scissors_character = 'Z';

// possible point values
const rock_points = 1;
const paper_points = 2;
const scissor_points = 3;

const draw_points = 3;
const win_points = 6;

const calculate_score_for_round = (p1, p2) => {
    let final_score = 0;

    if (p1 == p1_rock_character && p2 == p2_paper_character) {
        return paper_points + win_points;
    } else if (p1 == p1_paper_character && p2 == p2_scissors_character) {
        return scissor_points + win_points;
    } else if (p1 == p1_scissors_character && p2 == p2_rock_character) {
        return rock_points + win_points;
    }

    if (p1 == p1_rock_character && p2 == p2_scissors_character) {
        return scissor_points;
    } else if (p1 == p1_paper_character && p2 == p2_rock_character) {
        return rock_points;
    } else if (p1 == p1_scissors_character && p2 == p2_paper_character) {
        return paper_points;
    }

    // draw, work out points
    if (p1 == p1_rock_character) {
        final_score += rock_points;
    } else if (p1 == p1_paper_character) {
        final_score += paper_points;
    } else {
        final_score += scissor_points;
    }

    final_score += draw_points;

    return final_score;
}

// set up total score variable
let total_score = 0;

// console.log("Score for AY: " + calculate_score_for_round('A', 'Y'));
// console.log("Score for BX: " + calculate_score_for_round('B', 'X'));
// console.log("Score for CZ: " + calculate_score_for_round('C', 'Z'));

// read the file 
// first get the input file read
const filename = "input.txt";
const file = Bun.file(filename);
const text = await file.text();

// split the text into elements
const elements = text.split("\n");

for (let el of elements) {
    let tmp = el.split(" ");
    total_score += calculate_score_for_round(tmp[0], tmp[1]);
}
// total_score += calculate_score_for_round('A', 'Y'); console.log(total_score);
// total_score += calculate_score_for_round('B', 'X'); console.log(total_score);
// total_score += calculate_score_for_round('C', 'Z'); console.log(total_score);


console.log("The total score at the end would be " + total_score);