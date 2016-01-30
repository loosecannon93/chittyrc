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
    "fmt"
    "log"
    "io"
    "io/ioutil"
    "os"
)

/* Log levels */
type LogLevel int
const (
    QUIET    LogLevel = 10 * iota
    CRITICAL LogLevel 
    ERROR    LogLevel 
    WARNING  LogLevel 
    INFO     LogLevel 
    DEBUG    LogLevel 
    TRACE    LogLevel 
)

/* map levels to prefix strings */
const prefixes = map[LogLevel]String{
    QUIET   :"",
    CRITICAL:"CRITICAL: ",
    ERROR   :"ERROR: ",
    WARNING :"WARNING: ",
    INFO    :"INFO: ",
    DEBUG   :"DEBUG: ",
    TRACE   :"TRACE: "
}

/* Actual Logger lobjects to log to. 
 * Init initializes these, setting those below the log level
 * to discard their inputs using ioutil.DIscard 
*/
var loggers = map[LogLevel](*log.Logger){
    QUIET   :nil,
    CRITICAL:nil,
    ERROR   :nil,
    WARNING :nil,
    INFO    :nil,
    DEBUG   :nil,
    TRACE   :nil
}



func Init(level LogLevel) { 
    /* Lambda to determine what the output is */ 
    output := func(l LogLevel ) io.Writer { 
      if l > level { return ioutil.Discard } 
      else { return os.Stdout } 
    } 

    for l := range loggers { 
      loggers[l] = log.New(output(l), prefixes[l],
        log.Ldate|log.Ltime|log.Lshortfile)
    }
}
  

/*
 * Log - Print a log message
 *
 * level: Logging level of the message
 *
 * ...: Extra parameters to pass to Println
 *
 * Returns: nothing.
 */
func Log(level LogLevel, a...interface{}) { 
  loggers[level].Println(a...)
}

/* Logf - Print a fomatted message
 * 
 * Same as Log, but takes a format string and passes to 
 * Printf rather than Println
 */
func Logf(level, LogLevel, format String, a...interface{}) {
  loggers[level].Printf(format, a...)
}
  


