package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

var (
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
)

type SSHConfig struct {
	SSHUser        string
	SSHPassword    string
	SSHHost        string
	SSHPort        int
	PrivateKeyPath string
}

type DBConfig struct {
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     int
	DBName     string
}

type Job struct {
	tableName string
}

type Result struct {
	tableName string
	err       error
}

type WorkerPool struct {
	numWorkers int
	jobs       chan Job
	results    chan Result
	db         *sql.DB
	wg         *sync.WaitGroup
	jobRecord  *ProcessJobCount
}

type ProcessJobCount struct {
	totalJob   int
	successJob int
	failedJob  int
}

func NewWorkerPool(numWorkers int, db *sql.DB, jobRecord *ProcessJobCount) *WorkerPool {
	return &WorkerPool{
		numWorkers: numWorkers,
		jobs:       make(chan Job, 40),
		results:    make(chan Result, 40),
		db:         db,
		wg:         &sync.WaitGroup{},
		jobRecord:  jobRecord,
	}
}

func (wp *WorkerPool) AddJob(job Job) {
	wp.jobs <- job
}

func (wp *WorkerPool) Start() {
	for i := 0; i < wp.numWorkers; i++ {
		go wp.worker()
	}
	wp.wg.Add(wp.numWorkers)
}

func (wp *WorkerPool) worker() {
	defer wp.wg.Done()
	for job := range wp.jobs {
		// err := migrateData(wp.db, job.tableName, wp.jobRecord)
		fmt.Println("Task Done")
		wp.results <- Result{tableName: job.tableName, err: nil}
	}
}

// func createSSHTunnel(sshConfig SSHConfig, dbConfig DBConfig) (string, func(), error) {
// 	var authMethod ssh.AuthMethod

// 	if sshConfig.PrivateKeyPath != "" {
// 		key, err := os.ReadFile(sshConfig.PrivateKeyPath)
// 		if err != nil {
// 			return "", nil, fmt.Errorf("unable to read private key: %w", err)
// 		}
// 		signer, err := ssh.ParsePrivateKey(key)
// 		if err != nil {
// 			return "", nil, fmt.Errorf("unable to parse private key: %w", err)
// 		}
// 		authMethod = ssh.PublicKeys(signer)
// 	} else {
// 		authMethod = ssh.Password(sshConfig.SSHPassword)
// 	}

// 	sshClientConfig := &ssh.ClientConfig{
// 		User:            sshConfig.SSHUser,
// 		Auth:            []ssh.AuthMethod{authMethod},
// 		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
// 		Timeout:         10 * time.Second,
// 	}

// 	sshAddress := fmt.Sprintf("%s:%d", sshConfig.SSHHost, sshConfig.SSHPort)
// 	sshClient, err := ssh.Dial("tcp", sshAddress, sshClientConfig)
// 	if err != nil {
// 		return "", nil, fmt.Errorf("failed to connect to SSH server: %w", err)
// 	}

// 	listener, err := net.Listen("tcp", "127.0.0.1:0")
// 	if err != nil {
// 		return "", nil, fmt.Errorf("failed to listen on local port: %w", err)
// 	}

// 	go func() {
// 		for {
// 			localConn, err := listener.Accept()
// 			if err != nil {
// 				return
// 			}
// 			remoteConn, err := sshClient.Dial("tcp", fmt.Sprintf("%s:%d", dbConfig.DBHost, dbConfig.DBPort))
// 			if err != nil {
// 				log.Printf("failed to dial remote DB over SSH: %v", err)
// 				localConn.Close()
// 				continue
// 			}
// 			go copyConnection(localConn, remoteConn)
// 			go copyConnection(remoteConn, localConn)
// 		}
// 	}()

// 	localAddress := listener.Addr().String()
// 	closeTunnel := func() {
// 		listener.Close()
// 		sshClient.Close()
// 	}

// 	return localAddress, closeTunnel, nil
// }

// func copyConnection(src, dst net.Conn) {
// 	defer src.Close()
// 	defer dst.Close()

// 	_, err := io.Copy(src, dst)
// 	if err != nil {
// 		log.Printf("Error while copying data: %v", err)
// 	}
// }

// func migrateData(db *sql.DB, tableName string, pc *ProcessJobCount) error {
// 	query := fmt.Sprintf("SELECT record_id, %s.event_type_id, %s.event_category_id FROM %s WHERE 1=1", tableName, tableName, tableName)
// 	rows, err := db.Query(query)
// 	if err != nil {
// 		pc.failedJob++
// 		return fmt.Errorf("failed to query %s: %v", tableName, err)
// 	}
// 	defer rows.Close()

// 	pc.totalJob++
// 	pc.successJob++

// 	InfoLogger.Printf("Migrated data for %s", tableName)
// 	return nil
// }

func main() {
	// dbHost := flag.String("db_host", "", "database host")
	// dbPort := flag.Int("db_port", 3306, "database port")
	// dbUser := flag.String("db_user", "", "database user")
	// dbPass := flag.String("db_pw", "", "database password")
	// dbName := flag.String("db_name", "", "database name")
	// sshUser := flag.String("ssh_user", "", "ssh user")
	// sshHost := flag.String("ssh_host", "", "ssh host")
	// sshPort := flag.Int("ssh_port", 22, "ssh port")
	// sshPrivateKeyPath := flag.String("ssh_pvt_key_path", "", "ssh private key path")

	flag.Parse()

	// sshConfig := SSHConfig{SSHUser: *sshUser, SSHHost: *sshHost, SSHPort: *sshPort, PrivateKeyPath: *sshPrivateKeyPath}
	// dbConfig := DBConfig{DBUser: *dbUser, DBPassword: *dbPass, DBHost: *dbHost, DBPort: *dbPort, DBName: *dbName}

	// localAddress, closeTunnel, err := createSSHTunnel(sshConfig, dbConfig)
	// if err != nil {
	// 	log.Fatalf("Failed to create SSH tunnel: %v", err)
	// }
	// defer closeTunnel()

	// db, err := connectDB(localAddress, dbConfig)
	// if err != nil {
	// 	log.Fatalf("Failed to connect to the database: %v", err)
	// }
	// defer db.Close()

	jobRecord := &ProcessJobCount{}
	workerPool := NewWorkerPool(10, nil, jobRecord)

	// rows, err := db.Query("SELECT name FROM partitions")
	// if err != nil {
	// 	log.Fatalf("Failed to query partitions: %v", err)
	// }
	// defer rows.Close()

	for i := 0; i < 40; i++ {
		var name string
		// if err := rows.Scan(&name); err != nil {
		// 	ErrorLogger.Printf("Failed to scan partition name: %v", err)
		// 	continue
		// }
		workerPool.AddJob(Job{tableName: name})
	}

	close(workerPool.jobs)
	workerPool.Start()
	workerPool.wg.Wait()
}
