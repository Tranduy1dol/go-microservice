package domain

type Paragraph struct {
	ID        string     `bson:"_id"`
	Title     string     `bson:"title"`
	Content   string     `bson:"content"`
	JLPT      int        `bson:"jlpt"`
	Questions []Question `bson:"questions"`
	Tags      []string   `bson:"tags"`
	Source    string     `bson:"source"`
}
