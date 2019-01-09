package main

import (
	"github.com/julienschmidt/httprouter"
)

func addDebugHandlers(router *httprouter.Router) {
	router.GET("/debug/dmesg", makeRequestHandlerCommand(
		"kernel logs",
		"dmesg"))
	router.GET("/debug/processes", makeRequestHandlerCommand(
		"process list",
		"ps", "-eo", "user,pid,ppid,c,stime,tty,%cpu,%mem,vsz,rsz,cmd"))
	router.GET("/debug/logs", makeRequestHandlerCommand(
		"logs",
		"journalctl", "--reverse", "-b", "--no-pager", "-n", "50"))
	router.GET("/debug/meminfo", makeRequestHandlerFile(
		"memory data",
		"/proc/meminfo"))
}
