#!/bin/bash
set -o pipefail
> errors.txt
> run.log
GHA2DB_PROJECT=tuf IDB_DB=tuf PG_DB=tuf GHA2DB_LOCAL=1 ./structure 2>>errors.txt | tee -a run.log || exit 1
GHA2DB_EXCLUDE_REPOS='theupdateframework/notary' GHA2DB_PROJECT=tuf IDB_DB=tuf PG_DB=tuf GHA2DB_LOCAL=1 ./gha2db 2015-01-01 0 today now 'theupdateframework' 2>>errors.txt | tee -a run.log || exit 2
GHA2DB_PROJECT=tuf IDB_DB=tuf PG_DB=tuf GHA2DB_LOCAL=1 GHA2DB_EXACT=1 GHA2DB_OLDFMT=1 ./gha2db 2014-01-02 0 2014-12-31 23 'theupdateframework,tuf' 2>>errors.txt | tee -a run.log || exit 3
GHA2DB_PROJECT=tuf IDB_DB=tuf PG_DB=tuf GHA2DB_LOCAL=1 GHA2DB_MGETC=y GHA2DB_SKIPTABLE=1 GHA2DB_INDEX=1 ./structure 2>>errors.txt | tee -a run.log || exit 4
./grafana/influxdb_recreate.sh tuf
./tuf/setup_repo_groups.sh 2>>errors.txt | tee -a run.log || exit 5
./tuf/import_affs.sh 2>>errors.txt | tee -a run.log || exit 6
./tuf/setup_scripts.sh 2>>errors.txt | tee -a run.log || exit 7
./tuf/get_repos.sh 2>>errors.txt | tee -a run.log || exit 8
echo "All done. You should run ./tuf/reinit.sh script now."
