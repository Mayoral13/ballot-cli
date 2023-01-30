package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)
type Candidate struct{
	name string;
	id int;
	votes int;
	winner bool;
}
var started bool;
var ended bool;
var counter int;
var candidates[] Candidate;
var users = make(map[int] Candidate);

func main(){
	for true{
	fmt.Println("Welcome to the ballot program")
	fmt.Println("Enter 1 to enter ballot")
	fmt.Println("Enter 2 to view result")
	fmt.Println("Enter 3 to start/end election")
	fmt.Println("Enter 4 to poll votes")
	input := bufio.NewReader(os.Stdin)
    text,_ := input.ReadString('\n')
    value := strings.TrimSpace(text)
	if !(value == "1" || value == "2" || value == "3" || value =="4"){
     fmt.Println("Invalid input")
	}else if(value == "1"){
		AddUser();
		fmt.Println("User added")
	}else if(value == "2" && started){
		fmt.Println("The result are")
		ViewResult()
	}else if(value == "2" && !started){
		fmt.Println("Election has not started")
	}else if(value == "3"){
		if(started){
        started = false;
		fmt.Println("Election Ended")
		}else if(value == "4" && started && ended){
			EndVote()
			}else if(value == "4" && started && !ended){
				fmt.Println("NOT ENDED YET")
				}else if(value == "4" && !started){
				fmt.Println("Election not started")
				}else{
			started = true;
			fmt.Println("Election Started")
		}

	}else if(value == "4"){
		VoteUser()
		}else{
		fmt.Println("ERROR")
	}
}
}

func AddUser(){
if(!started){
fmt.Println("Enter name")
input := bufio.NewReader(os.Stdin)
text,_ := input.ReadString('\n')
value := strings.TrimSpace(text)
if(value != ""){
counter++;
candidates = append(candidates,Candidate{name:value,id:counter,votes:0,winner:false});
users[counter] = Candidate{name:value,id:counter,votes:0,winner:false}
}else{
	fmt.Println("ERROR INVALID INPUT")
}
}
}
func VoteUser(){
	if(started){
		for i:=0; i<len(candidates); i++{
			fmt.Println("Candidate names are : ",candidates[i].name)
			fmt.Println("Candidate id are : ",candidates[i].id)
		}
	}
	for true{
	fmt.Println("Enter Candidate's ID to register vote")
	input := bufio.NewReader(os.Stdin)
    text,_ := input.ReadString('\n')
    value := strings.TrimSpace(text)
	id,_ := strconv.Atoi(value)
	if !(UserExist(id)){
		fmt.Println("Unregistered candidate")
		EndVote();
	}else if(UserExist(id)){
	if votes,ok := users[id];ok{
		fmt.Println("You voted for candidate ", users[id].id)
		votes.votes += 1;
		users[id] = votes;
	}else{
		fmt.Println("NO VOTE")
	}
	}
	}
	
}
func ViewResult(){
	if(started){
		for i:=0; i<len(candidates); i++{
			fmt.Println("Candidate name is : ",candidates[i].name)
			fmt.Println("Candidate id is : ",candidates[i].id)
			fmt.Println("Candidate votes are : ",candidates[i].votes)
			fmt.Println("Candidate win ? : ",candidates[i].winner)
		}
	}else{
		fmt.Println("ELECTION NOT STARTED")
	}
}
func UserExist(id int) (bool){
if(users[id].id != 0){
 return true;
}else{
	fmt.Println("User not found")
	return false
}
}
func EndVote(){
	var winner int;
	if(started){
		fmt.Println("Voting ended")
		for i:=0; i<len(candidates) -1; i++{
			for j:=0; j<len(candidates); j++{
            if(candidates[i].votes > candidates[j].votes){
				winner = candidates[i].id;
			}else{
				winner = candidates[j].id;
			}
			}
		}
	}else{
		fmt.Println("VOTING NOT STARTED")
	}
	ended = true;
	fmt.Println("Winner is ", users[winner].name)
	fmt.Println("Wins with votes of ", users[winner].votes)
	os.Exit(0)
}