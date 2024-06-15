/**
 * External imports:
 */
import { Result } from 'js-cordyceps';
/**
 * Internal imports:
 */
import { firstDigit, lastDigit, parseFile, sum } from '../common/common.js';

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

const answerResult = parseFile('./data/data.txt', parseLine)
  .map(sum);

console.log(answerResult.unwrap());
