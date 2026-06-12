package labelname

import "strings"

type AliasRankRequest struct {
	TaskTimestamp     string
	ItemIndex         int
	TaxID             string
	SearchTerm        string
	Symbol            string
	ProteinID         string
	GeneID            string
	TranscriptID      string
	SequenceID        string
	LocusTag          string
	Aliases           []string
	Synonyms          []string
	DBXrefs           []string
	Chromosome        string
	MapLocation       string
	Description       string
	TypeOfGene        string
	SymbolAuthority   string
	FullNameAuthority string
	OtherDesignations []string
	FeatureType       string
}

type AliasRankResult struct {
	TaskTimestamp string
	ItemIndex     int
	RankedAliases []string
}

type rankedAlias struct {
	Text   string
	Score  int
	Family string
}

func RankAliases(request AliasRankRequest) AliasRankResult {
	ranked := rankAliasRequestItems(request)
	return AliasRankResult{
		TaskTimestamp: request.TaskTimestamp,
		ItemIndex:     request.ItemIndex,
		RankedAliases: rankedAliasTexts(ranked),
	}
}

func rankAliasRequestItems(request AliasRankRequest) []rankedAlias {
	if db, ok := openDefaultGeneDB(); ok {
		if ranked, handled := db.rank(request); handled {
			return ranked
		}
	}
	return nil
}

func rankedAliasTexts(items []rankedAlias) []string {
	out := make([]string, 0, len(items))
	seen := make(map[string]struct{}, len(items))
	for _, item := range items {
		text := strings.TrimSpace(item.Text)
		if text == "" {
			continue
		}
		key := normalizeAliasKey(text)
		if _, ok := seen[key]; ok {
			continue
		}
		seen[key] = struct{}{}
		out = append(out, text)
	}
	return out
}

func normalizeAliasKey(value string) string {
	return strings.ToUpper(strings.TrimSpace(value))
}

func symbolFamily(value string) string {
	value = strings.TrimSpace(value)
	if value == "" {
		return ""
	}
	runes := []rune(value)
	end := len(runes)
	for end > 0 {
		r := runes[end-1]
		if r >= '0' && r <= '9' {
			end--
			continue
		}
		break
	}
	for end > 0 {
		r := runes[end-1]
		if r == '-' || r == '_' || r == '.' {
			end--
			continue
		}
		break
	}
	prefix := strings.TrimSpace(string(runes[:end]))
	if prefix == "" || prefix == value {
		return ""
	}
	hasLetter := false
	for _, r := range prefix {
		if (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') {
			hasLetter = true
			break
		}
	}
	if !hasLetter {
		return ""
	}
	return strings.ToUpper(prefix)
}
