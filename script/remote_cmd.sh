#!/bin/bash

REMOTE_CMD_EXP=./remote_cmd.exp

function runcmd()
{
	if [[ $# -lt 2 ]]; then
		echo "Usage: runcmd host command"
		return
	fi

	local host="$1"
	local command="$2"

	echo ">>> begin do remote_cmd $host : $command, please wait"
	$REMOTE_CMD_EXP runcmd "$host" "$command"
}

function copyfile()
{
	if [[ $# -lt 2 ]]; then
		echo "Usage: copyfile from to"
		return
	fi

	local from="$1"
	local to="$2"

	echo ">>> begin scp file [$from-->$to] please wait ..."
	$REMOTE_CMD_EXP scp "$from" "$to"
}

function putfile()
{
	local host="$1"
	local from="$2"
	local to="$3"

	copyfile "$from" "$host:$to"
}
