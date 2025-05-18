package utils

import (
	"blogo/app/models"
	"blogo/app/repositories"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
)

func InsertTestData(userRepo repositories.UserRepo, postRepo repositories.PostRepo) {
	users_file, err := os.Open("app/utils/users.json")
	if err != nil {
		fmt.Println("Couldn't find users data file")
		return
	}
	defer users_file.Close()

	var users []models.User
	decoder := json.NewDecoder(users_file)
	if err := decoder.Decode(&users); err != nil {
		fmt.Println("Error importing user data: ", err.Error())
		return
	}

	for _, user := range users {
		userRepo.Add(&user)
	}

	posts_file, err := os.Open("app/utils/posts.json")
	if err != nil {
		fmt.Println("Couldn't find posts data file")
		return
	}
	defer posts_file.Close()

	var posts []models.Post
	decoder = json.NewDecoder(posts_file)
	if err := decoder.Decode(&posts); err != nil {
		fmt.Println("Error importing post data: ", err.Error())
		return
	}

	for _, post := range posts {
		post.UserID = uint(rand.Intn(len(users))) + 1
		postRepo.Add(&post)
	}

	fmt.Println("Testing data added")
}
