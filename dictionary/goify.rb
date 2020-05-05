# Generate the `words.go` script because pkger is not working very
# well...

words = File.read('words.txt').split("\n")
lines = ["package dictionary", "var Words = [#{words.size}]string{"]

words.each do |word|
  lines << %("#{word}",)
end

lines << "};"

File.write("words.go", lines.join("\n"))
