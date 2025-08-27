DATE=$(date +%Y%m%d_%H%M%S)
BACKUP_DIR="/backup/mattermost"
DB_NAME="mattermost"
DB_USER="root"
DB_HOST="localhost"

mkdir -p $BACKUP_DIR

# Backup database (via pg_dump)
pg_dump -U $DB_USER -h $DB_HOST $DB_NAME | gzip > $BACKUP_DIR/mattermost_db_$DATE.sql.gz

# Backup file storage (config, data, plugins)
tar -czf $BACKUP_DIR/mattermost_files_$DATE.tar.gz /opt/mattermost/{config,data,plugins}

# Hapus backup lebih dari 7 hari
find $BACKUP_DIR -type f -mtime +7 -delete