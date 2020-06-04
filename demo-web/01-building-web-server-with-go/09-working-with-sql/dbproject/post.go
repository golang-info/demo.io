package main

import "database/sql"

// Post: A record in Posts table
type Post struct {
	ID int
	Title string
	Content string
}

// PostDB type to wrap database connection
type PostDB struct {
	DB *sql.DB
}

// GetAllPosts Get All Posts possible only top 10
func (p *PostDB) GetAllPosts() ([]*Post, error) {
	stmt := `SELECT id, title, content FROM posts`

	rows, err := p.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	posts := []*Post{}
	for rows.Next() {
		p := &Post{}

		err = rows.Scan(&p.ID, &p.Title, &p.Content)
		if err != nil {
			return nil, err
		}
		posts = append(posts, p)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}
