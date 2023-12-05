const fs = require('fs');
const F = require('funky-lib');

const getId = F.pipe(
  F.split(' '),
  F.filter((str) => str.length),
  F.tail,
  (str) => parseInt(str),
);

const parseNumbers = F.pipe(
  F.split(' '),
  F.filter(F.prop('length')),
  F.map((s) => parseInt(s)),
);

const getWinningNumbers = F.pipe(F.head, parseNumbers);
const getPlayerNumbers = F.pipe(F.nth(1), parseNumbers);

const parse = F.pipe(
  F.split('\n'),
  F.transduce(
    F.compose(
      F.filter(F.identity),
      F.map(F.split(': ')),
      F.map((parts) => {
        const numParts = F.split(' | ', parts[1]);

        return {
          id: getId(parts[0]),
          winningNumbers: getWinningNumbers(numParts),
          playerNumbers: getPlayerNumbers(numParts),
        };
      }),
    ),
    (acc, curr) => F.concat(acc, [curr]),
    [],
  ),
);

const countWins = F.converge(
  F.pipe(F.intersection, F.size),
  [F.prop('winningNumbers'), F.prop('playerNumbers')]
);

const powOf2 = F.curry2(Math.pow)(2);

const calcPoints = F.pipe(
  F.of,
  F.ifElse(F.lt(1), F.always(0), F.pipe(F.subtract(1), powOf2)),
);

// ---- main ----

const contents = fs.readFileSync('./data/input.txt', { encoding: 'utf-8' });
const cards = parse(contents);
const points = F.map(
  F.pipe(countWins, calcPoints),
  cards
);

const part1Result = F.reduce(F.add, 0, points);

console.log(part1Result);

