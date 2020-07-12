package main

import "fmt"

var names = map[int]string{
	1:   "unary",
	2:   "binary",
	3:   "trinary",
	4:   "quaternary",
	5:   "quinary",
	10:  "seximal",
	11:  "septimal",
	12:  "octal",
	13:  "nonary",
	14:  "decimal",
	15:  "elevenary",
	20:  "dozenal",
	21:  "baker's dozenal",
	22:  "biseptimal",
	23:  "triquinary",
	24:  "hex",
	25:  "suboptimal",
	30:  "triseximal",
	31:  "untriseximal",
	32:  "vigesimal",
	33:  "triseptimal",
	34:  "bielevenary",
	35:  "unbielevenary",
	40:  "tetraseximal",
	41:  "pentaquinary",
	42:  "biker's dozenal",
	43:  "trinonary",
	44:  "tetraseptimal",
	45:  "untetraseptimal",
	50:  "pentaseximal",
	51:  "unpentaseximal",
	52:  "tetroctal",
	53:  "trielevenary",
	54:  "bisuboptimal",
	55:  "pentaseptimal",
	100: "niftimal",
}

func NameBase(base int) string {
	if val, ok := names[base]; ok {
		return val
	}
	return fmt.Sprintf("%dary", base)
}
