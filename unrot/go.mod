module github.com/tubbo/cipherutils/unrot

go 1.13

require (
	github.com/markbates/pkger v0.14.0
	github.com/tubbo/cipherutils/dictionary v0.0.0-00010101000000-000000000000
	github.com/yuya-takeyama/argf v0.0.0-20150217044922-3ff699b4b958
)

replace github.com/tubbo/cipherutils/dictionary => ../dictionary
