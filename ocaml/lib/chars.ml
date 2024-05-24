let explode str =
  let rec expl i accum =
    if i < 0 then accum else expl (i - 1) (str.[i] :: accum)
  in
  expl (String.length str - 1) []

let explode_rev str =
  let rec expl i accum =
    if i > String.length str - 1 then accum else expl (i + 1) (str.[i] :: accum)
  in
  expl 0 []

