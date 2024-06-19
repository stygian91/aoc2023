import { getSymbolNeighbours, parseFile } from "../common/common.js";

function part1(path) {
  const parseResult = parseFile(path);
  if (parseResult.isErr()) {
    console.error(parseResult.unwrap());
    return;
  }

  const schema = parseResult.unwrap();
  let sum = 0;

  for (const partNumber of schema.numbers) {
    const neighbours = getSymbolNeighbours(partNumber, schema.symbols)
    if (neighbours.length > 0) {
      sum += parseInt(partNumber.number);
    }
  }

  console.log(`Part 1: ${sum}`);
}

part1('./data/input.txt');
