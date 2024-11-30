# fastTrack Quiz
FastTrack CLI is a command-line tool that interacts with a server API. By using the program who will have to answer several different questions about Fast Track and in the end the program will display both the number of correct answers and how you compare to others that had already answered the quiz.
This README provides instructions for setting up and running the CLI.

## Prerequisites

- Ensure Go is installed (if not, follow instructions [here](https://go.dev/doc/install))



## Installation
1. Clone the repository and open folder
```
git clone https://github.com/lf-silva/fastTrack.git
cd fastTrack
```
2. Start the server by running the following command from the repository **root folder**:
```
make server
```
3. Open a second terminal window and, also from the repository **root folder**, install the application running these 2 commands:
```
make cli
```
4. Start the terminal application with the following instruction:
```
cli start
```



# Requirements
Fast Track Code Test Quiz - Instructions 
The task is to build a super simple quiz with a few questions and a few alternatives for each question. Each with one correct answer. 

**Preferred Stack:**
- Backend - Golang
- Database - Just in-memory, so no database 

**Preferred Components:**
- REST API or gRPC
- CLI that talks with the API, preferably using https://github.com/spf13/cobra ;( as CLI framework )

**User stories/Use cases:**
- User should be able to get questions with a number of answers
- User should be able to select just one answer per question.
- User should be able to answer all the questions and then post his/hers answers and get back how many correct answers they had, displayed to the user.
- User should see how well they compared to others who have taken the quiz, eg. "You were better than 60% of all quizzers"
