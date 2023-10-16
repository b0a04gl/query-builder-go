package model

type QueryRequest struct {
	Select  []string `json:"select"`
	From    string   `json:"from"`
	Where   string `json:"where"`
	OrderBy string   `json:"orderBy"`
}


/*
"select" : ["id","name","team","role"],
    "from" : "players",
    "where" : ["batsman"],
    "orderBy" : "tota"
*/