import { parseFile } from "../common/common.js";

function part1(path) {
  const parseResult = parseFile(path);
  if (parseResult.isErr()) {
    console.error(parseResult.unwrap());
    return;
  }

  const schema = parseResult.unwrap();
  let sum = 0;

  for (const partNumber of schema.numbers) {
    if (hasSymbolNeighbour(partNumber, schema.symbols)) {
      sum += parseInt(partNumber.number);
    }
  }

  console.log(`Part 1: ${sum}`);
}

function hasSymbolNeighbour(partNumber, symbols) {
  const start = { x: partNumber.start.x - 1, y: partNumber.start.y - 1 };
  const end = { x: partNumber.end.x + 1, y: partNumber.end.y + 1 };

  for (let x = start.x; x <= end.x; x++) {
    for (let y = start.y; y <= end.y; y++) {
      if (x >= partNumber.start.x && x <= partNumber.end.x && y >= partNumber.start.y && y <= partNumber.end.y) {
        continue;
      }

      const key = `${x}-${y}`;
      if (symbols.hasOwnProperty(key)) {
        return true;
      }
    }
  }

  return false;
}

part1('./data/input.txt');
