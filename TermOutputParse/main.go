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
	var text3 = `
yum isntall -\u007f\u007f\u007f\u007f\u007f\u007f\u007f\u007fnstall -y telnet\r
ifcofig\u007f\u007f\u007f\u007f\u007f\u007f\u007f\u007f\u007f\u007f\u007f\u007f\u007f\u007f\u007fifconfig\r
\x1b[A\x1b[A\x1b[A\r
ping 1114.\u007f\u007f\u007f4.1144.11\u007f\u007f\u007f\u007f.114.1144\r`

	// 1. Test github.com/robert-nix/ansihtml
	TestAnsi2Html(text1)
	TestAnsi2Html(text2)


	// 2. Test github.com/leaanthony/go-ansi-parser
	TestAnsiParser(text1)
	TestAnsiParser(text2)
	TestAnsiParser(text3)


}
