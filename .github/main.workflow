workflow "New workflow" {
  on = "push"
  resolves = ["latest"]
}

workflow "New workflow 1" {
  on = "push"
}

action "latest" {
  uses = "alpine"
  runs = "date"
}
