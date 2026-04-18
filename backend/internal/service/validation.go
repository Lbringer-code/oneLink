package service

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/Lbringer-code/oneLink/backend/internal/domain"
)

const (
	maxTitleLen = 120
	maxDisplayTitleLen = 200
	maxNoteLen = 500
	maxLinks = 50
)

func validateCreateRequest(req domain.CreateBundleRequest) error {

	title := strings.TrimSpace(req.Title)
	if title == "" {
		return fmt.Errorf("title cannot be empty")
	}
	if len(title) > maxTitleLen {
		return fmt.Errorf("title cannot exceed %d characters" , maxTitleLen)
	}

	if len(req.Links) == 0 {
		return fmt.Errorf("at least one link is required")
	}
	if len(req.Links) > maxLinks {
		return fmt.Errorf("number of links cannot exceed %d links per bundle" , maxLinks)
	}

	seen := make(map[string]bool)
	for i , link := range req.Links {
		trimmedURL := strings.TrimSpace(link.Url)
		if trimmedURL == "" {
			return fmt.Errorf("link %d: url cannot be empty" , i + 1)
		}

		parsed , err := url.ParseRequestURI(trimmedURL)
		if err != nil || ( parsed.Scheme != "http" && parsed.Scheme != "https" ) {
			return fmt.Errorf("link %d: invalid URL" , i + 1)
		}

		if seen[trimmedURL] {
			return fmt.Errorf("link %d: duplicate URL" , i + 1)
		}
		seen[trimmedURL] = true

		if link.DisplayText != nil && len(*link.DisplayText) > maxDisplayTitleLen {
			return fmt.Errorf("link %d: display text cannot exceed %d characters" , i + 1 , maxDisplayTitleLen)
		}

		if link.Note != nil && len(*link.Note) > maxNoteLen {
			return fmt.Errorf("link %d: note cannot exceed %d characters" , i + 1 , maxNoteLen)
		}

	}
	return nil
}
