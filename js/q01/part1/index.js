/**
 * External imports:
 */
import { Result, flipOptionList } from 'js-cordyceps';
/**
 * Internal imports:
 */
import { firstDigit, lastDigit, parseFile, sum } from '../common/common.js';

function parseLine(line) {
  const digitsOpt = flipOptionList([firstDigit(line), lastDigit(line)]);
  if (digitsOpt.isNone()) {
    return Result.makeErr(new Error('Did not find digit in line.'));
  }

  const [first, last] = digitsOpt.unwrap();
  return Result.makeOk(first * 10 + last);
}

const answerResult = parseFile('./data/data.txt', parseLine)
  .map(sum);

console.log(answerResult.unwrap());
