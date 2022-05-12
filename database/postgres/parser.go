package postgres

import (
	"github.com/k0kubun/sqldef/database"
	"github.com/k0kubun/sqldef/parser"
)

type PostgresParser struct {
	parser database.GenericParser
}

func NewParser() PostgresParser {
	return PostgresParser{
		parser: database.NewParser(parser.ParserModePostgres),
	}
}

func (p PostgresParser) Parse(sql string) ([]database.DDLStatement, error) {
	return p.parser.Parse(sql)
}
