# Features

- Check if the book is available
- Get all book status
- Get the info of most issued book
- List all the issued book 
- Get the info of top trending books
- Provide json format output

# How to use
Just start the program by running this command in the directory
`go run BookManagerMain.go`

# Endpoints
| Endpoint     | Description |
| --------- | -----:|
| localhost:8080/api/available | Check if the book is available |
| localhost:8080/api/status     |   Get all book status |
| localhost:8080/api/mostIssued      |   Get the info of most issued book |
| localhost:8080/api/top      |   Get the info of top trending books  |
| localhost:8080/api/issued      |    List all the issued book |

The result should be the json format.