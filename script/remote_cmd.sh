#!/bin/bash

default_passwd="#Bugsfor\$Linux"
REMOTE_CMD_EXP=./remote_cmd.exp

function runcmd()
{
	if [[ $# -lt 2 ]]; then
		echo "Usage: runcmd host command [password]"
		return
	fi

	local host="$1"
	local command="$2"
	local password=${3:-$default_passwd}

	echo ">>> begin do remote_cmd $host : $command, please wait"
	$REMOTE_CMD_EXP runcmd "$host" "$password" "$command"
}

function runcmdroot()
{
	if [[ $# -lt 2 ]]; then
		echo "Usage: runcmdroot host command"
		return
	fi

	local host="$1"
	local command="$2"

	echo ">>> begin do remote_cmd_root $host : $command, please wait"
	$REMOTE_CMD_EXP runcmdroot "$host" "$command"
}

function runinstall()
{
	if [[ $# -lt 2 ]]; then
		echo "Usage: runinstall host command"
		return
	fi

	local host="$1"
	local command="$2"

	echo ">>> begin do remote_install $host : $command, please wait"
	$REMOTE_CMD_EXP runinstall "$host" "$command"
}

function copyfile()
{
	if [[ $# -lt 2 ]]; then
		echo "Usage: copyfile from to [passwordfrom] [passwordto]"
		return
	fi

	local from="$1"
	local to="$2"
	local passwordfrom=${3:-"$default_passwd"}
	local passwordto=${4:-"$passwordfrom"}

	echo ">>> begin scp file [$from-->$to] please wait ..."
	$REMOTE_CMD_EXP scp "$from" "$to" "$passwordfrom" "$passwordto"
}

function getfile()
{
	local host="$1"
	local from="$2"
	local to="$3"

	copyfile "$host:$from" "$to"
}

function putfile()
{
	local host="$1"
	local from="$2"
	local to="$3"

	copyfile "$from" "$host:$to"
}

function rsync_ex()
{
	if [[ $# -lt 2 ]]; then
		echo "Usage: rsync_ex from to [password]"
		return
	fi

	local from="$1"
	local to="$2"
	local password=${3:-$default_passwd}

	$REMOTE_CMD_EXP rsync "$from" "$to" "$password"
}

function shell()
{
	if [[ $# -lt 1 ]]; then
		echo "Usage: shell ip [command] [password]"
		return
	fi

	local host="$1"
	local command="$2"
	local password=${3:-$default_passwd}

	$REMOTE_CMD_EXP ssh "$host" "$password" "$command"
}
