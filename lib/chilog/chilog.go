/*
 *  Logging functions
 *
 *  Use these functions to print log messages. Each message has an
 *  associated log level:
 *
 *  CRITICAL: A critical unrecoverable error
 *  ERROR: A recoverable error
 *  WARNING: A warning
 *  INFO: High-level information about the progress of the application
 *  DEBUG: Lower-level information
 *  TRACE: Very low-level information.
 *
 */

package chilog

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

/* Log levels */
type LogLevel int

const (
	QUIET    LogLevel = 10 * iota
	CRITICAL LogLevel = 10 * iota
	ERROR    LogLevel = 10 * iota
	WARNING  LogLevel = 10 * iota
	INFO     LogLevel = 10 * iota
	DEBUG    LogLevel = 10 * iota
	TRACE    LogLevel = 10 * iota
)

/* map levels to prefix strings */
var prefixes = map[LogLevel]string{
	QUIET:    "",
	CRITICAL: "CRITIC: ",
	ERROR:    "ERROR : ",
	WARNING:  "WARN  : ",
	INFO:     "INFO  : ",
	DEBUG:    "DEBUG : ",
	TRACE:    "TRACE : ",
}

/* Actual Logger lobjects to log to.
 * Init initializes these, setting those below the log level
 * to discard their inputs using ioutil.DIscard
 */
var (
	Critical *log.Logger
	Error    *log.Logger
	Warning  *log.Logger
	Info     *log.Logger
	Debug    *log.Logger
	Trace    *log.Logger
)

func Init(level LogLevel) {
	/* Lambda to determine what the output is */
	output := func(l LogLevel) io.Writer {
		if l > level {
			return ioutil.Discard
		} else {
			return os.Stdout
		}
	}

	Critical = log.New(output(CRITICAL), prefixes[CRITICAL],
		log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(output(ERROR), prefixes[ERROR],
		log.Ldate|log.Ltime|log.Lshortfile)
	Warning = log.New(output(WARNING), prefixes[WARNING],
		log.Ldate|log.Ltime|log.Lshortfile)
	Info = log.New(output(INFO), prefixes[INFO],
		log.Ldate|log.Ltime|log.Lshortfile)
	Debug = log.New(output(DEBUG), prefixes[DEBUG],
		log.Ldate|log.Ltime|log.Lshortfile)
	Trace = log.New(output(TRACE), prefixes[TRACE],
		log.Ldate|log.Ltime|log.Lshortfile)

	Info.Println("Initialized Logging system with level ", prefixes[level])
	Info.Println("Will Print log Messages at the following levels:")
	Critical.Println("Critical test")
	Error.Println("Error test")
	Warning.Println("Warning test")
	Info.Println("Info test")
	Debug.Println("Debug test")
	Trace.Println("Trace test")
	Info.Println("Use -v[v] to increase the logging level or -q to turn all off")
}
