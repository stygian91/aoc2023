let explode = Aoc2023.Chars.explode
let explode_rev = Aoc2023.Chars.explode_rev

let rec first_digit_lst chars =
  let to_int = Core.Char.to_int in
  let to_digit ch = to_int ch - to_int '0' in

  match chars with
  | [] -> None
  | h :: t ->
      if Core.Char.is_digit h then Some (to_digit h) else first_digit_lst t

let first_digit line = first_digit_lst @@ explode line
let last_digit line = first_digit_lst @@ explode_rev line

let parse_line line =
  let fopt = first_digit line in
  let lopt = last_digit line in
  match (fopt, lopt) with Some f, Some l -> Some ((f * 10) + l) | _ -> None

let process path =
  let lines = Core.In_channel.read_lines path in

  List.fold_right
    (fun v acc -> Option.value ~default:0 (parse_line v) + acc)
    lines 0

(* -------------------- *)

let path = "./data/q01/data.txt"
let res = process path
let () = Printf.printf "%d\n" res
