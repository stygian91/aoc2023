import { Option, Result, flipOptionList } from 'js-cordyceps';
import { parseFile, sum, toDigit, getSpeltDigit } from '../common/common.js';

function firstSpeltDigit(line) {
  let left = 0;
  let right = 3;

  while (left < line.length - 1) {
    const part = line.slice(left, right);
    const digitOpt = getSpeltDigit(part);
    if (digitOpt.isSome()) {
      return digitOpt.map(digit => ({ index: left, digit }));
    }

    right++;
    if (right > line.length) {
      left++;
      right = left + 3;
    }
  }

  return Option.makeNone();
}

function lastSpeltDigit(line) {
  let right = line.length;
  let left = right - 3;

  while (right >= 3) {
    const part = line.slice(left, right);
    const digitOpt = getSpeltDigit(part);
    if (digitOpt.isSome()) {
      return digitOpt.map(digit => ({ index: right - 1, digit }));
    }

    left--;
    if (left < 0) {
      right--;
      left = right - 3;
    }
  }

  return Option.makeNone();
}

function firstDigitChar(line) {
  for (let i = 0; i < line.length; i++) {
    const ch = line[i];
    const digitOpt = toDigit(ch);
    if (digitOpt.isSome()) {
      return digitOpt.map(digit => ({ index: i, digit }));
    }
  }

  return Option.makeNone();
}

function cmpDigitOptions(optA, optB, cmpFn) {
  if (optA.isNone() && optB.isNone()) {
    return Option.makeNone();
  } else if (optA.isNone()) {
    return optB;
  } else if (optB.isNone()) {
    return optA;
  }

  return optA.map(valueA => {
    const valueB = optB.unwrap();
    return cmpFn(valueA, valueB) ? valueA : valueB;
  });
}

function firstDigit(line) {
  const firstSpeltOpt = firstSpeltDigit(line);
  const firstDigitOpt = firstDigitChar(line);

  return cmpDigitOptions(firstSpeltOpt, firstDigitOpt, (optA, optB) => optA.index < optB.index);
}

function lastDigit(line) {
  const lastDigitOpt = firstDigitChar(line.split('').reverse().join(''))
    .map(found => ({
      digit: found.digit,
      index: line.length - 1 - found.index
    }));
  const lastSpeltOpt = lastSpeltDigit(line);

  return cmpDigitOptions(lastDigitOpt, lastSpeltOpt, (optA, optB) => optA.index > optB.index);
}

function parseLine(line) {
  let first = firstDigit(line);
  let last = lastDigit(line);

  if (first.isNone() || last.isNone()) {
    return Result.makeErr(new Error('No digit found.'));
  }

  first = first.unwrap();
  last = last.unwrap();

  return Result.makeOk((first.digit * 10) + last.digit);
}

console.log(parseFile('./data/data.txt', parseLine).map(sum).unwrap());
