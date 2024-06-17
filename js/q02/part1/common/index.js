import { readFileSync } from 'fs';
import { Result, flipResultList } from "js-cordyceps";

export function parseFile(path) {
  return Result.try(readFileSync, [path, { encoding: 'utf8' }])
    .andThen(content => {
      const parseGameResults = content
        .trim()
        .split('\n')
        .map((line, index) => {
          const result = parseGame(line);
          if (result.isErr()) {
            return result;
          }

          return Result.makeOk({
            number: index + 1,
            rounds: result.unwrap(),
          });
        });

      return flipResultList(parseGameResults);
    });
}

function parseGame(str) {
  const headerIdx = str.indexOf(': ');
  if (headerIdx === -1) {
    return Result.makeErr(`Did not find header index for game: ${str}`);
  }

  const parseRoundResults = str.slice(headerIdx + 2)
    .split('; ')
    .map(parseRound);

  return flipResultList(parseRoundResults);
}

function parseRound(str) {
  const parseCubeResults = str.trim()
    .split(',')
    .map(parseCubes);

  const result = flipResultList(parseCubeResults);
  if (result.isErr()) {
    return result;
  }

  let out = {};
  for (const entry of result.unwrap()) {
    out = {...out, ...entry};
  }

  return Result.makeOk(out);
}

function parseCubes(str) {
  const parts = str.trim().split(' ');
  if (parts.length !== 2) {
    return Result.makeErr(`Incorrect cube number string: ${str}`);
  }

  const amount = parseInt(parts[0], 10);
  if (Number.isNaN(amount)) {
    return Result.makeErr(`Error while parsing number: ${parts[0]}`);
  }

  const color = parts[1];
  if (!['red', 'green', 'blue'].includes(color)) {
    return Result.makeErr(`Invalid color: ${color}`);
  }

  return Result.makeOk({
    [color]: amount,
  });
}
