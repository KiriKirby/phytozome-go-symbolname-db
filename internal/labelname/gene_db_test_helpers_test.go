package labelname

import (
	"compress/gzip"
	"net/http"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func buildTestGeneInfoDB(t testing.TB, content string) string {
	t.Helper()
	dir := t.TempDir()
	gzPath := filepath.Join(dir, "gene_info.gz")
	file, err := os.Create(gzPath)
	if err != nil {
		t.Fatalf("create gzip: %v", err)
	}
	gz := gzip.NewWriter(file)
	if _, err := gz.Write([]byte(content)); err != nil {
		t.Fatalf("write gzip: %v", err)
	}
	if err := gz.Close(); err != nil {
		t.Fatalf("close gzip: %v", err)
	}
	if err := file.Close(); err != nil {
		t.Fatalf("close file: %v", err)
	}
	dbPath := filepath.Join(dir, "symbolname.pgd")
	lastModified := time.Date(2026, 6, 10, 0, 0, 0, 0, time.UTC)
	if err := buildGeneInfoDatabaseFromGZ(gzPath, dbPath, GeneInfoMetadata{
		URL:             GeneInfoURL,
		LastModified:    lastModified,
		LastModifiedRaw: lastModified.Format(http.TimeFormat),
		ContentLength:   int64(len(content)),
	}, DownloadOptions{}); err != nil {
		t.Fatalf("build gene db: %v", err)
	}
	return dbPath
}

func stringsJoinLines(lines ...string) string {
	out := ""
	for _, line := range lines {
		out += line + "\n"
	}
	return out
}
