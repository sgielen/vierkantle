-- name: SetScore :exec
INSERT INTO vierkantle.scores (board_name, anonymous_id, user_id, team_size, words, seconds, started_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, NOW(), NOW())
ON CONFLICT (board_name, anonymous_id) DO UPDATE SET user_id=$3, team_size=$4, words=$5, seconds=$6, updated_at=NOW();

-- name: GetScores :many
SELECT anonymous_id, users.username, team_size, words, seconds, COUNT(*) OVER() AS full_count
FROM vierkantle.scores
LEFT JOIN vierkantle.users ON users.id=scores.user_id
WHERE board_name=$1
ORDER BY words DESC, seconds ASC, scores.ctid ASC
LIMIT $2 OFFSET $3;

-- name: GetMyScore :one
SELECT my_score.team_size, my_score.words, my_score.seconds, (
	SELECT COUNT(their_score.ctid) FROM vierkantle.scores their_score
	WHERE their_score.board_name=$1
	AND (their_score.words > my_score.words
	  OR (their_score.words = my_score.words AND their_score.seconds < my_score.seconds)
	  OR (their_score.words = my_score.words AND their_score.seconds = my_score.seconds AND their_score.ctid < my_score.ctid)
	)
) AS rank
FROM vierkantle.scores my_score
WHERE my_score.board_name=$1 AND my_score.anonymous_id=$2;

-- name: RegisterUser :one
INSERT INTO vierkantle.users (username, email, registered_at, last_login_at)
VALUES ($1, $2, NOW(), NULL) RETURNING id;

-- name: LoginUser :one
UPDATE vierkantle.users SET last_login_at=NOW() WHERE id=$1 RETURNING username;