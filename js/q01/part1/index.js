import { Result } from 'js-cordyceps';
import { readFileSync } from 'fs';

import { firstDigit, lastDigit } from '../common/common.js';

function parseLine(line) {
  const firstOpt = firstDigit(line);
  const lastOpt = lastDigit(line);

  if (firstOpt.isNone() || lastOpt.isNone()) {
    return Result.makeErr(
      new Error(`Did not find digit in line. Line: ${line}; first: ${firstOpt.unwrap()}; last: ${lastOpt.unwrap()}`)
    );
  }

  return Result.makeOk(firstOpt.map(x => x * 10).unwrap() + lastOpt.unwrap());
}

function parseFile(path) {
  return Result.try(() => readFileSync(path, { encoding: 'utf8' }))
    .andThen((contents) => {
      const lines = contents.split('\n');
      const numbers = [];

      for (const line of lines) {
        if (line.length === 0) {
          continue;
        }

        const result = parseLine(line);
        if (result.isErr()) {
          return result;
        }

        numbers.push(result.unwrap());
      }

      return Result.makeOk(numbers);
    });
}

function sum(numbers) {
  return numbers.reduce((acc, curr) => acc + curr, 0);
}

const answerResult = parseFile('./data/data.txt')
  .map(sum);

console.log(answerResult.unwrap());
