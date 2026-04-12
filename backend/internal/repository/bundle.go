package repository

import (
	"github.com/Lbringer-code/oneLink/backend/internal/domain"
)

func (r *Repository) CreateBundleWithLinks(bundle domain.BundleDB , links []domain.LinkDB) (error){
	tx , err := r.db.Beginx()
	if err != nil{
		return err
	}
	
	defer tx.Rollback()

	_ , err = tx.NamedExec(
		`INSERT INTO bundle ( slug , title , created_at , last_accessed )
		VALUES ( :slug , :title , :created_at , :last_accessed )`, 
		bundle ,
	)
	if err != nil {
		return err
	}

	for _ , link := range links {
		_ , err = tx.NamedExec(
			`INSERT INTO link ( bundle_slug , url , note , display_text , created_at )
			VALUES ( :bundle_slug , :url , :note , :display_text )` , 
			link,
		)
		if err != nil{
			return err
		}
	}

	return tx.Commit()
}

func (r *Repository) GetBundleWithLinks(slug string) (*domain.BundleDB , *[]domain.LinkDB , error ){
	var bundle domain.BundleDB
	err := r.db.Get(
		&bundle ,
		`SELECT * FROM bundle
		WHERE slug = $1` ,
		slug ,
	)
	if err != nil {
		return nil , nil , err
	}
	
	var links []domain.LinkDB
	err = r.db.Select(
		&links , 
		`SELECT * FROM link
		WHERE bundle_slug = $1` , 
		slug , 
	)
	if err != nil {
		return nil , nil , err
	}

	return &bundle , &links , nil
}