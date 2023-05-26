//
// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// This is a commercial product and requires a license to operate.
// A trial license can be obtained at https://unidoc.io
//
// DO NOT EDIT: generated by unitwist Go source code obfuscator.
//
// Use of this source code is governed by the UniDoc End User License Agreement
// terms that can be accessed at https://unidoc.io/eula/

package logger ;import (_a "fmt";_c "io";_g "os";_af "path/filepath";_e "runtime";);

// Notice logs notice message.
func (_fc ConsoleLogger )Notice (format string ,args ...interface{}){if _fc .LogLevel >=LogLevelNotice {_ge :="\u005bN\u004f\u0054\u0049\u0043\u0045\u005d ";_fc .output (_g .Stdout ,_ge ,format ,args ...);};};

// Error does nothing for dummy logger.
func (DummyLogger )Error (format string ,args ...interface{}){};

// IsLogLevel returns true from dummy logger.
func (DummyLogger )IsLogLevel (level LogLevel )bool {return true };

// NewConsoleLogger creates new console logger.
func NewConsoleLogger (logLevel LogLevel )*ConsoleLogger {return &ConsoleLogger {LogLevel :logLevel }};

// WriterLogger is the logger that writes data to the Output writer
type WriterLogger struct{LogLevel LogLevel ;Output _c .Writer ;};

// Logger is the interface used for logging in the unipdf package.
type Logger interface{Error (_f string ,_gg ...interface{});Warning (_fb string ,_ee ...interface{});Notice (_cg string ,_gf ...interface{});Info (_cd string ,_b ...interface{});Debug (_gff string ,_cdc ...interface{});Trace (_ac string ,_cb ...interface{});IsLogLevel (_gfc LogLevel )bool ;};

// Warning logs warning message.
func (_da ConsoleLogger )Warning (format string ,args ...interface{}){if _da .LogLevel >=LogLevelWarning {_ef :="\u005b\u0057\u0041\u0052\u004e\u0049\u004e\u0047\u005d\u0020";_da .output (_g .Stdout ,_ef ,format ,args ...);};};

// NewWriterLogger creates new 'writer' logger.
func NewWriterLogger (logLevel LogLevel ,writer _c .Writer )*WriterLogger {logger :=WriterLogger {Output :writer ,LogLevel :logLevel };return &logger ;};

// Debug does nothing for dummy logger.
func (DummyLogger )Debug (format string ,args ...interface{}){};func _dg (_db _c .Writer ,_fdc string ,_dbb string ,_gda ...interface{}){_ ,_gbg ,_gea ,_ga :=_e .Caller (3);if !_ga {_gbg ="\u003f\u003f\u003f";_gea =0;}else {_gbg =_af .Base (_gbg );};_df :=_a .Sprintf ("\u0025s\u0020\u0025\u0073\u003a\u0025\u0064 ",_fdc ,_gbg ,_gea )+_dbb +"\u000a";_a .Fprintf (_db ,_df ,_gda ...);};

// Warning logs warning message.
func (_gb WriterLogger )Warning (format string ,args ...interface{}){if _gb .LogLevel >=LogLevelWarning {_acd :="\u005b\u0057\u0041\u0052\u004e\u0049\u004e\u0047\u005d\u0020";_gb .logToWriter (_gb .Output ,_acd ,format ,args ...);};};

// Info logs info message.
func (_eb WriterLogger )Info (format string ,args ...interface{}){if _eb .LogLevel >=LogLevelInfo {_ecf :="\u005bI\u004e\u0046\u004f\u005d\u0020";_eb .logToWriter (_eb .Output ,_ecf ,format ,args ...);};};

// IsLogLevel returns true if log level is greater or equal than `level`.
// Can be used to avoid resource intensive calls to loggers.
func (_gd ConsoleLogger )IsLogLevel (level LogLevel )bool {return _gd .LogLevel >=level };

// Info logs info message.
func (_cc ConsoleLogger )Info (format string ,args ...interface{}){if _cc .LogLevel >=LogLevelInfo {_bc :="\u005bI\u004e\u0046\u004f\u005d\u0020";_cc .output (_g .Stdout ,_bc ,format ,args ...);};};

// Debug logs debug message.
func (_ecc WriterLogger )Debug (format string ,args ...interface{}){if _ecc .LogLevel >=LogLevelDebug {_dae :="\u005b\u0044\u0045\u0042\u0055\u0047\u005d\u0020";_ecc .logToWriter (_ecc .Output ,_dae ,format ,args ...);};};var Log Logger =DummyLogger {};

// DummyLogger does nothing.
type DummyLogger struct{};

// Trace logs trace message.
func (_gfcc ConsoleLogger )Trace (format string ,args ...interface{}){if _gfcc .LogLevel >=LogLevelTrace {_ad :="\u005b\u0054\u0052\u0041\u0043\u0045\u005d\u0020";_gfcc .output (_g .Stdout ,_ad ,format ,args ...);};};

// ConsoleLogger is a logger that writes logs to the 'os.Stdout'
type ConsoleLogger struct{LogLevel LogLevel ;};func (_ce ConsoleLogger )output (_acc _c .Writer ,_bb string ,_ec string ,_eef ...interface{}){_dg (_acc ,_bb ,_ec ,_eef ...);};

// Warning does nothing for dummy logger.
func (DummyLogger )Warning (format string ,args ...interface{}){};

// Notice does nothing for dummy logger.
func (DummyLogger )Notice (format string ,args ...interface{}){};

// Error logs error message.
func (_efe WriterLogger )Error (format string ,args ...interface{}){if _efe .LogLevel >=LogLevelError {_ace :="\u005b\u0045\u0052\u0052\u004f\u0052\u005d\u0020";_efe .logToWriter (_efe .Output ,_ace ,format ,args ...);};};func (_dag WriterLogger )logToWriter (_aec _c .Writer ,_afb string ,_be string ,_ccg ...interface{}){_dg (_aec ,_afb ,_be ,_ccg );};

// Info does nothing for dummy logger.
func (DummyLogger )Info (format string ,args ...interface{}){};

// SetLogger sets 'logger' to be used by the unidoc unipdf library.
func SetLogger (logger Logger ){Log =logger };

// Debug logs debug message.
func (_fd ConsoleLogger )Debug (format string ,args ...interface{}){if _fd .LogLevel >=LogLevelDebug {_de :="\u005b\u0044\u0045\u0042\u0055\u0047\u005d\u0020";_fd .output (_g .Stdout ,_de ,format ,args ...);};};

// Error logs error message.
func (_ag ConsoleLogger )Error (format string ,args ...interface{}){if _ag .LogLevel >=LogLevelError {_ae :="\u005b\u0045\u0052\u0052\u004f\u0052\u005d\u0020";_ag .output (_g .Stdout ,_ae ,format ,args ...);};};

// LogLevel is the verbosity level for logging.
type LogLevel int ;

// Trace logs trace message.
func (_aa WriterLogger )Trace (format string ,args ...interface{}){if _aa .LogLevel >=LogLevelTrace {_ea :="\u005b\u0054\u0052\u0041\u0043\u0045\u005d\u0020";_aa .logToWriter (_aa .Output ,_ea ,format ,args ...);};};

// Notice logs notice message.
func (_dec WriterLogger )Notice (format string ,args ...interface{}){if _dec .LogLevel >=LogLevelNotice {_eec :="\u005bN\u004f\u0054\u0049\u0043\u0045\u005d ";_dec .logToWriter (_dec .Output ,_eec ,format ,args ...);};};const (LogLevelTrace LogLevel =5;LogLevelDebug LogLevel =4;LogLevelInfo LogLevel =3;LogLevelNotice LogLevel =2;LogLevelWarning LogLevel =1;LogLevelError LogLevel =0;);

// Trace does nothing for dummy logger.
func (DummyLogger )Trace (format string ,args ...interface{}){};

// IsLogLevel returns true if log level is greater or equal than `level`.
// Can be used to avoid resource intensive calls to loggers.
func (_fdg WriterLogger )IsLogLevel (level LogLevel )bool {return _fdg .LogLevel >=level };