-- name: CreateFeedFollow :one
WITH feed_follows AS (
    INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
    VALUES (
        $1,
        $2,
        $3,
        $4,
        $5
    )
    RETURNING *
)
SELECT 
    feed_follows.*,
    users.name AS user_name,
    feeds.name AS feed_name,
    feeds.url
FROM feed_follows
INNER JOIN users ON users.id = feed_follows.user_id
INNER JOIN feeds ON feeds.id = feed_follows.feed_id;

-- name: GetFeedFollowsForUser :many
SELECT
    feed_follows.*,
    users.name AS user_name,
    feeds.name AS feed_name,
    feeds.url
FROM users
INNER JOIN feed_follows ON users.id = feed_follows.user_id
INNER JOIN feeds ON feeds.id = feed_follows.feed_id
WHERE users.id = $1;

-- name: DeleteFeedFollow :one
DELETE FROM feed_follows
WHERE user_id = $1 AND feed_id = $2
RETURNING *;
