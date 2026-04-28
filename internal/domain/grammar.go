package domain

type Grammar struct {
	ID        string           `bson:"_id"`
	Pattern   string           `bson:"pattern"`
	Meaning   string           `bson:"meaning"`
	Formation string           `bson:"formation"`
	JLPT      int              `bson:"jlpt"`
	Example   []GrammarExample `bson:"examples"`
	Notes     string           `bson:"notes"`
	Source    string           `bson:"source"`
}

type GrammarExample struct {
	Japanese    string `bson:"japanese"`
	Reading     string `bson:"reading"`
	Translation string `bson:"translation"`
}
