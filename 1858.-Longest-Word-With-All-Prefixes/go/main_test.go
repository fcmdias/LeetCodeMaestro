package main

import "testing"

var tests = []struct {
	words []string
	want  string
}{
	{
		words: []string{"abc", "bc", "ab", "qwe"},
		want:  "",
	},
	{
		words: []string{"abc", "bc", "ab", "qwe", "a"},
		want:  "abc",
	},
	{
		words: []string{"oinjnallhkurmqjfahpw", "oinjnallhkurmqjfahpwv", "o", "oinjnallhkurmq", "oin", "oinjnallhkurmqjfah", "oinj", "oinjnallhkurmqjfahp", "oinjnall", "oinjnallhkurmqjfa", "oinjnallhkur", "oinjnallhkurmqjf", "oinjn", "oinjnallhkurm", "oinjnallhku", "oinjnallh", "oinjna", "oinjnallhk", "oi", "oinjnallhkurmqj", "oinjnal", "oinjnallhkurmqjfahpw", "oinjnallhkurmqjfahpwv", "o", "oinjnallhkurmq", "oin", "oinjnallhkurmqjfah", "oinj", "oinjnallhkurmqjfahp", "oinjnall", "oinjnallhkurmqjfa", "oinjnallhkur", "oinjnallhkurmqjf", "oinjn", "oinjnallhkurm", "oinjnallhku", "oinjnallh", "oinjna", "oinjnallhk", "oi", "oinjnallhkurmqj", "oinjnal"},
		want:  "oinjnallhkurmqjfahpwv",
	},
	{
		words: []string{"pmvgibowhihgrdejdgehqdxyjfcrjajkdtfmfegfyazbzi", "o", "pmvgibowhihgrdejdgehqdxyjfcrjajkdtf", "pmvgibowhihgrdejdgehqdxyjfcrjajkdtfmfegfyazbzixcmvbpsomdsyzw", "pmvgibowhihgrdejdgehqdxyjfcrjajkdtfmfegfyazbzixc", "pmvgibowhihgrdejdgehqdxyjfcrjajkdtfmfegfyazbzixcmvb", "ozjivmgyvwmmeewmtrwgolluevnvvtxxoxnzynfwhijgprevvtz", "ozjivmgyvwmmeewmtrwgolluevnvvtxxoxnzynfw", "pmvgibowhihgrdejdgehqdxyjfcrjajkdtfmfegfyazbzixcmvbps", "pmvgibowhihgrdejdgehqdxyjfcrjajk", "pmvgibowhihgrdejdgehqdxyjfcrjajkdtfmfegfyazbzixcmvbpsomdsyzw", "pmvgibow", "ozjivmgyvwmmeew", "pmvgibowhihgrdejdgehqdxyjfcrjajkdtfmfegfyazbzixcmvbpsomdsy", "pmvgibowhihgrdejdgehqdxyjfcrjajkdtfmfegfyazbzixcm", "ozjivmgyvwmmeewmtrwgolluevnvvtxxoxnzynfwhijgprevvtzrewclrj", "ozjivmgyvwmmeewmtrwgolluevnvvtxxoxnzynfwhijgp", "ozjivmgyvwmmeewmtrwgolluevnvvtxxoxnzynfwhijgprevvtzrew", "odakvhi", "p", "fxxjonexqt", "ozjivmgyvwmmeewmtrwgolluevnv", "pmvgibowhihgrdejdgeh", "od", "pmvgibowhihgrdejdgehqdxyjfcrjajkdtfm", "pmvgibowhihgrdejdgehqdxyjfcrj", "ozjivmgyvwmmeewmtrwgolluevnvvtxxoxnzynfwhijgprevvtzr", "tpgz", "fxxjo", "ozjivmgyvwmmeewmtrwgoll", "ozjivmgyvwm", "ozjivm", "pmvgibowhihgrdejdg", "ozjivmgyvwmmeewmtrwgollue", "ozjivmgyvwmmeewmtrwgolluevnvvtxxoxnzynfwhijgprevvtzre", "fxxjonexqtpfaqeliviegkmn", "pmvgibowhihgrd", "odakvhixlqm", "t", "fxxj", "ozjivmgyvwmmeewmtrwgolluevnvvtxxoxnz", "fxxjonexq", "fxxjonex", "pmv", "fxxjonexqtpfaqeliviegkmndpam", "fxxjonexqtpfa", "fxxjonexqtpfaqeliviegkmndp", "pmvgibowhihgrdejdgehqdxyjfcr", "ozjivmgyvwmmeewmtrwgolluevnvvtxxoxnzynfwhi", "pm", "pmvgibowhihgrdejd", "ozjivmgyvwmmeewmtrwgol", "pmvgibowhihgrdejdgehqdxy", "fxxjonexqtpfaqel", "ozjivmgyvwmm", "pmvgibowhihgrde", "fxxjonexqtpfaqe", "fxxjonexqtpfaqelivi", "pmvgibowhihgrdejdgehqdx", "fxxjonexqtpfaqeliviegk", "fxxjon", "pmvgibowh", "ozjivmgyvwmmeewmtrwgolluevnvvtxxoxnzynfwhijgprevvtzrewcl", "ozjivmgyv", "odak", "odakvhixl", "pmvgibowhihgrdejdgehqdxyj", "pmvgibowhihgrdejdgehqdxyjfcrjajkdtfmfegfyazbzixcmvbpsom", "oda", "odakvhixlq", "pmvgibowhi", "ozjivmgyvwmmeewmtrwg", "fxxjonexqtpfaqeliviegkmndpa", "pmvgibowhihgrdejdgehqdxyjfcrjajkdtfmfeg", "pmvgibowhihgrdejdge", "ozjivmgyvwmmeewmtrwgolluev", "ozjivmgyvwmmeewmtrw", "ozjivmgyvwmmeewmtrwgolluevnvvt", "fxxjonexqtpfaqelivie", "pmvgibowhihgrdejdgehqdxyjfcrjajkdtfmfegfya", "ozjivmgyvwmmeewmtrwgolluevnvvtxxoxnzynfwhijgprevvtzrewclrjbn", "oz", "odakv", "ozjivmgyvwmmeewmtrwgolluevn", "fxxjonexqtpfaqeliviegkmnd", "pmvgibowhihgr", "ozjivmgyvwmmeewmtrwgolluevnvvtxxoxnzynfwhijgprevvtzrewc", "ozjivmgyvw", "pmvgibowhihgrdejdgehqdxyjfcrjajkdtfmfegfyazbzixcmvbp", "pmvgibowhihgrdej", "pmvgibowhihgrdejdgehq", "ozji", "ozjiv", "ozjivmgyvwmmeewmtrwgolluevnvvtxxoxn", "ozjivmgyvwmmeewmtrwgolluevnvvtxxoxnzynfwhijgprev", "fxxjonexqtpfaqeliv", "ozjivmgyvwmmeewmtrwgolluevnvvtxxox", "pmvgibowhihgrdejdgehqdxyjfcrjajkdtfmfegfyazbzixcmvbpsomd", "fxxjonexqtp", "fxxjonexqtpfaqeli", "ozjivmgyvwmmeewmtrwgollu", "tpgz", "ozj", "odakvhixlqm", "ozjivmgyvwmmee", "pmvgi", "pmvg", "pmvgibowhihgrdejdgehqdxyjfcrjajkdtfmfegfyazbzix", "fxxjone", "pmvgibowhihgrdejdgehqdxyjfcrjajkdtfmfegfyazbzixcmvbpsomdsyz", "tpg", "pmvgibo", "fxx", "ozjivmgyvwmmeewmtr", "pmvgibowhihgrdejdgehqdxyjfcrjajkdtfmfegfyazbzixcmvbpso", "ozjivmgyvwmmeewm", "pmvgibowhihgrdejdgehqdxyjfcrjajkdtfmfegf", "ozjivmg", "pmvgibowhihg", "ozjivmgyvwmmeewmtrwgolluevnvvtxxoxnzynfwh", "ozjivmgy", "ozjivmgyvwmmeewmtrwgolluevnvvtxxoxnzynfwhijgprevvtzrewclrjbn", "pmvgibowhihgrdejdgehqdxyjfcrjajkdtfmfegfyazbzixcmv", "fxxjonexqtpfaqeliviegkmndpam", "ozjivmgyvwmmeewmtrwgo", "pmvgibowhih", "fxxjonexqtpfaqeliviegkm", "pmvgibowhihgrdejdgehqdxyjfcrjajkdtfmfegfyazbzixcmvbpsomds", "pmvgibowhihgrdejdgehqdxyjfcrjajkd", "ozjivmgyvwmmeewmtrwgolluevnvv", "ozjivmgyvwmmeewmtrwgolluevnvvtxxoxnzynf", "pmvgibowhihgrdejdgehqdxyjf", "pmvgibowhihgrdejdgehqdxyjfcrjaj", "odakvhix", "ozjivmgyvwmmeewmtrwgolluevnvvtxxoxnzynfwhijg", "ozjivmgyvwmmeewmtrwgolluevnvvtxx", "odakvh", "pmvgibowhihgrdejdgehqdxyjfcrjajkdtfmfegfyazb", "ozjivmgyvwmmeewmtrwgolluevnvvtxxoxnzynfwhijgprevvtzrewclrjb", "pmvgib", "pmvgibowhihgrdejdgehqdxyjfcrja", "fxxjonexqtpf", "fx", "f", "pmvgibowhihgrdejdgehqdxyjfcrjajkdt", "pmvgibowhihgrdejdgehqdxyjfcrjajkdtfmfegfyazbz", "ozjivmgyvwmmeewmtrwgolluevnvvtxxoxnzyn", "ozjivmgyvwmmeewmtrwgolluevnvvtxxoxnzynfwhijgprevvtzrewclr", "ozjivmgyvwmmeewmtrwgolluevnvvtxxoxnzynfwhij", "ozjivmgyvwmmeewmtrwgolluevnvvtxxoxnzynfwhijgprevvt", "ozjivmgyvwmme", "ozjivmgyvwmmeewmtrwgolluevnvvtxxoxnzy", "pmvgibowhihgrdejdgehqdxyjfc", "ozjivmgyvwmmeewmtrwgolluevnvvtxxo", "fxxjonexqtpfaqelivieg", "o", "pmvgibowhihgrdejdgehqdxyjfcrjajkdtfmfegfyaz", "ozjivmgyvwmmeewmtrwgolluevnvvtxxoxnzynfwhijgpre", "ozjivmgyvwmmeewmtrwgolluevnvvtx", "ozjivmgyvwmmeewmtrwgolluevnvvtxxoxnzynfwhijgprevv", "ozjivmgyvwmmeewmt", "pmvgibowhihgrdejdgehqdxyjfcrjajkdtfmfe", "tp", "pmvgibowhihgrdejdgehqd", "ozjivmgyvwmmeewmtrwgolluevnvvtxxoxnzynfwhijgpr", "fxxjonexqtpfaq", "pmvgibowhihgrdejdgehqdxyjfcrjajkdtfmf", "pmvgibowhihgrdejdgehqdxyjfcrjajkdtfmfegfy"},
		want:  "ozjivmgyvwmmeewmtrwgolluevnvvtxxoxnzynfwhijgprevvtzrewclrjbn",
	},
	{
		words: []string{"abc", "bc", "qwe", "a"},
		want:  "a",
	},
}

func TestAlgoOne(t *testing.T) {
	for _, test := range tests {
		got := AlgoOne(test.words)
		if got != test.want {
			t.Fatalf("got: %v want: %v", got, test.want)
		}
	}
}

func BenchmarkAlgoOneTest0(b *testing.B) {
	for n := 0; n < b.N; n++ {
		AlgoOne(tests[0].words)
	}
}
func BenchmarkAlgoOneTest1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		AlgoOne(tests[1].words)
	}
}
func BenchmarkAlgoOneTest3(b *testing.B) {
	for n := 0; n < b.N; n++ {
		AlgoOne(tests[3].words)
	}
}
