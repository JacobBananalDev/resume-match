package analyze

/*
SkillBank is the list of canonical skills
recognized by the analyzer.
*/
var SkillBank = []string{
	"go",
	"docker",
	"kubernetes",
	"aws",
	"postgres",
	"react",
	"python",
	"terraform",
	"javascript",
	"typescript",
	"sql",
	"api",
	"distributed",
	"database",
}

/*
SkillSynonyms maps alternate skill names
to canonical names.
*/
var SkillSynonyms = map[string]string{
	"golang": "go",
	"k8s": "kubernetes",
	"js": "javascript",
	"ts": "typescript",

	"postgresql": "postgres",
	"psql": "postgres",
}

/*
SemanticSkills expands related technologies.
*/
var SemanticSkills = map[string][]string{
	"postgres": {"sql", "database"},
	"mysql": {"sql", "database"},
	"redis": {"database", "cache"},
	"grpc": {"api"},
	"rest": {"api"},
	"microservices": {"distributed"},
	"docker": {"containers"},
	"kubernetes": {"containers"},
}

/*
NormalizeSkill converts synonyms into
canonical skill names.
*/
func NormalizeSkill(skill string) string {

	if normalized, exists := SkillSynonyms[skill]; exists {
		return normalized
	}

	return skill
}