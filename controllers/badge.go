package controllers

import (
	"fmt"
	"github.com/ArchCI/archci/models"
	log "github.com/Sirupsen/logrus"
)

// GetProjectBadge takes project id and return the badge image of its status.
func (c *ApiController) GetProjectBadge() {
	log.Info("Get project badge")

	// TODO(tobe): need more svg images.
	success_svg := []byte(`<svg xmlns="http://www.w3.org/2000/svg" width="104" height="20"><linearGradient id="b" x2="0" y2="100%"><stop offset="0" stop-color="#bbb" stop-opacity=".1"/><stop offset="1" stop-opacity=".1"/></linearGradient><mask id="a"><rect width="104" height="20" rx="3" fill="#fff"/></mask><g mask="url(#a)"><path fill="#555" d="M0 0h54v20H0z"/><path fill="#4c1" d="M54 0h50v20H54z"/><path fill="url(#b)" d="M0 0h104v20H0z"/></g><g fill="#fff" text-anchor="middle" font-family="DejaVu Sans,Verdana,Geneva,sans-serif" font-size="11"><text x="28" y="15" fill="#010101" fill-opacity=".3">solution</text><text x="28" y="14">solution</text><text x="78" y="15" fill="#010101" fill-opacity=".3">correct</text><text x="78" y="14">correct</text></g></svg>`)
	fail_svg := []byte(`<svg xmlns="http://www.w3.org/2000/svg" width="99" height="20"><linearGradient id="b" x2="0" y2="100%"><stop offset="0" stop-color="#bbb" stop-opacity=".1"/><stop offset="1" stop-opacity=".1"/></linearGradient><mask id="a"><rect width="99" height="20" rx="3" fill="#fff"/></mask><g mask="url(#a)"><path fill="#555" d="M0 0h54v20H0z"/><path fill="#e05d44" d="M54 0h45v20H54z"/><path fill="url(#b)" d="M0 0h99v20H0z"/></g><g fill="#fff" text-anchor="middle" font-family="DejaVu Sans,Verdana,Geneva,sans-serif" font-size="11"><text x="28" y="15" fill="#010101" fill-opacity=".3">solution</text><text x="28" y="14">solution</text><text x="75.5" y="15" fill="#010101" fill-opacity=".3">wrong</text><text x="75.5" y="14">wrong</text></g></svg>`)

	var badge []byte

	projectId, _ := c.GetInt64(":projectId")
	project := models.GetProjectWithId(projectId)

	if project.Status == models.PROJECT_STATUS_NEED_TEST {
		badge = success_svg
	} else if project.Status == models.PROJECT_STATUS_TESTING {
		badge = success_svg
	} else if project.Status == models.PROJECT_STATUS_SUCCESS {
		badge = success_svg
	} else {
		badge = fail_svg
	}

	header := c.Ctx.ResponseWriter.Header()
	header.Set("Content-Type", "image/svg+xml")
	header.Set("Cache-Control", "no-cache, no-store, must-revalidate")

	c.Ctx.ResponseWriter.Write(badge)
}

// GetProjectBadgeUrl returns the url of the status badge.
func (c *ApiController) GetProjectBadgeUrl() {
	log.Debug("Get project badge url")

	// TODO(tobe): use the address of server instead of 127.0.0.1
	// The link should look like "http://127.0.0.1:10010/v1/badge/1)"
	projectId, _ := c.GetInt64(":projectId")
	link := fmt.Sprintf("http://127.0.0.1:10010/v1/badge/%d", projectId)

	c.Ctx.WriteString(link)
}

// GetProjectBadgeMarkdown returns the markdown link of the status badge.
func (c *ApiController) GetProjectBadgeMarkdown() {
	log.Debug("Get project badge with markdown link")

	// The link should look like "[![ArchCI](http://127.0.0.1:10010/v1/badge/1)](http://127.0.0.1:10010/projects/1)"
	projectId, _ := c.GetInt64(":projectId")
	link := fmt.Sprintf("[![ArchCI](http://127.0.0.1:10010/v1/badge/%d)](http://127.0.0.1:10010/projects/%d)", projectId, projectId)

	c.Ctx.WriteString(link)
}
