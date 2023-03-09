echo "Check if exists a .env"
FILE=../.env
if [ -f "$FILE" ]; then

echo "File exists, removing"
    rm $FILE
fi

echo "Creating an empty .env"
touch ../.env

echo "Populanting with configs"
echo DB_NAME=$DB_NAME >> ../.env
echo DB_HOST=$DB_HOST >> ../.env
echo DB_USER=$DB_USER >> ../.env
echo DB_PORT=$DB_PORT >> ../.env
echo DB_PASS=$DB_PASS >> ../.env

echo "Showing content"
cat ../.env