package db

import (
	"context"
	"fmt"
	"github.com/Tranduy1dol/go-microservice/internal/store"
	"log"
	"math/rand"
)

var usernames = []string{
	"alice", "bob", "charlie", "dave", "eve", "frank", "grace", "heidi", "ivan", "judy",
	"mallory", "nina", "olivia", "peter", "quinn", "rachel", "sam", "trudy", "victor", "wendy",
	"xander", "yara", "zane", "aaron", "bella", "carl", "diana", "eric", "fiona", "george",
	"harry", "isla", "jack", "karen", "leo", "mona", "nathan", "oliver", "paula", "quincy",
	"robert", "sophia", "tina", "ursula", "vicky", "will", "xena", "yasmine", "zach",
	"aaron1", "bella2", "carl3", "diana4", "eric5", "fiona6", "george7", "harry8", "isla9",
	"jack10", "karen11", "leo12", "mona13", "nathan14", "oliver15", "paula16", "quincy17",
}

var titles = []string{
	"Introduction to Go", "Advanced Go Programming", "Go Concurrency Patterns",
	"Building Web Applications with Go", "Go Microservices Architecture",
	"Testing in Go", "Go for Data Science", "Go and Cloud Computing",
	"Go for DevOps", "Go Performance Optimization", "Go Security Best Practices",
	"Go Design Patterns", "Go and Machine Learning", "Go for IoT",
	"Go and Blockchain", "Go for Game Development", "Go for Mobile Development",
}

var tags = []string{
	"golang", "microservices", "web development", "cloud computing",
	"devops", "data science", "machine learning", "iot", "blockchain",
	"game development", "mobile development", "security", "performance",
	"design patterns", "testing", "concurrency", "architecture",
}

var contents = []string{
	"This is a sample content for the post. It can be about anything related to Go programming.",
	"Go is a statically typed, compiled programming language designed for simplicity and efficiency.",
	"Concurrency in Go is achieved through goroutines and channels, making it easy to write concurrent programs.",
	"Go's standard library provides a rich set of packages for building web applications, handling HTTP requests, and more.",
	"Microservices architecture in Go allows developers to build scalable and maintainable applications.",
	"Testing in Go is straightforward with its built-in testing package, making it easy to write unit tests.",
	"Go is increasingly being used in data science for its performance and ease of use with large datasets.",
	"Cloud computing with Go allows developers to build applications that can scale seamlessly across distributed systems.",
	"DevOps practices can be enhanced with Go, enabling automation and efficient deployment pipelines.",
	"Data science in Go is gaining popularity due to its performance and rich ecosystem of libraries.",
	"Machine learning in Go is supported by libraries like Gorgonia and GoLearn, making it easier to implement ML algorithms.",
	"Go's simplicity and performance make it a great choice for IoT applications, enabling efficient resource usage.",
	"Blockchain development in Go is facilitated by libraries like go-ethereum, allowing developers to build decentralized applications.",
	"Game development in Go is becoming more popular with libraries like Ebiten and Pixel, enabling 2D game development.",
	"Mobile development in Go is promising with projects like Gomobile, allowing developers to build Android and iOS apps.",
}

var comments = []string{
	"Thanks for sharing this information. It's very helpful for beginners.",
	"I love Go's concurrency model. It makes writing concurrent programs so much easier.",
	"Go's standard library is amazing. It has everything you need to build web applications.",
	"Microservices architecture in Go is a game changer. It allows for better scalability and maintainability.",
	"Testing in Go is so straightforward. I appreciate how easy it is to write unit tests.",
	"Data science in Go is fascinating. I can't wait to explore more libraries and tools.",
	"Cloud computing with Go is powerful. It enables seamless scaling across distributed systems.",
	"DevOps practices in Go are efficient. I love how it simplifies automation and deployment.",
	"Data science in Go is gaining traction. The performance and ease of use are impressive.",
	"Machine learning in Go is exciting. I look forward to implementing some algorithms.",
	"Go's simplicity is a breath of fresh air. It makes programming enjoyable.",
	"IoT applications in Go are efficient. The resource usage is optimized for low-power devices.",
	"Blockchain development in Go is interesting. I appreciate the libraries available for building decentralized apps.",
	"Game development in Go is fun. The libraries make it easy to create 2D games.",
	"Mobile development in Go is promising. I can't wait to see more apps built with it.",
}

func Seed(store store.Storage) {
	ctx := context.Background()

	users := generateUser(100)
	for _, user := range users {
		if err := store.Users.Create(ctx, user); err != nil {
			log.Printf("Error creating user: %v", err)
			return
		}
		log.Printf("User created: %s", user.Username)
	}

	posts := generatePosts(100, users)
	for _, post := range posts {
		if err := store.Posts.Create(ctx, post); err != nil {
			log.Printf("Error creating post: %v", err)
			return
		}
	}

	comments := generateComments(100, posts, users)
	for _, comment := range comments {
		if err := store.Comments.Create(ctx, comment); err != nil {
			log.Printf("Error creating comment: %v", err)
			return
		}
	}
	log.Printf("Seeding completed successfully!")
}

func generateUser(num int) []*store.User {
	users := make([]*store.User, num)

	for i := 0; i < num; i++ {
		users[i] = &store.User{
			Username: usernames[i%len(usernames)] + fmt.Sprintf("%d", i),
			Email:    usernames[i%len(usernames)] + fmt.Sprintf("%d", i) + "@example.com",
			Password: "password",
		}
	}

	return users
}

func generatePosts(num int, users []*store.User) []*store.Post {
	posts := make([]*store.Post, num)
	for i := 0; i < num; i++ {
		user := users[rand.Intn(len(users))]

		posts[i] = &store.Post{
			Title:   titles[rand.Intn(len(titles))],
			Content: contents[rand.Intn(len(contents))],
			UserID:  user.ID,
			Tags: []string{
				tags[rand.Intn(len(tags))],
				tags[rand.Intn(len(tags))],
			},
		}
	}

	return posts
}

func generateComments(num int, posts []*store.Post, users []*store.User) []*store.Comment {
	cmts := make([]*store.Comment, num)
	for i := 0; i < num; i++ {
		post := posts[rand.Intn(len(posts))]
		user := users[rand.Intn(len(users))]

		cmts[i] = &store.Comment{
			Content: comments[rand.Intn(len(comments))],
			UserID:  user.ID,
			PostID:  post.ID,
		}
	}

	return cmts
}
