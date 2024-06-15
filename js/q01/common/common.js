import { Option } from 'js-cordyceps';

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
