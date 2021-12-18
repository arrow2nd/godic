package godic

// 記法オプション
const (
	PascalCase         = "pascal"           // PascalCase
	CamelCase          = "camel"            // camelCase
	SnakeCase          = "lower underscore" // snake_case
	ScreamingSnakeCase = "upper underscore" // SNAKE_CASE
	KebabCase          = "hyphen"           // kabab-case
)

// 接頭語オプション
const (
	LiteralPrefix     = "literal"              // リテラル（変換なし）
	MSPrefix          = "MS naming guidelines" // MSネーミングガイドライン準拠
	CamelStrictPrefix = "camel strict"         // CamelCase準拠
)
