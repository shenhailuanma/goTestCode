package main

func main()  {

	var text1 = `
[Pipeline] sh
+ terraform apply -input=false tfplan
[0m[1m[32m
Apply complete! Resources: 0 added, 0 changed, 0 destroyed.[0m
[Pipeline] }
`

	var text2 = "\x1b[33mThis text is yellow.\x1b[m"

	// 1. Test github.com/robert-nix/ansihtml
	TestAnsi2Html(text1)
	TestAnsi2Html(text2)


	// 2. Test github.com/leaanthony/go-ansi-parser
	TestAnsiParser(text1)
	TestAnsiParser(text2)

}
