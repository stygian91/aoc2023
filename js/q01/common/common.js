import { Option, Result  } from 'js-cordyceps';
import { readFileSync } from 'fs';

export function firstDigit(line) {
  for (let i = 0; i < line.length; i++) {
    const ch = line[i];
    const digitOpt = toDigit(ch);
    if (digitOpt.isSome()) {
      return digitOpt;
    }
  }

  return Option.makeNone();
}

export function lastDigit(line) {
  for (let i = line.length - 1; i >= 0; i--) {
    const ch = line[i];
    const digitOpt = toDigit(ch);
    if (digitOpt.isSome()) {
      return digitOpt;
    }
  }

  return Option.makeNone();
}

export function toDigit(ch) {
  switch (ch) {
    case '0':
      return Option.make(0);
    case '1':
      return Option.make(1);
    case '2':
      return Option.make(2);
    case '3':
      return Option.make(3);
    case '4':
      return Option.make(4);
    case '5':
      return Option.make(5);
    case '6':
      return Option.make(6);
    case '7':
      return Option.make(7);
    case '8':
      return Option.make(8);
    case '9':
      return Option.make(9);

    default:
      return Option.makeNone();
  }
}

export function parseFile(path, parseLine) {
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

export function sum(numbers) {
  return numbers.reduce((acc, curr) => acc + curr, 0);
}

const digitMap = {
  zero: 0,
  one: 1,
  two: 2,
  three: 3,
  four: 4,
  five: 5,
  six: 6,
  seven: 7,
  eight: 8,
  nine: 9,
};

export function getSpeltDigit(str) {
  return Option.make(digitMap[str]);
}
