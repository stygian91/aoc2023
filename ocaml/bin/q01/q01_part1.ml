let explode str =
  let rec expl i accum = if i < 0 then
    accum
  else
    expl (i - 1) (str.[i] :: accum)
  in expl (String.length str - 1) []

let explode_rev str =
  let rec expl i accum =
    if i > String.length(str) - 1 then accum else expl (i + 1) (str.[i] :: accum) in
  expl (0) []

let rec first_digit_lst chars =
  let to_digit ch = match ch with
  | '0' -> Some 0
  | '1' -> Some 1
  | '2' -> Some 2
  | '3' -> Some 3
  | '4' -> Some 4
  | '5' -> Some 5
  | '6' -> Some 6
  | '7' -> Some 7
  | '8' -> Some 8
  | '9' -> Some 9
  | _ -> None

  in match chars with
    | [] -> None
    | h :: t -> if Core.Char.is_digit h then to_digit h else first_digit_lst t

let first_digit line = first_digit_lst @@ explode line
let last_digit line = first_digit_lst @@ explode_rev line

let parse_line line =
  let fopt = first_digit line in
  let lopt = last_digit line in
  match (fopt, lopt) with
    | (Some f, Some l) -> Some ((f * 10) + l)
    | _ -> None

let process path =
    let lines = Core.In_channel.read_lines path in

    let or_zero x = match parse_line x with
    | None -> 0
    | Some x -> x in

    List.fold_right (fun v acc -> or_zero v + acc) lines 0

(* -------------------- *)

let path = "./data/q01/data.txt"
let res = process path
let () = Printf.printf "%d\n" res
