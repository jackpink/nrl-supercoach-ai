
-- name: GetRookies :many
SELECT PlayerScore.playerID FROM PlayerScore INNER JOIN Round 
ON (PlayerScore.round = Round.round AND PlayerScore.year_ = Round.year_) 
WHERE Round.round = $1 
AND Round.year_ = $2 
AND PlayerScore.price = $3 
AND jersey<=13;