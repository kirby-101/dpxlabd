#!/bin/sh
#
# PROVIDE: dpxlabd
# REQUIRE: DAEMON NETWORKING
# KEYWORD: shutdown

#
# Configuration (rc.conf)
#
# hbsdsrv_www_enable (bool):    Set to NO by default.
#
# hbsdsrv_www_flags: (string):  Custom Argv
#
# hbsdsrv_www_config (string):  Config File
#
# hbsdsrv_www_logfile (string): Log File

. /etc/rc.subr

name=dpxlabd
rcvar=dpxlabd_enable


# 
dpxlabd_enable=${dpxlabd_enable-:"NO"}
dpxlabd_flags=${dpxlabd_flags:-""}

dpxlabd_configfile=${dpxlabd_configfile:-"/www/${name}.yml"}
dpxlabd_pidfile=${dpxlabd_pidfile-:"/var/run/${name}.pid"}

#required_files=

# cmds + argv
command="/www/dpxlabd"
command_args="--configfile ${dpxlabd_configfile}"


load_rc_config $name

# 
#pidfile=""
#procname
run_rc_command "$1"
