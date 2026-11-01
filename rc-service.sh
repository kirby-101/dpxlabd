#!/bin/sh
#
# PROVIDE: dpxlabd
# REQUIRE: DAEMON NETWORKING
# KEYWORD: shutdown

#
# Configuration (rc.conf)
#
# dpxlabd_enable (bool)      : Set to NO by default.
# dpxlabd_flags: (string)    : Custom Argv
# dpxlabd_configfile (string): Config File
# dpxlabd_pidfile (string)   : Log File

. /etc/rc.subr

name=dpxlabd
rcvar=dpxlabd_enable

${dpxlabd_enable:-"NO"}
${dpxlabd_flags:-""}

${dpxlabd_configfile:-"/usr/local/etc/dpxlabd/config.yml"}
${dpxlabd_pidfile:-"/var/run/${name}.pid"}

#required_files=

# cmds + argv
command="/www/dpxlabd"
command_args="--configfile ${dpxlabd_configfile}"

load_rc_config $name

# 
#pidfile=""
#procname
run_rc_command "$1"
