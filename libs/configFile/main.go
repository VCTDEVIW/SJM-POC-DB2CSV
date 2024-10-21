package project

type Config struct {
    Configuration struct {
        OutputPath string `json:"outputPath"`
    } `json:"configuration"`
    SQLSrv struct {
        SQLDBHost     string `json:"sql-db-host"`
        SQLDBPort     int    `json:"sql-db-port"`
        SQLDBUsername string `json:"sql-db-username"`
        SQLDBPassword string `json:"sql-db-password"`
        SQLDBDbname   string `json:"sql-db-dbname"`
        SQLDBStname   string `json:"sql-db-stname"`
    } `json:"sqlsrv"`
    MongoDB interface{} `json:"mongodb"` // Empty object
    Misc    interface{} `json:"misc"`    // Empty object
}
