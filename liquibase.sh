
PROJ_DIR=$(pwd)/db

DB_HOST=localhost
DB_PORT=3306
DB_NAME=todoapp

DB_USER=root
DB_PASSWORD=password

PROPERTY_FILE=liquibase.properties
CHANGELOG_FILE=changelog-master

FORMAT=xml

URL=jdbc:mysql://$DB_HOST:$DB_PORT/$DB_NAME

#echo url: $URL

#liquibase init project \
#  --project-dir="$PROJ_DIR" \
#  --changelog-file="$CHANGELOG_FILE" \
#  --format=$FORMAT \
#  --project-defaults-file=$PROPERTY_FILE \
#  --url="$URL" \
#  --username=$DB_USER \
#  --password=$DB_PASSWORD \
#  --log-level info

liquibase changelog-sync  --project-dir="$PROJ_DIR"