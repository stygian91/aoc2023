import { readFileSync } from 'fs';
import { Result } from 'js-cordyceps';

function isDigit(ch) {
  const chCode = ch.charCodeAt(0);

  return chCode >= '0'.charCodeAt(0) && chCode <= '9'.charCodeAt(0);
}

export function parseFile(path) {
  const readResult = Result.try(readFileSync, [path, { encoding: 'utf8' }]);
  if (readResult.isErr()) {
    return readResult;
  }

  const numbers = [];
  const symbols = {};
  const lines = readResult.unwrap().trim().split('\n');
  let number = '';
  let partNumber = {};

  function addNum() {
    partNumber.number = number;
    numbers.push(partNumber);
    partNumber = {};
    number = '';
  }

  for (let i = 0; i < lines.length; i++) {
    const line = lines[i];

    for (let j = 0; j < line.length; j++) {
      const ch = line.charAt(j);

      if (isDigit(ch)) {
        if (number.length === 0) {
          partNumber = { start: { x: j, y: i } };
        }

        number += ch;
      } else {
        if (ch !== '.') {
          symbols[`${j}-${i}`] = ch;
        }

        if (number.length > 0) {
          partNumber.end = { x: j - 1, y: i };
          addNum();
        }
      }

      if (j === line.length - 1 && number.length > 0) {
        partNumber.end = { x: j - 1, y: i };
        addNum();
      }
    }
  }

  return Result.makeOk({ numbers, symbols });
}
