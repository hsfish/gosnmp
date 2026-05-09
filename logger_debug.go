//go:build !gosnmp_nodebug
// +build !gosnmp_nodebug

package gosnmp

func (l *Logger) Print(v ...interface{}) {
	if l.logger != nil {
		l.logger.Print(v...)
	}
}

func (l *Logger) Printf(format string, v ...interface{}) {
	if l.logger != nil {
		l.logger.Printf(format, v...)
	}
}

func (l *Logger) Warn(format string, v ...interface{}) {
	if l.logger != nil {
		l.logger.Printf(format, v...)
	}
}

func (l *Logger) Error(format string, v ...interface{}) {
	if l.logger != nil {
		l.logger.Printf(format, v...)
	}
}
