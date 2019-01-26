workflow "New workflow" {
  on = "push"
  resolves = ["find"]
}

workflow "New workflow 1" {
  on = "push"
}

action "find" {
  uses = "find"
  runs = "date"
}
