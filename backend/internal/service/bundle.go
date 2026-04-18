package service

import (
	"database/sql"
	"errors"
	"strings"
	"time"

	"github.com/Lbringer-code/oneLink/backend/internal/domain"
)

var (
	ErrValidation = errors.New("validation error")
	ErrNotFound = errors.New("bundle not found")
	ErrInternal = errors.New("internal server error")
)

func (s *Service) CreateBundle( req domain.CreateBundleRequest ) ( *domain.CreateBundleResponse , error ) {
	err := validateCreateRequest(req)
	if err != nil {
		return nil , errors.Join( ErrValidation , err ) 
	}

	slug , err := generateSlug(req.Title)
	if err != nil {
		return nil , errors.Join( ErrInternal , err )
	}

	now := time.Now()
	
	bundle := domain.BundleDB{
		Slug: slug,
		Title: req.Title,
		CreatedAt: now,
		LastAccessed: now,
	}

	links := make( []domain.LinkDB , len(req.Links) )
	for i , l := range req.Links {
		links[i] = domain.LinkDB{
			BundleSlug: slug,
			Url: strings.TrimSpace(l.Url),
			Note: l.Note,
			DisplayText: l.DisplayText,
			CreatedAt: now,
		}
	}

	err = s.repo.CreateBundleWithLinks(bundle , links)
	if err != nil {
		return nil , errors.Join( ErrInternal , err )
	}

	return &domain.CreateBundleResponse{
		CreatedAt: now,
		Slug: slug,
		Title: req.Title,
	} , nil
}

func (s *Service) GetBundle( slug string) ( *domain.GetBundleResponse , error ) {
	slug = strings.TrimSpace(slug)
	if slug == "" {
		return nil , errors.Join( ErrValidation , errors.New("slug cannot be empty") )
	}

	bundle , links , err := s.repo.GetBundleWithLinks( slug )
	if err != nil {
		if errors.Is( err , sql.ErrNoRows ) {
			return nil , ErrNotFound
		}
		return nil , errors.Join( ErrInternal , err )
	}

	_ = s.repo.UpdateLastAccessed( slug , time.Now() )

	apiLinks := make([]domain.Link , len(links))
	for i , l := range links {
		apiLinks[i] = domain.Link{
			DisplayText: l.DisplayText,
			Note: l.Note,
			Url: l.Url,
		}
	}

	return &domain.GetBundleResponse{
		CreatedAt: bundle.CreatedAt,
		Links: apiLinks,
		Slug: bundle.Slug,
		Title: bundle.Title,
	} , nil
}

func (s *Service) CleanupStaleBundles( cutoff time.Time ) (int64 , error) {
	count , err := s.repo.DeleteStaleBundles(cutoff)
	if err != nil {
		return 0 , errors.Join( ErrInternal , err )
	}
	return count , nil
}