workflow "New workflow" {
  on = "push"
  resolves = ["latest"]
}

action "latest" {
  uses = "docker://alpine"
  runs = "date"
}
