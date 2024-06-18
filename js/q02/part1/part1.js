import { parseFile, getGameMaxes } from "../common/index.js";

const maxCubes = {
  red: 12,
  green: 13,
  blue: 14,
};

function isValid(gameMaxes) {
  return gameMaxes.red <= maxCubes.red
    && gameMaxes.green <= maxCubes.green
    && gameMaxes.blue <= maxCubes.blue;
}

function part1(path) {
  const parseFileResult = parseFile(path);
  if (parseFileResult.isErr()) {
    console.error(`Parse file error: ${parseFileResult.unwrap()}`);
    return;
  }

  const validGames = parseFileResult.unwrap().filter(game => {
    const maxes = getGameMaxes(game);
    return isValid(maxes);
  });

  const idSum = validGames.reduce((acc, curr) => acc + curr.number, 0);
  console.log(`Part 1: ${idSum}`);
}

part1('./data/input.txt');
