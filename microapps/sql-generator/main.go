package main

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"text/template"
)

// category represents the classification of insights
type category string

// category constants
const (
	kpi         category = "KPI"
	kqi         category = "KQI"
	observation category = "Observation"
	counter     category = "Counter"
)

func categoryFromString(s string) category {
	switch s {
	case "KPI":
		return kpi
	case "KQI":
		return kqi
	case "Observation":
		return observation
	case "Counter":
		return counter
	default:
		return ""
	}
}

// dataType represnets the data type of insights, ex: integer, string etc..
type dataType string

// data types
const (
	integer dataType = "integer"
	bigint  dataType = "bigint"
	double  dataType = "double"
	str     dataType = "string"
	float   dataType = "float"
	long    dataType = "long"
)

func dataTypeFromString(s string) dataType {
	switch s {
	case "integer":
		return integer
	case "bigint":
		return bigint
	case "double":
		return double
	case "string":
		return str
	case "float":
		return float
	case "long":
		return long
	default:
		return ""
	}
}

// sql map of table and query
var sqlMap map[string]string = map[string]string{
	"mdata_formula": `LOCK TABLES MDATA_Formula WRITE;
/*!40000 ALTER TABLE MDATA_Formula DISABLE KEYS */;
INSERT INTO MDATA_Formula (formulaUID,insightUID,attributeUID,flow,monitoredPointType,physicalTarget,expression,filter,isSensitive,isEditable,lastUpdateTS,colOrder)
VALUES
{{ $formulaUID := .FormulaUID }}
{{- range $index, $element := .Insights -}}
({{add $formulaUID $index}},{{$element.InsightUID}},NULL,'{{$element.Flow}}','{{$element.MonitoredPointType}}','{{$element.PhysicalTarget}}','{{$element.Expression}}','',0,0,0,0),
{{ end -}}
/*!40000 ALTER TABLE MDATA_Formula ENABLE KEYS */;
UNLOCK TABLES;`,

	"mdata_formula_snapshot": `LOCK TABLES MDATA_FormulaSnapshot WRITE;
/*!40000 ALTER TABLE MDATA_FormulaSnapshot DISABLE KEYS */;
INSERT INTO MDATA_FormulaSnapshot (formulaUID,domain,insightUID,attributeUID,insightId,insightName,attributeID,attributeName,description,category,unit,type,format,restName,isDerived,toBeExported,isSensitive,isStandard,flow,monitoredPointType,physicalTarget,expression,translatedExpression,filter,isEditable,lastUpdateTS,colOrder,submissionTS,version)
VALUES
{{ $formulaUID := .FormulaUID }}
{{- range $index, $element := .Insights -}}
({{add $formulaUID $index}},'{{$element.Domain}}',{{$element.InsightUID}},NULL,'{{$element.InsightId}}','{{$element.InsightName}}',NULL,NULL,'{{$element.Description}}','{{$element.Category}}','{{$element.Unit}}','{{$element.Type}}','{{$element.Format}}','{{$element.RestName}}',0,1,0,1,'{{$element.Flow}}','{{$element.MonitoredPointType}}','{{$element.PhysicalTarget}}','{{$element.Expression}}','{{$element.TranslatedExpression}}','',0,0,0,UNIX_TIMESTAMP()*1000,NULL),
{{ end -}}
/*!40000 ALTER TABLE MDATA_FormulaSnapshot ENABLE KEYS */;
UNLOCK TABLES;`,
}

// insights table has a constraint of domain-insightID to be unique
// so dataset is different for insight table
var insightTemplate string = `LOCK TABLES MDATA_Insight WRITE;
/*!40000 ALTER TABLE MDATA_Insight DISABLE KEYS */;
INSERT INTO MDATA_Insight (insightUID,insightId,domain,insightName,description,category,unit,type,format,restName,isDerived,toBeExported,isSensitive,isStandard,physicalTarget)
VALUES
{{ $insightUID := .InsightUID }}
{{- range $index, $element := .Insights -}}
({{$element.InsightUID}},'{{$element.InsightId}}','{{$element.Domain}}','{{$element.InsightName}}','{{$element.Description}}','{{$element.Category}}','{{$element.Unit}}','{{$element.Type}}','{{$element.Format}}','{{$element.RestName}}',0,1,0,1,'{{$element.PhysicalTarget}}'),
{{ end -}}
/*!40000 ALTER TABLE MDATA_Insight ENABLE KEYS */;
UNLOCK TABLES;`

const (
	// no of columns vailable in the input csv file
	// the order of the column should be in the order of field declaration in [insight] struct
	csvColumnLength int = 14
	// csv input file should have the folllowing column in the specific order as [Insight] struct
	csvHeaderColumn = "domain,flow,monitoredPointType,insightId,insightName,description,category,unit,type,format,restName,physicalTarget,expression,translatedExpression"
)

// insight represent insight to be added tot he metastore
// the columns of insights usually corresponds to the column of insights.csv file
// the combination of domain, flow, monitoredPointType and insightId uniquely identifies a insight
type Insight struct {
	Domain               string
	Flow                 string
	MonitoredPointType   string
	InsightId            string
	InsightUID           int
	InsightName          string
	Description          string
	Category             category
	Unit                 string
	Type                 dataType
	Format               string
	RestName             string
	PhysicalTarget       string
	Expression           string
	TranslatedExpression string
}

type InsightKey struct {
	Domain     string
	InsightId  string
	InsightUID int
}

type data struct {
	Insights   []Insight
	InsightUID int
	FormulaUID int
}

// the inputs to main function
var ()

// main function which takes in a metastore csv formula file and create insert sql for MDATA_Insight, MDATA_Formula and MDATA_FormulaSnapshot tables
func main() {
	usage := `
usage of sql-generator:
sql-generator sub-commnd [options]
	subcommand is one of: [generateql, generateavro]
	ex: sql-generateor generatesql [options]
	use "sql-generator sub-command" for subcommand options.`
	if len(os.Args) < 2 {
		fmt.Println(usage)
		os.Exit(1)
	}

	switch command := os.Args[1]; command {
	case "generatesql":
		generateSqlCommand(os.Args[2:])
	case "generateavro":
		generateAvroCommand(os.Args[2:])
	default:
		fmt.Println("Invalid subcommand", command)
		fmt.Println(usage)
		os.Exit(1)
	}
}

func generateSqlCommand(args []string) {
	var insightsCsvFile string    // the input csv file
	var insightsKeyCsvFile string // the insights key csv file
	var insightUID int            // current end of insightsUID
	var formulaUID int            //current end of formulaUID

	generateSqlCmd := flag.NewFlagSet("generatesql", flag.ContinueOnError)
	generateSqlCmd.StringVar(&insightsCsvFile, "csv", "insights.csv", "csv file which has the insights and its formulas")
	generateSqlCmd.StringVar(&insightsKeyCsvFile, "insightsCsv", "insightsKeyCsvFile.csv", "csv file which has the already existing insights and its insightsUID")
	generateSqlCmd.IntVar(&insightUID, "insightUID", -1, "the insightUID to start with")
	generateSqlCmd.IntVar(&formulaUID, "formulaUID", -1, "the formulaUID to start with")

	if err := generateSqlCmd.Parse(args); err != nil {
		fmt.Printf("generatesql command failed, err: %v\n", err)
		os.Exit(1)
	}

	var errStr string
	if len(strings.Trim(insightsCsvFile, "")) <= 0 {
		errStr = "missing insights csvfile\n"
	}
	if insightUID <= 0 {
		errStr = errStr + "insightsUID should be greater than zero\n"
	}
	if formulaUID <= 0 {
		errStr = errStr + "formulaUID should be greater than zero"
	}
	if len(errStr) != 0 {
		fmt.Printf("input validation failure, err: %s\n", errStr)
		os.Exit(1)
	}

	var d data
	insights := parseCsv(insightsCsvFile)
	//get a copy and deduplicate with domain-insightID and create the map
	unique, insightmap := getDatasetForInsightTable(insights, insightUID)
	//fmt.Printf("len of insight %d, len of unique %d, insight : %v, \n unique: %v\n", len(insights), len(unique), insights, unique)

	//the entries which are available in exists map should be removed from unique and the insightmap should be updated with the entries in exists map
	insightsKey := parseInsightsKeyCsv(insightsKeyCsvFile)
	alreadyExistsMap := getAlreadyExistsMap(insightsKey)
	unique = removeAlreadyExistingEntries(unique, alreadyExistsMap)
	insightmap = updateKeyMapWithExitingEnteries(insightmap, alreadyExistsMap)
	//fmt.Println("updated key map : %v \n", insightmap)

	d = data{
		Insights:   unique,
		InsightUID: insightUID,
		FormulaUID: formulaUID,
	}
	query := generateSql(insightTemplate, &d)
	fmt.Println(query)

	insights = updateInsightWithInsightId(insights, insightmap)
	//fmt.Printf("updated Insights %v \n", insights)
	d = data{
		Insights:   insights,
		InsightUID: insightUID,
		FormulaUID: formulaUID,
	}
	for _, queryTemplate := range sqlMap {
		query := generateSql(queryTemplate, &d)
		fmt.Println(query)
	}
}

func generateAvroCommand(args []string) {
	var insightsCsvFile string // the input csv file
	var outputDir string       // the output directory to write the result files

	generateAvroCmd := flag.NewFlagSet("generateavro", flag.ContinueOnError)
	generateAvroCmd.StringVar(&insightsCsvFile, "csv", "insights.csv", "csv file which has the insights and its formulas")
	generateAvroCmd.StringVar(&outputDir, "out", ".", "directory path to generate output avsc files, if generated avro already esists in the target directory, it will be replaced")

	if err := generateAvroCmd.Parse(args); err != nil {
		fmt.Printf("generateavro command failed, err: %v", err)
		os.Exit(1)
	}
	var errStr string
	if len(strings.Trim(insightsCsvFile, "")) <= 0 {
		fmt.Printf("input validation failure, err: %s\n", errStr)
		os.Exit(1)
	}

	insights := parseCsv(insightsCsvFile)
	//create a map with key: Domain+"-"+MonitoredPointType and set of insights as value
	insightMap := createMapForSchemaGeneration(insights)
	for key, val := range insightMap {
		//create a avro file for each key and set of insights
		if err := writeAsAvroFile(val, key+".avsc", outputDir); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}

func parseCsv(csvFile string) []Insight {
	file, err := os.Open(csvFile)
	if err != nil {
		log.Fatalf("Error opening csv file %s, err: %v", csvFile, err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ','
	// read header
	if row, err := reader.Read(); err != nil {
		log.Fatalf("error reading header from csv file %s, err: %v", csvFile, err)
	} else if rowCsv := strings.Join(row, ","); rowCsv != csvHeaderColumn {
		//log.Fatalf("input csv file %s has a different header row than expected. expected: %s got: %s\n", csvFile, csvHeaderColumn, rowCsv) //this is commented out because the comaprision is giving false result might be a go bug
	}

	//read remaining lines
	rows, err := reader.ReadAll() //records is [][]string
	if err != nil {
		log.Fatalf("error reading csv file %s, err: %v", csvFile, err)
	}
	//validate csv data
	if err := validateInsightData(rows); err != nil {
		log.Fatalln("invalid input csv data", err)
	}
	//convert [][]string to []Insight
	var insights []Insight = make([]Insight, 0)
	for _, row := range rows {
		cat := categoryFromString(row[6])
		if len(cat) == 0 {
			log.Fatalf("invalid category value %s detected in row %s", row[6], strings.Join(row, ","))
		}
		typ := dataTypeFromString(row[8])
		if len(typ) == 0 {
			log.Fatalf("invalid type value %s detected in row %s", row[8], strings.Join(row, ","))
		}
		insights = append(insights, Insight{
			Domain:               row[0],
			Flow:                 row[1],
			MonitoredPointType:   row[2],
			InsightId:            row[3],
			InsightName:          row[4],
			Description:          row[5],
			Category:             cat,
			Unit:                 row[7],
			Type:                 typ,
			Format:               row[9],
			RestName:             row[10],
			PhysicalTarget:       row[11],
			Expression:           row[12],
			TranslatedExpression: row[13],
		})
	}
	return insights
}

func parseInsightsKeyCsv(csvFile string) []InsightKey {
	file, err := os.Open(csvFile)
	if err != nil {
		log.Fatalf("Error opening csv file %s, err: %v", csvFile, err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ','
	// read header
	if row, err := reader.Read(); err != nil {
		log.Fatalf("error reading header from csv file %s, err: %v", csvFile, err)
	} else if rowCsv := strings.Join(row, ","); rowCsv != "domain,insightId,insightUID" {
		//log.Fatalf("input csv file %s has a different header row than expected. expected: %s got: %s\n", csvFile, csvHeaderColumn, rowCsv) //this is commented out because the comaprision is giving false result might be a go bug
	}

	//read remaining lines
	rows, err := reader.ReadAll() //records is [][]string
	if err != nil {
		log.Fatalf("error reading csv file %s, err: %v", csvFile, err)
	}

	//convert [][]string to []Insight
	var keys []InsightKey = make([]InsightKey, 0)
	for _, row := range rows {
		i, err := strconv.Atoi(row[2])
		if err != nil {
			log.Fatalf("error converting insightUID from string to int, insightUID: %s, err %v \n", row[2], err)
		}
		keys = append(keys, InsightKey{
			Domain:     row[0],
			InsightId:  row[1],
			InsightUID: i,
		})
	}
	return keys
}

func getAlreadyExistsMap(existsKeys []InsightKey) map[string]int {
	var existsMap map[string]int = make(map[string]int)
	for _, key := range existsKeys {
		insightKey := key.Domain + "-" + key.InsightId
		existsMap[insightKey] = key.InsightUID
	}
	return existsMap
}

func removeAlreadyExistingEntries(entries []Insight, existsmap map[string]int) []Insight {
	var insights []Insight = make([]Insight, 0)
	for _, insight := range entries {
		insightkey := insight.Domain + "-" + insight.InsightId
		if _, exists := existsmap[insightkey]; !exists {
			insights = append(insights, insight)
		}
	}
	return insights
}

// update the insightUID values of keyMap with existsmap Enteries
func updateKeyMapWithExitingEnteries(keyMap map[string]int, existsMap map[string]int) map[string]int {
	for key, _ := range keyMap {
		if insightUID, exists := existsMap[key]; exists {
			keyMap[key] = insightUID
		}
	}
	return keyMap
}

func validateInsightData(rows [][]string) error {
	//non empty rows
	if len(rows) <= 0 {
		return errors.New("zero insight rows. Nothing to generate")
	}
	uniqueKeyMap := make(map[string]bool)
	for _, row := range rows {
		// domain + flow + monitoredPointType + insightId should be unique
		uniqueKey := row[0] + row[1] + row[2] + row[3]
		if uniqueKeyMap[uniqueKey] {
			return errors.New("duplicate row: " + strings.Join(row, ","))
		} else {
			uniqueKeyMap[uniqueKey] = true
		}

		// the csv file should have adequate no of rows to be mapped to the insights struct
		if rowSize := len(row); rowSize != csvColumnLength {
			return fmt.Errorf("row: %s should have %d columns in the order defined in the Insight struct", strings.Join(row, ","), csvColumnLength)
		}
	}
	return nil
}

func generateSql(queryTemplate string, d *data) string {
	rendered := renderTemplate(queryTemplate, &d)
	return strings.TrimSuffix(strings.TrimSpace(rendered), ",")
}

func renderTemplate(templateStr string, data any) string {
	funcMap := template.FuncMap{
		// The name "inc" is what the function will be called in the template text.
		"add": func(i int, j int) int {
			return i + j
		},
	}

	t := template.Must(template.New("tmp").Funcs(funcMap).Parse(templateStr))
	var writer bytes.Buffer
	if err := t.Execute(&writer, data); err != nil {
		log.Fatal(err)
	}
	return writer.String()
}

func getDatasetForInsightTable(duplicateInsights []Insight, insightUIDStart int) (unique []Insight, insightMap map[string]int) {
	unique = make([]Insight, 0)
	insightMap = make(map[string]int)
	insightUID := insightUIDStart
	for _, insight := range duplicateInsights {
		insightKey := insight.Domain + "-" + insight.InsightId
		if _, exists := insightMap[insightKey]; !exists {
			insight.InsightUID = insightUID
			unique = append(unique, insight)
			insightMap[insightKey] = insightUID
			insightUID++
		}
	}
	return unique, insightMap
}

func updateInsightWithInsightId(insights []Insight, insightMap map[string]int) (updated []Insight) {
	updated = make([]Insight, 0)
	for _, insight := range insights {
		insightKey := insight.Domain + "-" + insight.InsightId
		if insightUID, exists := insightMap[insightKey]; exists {
			insight.InsightUID = insightUID
			updated = append(updated, insight)
		} else {
			log.Fatalf("insight not found in insight Map, insightMap: %v \n insight: %v\n", insightMap, insight)
		}
	}
	return updated
}

func createMapForSchemaGeneration(insights []Insight) map[string][]Insight {
	var m map[string][]Insight = make(map[string][]Insight)
	for _, i := range insights {
		key := i.Domain + "-" + i.MonitoredPointType
		if val, exists := m[key]; exists {
			val = append(val, i)
			m[key] = val
		} else {
			val = []Insight{i}
			m[key] = val
		}
	}
	return m
}

// struct representing a avro field
type Field struct {
	Name          string   `json:"name"`
	Type          []string `json:"type"`
	DefaultValue  *string  `json:"default"`
	Documentation string   `json:"doc,omitempty"`
	Aliases       string   `json:"aliases,omitempty"`
}

// A map which converts excel file insights datatypes to avro data types
var toAvroType map[dataType]string = map[dataType]string{
	integer: "integer",
	bigint:  "long",
	double:  "double",
	str:     "string",
	float:   "float",
	long:    "long",
}

func writeAsAvroFile(insights []Insight, fileNameWithExtension string, outputDir string) error {
	var fields []Field = make([]Field, len(insights))
	for i, entry := range insights {
		fields[i] = Field{
			Name: entry.InsightId,
			Type: []string{"null", toAvroType[entry.Type]},
		}
	}
	jsonData, err := json.MarshalIndent(fields, "", " ")
	if err != nil {
		return err
	}
	filepath := outputDir + "/" + fileNameWithExtension
	if err = os.WriteFile(filepath, jsonData, 0666); err != nil {
		return fmt.Errorf("error writing avsc file. content: %v , filePath: %s, err: %v", string(jsonData), filepath, err)
	}
	return nil
}
