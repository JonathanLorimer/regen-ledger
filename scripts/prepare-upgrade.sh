#!/usr/bin/env bash
set -x
if  [ $UPGRADE_COMMIT != "null" ]; then
    echo "Doing upgrade to $UPGRADE_COMMIT"
    cd $REGEN_LEDGER_REPO
    git fetch
    git clean -f
    git checkout -f $UPGRADE_COMMIT
    nixos-rebuild build
fi
