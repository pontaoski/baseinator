package main

import (
	"baseinator/bases"
	"errors"
	"fmt"
	"strconv"

	"github.com/manifoldco/promptui"
)

var (
	maximumBase int64
	minimumBase int64
)

func main() {
	println("hi, this is libre software, licensed under the AGPL. the source is available at https://github.com/pontaoski/baseinator")
	println("welcome to baser, a program that helps you figure out what base is best for your needs!")
	println("before we begin, let's set a few boundaries. what is the largest base you'd want to use?")
	res, _ := (&promptui.Prompt{
		Label: "maximum base (in decimal)",
		Validate: func(input string) error {
			num, err := strconv.ParseInt(input, 10, 64)
			if err != nil {
				return errors.New("that's not a number you bonk")
			}
			if num > 720720 {
				return errors.New("ähm... why the heck do you want a base that big? go smaller")
			} else if num < 2 {
				return errors.New("seriously, i said largest")
			}
			return nil
		},
	}).Run()
	maximumBase, _ = strconv.ParseInt(res, 10, 64)
	println("now, what about the smallest?")
	res, _ = (&promptui.Prompt{
		Label: "smallest base (in decimal)",
		Validate: func(input string) error {
			num, err := strconv.ParseInt(input, 10, 64)
			if err != nil {
				return errors.New("that's not a number you twat")
			}
			if num > maximumBase {
				return errors.New("this is bigger than your maximum, you doofus")
			}
			if num > 720720 {
				return errors.New("but y tho")
			} else if num < 2 {
				return errors.New("how do you even use a base that small")
			}
			return nil
		},
	}).Run()
	minimumBase, _ = strconv.ParseInt(res, 10, 64)
	pool := bases.PoolForRange(minimumBase, maximumBase)
	denom := 2
	for len(pool) > 1 {
		fmt.Printf("there are %d bases in your pool\n", len(pool))
		fmt.Printf("the smallest base in your pool is base %d, %s\n", pool.Smallest(), pool.Smallest().Name())
		amt, length, recurring := pool.LargestExpansionFor(denom)
		if recurring {
			fmt.Printf("%d bases in your pool represent 1/%d with a repeating decimal\n", len(amt), denom)
		} else {
			fmt.Printf("%d bases in your pool represent 1/%d with %d digit(s)\n", len(amt), denom, length)
		}
		_, res, _ := (&promptui.Select{
			Label: "remove these bases?",
			Items: []string{"no", "yes"},
		}).Run()
		if res == "yes" {
			pool.RemoveLongestExpansionFor(denom)
		}

		denom++
	}
	if len(pool) == 1 {
		best := pool.Smallest()
		fmt.Printf("oh wow, it looks like base %d, %s, is the best for you!\n", best, best.Name())
	} else {
		println("oh no, it doesn't look like any bases are good for you")
	}
}
