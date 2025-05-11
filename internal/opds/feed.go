package opds

import (
	"fmt"
	"time"
)

// GenerateRootCatalogFeed returns the XML for the root OPDS catalog feed.
func GenerateRootCatalogFeed() string {
	now := time.Now().UTC().Format(time.RFC3339)
	return fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8"?>
<feed xmlns="http://www.w3.org/2005/Atom" xmlns:opds="http://opds-spec.org/2010/catalog">
  <title>OPDS Root Catalog</title>
  <id>urn:uuid:root-catalog</id>
  <updated>%s</updated>
  <link rel="self" href="/opds" type="application/atom+xml;profile=opds-catalog;kind=acquisition"/>
  <link rel="search" href="/opds/search" type="application/atom+xml;profile=opds-catalog;kind=acquisition"/>
</feed>`, now)
}

// GenerateSearchResultsFeed returns the XML for the OPDS search results feed based on the provided query.
func GenerateSearchResultsFeed(query string) string {
	now := time.Now().UTC().Format(time.RFC3339)
	return fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8"?>
<feed xmlns="http://www.w3.org/2005/Atom" xmlns:opds="http://opds-spec.org/2010/catalog">
  <title>Search Results for: %s</title>
  <id>urn:uuid:search-results</id>
  <updated>%s</updated>
  <link rel="self" href="/opds/search?q=%s" type="application/atom+xml;profile=opds-catalog;kind=acquisition"/>
  <entry>
    <title>Sample Book</title>
    <author>
      <name>Sample Author</name>
    </author>
    <id>urn:uuid:sample-book</id>
    <updated>%s</updated>
    <link rel="http://opds-spec.org/acquisition" href="/opds/book/1" type="application/epub+zip" title="Download EPUB"/>
  </entry>
</feed>`, query, now, query, now)
} 