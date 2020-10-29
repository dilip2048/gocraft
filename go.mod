module github.com/conversenow/balrog

require (
	github.com/rs/zerolog v1.17.2
	k8s.io/utils v0.0.0-20191114184206-e782cd3c129f
)

replace (
	//github.com/conversenow/elrond => ../elrond
	//github.com/conversenow/gimli => ../gimli
	//github.com/conversenow/lembas => ../lembas
	github.com/jesseward/azuretexttospeech => github.com/akshaylb/azuretexttospeech v0.0.0-20190904144142-42c3c7cee5d6
	github.com/nlopes/slack => github.com/akshaylb/slack v0.5.1-0.20190227121055-56a7add8c63f
)

go 1.14
