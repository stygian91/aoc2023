import { parseFile, getGameMaxes } from "../common/index.js";

function part2(path) {
  const parseFileResult = parseFile(path);
  if (parseFileResult.isErr()) {
    console.error(`Parse file error: ${parseFileResult.unwrap()}`);
    return;
  }

  const result = parseFileResult.unwrap().reduce((accumulator, game) => {
    const maxes = getGameMaxes(game);
    return accumulator + (maxes.red * maxes.green * maxes.blue);
  }, 0);

  console.log(`Part 2: ${result}`);
}

part2('./data/input.txt');
