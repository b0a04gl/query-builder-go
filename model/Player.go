package model
 
type Player struct {
	Id  int `json:"id"`
	Name    string   `json:"name"`
	Team    string   `json:"team"`
	Role    string   `json:"role"`
	Total_runs    int   `json:"total_runs"`
	Total_wickets   int `json:"total_wickets"`
}

/*

 id | name  | team  |  role   | total_runs | total_wickets
----+-------+-------+---------+------------+---------------
  1 | virat | india | batsman |      15000 |             5

*/
