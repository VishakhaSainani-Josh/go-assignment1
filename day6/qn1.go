/*1. Given a string containing conversation between alice and bob. In the string, if it reaches $, it means end of alice message and if it reaches #, it means end of bob's message. and if it reaches ^,
it means end of conversation ignore string after that.

Note: given string doesn't contain any spaces. If message contains two continuous conversations from one person it should be printed one after another as given in the example.

write a program to separate out messages from alice and bob. Write messages from alice and bob on seperate channels, whenever a message from alice/bob, print it in front of their name as shown in the example below:

Note: there is a single space before and after colon(:) and no space before and after comma.

e.g: "helloBob$helloalice#howareyou?#Iamgood.howareyou?$^"
output: alice : helloBob,bob : helloalice,bob : howareyou?,alice : Iamgood.howareyou?*/

package main

import "fmt"

func main() {
	var conversation string = "helloBob$helloalice#howareyou?#Iamgood.howareyou?$^"
	var dialogue string
	alice := make(chan string)
	bob := make(chan string)
	for _, i := range conversation {

		if i == '$' {
			go func() {
				alice <- dialogue
			}()

			fmt.Printf("alice : %s,", <-alice)

			dialogue = ""
			continue
		} else if i == '#' {
			
			go func() {
				bob <- dialogue
			}()

			fmt.Printf("bob : %s,", <-bob)

			dialogue = ""
			continue
		} else if i=='^'{
			break
		}
		dialogue = dialogue + string(i)

	}

}
