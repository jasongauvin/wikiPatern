package models

// LoadFixtures of articles in database
func LoadFixtures() {

	// Article Fixture
	article := Article{Title: "Test article 1", Content: "This is the content of the test article number 1, please comment if you liked it or if you have any questions."}
	db.Create(&article)

	// Comment Fixture
	comment := Comment{Text: "This is an awesome comment!"}
	db.Create(&comment)
}
