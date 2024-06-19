import { parseFile } from "../common/common.js";

function isNeighbour(partNumber, pos) {
  const posParts = pos.split('-');
  const posX = parseInt(posParts[0]);
  const posY = parseInt(posParts[1]);

  for (let y = posY - 1; y <= posY + 1; y++) {
    if (y !== partNumber.start.y) {
      continue;
    }

    for (let x = posX - 1; x <= posX + 1; x++) {
      if (x >= partNumber.start.x && x <= partNumber.end.x) {
        return true;
      }
    }
  }

  return false;
}

function part2(path) {
  const parseResult = parseFile(path);
  if (parseResult.isErr()) {
    console.error(parseResult.unwrap());
    return;
  }

  let sum = 0;

  const schema = parseResult.unwrap();
  for (const entry of Object.entries(schema.symbols)) {
    const pos = entry[0];
    const symbol = entry[1];

    if (symbol !== '*') {
      continue;
    }

    const neighbours = [];

    for (const partNumber of schema.numbers) {
      if (isNeighbour(partNumber, pos)) {
        neighbours.push(partNumber);
      }
    }

    if (neighbours.length !== 2) {
      continue;
    }

    const power = parseInt(neighbours[0].number) * parseInt(neighbours[1].number);
    sum += power;
  }

  console.log(`Part 2: ${sum}`);
}

part2('./data/input.txt');
