package repositories

import (
	"database/sql"
	"fmt"
	"log"

	"forum/backend/models"
)

func GetReactions(db *sql.DB, id int, react string) ([]models.Reaction, error) {
	query := `
		SELECT * FROM tblReactions
		WHERE post_id = ? AND reaction = ? AND reaction_status = 'clicked'
	`
	rows, err := db.Query(query, id, react)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	var Reactions []models.Reaction

	for rows.Next() {
		reaction := models.Reaction{}

		err := rows.Scan(&reaction.ID, &reaction.Reaction, &reaction.ReactionStatus, &reaction.UserID, &reaction.PostID)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		Reactions = append(Reactions, reaction)
	}

	// Check for errors after iteration
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %w", err)
	}

	return Reactions, nil
}

// write a function that retrieves the number of likes and dislikes 
func GetLikesAndDislikes(db *sql.DB, postId int) (int, int, error) {
	query := `
		SELECT 
			SUM(CASE WHEN reaction = 'Like' THEN 1 ELSE 0 END) AS likes,
			SUM(CASE WHEN reaction = 'Dislike' THEN 1 ELSE 0 END) AS dislikes
		FROM tblReactions
		WHERE post_id = ? AND reaction_status = 'clicked'
	`
	row := db.QueryRow(query, postId)

	var likes, dislikes int
	err := row.Scan(&likes, &dislikes)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to execute query: %w", err)
	}

	return likes, dislikes, nil
}

func CheckReactions(db *sql.DB, userId, postId int) (bool, string) {
	query := `
    SELECT * FROM tblReactions
    WHERE post_id = ? AND user_id = ?
`

	rows, err := db.Query(query, postId, userId)
	if err != nil {
		log.Println("error executing query:", err)
		return false, ""
	}
	defer rows.Close()

	if !rows.Next() {
		log.Println("No matching rows found")
		return false, ""
	}

	reaction := models.Reaction{}

	err = rows.Scan(&reaction.ID, &reaction.Reaction, &reaction.ReactionStatus, &reaction.UserID, &reaction.PostID)
	if err != nil {
		return false, ""
	}

	return true, reaction.Reaction
}

func UpdateReaction(db *sql.DB, reaction string, userId, postId int) error {
	query := `
	UPDATE tblReactions
	SET reaction = ?
	 WHERE post_id = ? AND user_id = ?
	`
	_, err := db.Exec(query, reaction, postId, userId)
	if err != nil {
		log.Println("error executing query:", err)
		return err
	}
	return nil
}

func UpdateReactionStatus(db *sql.DB, userId, postId int) error {
	query := `
    UPDATE tblReactions
    SET reaction_status = 
        CASE 
            WHEN reaction_status = 'clicked' THEN 'unclicked'
            ELSE 'clicked'
        END
    WHERE post_id = ? AND user_id = ?
`

	_, err := db.Exec(query, postId, userId)
	if err != nil {
		log.Println("Error executing query:", err)
		return err
	}
	return nil
}

// AddReaction adds a reaction to the database

func InsertReaction(db *sql.DB, reaction models.Reaction) error {
	query := `
	INSERT INTO tblReactions (reaction, reaction_status, user_id, post_id)
	VALUES (?, 'clicked', ?, ?)
`

	_, err := db.Exec(query, reaction.Reaction, reaction.UserID, reaction.PostID)
	if err != nil {
		log.Println("Error executing query:", err)
		return err
	}
	return nil
}