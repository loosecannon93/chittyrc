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
    "log"
    "io"
    "io/ioutil"
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
    QUIET   :"",
    CRITICAL:"CRITICAL: ",
    ERROR   :"ERROR: ",
    WARNING :"WARNING: ",
    INFO    :"INFO: ",
    DEBUG   :"DEBUG: ",
    TRACE   :"TRACE: ",
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
    TRACE   :nil,
}



func Init(level LogLevel) { 
    /* Lambda to determine what the output is */ 
    output := func(l LogLevel ) io.Writer { 
      if l > level { 
        return ioutil.Discard 
      } else { 
        return os.Stdout 
      } 
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
func Logf(level LogLevel, LogLevel, format string, a...interface{}) {
  loggers[level].Printf(format, a...)
}
  


