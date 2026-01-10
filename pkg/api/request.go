package api

// DBCredentials defines the fields required to establish
// a connection to a database.
type DBCredentials struct {
	Host     string
	Port     int
	DBName   string
	UserName string
	Password string
}

// BackupRequest defines the fields required to perform
// a database backup operation
type BackupRequest struct {
	DBCredentials DBCredentials // info to connect with db
	TargetPath    string        // where to store the backup files
}

// RestoreRequest defines the fields required to perform
// a database restoration operation
type RestoreRequest struct {
	DBCredentials DBCredentials // info to connect with db
	SourcePath    string        // where to retrieve backup files from
}
