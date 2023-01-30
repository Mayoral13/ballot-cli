package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)
type Candidate struct{ //TO STORE CANDIDATES DETAILS
	name string;
	id int;
	votes int;
	winner bool;
}
var started bool; //TO CHECK IF VOTING HAS STARTED
var ended bool; //TO CHECK IF VOTING HAS ENDED
var counter int; //COUNTER TO KEEP COUNT OF CANDIDATE'S ID
var candidates[] Candidate; //SLICE TO STORE ALL CANDIDATES
var users = make(map[int] Candidate); //MAPPING OF CANDIDATE'S ID TO THE CANDIDATE STRUCT

func main(){ // MAIN PROGRAM CODE SEEMS SELF EXPLANATORY
	for true{
	fmt.Println("Welcome to the ballot program")
	fmt.Println("Enter 1 to enter ballot")
	fmt.Println("Enter 2 to end voting")
	fmt.Println("Enter 3 to start/end election")
	fmt.Println("Enter 4 to vote") //DURING VOTING TO POLL RESULT INPUT AN ID THAT DOES NOT EXIST
	input := bufio.NewReader(os.Stdin)
    text,_ := input.ReadString('\n')
    value := strings.TrimSpace(text)
	if !(value == "1" || value == "2" || value == "3" || value =="4"){
     fmt.Println("Invalid input")
	}else if(value == "2"){
		EndVote()
		}else if(value == "1"){
		AddUser();
		fmt.Println("User added")
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

func AddUser(){ //TO REGISTER CANDIDATES FOR ELECTION
if(!started){ // CANNOT EXECUTE UNLESS STARTED IS TRUE
fmt.Println("Enter name")
input := bufio.NewReader(os.Stdin)
text,_ := input.ReadString('\n')
value := strings.TrimSpace(text)
if(value != ""){
counter++;
candidates = append(candidates,Candidate{name:value,id:counter,votes:0,winner:false});// STORES CANDIDATES DETAILS
users[counter] = Candidate{name:value,id:counter,votes:0,winner:false} 
}else{
	fmt.Println("ERROR INVALID INPUT")
}
}
}
func VoteUser(){ // TO VOTE FOR SPECIFIC USER BASED ON USER INPUT
	if(started){
		for i:=0; i<len(candidates); i++{//DISPLAYS ALL CANDIDATES NAMES AND THEIR ID 
			fmt.Println("Candidate names are : ",candidates[i].name)
			fmt.Println("Candidate id are : ",candidates[i].id)
		}
	}
	fmt.Println("Enter Candidate's ID to register vote")
	input := bufio.NewReader(os.Stdin)
    text,_ := input.ReadString('\n')
    value := strings.TrimSpace(text)
	id,_ := strconv.Atoi(value)
	if !(UserExist(id)){ // A FUNCTION CHECKS TO SEE IF A USER EXISTS BY PASSING ITS ID AS A PARAMETER  
		fmt.Println("NOT FOUND")
	}else if(UserExist(id)){
	if votes,ok := users[id];ok{ // WORKAROUND TO REINITIALIZE CANDIDATE STRUCT DATA
		votes.votes++;
		users[id] = votes;
		fmt.Println("You voted for candidate ", users[id].id)
	}else{
		fmt.Println("NO VOTE")
	}
	}
}

func UserExist(id int) (bool){ //FUNCTION TO CHECK IF A USER EXISTS
if(users[id].id != 0){
 return true;
}else{
	return false
}
}
func EndVote(){ // FUNCTION TO POLL ALL VOTES AND DECLARE WINNER
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
		main()
		fmt.Println("VOTING NOT STARTED")
	}
	ended = true;
	fmt.Println("Winner is ", users[winner].name)
	fmt.Println("Wins with votes of ", users[winner].votes)
	os.Exit(0)
}
